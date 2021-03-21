package authn

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hako/branca"

	"github.com/markbates/goth/gothic"

	"github.com/adnaan/authn/models"
	"github.com/google/uuid"
)

type Account struct {
	acc          *models.Account
	req          *http.Request
	cfg          Config
	sessionStore SessionsStore
	client       *models.Client
	branca       *branca.Branca
}

func (a *API) newAccount(ctx context.Context, email, password string, provider string, attributes map[string]interface{}, sendConfirm bool) (*models.Account, error) {
	var acc *models.Account
	var err error

	err = withTx(ctx, a.client, func(tx *models.Tx) error {
		// create user
		createAccount := tx.Account.Create().
			SetEmail(email).
			SetPassword(password).
			SetProvider(provider).
			SetAttributes(attributes)

		// send confirmation email
		if sendConfirm {
			confirmationToken := uuid.New().String()
			err = a.cfg.SendMail(Confirmation, confirmationToken, email, attributes)
			if err != nil {
				return fmt.Errorf("err sending confirmation email %v", err)
			}

			createAccount = createAccount.SetConfirmationToken(confirmationToken).SetConfirmationSentAt(time.Now())

		} else {
			if provider == "email_signup" {
				return fmt.Errorf("provider is email but sendEmailFunc is nil")
			}
			createAccount = createAccount.SetConfirmed(true)
		}

		// create user
		acc, err = createAccount.Save(ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return acc, err
	}

	return acc, nil
}

func logout(sessionStore SessionsStore, w http.ResponseWriter, r *http.Request) {
	session, err := sessionStore.Get(r, sessionStoreKey)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	session.Values = nil
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	_ = gothic.Logout(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *Account) Logout(w http.ResponseWriter, r *http.Request) {
	logout(a.sessionStore, w, r)
}

func (a *Account) ChangeEmail(newEmail string) error {
	if !isEmailValid(newEmail) {
		return fmt.Errorf("%w", ErrInvalidEmail)
	}

	return withTx(a.req.Context(), a.client, func(tx *models.Tx) error {
		var err error
		token := uuid.New().String()
		a.acc, err = tx.Account.UpdateOneID(a.acc.ID).
			SetEmailChange(newEmail).
			SetEmailChangeToken(token).Save(a.req.Context())
		if err != nil {
			return fmt.Errorf("%v %w", err, ErrInternal)
		}

		err = a.cfg.SendMail(ChangeEmail, token, a.acc.Email, a.acc.Attributes)
		if err != nil {
			return fmt.Errorf("%v %w", err, ErrInternal)
		}
		a.acc, err = tx.Account.UpdateOneID(a.acc.ID).
			SetEmailChangeSentAt(time.Now()).Save(a.req.Context())
		if err != nil {
			return fmt.Errorf("%v %w", err, ErrInternal)
		}
		return nil
	})
}

func (a *Account) ConfirmEmailChange(token string) error {

	if *a.acc.EmailChangeToken != token {
		return fmt.Errorf("%w", ErrInvalidToken)
	}

	err := a.acc.Update().
		SetEmail(a.acc.EmailChange).
		ClearEmailChangeToken().
		ClearEmailChangeSentAt().
		ClearEmailChange().
		Exec(a.req.Context())

	if err != nil {
		return fmt.Errorf("%v %w", err, ErrUpdatingUser)
	}

	return nil
}

func (a *Account) Delete() error {
	err := a.client.Account.DeleteOne(a.acc).Exec(a.req.Context())
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}
	return nil
}

func (a *Account) Attributes() *Attributes {
	return &Attributes{
		req:          a.req,
		acc:          a.acc,
		sessionStore: a.sessionStore,
		sessionAttributes: &SessionAttributes{
			req:          a.req,
			acc:          a.acc,
			sessionStore: a.sessionStore,
		},
		sensitiveAttributes: &SensitiveAttributes{
			req:          a.req,
			acc:          a.acc,
			sessionStore: a.sessionStore,
			branca:       a.branca,
		},
	}
}

type SessionAttributes struct {
	req          *http.Request
	acc          *models.Account
	sessionStore SessionsStore
}

// Get gets an attribute from the temporary session store for the account
func (a *SessionAttributes) Get(key string) (interface{}, error) {

	session, err := a.sessionStore.Get(a.req, sessionStoreKey)
	if err != nil {
		return nil, fmt.Errorf("err: %v, %w", err, ErrLoginSessionNotFound)
	}

	val, ok := session.Values[key]
	if !ok {
		return nil, fmt.Errorf("err: %v, %w", err, ErrAttributeNotFound)
	}

	return val, nil
}

// Set sets an attribute in the temporary session store for the account
func (a *SessionAttributes) Set(w http.ResponseWriter, key string, val interface{}) error {

	session, err := a.sessionStore.Get(a.req, sessionStoreKey)
	if err != nil {
		return fmt.Errorf("err: %v, %w", err, ErrLoginSessionNotFound)
	}

	session.Values[key] = val
	err = session.Save(a.req, w)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}

	return nil
}

// Del deletes an attribute from the temporary session store for the account
func (a *SessionAttributes) Del(w http.ResponseWriter, key string) error {

	session, err := a.sessionStore.Get(a.req, sessionStoreKey)
	if err != nil {
		return fmt.Errorf("err: %v, %w", err, ErrLoginSessionNotFound)
	}
	delete(session.Values, key)
	err = session.Save(a.req, w)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}
	return nil
}

type Attributes struct {
	req                 *http.Request
	acc                 *models.Account
	sessionStore        SessionsStore
	sessionAttributes   *SessionAttributes
	sensitiveAttributes *SensitiveAttributes
}

// Get gets an attribute from the database for the account
func (a *Attributes) Get(key string) (interface{}, error) {
	v, ok := a.acc.Attributes[key]
	if !ok {
		return nil, fmt.Errorf("%w", ErrAttributeNotFound)
	}
	return v, nil
}

// Set stores an attribute in the database for the account
func (a *Attributes) Set(key string, val interface{}) error {
	currentAttributes := a.acc.Attributes
	currentAttributes[key] = val

	err := a.acc.Update().SetAttributes(currentAttributes).Exec(a.req.Context())
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrUpdatingUser)
	}
	return nil
}

// Del deletes an attribute from the database for the account
func (a *Attributes) Del(key string) error {
	currentAttributes := a.acc.Attributes
	delete(currentAttributes, key)

	err := a.acc.Update().SetAttributes(currentAttributes).Exec(a.req.Context())
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrUpdatingUser)
	}
	return nil
}

// All gets all attributes from the database for the account
func (a *Attributes) All() map[string]interface{} {
	return a.acc.Attributes
}

// SetMany stores  many attributes in the database for the account
func (a *Attributes) SetMany(attrs map[string]interface{}) error {
	currentAttributes := a.acc.Attributes
	for k, v := range attrs {
		currentAttributes[k] = v
	}
	err := a.acc.Update().SetAttributes(currentAttributes).Exec(a.req.Context())
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrUpdatingUser)
	}
	return nil
}

// DelMany deletes many attributes from the database for the account
func (a *Attributes) DelMany(keys []string) error {
	currentAttributes := a.acc.Attributes
	for _, k := range keys {
		delete(currentAttributes, k)
	}
	err := a.acc.Update().SetAttributes(currentAttributes).Exec(a.req.Context())
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrUpdatingUser)
	}
	return nil
}

// SetBytes stores byte data in the database for the account
func (a *Attributes) SetBytes(attr []byte) error {
	err := a.acc.Update().SetAttributeBytes(attr).Exec(a.req.Context())
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrUpdatingUser)
	}
	return nil
}

// GetBytes gets byte data from the database for the account
func (a *Attributes) GetBytes() []byte {
	return a.acc.AttributeBytes
}

// DelBytes delete byte data from the database for the account
func (a *Attributes) DelBytes() error {
	err := a.acc.Update().ClearAttributeBytes().Exec(a.req.Context())
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrUpdatingUser)
	}
	return nil
}

func (a *Attributes) Session() *SessionAttributes {
	return a.sessionAttributes
}

func (a *Attributes) Sensitive() *SensitiveAttributes {
	return a.sensitiveAttributes
}

// SensitiveAttributes provides get, set, del methods for sensitive attributes for an account
type SensitiveAttributes struct {
	req          *http.Request
	acc          *models.Account
	sessionStore SessionsStore
	branca       *branca.Branca
}

// Get gets a sensitive attribute from the database for the account
// It uses https://github.com/hako/branca to encode/decode strings
func (a *SensitiveAttributes) Get(key string) (string, error) {
	v, ok := a.acc.SensitiveAttributes[key]
	if !ok {
		return "", fmt.Errorf("%w", ErrAttributeNotFound)
	}
	return a.branca.DecodeToString(v)
}

// Set stores a sensitive attribute in the database for the account
// It uses https://github.com/hako/branca to encode/decode strings
func (a *SensitiveAttributes) Set(key string, val string) error {
	currentAttributes := a.acc.SensitiveAttributes
	valx, err := a.branca.EncodeToString(val)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}
	currentAttributes[key] = valx

	err = a.acc.Update().SetSensitiveAttributes(currentAttributes).Exec(a.req.Context())
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrUpdatingUser)
	}
	return nil
}

// Del deletes a sensitive attribute from the database for the account
func (a *SensitiveAttributes) Del(key string) error {
	currentAttributes := a.acc.SensitiveAttributes
	delete(currentAttributes, key)

	err := a.acc.Update().SetSensitiveAttributes(currentAttributes).Exec(a.req.Context())
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrUpdatingUser)
	}
	return nil
}
