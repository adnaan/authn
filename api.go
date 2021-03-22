package authn

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fatih/structs"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"

	"github.com/google/uuid"
	"github.com/markbates/goth/gothic"

	"github.com/adnaan/authn/models/account"

	"golang.org/x/crypto/bcrypt"

	"github.com/adnaan/authn/models"

	"github.com/markbates/goth"
)

const defaultMaxAge = 60 * 60 * 24 * 30 // 30 days
const defaultPath = "/"
const (
	formContentType = "application/x-www-form-urlencoded"
	CtxUserIdKey    = "key_user_id"
	sessionStoreKey = "auth"
)

type Config struct {
	Driver        string
	Datasource    string
	SessionSecret string
	SendMail      SendMailFunc
	SessionMaxAge int
	SessionPath   string
	GothProviders []goth.Provider
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
	sessionStore SessionsStore

	// database client
	client *models.Client
}

// New authn client
func New(ctx context.Context, cfg Config) *API {

	client, err := models.Open(cfg.Driver, cfg.Datasource)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}
	if cfg.SessionMaxAge == 0 {
		cfg.SessionMaxAge = defaultMaxAge
	}
	if cfg.SessionPath == "" {
		cfg.SessionPath = defaultPath
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
		sessionStore: sessionStore,
		client:       client,
	}
}

// Signup a new account with email and password
func (a *API) Signup(ctx context.Context, email, password string, attributes map[string]interface{}) error {
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

	exists, err := a.client.Account.Query().Where(account.Email(email)).Exist(ctx)
	if err != nil {
		return fmt.Errorf("%w", ErrInternal)
	}
	if exists {
		return fmt.Errorf("%w", ErrUserExists)
	}

	_, err = a.newAccount(ctx, email, hashedPassword, "email_signup", attributes, true)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	return nil
}

// ConfirmSignupEmail confirms the email used in signup
func (a *API) ConfirmSignupEmail(ctx context.Context, token string) error {
	acc, err := a.client.Account.Query().Where(account.ConfirmationToken(token)).Only(ctx)
	if err != nil {
		return fmt.Errorf("%w", ErrUserNotFound)
	}
	acc, err = acc.Update().
		SetConfirmed(true).
		ClearConfirmationToken().
		ClearConfirmationSentAt().
		Save(ctx)
	if err != nil {
		return fmt.Errorf("%w", ErrInternal)
	}
	return nil
}

// Login logs in the account with  email and password
func (a *API) Login(w http.ResponseWriter, r *http.Request, email, password string) error {

	if len(password) < 8 {
		return fmt.Errorf("%w", ErrInvalidPassword)
	}

	if !isEmailValid(email) {
		return fmt.Errorf("%w", ErrInvalidEmail)
	}

	acc, err := a.client.Account.Query().Where(account.Email(email)).Only(r.Context())
	if err != nil {
		return fmt.Errorf("%w", ErrUserNotFound)
	}

	if !acc.Confirmed {
		return fmt.Errorf("%w", ErrEmailNotConfirmed)
	}

	err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password))
	if err != nil {
		return fmt.Errorf("err: %v, %w", err, ErrWrongPassword)
	}

	// set provider
	acc, err = acc.Update().SetProvider("email_signup").Save(r.Context())
	if err != nil {
		return err
	}

	session, err := a.sessionStore.Get(r, sessionStoreKey)
	if err != nil {
		return fmt.Errorf("err: %v, %w", err, ErrLoginSessionNotFound)
	}

	session.Values["id"] = acc.ID.String()

	err = session.Save(r, w)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	return nil
}

// SendPasswordlessToken sends a passwordless login token to a previously signed up email
func (a *API) SendPasswordlessToken(ctx context.Context, email string) error {

	if !isEmailValid(email) {
		return fmt.Errorf("%w", ErrInvalidEmail)
	}

	acc, err := a.client.Account.Query().Where(account.Email(email)).Only(ctx)
	if err != nil {
		return fmt.Errorf("%w", ErrUserNotFound)
	}

	otp := uuid.New().String()

	err = withTx(ctx, a.client, func(tx *models.Tx) error {

		acc, err = acc.Update().SetOtp(otp).Save(ctx)
		if err != nil {
			return err
		}

		err = a.cfg.SendMail(Passwordless, otp, email, acc.Attributes)
		if err != nil {
			return err
		}

		acc, err = acc.Update().SetOtpSentAt(time.Now()).Save(ctx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	return nil
}

// LoginWithPasswordlessToken authenticates the token sent to email previously with SendPasswordlessLoginLink and logs in the account
func (a *API) LoginWithPasswordlessToken(w http.ResponseWriter, r *http.Request, loginToken string) error {
	acc, err := a.client.Account.Query().Where(account.Otp(loginToken)).Only(r.Context())
	if err != nil {
		return fmt.Errorf("%w", ErrUserNotFound)
	}

	session, err := a.sessionStore.Get(r, sessionStoreKey)
	if err != nil {
		return fmt.Errorf("err: %v, %w", err, ErrLoginSessionNotFound)
	}

	session.Values["id"] = acc.ID.String()

	err = session.Save(r, w)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	acc, err = acc.Update().ClearOtp().ClearOtpSentAt().Save(r.Context())
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	return nil
}

// LoginWithProvider attempts to log in an existing account. If not successful, redirects to the provider signup flow for authentication.
// The callback from the provider redirect is to be handled by `LoginProviderCallback`
func (a *API) LoginWithProvider(w http.ResponseWriter, r *http.Request) error {
	// try to get the user without re-authenticating
	if usr, err := gothic.CompleteUserAuth(w, r); err == nil {
		acc, err := a.client.Account.Query().Where(account.Email(usr.Email)).Only(r.Context())
		if err != nil {
			return fmt.Errorf("%w", ErrUserNotFound)
		}
		session, err := a.sessionStore.Get(r, sessionStoreKey)
		if err != nil {
			return fmt.Errorf("err: %v, %w", err, ErrLoginSessionNotFound)
		}

		session.Values["id"] = acc.ID.String()

		err = session.Save(r, w)
		if err != nil {
			return fmt.Errorf("%v %w", err, ErrInternal)
		}
	} else {
		gothic.BeginAuthHandler(w, r)
	}
	return nil
}

// LoginProviderCallback handles the callback from a provider. It authenticates an account using a third-party provider(https://github.com/markbates/goth).
// A new account is created if not present.
func (a *API) LoginProviderCallback(w http.ResponseWriter, r *http.Request, attributes map[string]interface{}) error {
	usr, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
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
	acc, err = a.client.Account.Query().Where(account.Email(usr.Email)).Only(r.Context())
	if err != nil {
		// err is not nil and not found
		if !models.IsNotFound(err) {
			return fmt.Errorf("%w", ErrInternal)
		}

		// user not found, create a user
		acc, err = a.newAccount(r.Context(), usr.Email, usr.AccessToken, usr.Provider, attributes, false)
		if err != nil {
			return err
		}
	} else {
		// just update the provider and attributes
		_, err = acc.Update().SetProvider(usr.Provider).SetAttributes(attributes).Save(r.Context())
		if err != nil {
			return err
		}
	}

	session, err := a.sessionStore.Get(r, sessionStoreKey)
	if err != nil {
		return fmt.Errorf("err: %v, %w", err, ErrLoginSessionNotFound)
	}

	session.Values["id"] = acc.ID.String()
	err = session.Save(r, w)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	return nil
}

func (a *API) Recovery(ctx context.Context, email string) error {

	if !isEmailValid(email) {
		return fmt.Errorf("%w", ErrInvalidEmail)
	}

	acc, err := a.client.Account.Query().Where(account.Email(email)).Only(ctx)
	if err != nil {
		return fmt.Errorf("%w", ErrUserNotFound)
	}

	recoveryToken := uuid.New().String()
	err = withTx(ctx, a.client, func(tx *models.Tx) error {

		acc, err = acc.Update().SetRecoveryToken(recoveryToken).Save(ctx)
		if err != nil {
			return err
		}

		err = a.cfg.SendMail(Recovery, recoveryToken, email, acc.Attributes)
		if err != nil {
			return err
		}

		acc, err = acc.Update().SetRecoverySentAt(time.Now()).Save(ctx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	return nil
}

func (a *API) ConfirmRecovery(ctx context.Context, token, password string) error {

	if len(password) < 8 {
		return fmt.Errorf("%w", ErrInvalidPassword)
	}

	acc, err := a.client.Account.Query().Where(account.RecoveryToken(token)).Only(ctx)
	if err != nil {
		return fmt.Errorf("%w", ErrUserNotFound)
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	err = acc.Update().SetPassword(hashedPassword).Exec(ctx)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	return nil
}

// CurrentAccount retrieves the current logged in current
func (a *API) CurrentAccount(r *http.Request) (*Account, error) {
	session, err := a.sessionStore.Get(r, sessionStoreKey)
	if err != nil {
		return nil, fmt.Errorf("%v, %w", err, ErrLoginSessionNotFound)
	}

	id, ok := session.Values["id"]
	if !ok {
		return nil, fmt.Errorf("%w", ErrUserNotLoggedIn)
	}

	userID, ok := id.(string)
	if !ok {
		return nil, fmt.Errorf("%v %w", err, ErrUserNotLoggedIn)
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("%v %w", err, ErrLoginSessionNotFound)
	}

	acc, err := a.client.Account.Get(r.Context(), uid)
	if err != nil {
		return nil, fmt.Errorf("%v %w", err, ErrUserNotFound)
	}

	return &Account{
		acc:          acc,
		sessionStore: a.sessionStore,
		req:          r,
		client:       a.client,
	}, nil
}

// IsAuthenticated is an authentication middleware
// It assumes form content type. For API authentication, please use CurrentAccount to implement your own mw.
func (a *API) IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-type")
		if contentType == "" {
			contentType = formContentType
		}

		redirectTo := fmt.Sprintf("/login?from=%s", r.URL.Path)

		session, err := a.sessionStore.Get(r, sessionStoreKey)
		if err != nil {
			switch contentType {
			case formContentType:
				http.Redirect(w, r, redirectTo, http.StatusSeeOther)
				return

			}
			return
		}

		id, ok := session.Values["id"]
		if !ok {
			switch contentType {
			case formContentType:
				http.Redirect(w, r, redirectTo, http.StatusSeeOther)
				return
			}
		} else {
			ctx := r.Context()
			ctx = context.WithValue(ctx, CtxUserIdKey, id)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
