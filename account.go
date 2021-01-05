package authzen

import (
	"fmt"
	"time"

	"github.com/adnaan/authzen/models"
	"github.com/adnaan/authzen/models/workspaceinvitation"
	"github.com/google/uuid"
)

type Account struct {
	*models.Account
}

func (a *API) newAccount(email, password, role, provider string, attributes map[string]interface{}, sendConfirm bool) (*models.Account, error) {
	var wrkspace *models.Workspace
	var acc *models.Account
	var err error

	err = withTx(a.ctx, a.c, func(tx *models.Tx) error {
		// create user
		createAccount := tx.Account.Create().
			SetEmail(email).
			SetPassword(password).
			SetRoles([]string{role}).
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
		acc, err = createAccount.Save(a.ctx)
		if err != nil {
			return err
		}

		var name string
		nameVal, ok := attributes["name"]
		if ok {
			nameStr, ok := nameVal.(string)
			if ok {
				name = nameStr
			}
		}

		// create default personal workspace
		wrkspace, err = tx.Workspace.Create().
			SetName(fmt.Sprintf("%s's Personal", name)).
			SetDescription(fmt.Sprintf("%s Personal Plan", email)).
			SetOwner(acc).
			SetPlan("default").
			Save(a.ctx)
		if err != nil {
			return err
		}
		// create role for personal workspace
		_, err = tx.WorkspaceRole.Create().
			SetName("owner").
			SetWorkspaceID(wrkspace.ID).
			SetWorkspaces(wrkspace).
			SetAccountID(acc.ID).
			SetAccounts(acc).
			Save(a.ctx)
		if err != nil {
			return err
		}

		invitations, err := tx.WorkspaceInvitation.Query().Where(workspaceinvitation.Email(email)).All(a.ctx)
		if err == nil {
			for _, invitation := range invitations {
				_, err = a.c.WorkspaceRole.Create().
					SetName(invitation.Role).
					SetAccountID(acc.ID).
					SetWorkspaceID(invitation.WorkspaceID).
					Save(a.ctx)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
	if err != nil {
		return acc, err
	}

	return acc, nil
}
