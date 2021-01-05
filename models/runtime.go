// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"time"

	"github.com/adnaan/authzen/models/account"
	"github.com/adnaan/authzen/models/accountrole"
	"github.com/adnaan/authzen/models/group"
	"github.com/adnaan/authzen/models/grouprole"
	"github.com/adnaan/authzen/models/permission"
	"github.com/adnaan/authzen/models/session"
	"github.com/adnaan/authzen/models/workspace"
	"github.com/adnaan/authzen/models/workspaceinvitation"
	"github.com/adnaan/authzen/models/workspacerole"
	"github.com/adnaan/authzen/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescBillingID is the schema descriptor for billing_id field.
	accountDescBillingID := accountFields[1].Descriptor()
	// account.BillingIDValidator is a validator for the "billing_id" field. It is called by the builders before save.
	account.BillingIDValidator = accountDescBillingID.Validators[0].(func(string) error)
	// accountDescProvider is the schema descriptor for provider field.
	accountDescProvider := accountFields[2].Descriptor()
	// account.ProviderValidator is a validator for the "provider" field. It is called by the builders before save.
	account.ProviderValidator = accountDescProvider.Validators[0].(func(string) error)
	// accountDescEmail is the schema descriptor for email field.
	accountDescEmail := accountFields[3].Descriptor()
	// account.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	account.EmailValidator = accountDescEmail.Validators[0].(func(string) error)
	// accountDescPassword is the schema descriptor for password field.
	accountDescPassword := accountFields[4].Descriptor()
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
	// accountDescAPIKey is the schema descriptor for api_key field.
	accountDescAPIKey := accountFields[5].Descriptor()
	// account.APIKeyValidator is a validator for the "api_key" field. It is called by the builders before save.
	account.APIKeyValidator = accountDescAPIKey.Validators[0].(func(string) error)
	// accountDescConfirmed is the schema descriptor for confirmed field.
	accountDescConfirmed := accountFields[6].Descriptor()
	// account.DefaultConfirmed holds the default value on creation for the confirmed field.
	account.DefaultConfirmed = accountDescConfirmed.Default.(bool)
	// accountDescConfirmationToken is the schema descriptor for confirmation_token field.
	accountDescConfirmationToken := accountFields[8].Descriptor()
	// account.ConfirmationTokenValidator is a validator for the "confirmation_token" field. It is called by the builders before save.
	account.ConfirmationTokenValidator = accountDescConfirmationToken.Validators[0].(func(string) error)
	// accountDescRecoveryToken is the schema descriptor for recovery_token field.
	accountDescRecoveryToken := accountFields[10].Descriptor()
	// account.RecoveryTokenValidator is a validator for the "recovery_token" field. It is called by the builders before save.
	account.RecoveryTokenValidator = accountDescRecoveryToken.Validators[0].(func(string) error)
	// accountDescOtp is the schema descriptor for otp field.
	accountDescOtp := accountFields[12].Descriptor()
	// account.OtpValidator is a validator for the "otp" field. It is called by the builders before save.
	account.OtpValidator = accountDescOtp.Validators[0].(func(string) error)
	// accountDescEmailChange is the schema descriptor for email_change field.
	accountDescEmailChange := accountFields[13].Descriptor()
	// account.EmailChangeValidator is a validator for the "email_change" field. It is called by the builders before save.
	account.EmailChangeValidator = accountDescEmailChange.Validators[0].(func(string) error)
	// accountDescEmailChangeToken is the schema descriptor for email_change_token field.
	accountDescEmailChangeToken := accountFields[15].Descriptor()
	// account.EmailChangeTokenValidator is a validator for the "email_change_token" field. It is called by the builders before save.
	account.EmailChangeTokenValidator = accountDescEmailChangeToken.Validators[0].(func(string) error)
	// accountDescCreatedAt is the schema descriptor for created_at field.
	accountDescCreatedAt := accountFields[19].Descriptor()
	// account.DefaultCreatedAt holds the default value on creation for the created_at field.
	account.DefaultCreatedAt = accountDescCreatedAt.Default.(func() time.Time)
	// accountDescUpdatedAt is the schema descriptor for updated_at field.
	accountDescUpdatedAt := accountFields[20].Descriptor()
	// account.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	account.DefaultUpdatedAt = accountDescUpdatedAt.Default.(func() time.Time)
	// account.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	account.UpdateDefaultUpdatedAt = accountDescUpdatedAt.UpdateDefault.(func() time.Time)
	// accountDescID is the schema descriptor for id field.
	accountDescID := accountFields[0].Descriptor()
	// account.DefaultID holds the default value on creation for the id field.
	account.DefaultID = accountDescID.Default.(func() uuid.UUID)
	accountroleFields := schema.AccountRole{}.Fields()
	_ = accountroleFields
	// accountroleDescName is the schema descriptor for name field.
	accountroleDescName := accountroleFields[1].Descriptor()
	// accountrole.NameValidator is a validator for the "name" field. It is called by the builders before save.
	accountrole.NameValidator = accountroleDescName.Validators[0].(func(string) error)
	// accountroleDescCreatedAt is the schema descriptor for created_at field.
	accountroleDescCreatedAt := accountroleFields[4].Descriptor()
	// accountrole.DefaultCreatedAt holds the default value on creation for the created_at field.
	accountrole.DefaultCreatedAt = accountroleDescCreatedAt.Default.(func() time.Time)
	// accountroleDescUpdatedAt is the schema descriptor for updated_at field.
	accountroleDescUpdatedAt := accountroleFields[5].Descriptor()
	// accountrole.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	accountrole.DefaultUpdatedAt = accountroleDescUpdatedAt.Default.(func() time.Time)
	// accountrole.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	accountrole.UpdateDefaultUpdatedAt = accountroleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// accountroleDescID is the schema descriptor for id field.
	accountroleDescID := accountroleFields[0].Descriptor()
	// accountrole.DefaultID holds the default value on creation for the id field.
	accountrole.DefaultID = accountroleDescID.Default.(func() uuid.UUID)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[1].Descriptor()
	// group.NameValidator is a validator for the "name" field. It is called by the builders before save.
	group.NameValidator = groupDescName.Validators[0].(func(string) error)
	// groupDescDescription is the schema descriptor for description field.
	groupDescDescription := groupFields[2].Descriptor()
	// group.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	group.DescriptionValidator = groupDescDescription.Validators[0].(func(string) error)
	// groupDescCreatedAt is the schema descriptor for created_at field.
	groupDescCreatedAt := groupFields[4].Descriptor()
	// group.DefaultCreatedAt holds the default value on creation for the created_at field.
	group.DefaultCreatedAt = groupDescCreatedAt.Default.(func() time.Time)
	// groupDescUpdatedAt is the schema descriptor for updated_at field.
	groupDescUpdatedAt := groupFields[5].Descriptor()
	// group.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	group.DefaultUpdatedAt = groupDescUpdatedAt.Default.(func() time.Time)
	// group.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	group.UpdateDefaultUpdatedAt = groupDescUpdatedAt.UpdateDefault.(func() time.Time)
	// groupDescID is the schema descriptor for id field.
	groupDescID := groupFields[0].Descriptor()
	// group.DefaultID holds the default value on creation for the id field.
	group.DefaultID = groupDescID.Default.(func() uuid.UUID)
	grouproleFields := schema.GroupRole{}.Fields()
	_ = grouproleFields
	// grouproleDescName is the schema descriptor for name field.
	grouproleDescName := grouproleFields[1].Descriptor()
	// grouprole.NameValidator is a validator for the "name" field. It is called by the builders before save.
	grouprole.NameValidator = grouproleDescName.Validators[0].(func(string) error)
	// grouproleDescCreatedAt is the schema descriptor for created_at field.
	grouproleDescCreatedAt := grouproleFields[5].Descriptor()
	// grouprole.DefaultCreatedAt holds the default value on creation for the created_at field.
	grouprole.DefaultCreatedAt = grouproleDescCreatedAt.Default.(func() time.Time)
	// grouproleDescUpdatedAt is the schema descriptor for updated_at field.
	grouproleDescUpdatedAt := grouproleFields[6].Descriptor()
	// grouprole.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	grouprole.DefaultUpdatedAt = grouproleDescUpdatedAt.Default.(func() time.Time)
	// grouprole.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	grouprole.UpdateDefaultUpdatedAt = grouproleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// grouproleDescID is the schema descriptor for id field.
	grouproleDescID := grouproleFields[0].Descriptor()
	// grouprole.DefaultID holds the default value on creation for the id field.
	grouprole.DefaultID = grouproleDescID.Default.(func() uuid.UUID)
	permissionFields := schema.Permission{}.Fields()
	_ = permissionFields
	// permissionDescAction is the schema descriptor for action field.
	permissionDescAction := permissionFields[2].Descriptor()
	// permission.ActionValidator is a validator for the "action" field. It is called by the builders before save.
	permission.ActionValidator = permissionDescAction.Validators[0].(func(string) error)
	// permissionDescTarget is the schema descriptor for target field.
	permissionDescTarget := permissionFields[3].Descriptor()
	// permission.TargetValidator is a validator for the "target" field. It is called by the builders before save.
	permission.TargetValidator = permissionDescTarget.Validators[0].(func(string) error)
	// permissionDescCreatedAt is the schema descriptor for created_at field.
	permissionDescCreatedAt := permissionFields[4].Descriptor()
	// permission.DefaultCreatedAt holds the default value on creation for the created_at field.
	permission.DefaultCreatedAt = permissionDescCreatedAt.Default.(func() time.Time)
	// permissionDescUpdatedAt is the schema descriptor for updated_at field.
	permissionDescUpdatedAt := permissionFields[5].Descriptor()
	// permission.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	permission.DefaultUpdatedAt = permissionDescUpdatedAt.Default.(func() time.Time)
	// permission.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	permission.UpdateDefaultUpdatedAt = permissionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// permissionDescID is the schema descriptor for id field.
	permissionDescID := permissionFields[0].Descriptor()
	// permission.DefaultID holds the default value on creation for the id field.
	permission.DefaultID = permissionDescID.Default.(func() uuid.UUID)
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
	workspaceFields := schema.Workspace{}.Fields()
	_ = workspaceFields
	// workspaceDescName is the schema descriptor for name field.
	workspaceDescName := workspaceFields[1].Descriptor()
	// workspace.NameValidator is a validator for the "name" field. It is called by the builders before save.
	workspace.NameValidator = workspaceDescName.Validators[0].(func(string) error)
	// workspaceDescPlan is the schema descriptor for plan field.
	workspaceDescPlan := workspaceFields[2].Descriptor()
	// workspace.PlanValidator is a validator for the "plan" field. It is called by the builders before save.
	workspace.PlanValidator = workspaceDescPlan.Validators[0].(func(string) error)
	// workspaceDescDescription is the schema descriptor for description field.
	workspaceDescDescription := workspaceFields[3].Descriptor()
	// workspace.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	workspace.DescriptionValidator = workspaceDescDescription.Validators[0].(func(string) error)
	// workspaceDescCreatedAt is the schema descriptor for created_at field.
	workspaceDescCreatedAt := workspaceFields[5].Descriptor()
	// workspace.DefaultCreatedAt holds the default value on creation for the created_at field.
	workspace.DefaultCreatedAt = workspaceDescCreatedAt.Default.(func() time.Time)
	// workspaceDescUpdatedAt is the schema descriptor for updated_at field.
	workspaceDescUpdatedAt := workspaceFields[6].Descriptor()
	// workspace.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	workspace.DefaultUpdatedAt = workspaceDescUpdatedAt.Default.(func() time.Time)
	// workspace.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	workspace.UpdateDefaultUpdatedAt = workspaceDescUpdatedAt.UpdateDefault.(func() time.Time)
	// workspaceDescID is the schema descriptor for id field.
	workspaceDescID := workspaceFields[0].Descriptor()
	// workspace.DefaultID holds the default value on creation for the id field.
	workspace.DefaultID = workspaceDescID.Default.(func() uuid.UUID)
	workspaceinvitationFields := schema.WorkspaceInvitation{}.Fields()
	_ = workspaceinvitationFields
	// workspaceinvitationDescRole is the schema descriptor for role field.
	workspaceinvitationDescRole := workspaceinvitationFields[2].Descriptor()
	// workspaceinvitation.RoleValidator is a validator for the "role" field. It is called by the builders before save.
	workspaceinvitation.RoleValidator = workspaceinvitationDescRole.Validators[0].(func(string) error)
	// workspaceinvitationDescEmail is the schema descriptor for email field.
	workspaceinvitationDescEmail := workspaceinvitationFields[3].Descriptor()
	// workspaceinvitation.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	workspaceinvitation.EmailValidator = workspaceinvitationDescEmail.Validators[0].(func(string) error)
	// workspaceinvitationDescCreatedAt is the schema descriptor for created_at field.
	workspaceinvitationDescCreatedAt := workspaceinvitationFields[4].Descriptor()
	// workspaceinvitation.DefaultCreatedAt holds the default value on creation for the created_at field.
	workspaceinvitation.DefaultCreatedAt = workspaceinvitationDescCreatedAt.Default.(func() time.Time)
	// workspaceinvitationDescID is the schema descriptor for id field.
	workspaceinvitationDescID := workspaceinvitationFields[0].Descriptor()
	// workspaceinvitation.DefaultID holds the default value on creation for the id field.
	workspaceinvitation.DefaultID = workspaceinvitationDescID.Default.(func() uuid.UUID)
	workspaceroleFields := schema.WorkspaceRole{}.Fields()
	_ = workspaceroleFields
	// workspaceroleDescName is the schema descriptor for name field.
	workspaceroleDescName := workspaceroleFields[1].Descriptor()
	// workspacerole.NameValidator is a validator for the "name" field. It is called by the builders before save.
	workspacerole.NameValidator = workspaceroleDescName.Validators[0].(func(string) error)
	// workspaceroleDescCreatedAt is the schema descriptor for created_at field.
	workspaceroleDescCreatedAt := workspaceroleFields[5].Descriptor()
	// workspacerole.DefaultCreatedAt holds the default value on creation for the created_at field.
	workspacerole.DefaultCreatedAt = workspaceroleDescCreatedAt.Default.(func() time.Time)
	// workspaceroleDescUpdatedAt is the schema descriptor for updated_at field.
	workspaceroleDescUpdatedAt := workspaceroleFields[6].Descriptor()
	// workspacerole.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	workspacerole.DefaultUpdatedAt = workspaceroleDescUpdatedAt.Default.(func() time.Time)
	// workspacerole.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	workspacerole.UpdateDefaultUpdatedAt = workspaceroleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// workspaceroleDescID is the schema descriptor for id field.
	workspaceroleDescID := workspaceroleFields[0].Descriptor()
	// workspacerole.DefaultID holds the default value on creation for the id field.
	workspacerole.DefaultID = workspaceroleDescID.Default.(func() uuid.UUID)
}
