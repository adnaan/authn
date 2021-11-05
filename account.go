package authn

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/markbates/goth/gothic"

	"github.com/adnaan/authn/models"
	"github.com/google/uuid"
)

type Account struct {
	acc          *models.Account
	req *http.Request
	cfg          Config
	sessionStore SessionsStore
	client       *models.Client
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

func (a *Account) Logout(w http.ResponseWriter, r *http.Request) error {
	if a.sessionStore == nil {
		return fmt.Errorf("session is missing")
	}
	logout(a.sessionStore, w, r)
	return nil
}

func (a *Account) ChangeEmail(ctx context.Context, newEmail string) error {
	if !isEmailValid(newEmail) {
		return fmt.Errorf("%w", ErrInvalidEmail)
	}

	return withTx(ctx, a.client, func(tx *models.Tx) error {
		var err error
		token := uuid.New().String()
		a.acc, err = tx.Account.UpdateOneID(a.acc.ID).
			SetEmailChange(newEmail).
			SetEmailChangeToken(token).Save(ctx)
		if err != nil {
			return fmt.Errorf("%v %w", err, ErrInternal)
		}

		err = a.cfg.SendMail(ChangeEmail, token, a.acc.Email, a.acc.Attributes)
		if err != nil {
			return fmt.Errorf("%v %w", err, ErrInternal)
		}
		a.acc, err = tx.Account.UpdateOneID(a.acc.ID).
			SetEmailChangeSentAt(time.Now()).Save(ctx)
		if err != nil {
			return fmt.Errorf("%v %w", err, ErrInternal)
		}
		return nil
	})
}

func (a *Account) ConfirmEmailChange(ctx context.Context, token string) error {
	if *a.acc.EmailChangeToken != token {
		return fmt.Errorf("%w", ErrInvalidToken)
	}
	err := a.acc.Update().
		SetEmail(a.acc.EmailChange).
		ClearEmailChangeToken().
		ClearEmailChangeSentAt().
		ClearEmailChange().
		Exec(ctx)

	if err != nil {
		return fmt.Errorf("%v %w", err, ErrUpdatingUser)
	}

	return nil
}

func (a *Account) Delete(ctx context.Context) error {
	err := a.client.Account.DeleteOne(a.acc).Exec(ctx)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrInternal)
	}
	return nil
}

func (a *Account) Email() string {
	return a.acc.Email
}

func (a *Account) ID() uuid.UUID {
	return a.acc.ID
}
