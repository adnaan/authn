package authzen

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"

	"github.com/fatih/structs"
	"github.com/google/uuid"
	"github.com/markbates/goth/gothic"

	"github.com/adnaan/authzen/models/account"

	"golang.org/x/crypto/bcrypt"

	"github.com/adnaan/authzen/models"

	"github.com/hako/branca"
	"github.com/markbates/goth"
	"github.com/zpatrick/rbac"
)

const defaultMaxAge = 60 * 60 * 24 * 30 // 30 days
const defaultPath = "/"

type PermissionMatcher struct {
	ActionMatcher rbac.Matcher
	TargetMatcher func(userID string) rbac.Matcher
}

type Config struct {
	Driver          string
	Datasource      string
	SessionSecret   string
	SendMail        SendMailFunc
	APIMasterSecret string
	GothProviders   []goth.Provider
	Roles           map[string][]PermissionMatcher
}

// hashPassword generates a hashed password from a plaintext string
func hashPassword(password string) (string, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(pw), nil
}

type API struct {
	cfg          Config
	ctx          context.Context
	branca       *branca.Branca
	sessionStore SessionsStore

	// database client
	c *models.Client
}

func New(ctx context.Context, cfg Config) *API {

	client, err := models.Open(cfg.Driver, cfg.Datasource)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}
	sessionStore := &defaultSessionStore{
		Ctx:    ctx,
		Client: client,
		Codecs: securecookie.CodecsFromPairs([]byte(cfg.SessionSecret)),
		SessionOpts: &sessions.Options{
			Path:   defaultPath,
			MaxAge: defaultMaxAge,
		}}

	if len(cfg.GothProviders) > 0 {
		goth.UseProviders(cfg.GothProviders...)
		gothic.Store = sessionStore
	}
	return &API{
		cfg:          cfg,
		ctx:          ctx,
		sessionStore: sessionStore,
		c:            client,
		branca:       branca.NewBranca(cfg.APIMasterSecret)}
}

func (a *API) NewEmailPasswordAccount(email, password, role string, attributes map[string]interface{}) error {
	if len(password) < 8 {
		return fmt.Errorf("%w", ErrInvalidPassword)
	}

	if !isEmailValid(email) {
		return fmt.Errorf("%w", ErrInvalidEmail)
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	exists, err := a.c.Account.Query().Where(account.Email(email)).Exist(a.ctx)
	if err != nil {
		return fmt.Errorf("%w", ErrInternal)
	}
	if exists {
		return fmt.Errorf("%w", ErrUserExists)
	}

	_, err = a.newAccount(email, hashedPassword, role, "email_signup", attributes, true)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	return nil
}

func (a *API) NewProviderAccount(w http.ResponseWriter, r *http.Request, role string, attributes map[string]interface{}) error {
	usr, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		return err
	}

	type User struct {
		RawData           map[string]interface{}
		Provider          string
		Email             string
		Name              string
		FirstName         string
		LastName          string
		NickName          string
		Description       string
		UserID            string
		AvatarURL         string
		Location          string
		AccessToken       string
		AccessTokenSecret string
		RefreshToken      string
		ExpiresAt         time.Time
		IDToken           string
	}

	if attributes == nil {
		attributes = map[string]interface{}{
			"name": usr.Name,
		}
	} else {
		attributes["name"] = usr.Name
	}

	userMap := structs.Map(usr)

	for k, v := range userMap {
		attributes[k] = v
	}

	var acc *models.Account
	// find existing account or create one
	acc, err = a.c.Account.Query().Where(account.Email(usr.Email)).Only(a.ctx)
	if err != nil {
		// err is not nil and not found
		if !models.IsNotFound(err) {
			return fmt.Errorf("%w", ErrInternal)
		}

		// user not found, create a user
		acc, err = a.newAccount(usr.Email, usr.AccessToken, role, usr.Provider, attributes, false)
		if err != nil {
			return err
		}
	} else {
		// just update the provider and attributes
		_, err = acc.Update().SetProvider(usr.Provider).SetAttributes(attributes).Save(a.ctx)
		if err != nil {
			return err
		}
	}

	session, err := a.sessionStore.Get(r, "auth-session")
	if err != nil {
		return fmt.Errorf("err: %v, %w", err, ErrLoginSessionNotFound)
	}

	token := uuid.New().String()
	session.Values["id"] = acc.ID.String()
	session.Values["token"] = token
	err = session.Save(r, w)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	return nil
}
