// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"time"

	"github.com/adnaan/authn/models/account"
	"github.com/adnaan/authn/models/session"
	"github.com/adnaan/authn/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescProvider is the schema descriptor for provider field.
	accountDescProvider := accountFields[1].Descriptor()
	// account.ProviderValidator is a validator for the "provider" field. It is called by the builders before save.
	account.ProviderValidator = accountDescProvider.Validators[0].(func(string) error)
	// accountDescEmail is the schema descriptor for email field.
	accountDescEmail := accountFields[2].Descriptor()
	// account.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	account.EmailValidator = accountDescEmail.Validators[0].(func(string) error)
	// accountDescPassword is the schema descriptor for password field.
	accountDescPassword := accountFields[3].Descriptor()
	// account.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	account.PasswordValidator = func() func(string) error {
		validators := accountDescPassword.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(password string) error {
			for _, fn := range fns {
				if err := fn(password); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// accountDescLocked is the schema descriptor for locked field.
	accountDescLocked := accountFields[4].Descriptor()
	// account.DefaultLocked holds the default value on creation for the locked field.
	account.DefaultLocked = accountDescLocked.Default.(bool)
	// accountDescConfirmed is the schema descriptor for confirmed field.
	accountDescConfirmed := accountFields[5].Descriptor()
	// account.DefaultConfirmed holds the default value on creation for the confirmed field.
	account.DefaultConfirmed = accountDescConfirmed.Default.(bool)
	// accountDescConfirmationToken is the schema descriptor for confirmation_token field.
	accountDescConfirmationToken := accountFields[7].Descriptor()
	// account.ConfirmationTokenValidator is a validator for the "confirmation_token" field. It is called by the builders before save.
	account.ConfirmationTokenValidator = accountDescConfirmationToken.Validators[0].(func(string) error)
	// accountDescRecoveryToken is the schema descriptor for recovery_token field.
	accountDescRecoveryToken := accountFields[9].Descriptor()
	// account.RecoveryTokenValidator is a validator for the "recovery_token" field. It is called by the builders before save.
	account.RecoveryTokenValidator = accountDescRecoveryToken.Validators[0].(func(string) error)
	// accountDescOtp is the schema descriptor for otp field.
	accountDescOtp := accountFields[11].Descriptor()
	// account.OtpValidator is a validator for the "otp" field. It is called by the builders before save.
	account.OtpValidator = accountDescOtp.Validators[0].(func(string) error)
	// accountDescEmailChange is the schema descriptor for email_change field.
	accountDescEmailChange := accountFields[12].Descriptor()
	// account.EmailChangeValidator is a validator for the "email_change" field. It is called by the builders before save.
	account.EmailChangeValidator = accountDescEmailChange.Validators[0].(func(string) error)
	// accountDescEmailChangeToken is the schema descriptor for email_change_token field.
	accountDescEmailChangeToken := accountFields[14].Descriptor()
	// account.EmailChangeTokenValidator is a validator for the "email_change_token" field. It is called by the builders before save.
	account.EmailChangeTokenValidator = accountDescEmailChangeToken.Validators[0].(func(string) error)
	// accountDescCreatedAt is the schema descriptor for created_at field.
	accountDescCreatedAt := accountFields[18].Descriptor()
	// account.DefaultCreatedAt holds the default value on creation for the created_at field.
	account.DefaultCreatedAt = accountDescCreatedAt.Default.(func() time.Time)
	// accountDescUpdatedAt is the schema descriptor for updated_at field.
	accountDescUpdatedAt := accountFields[19].Descriptor()
	// account.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	account.DefaultUpdatedAt = accountDescUpdatedAt.Default.(func() time.Time)
	// account.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	account.UpdateDefaultUpdatedAt = accountDescUpdatedAt.UpdateDefault.(func() time.Time)
	// accountDescID is the schema descriptor for id field.
	accountDescID := accountFields[0].Descriptor()
	// account.DefaultID holds the default value on creation for the id field.
	account.DefaultID = accountDescID.Default.(func() uuid.UUID)
	sessionFields := schema.Session{}.Fields()
	_ = sessionFields
	// sessionDescCreatedAt is the schema descriptor for created_at field.
	sessionDescCreatedAt := sessionFields[2].Descriptor()
	// session.DefaultCreatedAt holds the default value on creation for the created_at field.
	session.DefaultCreatedAt = sessionDescCreatedAt.Default.(func() time.Time)
	// sessionDescUpdatedAt is the schema descriptor for updated_at field.
	sessionDescUpdatedAt := sessionFields[3].Descriptor()
	// session.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	session.DefaultUpdatedAt = sessionDescUpdatedAt.Default.(func() time.Time)
	// session.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	session.UpdateDefaultUpdatedAt = sessionDescUpdatedAt.UpdateDefault.(func() time.Time)
}
