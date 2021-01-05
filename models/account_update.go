// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"time"

	"github.com/adnaan/authzen/models/account"
	"github.com/adnaan/authzen/models/accountrole"
	"github.com/adnaan/authzen/models/grouprole"
	"github.com/adnaan/authzen/models/predicate"
	"github.com/adnaan/authzen/models/workspace"
	"github.com/adnaan/authzen/models/workspacerole"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// Where adds a new predicate for the builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.mutation.predicates = append(au.mutation.predicates, ps...)
	return au
}

// SetBillingID sets the billing_id field.
func (au *AccountUpdate) SetBillingID(s string) *AccountUpdate {
	au.mutation.SetBillingID(s)
	return au
}

// SetNillableBillingID sets the billing_id field if the given value is not nil.
func (au *AccountUpdate) SetNillableBillingID(s *string) *AccountUpdate {
	if s != nil {
		au.SetBillingID(*s)
	}
	return au
}

// ClearBillingID clears the value of billing_id.
func (au *AccountUpdate) ClearBillingID() *AccountUpdate {
	au.mutation.ClearBillingID()
	return au
}

// SetProvider sets the provider field.
func (au *AccountUpdate) SetProvider(s string) *AccountUpdate {
	au.mutation.SetProvider(s)
	return au
}

// SetEmail sets the email field.
func (au *AccountUpdate) SetEmail(s string) *AccountUpdate {
	au.mutation.SetEmail(s)
	return au
}

// SetPassword sets the password field.
func (au *AccountUpdate) SetPassword(s string) *AccountUpdate {
	au.mutation.SetPassword(s)
	return au
}

// SetAPIKey sets the api_key field.
func (au *AccountUpdate) SetAPIKey(s string) *AccountUpdate {
	au.mutation.SetAPIKey(s)
	return au
}

// SetNillableAPIKey sets the api_key field if the given value is not nil.
func (au *AccountUpdate) SetNillableAPIKey(s *string) *AccountUpdate {
	if s != nil {
		au.SetAPIKey(*s)
	}
	return au
}

// ClearAPIKey clears the value of api_key.
func (au *AccountUpdate) ClearAPIKey() *AccountUpdate {
	au.mutation.ClearAPIKey()
	return au
}

// SetConfirmed sets the confirmed field.
func (au *AccountUpdate) SetConfirmed(b bool) *AccountUpdate {
	au.mutation.SetConfirmed(b)
	return au
}

// SetNillableConfirmed sets the confirmed field if the given value is not nil.
func (au *AccountUpdate) SetNillableConfirmed(b *bool) *AccountUpdate {
	if b != nil {
		au.SetConfirmed(*b)
	}
	return au
}

// ClearConfirmed clears the value of confirmed.
func (au *AccountUpdate) ClearConfirmed() *AccountUpdate {
	au.mutation.ClearConfirmed()
	return au
}

// SetConfirmationSentAt sets the confirmation_sent_at field.
func (au *AccountUpdate) SetConfirmationSentAt(t time.Time) *AccountUpdate {
	au.mutation.SetConfirmationSentAt(t)
	return au
}

// SetNillableConfirmationSentAt sets the confirmation_sent_at field if the given value is not nil.
func (au *AccountUpdate) SetNillableConfirmationSentAt(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetConfirmationSentAt(*t)
	}
	return au
}

// ClearConfirmationSentAt clears the value of confirmation_sent_at.
func (au *AccountUpdate) ClearConfirmationSentAt() *AccountUpdate {
	au.mutation.ClearConfirmationSentAt()
	return au
}

// SetConfirmationToken sets the confirmation_token field.
func (au *AccountUpdate) SetConfirmationToken(s string) *AccountUpdate {
	au.mutation.SetConfirmationToken(s)
	return au
}

// SetNillableConfirmationToken sets the confirmation_token field if the given value is not nil.
func (au *AccountUpdate) SetNillableConfirmationToken(s *string) *AccountUpdate {
	if s != nil {
		au.SetConfirmationToken(*s)
	}
	return au
}

// ClearConfirmationToken clears the value of confirmation_token.
func (au *AccountUpdate) ClearConfirmationToken() *AccountUpdate {
	au.mutation.ClearConfirmationToken()
	return au
}

// SetRecoverySentAt sets the recovery_sent_at field.
func (au *AccountUpdate) SetRecoverySentAt(t time.Time) *AccountUpdate {
	au.mutation.SetRecoverySentAt(t)
	return au
}

// SetNillableRecoverySentAt sets the recovery_sent_at field if the given value is not nil.
func (au *AccountUpdate) SetNillableRecoverySentAt(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetRecoverySentAt(*t)
	}
	return au
}

// ClearRecoverySentAt clears the value of recovery_sent_at.
func (au *AccountUpdate) ClearRecoverySentAt() *AccountUpdate {
	au.mutation.ClearRecoverySentAt()
	return au
}

// SetRecoveryToken sets the recovery_token field.
func (au *AccountUpdate) SetRecoveryToken(s string) *AccountUpdate {
	au.mutation.SetRecoveryToken(s)
	return au
}

// SetNillableRecoveryToken sets the recovery_token field if the given value is not nil.
func (au *AccountUpdate) SetNillableRecoveryToken(s *string) *AccountUpdate {
	if s != nil {
		au.SetRecoveryToken(*s)
	}
	return au
}

// ClearRecoveryToken clears the value of recovery_token.
func (au *AccountUpdate) ClearRecoveryToken() *AccountUpdate {
	au.mutation.ClearRecoveryToken()
	return au
}

// SetOtpSentAt sets the otp_sent_at field.
func (au *AccountUpdate) SetOtpSentAt(t time.Time) *AccountUpdate {
	au.mutation.SetOtpSentAt(t)
	return au
}

// SetNillableOtpSentAt sets the otp_sent_at field if the given value is not nil.
func (au *AccountUpdate) SetNillableOtpSentAt(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetOtpSentAt(*t)
	}
	return au
}

// ClearOtpSentAt clears the value of otp_sent_at.
func (au *AccountUpdate) ClearOtpSentAt() *AccountUpdate {
	au.mutation.ClearOtpSentAt()
	return au
}

// SetOtp sets the otp field.
func (au *AccountUpdate) SetOtp(s string) *AccountUpdate {
	au.mutation.SetOtp(s)
	return au
}

// SetNillableOtp sets the otp field if the given value is not nil.
func (au *AccountUpdate) SetNillableOtp(s *string) *AccountUpdate {
	if s != nil {
		au.SetOtp(*s)
	}
	return au
}

// ClearOtp clears the value of otp.
func (au *AccountUpdate) ClearOtp() *AccountUpdate {
	au.mutation.ClearOtp()
	return au
}

// SetEmailChange sets the email_change field.
func (au *AccountUpdate) SetEmailChange(s string) *AccountUpdate {
	au.mutation.SetEmailChange(s)
	return au
}

// SetNillableEmailChange sets the email_change field if the given value is not nil.
func (au *AccountUpdate) SetNillableEmailChange(s *string) *AccountUpdate {
	if s != nil {
		au.SetEmailChange(*s)
	}
	return au
}

// ClearEmailChange clears the value of email_change.
func (au *AccountUpdate) ClearEmailChange() *AccountUpdate {
	au.mutation.ClearEmailChange()
	return au
}

// SetEmailChangeSentAt sets the email_change_sent_at field.
func (au *AccountUpdate) SetEmailChangeSentAt(t time.Time) *AccountUpdate {
	au.mutation.SetEmailChangeSentAt(t)
	return au
}

// SetNillableEmailChangeSentAt sets the email_change_sent_at field if the given value is not nil.
func (au *AccountUpdate) SetNillableEmailChangeSentAt(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetEmailChangeSentAt(*t)
	}
	return au
}

// ClearEmailChangeSentAt clears the value of email_change_sent_at.
func (au *AccountUpdate) ClearEmailChangeSentAt() *AccountUpdate {
	au.mutation.ClearEmailChangeSentAt()
	return au
}

// SetEmailChangeToken sets the email_change_token field.
func (au *AccountUpdate) SetEmailChangeToken(s string) *AccountUpdate {
	au.mutation.SetEmailChangeToken(s)
	return au
}

// SetNillableEmailChangeToken sets the email_change_token field if the given value is not nil.
func (au *AccountUpdate) SetNillableEmailChangeToken(s *string) *AccountUpdate {
	if s != nil {
		au.SetEmailChangeToken(*s)
	}
	return au
}

// ClearEmailChangeToken clears the value of email_change_token.
func (au *AccountUpdate) ClearEmailChangeToken() *AccountUpdate {
	au.mutation.ClearEmailChangeToken()
	return au
}

// SetMetadata sets the metadata field.
func (au *AccountUpdate) SetMetadata(m map[string]interface{}) *AccountUpdate {
	au.mutation.SetMetadata(m)
	return au
}

// SetRoles sets the roles field.
func (au *AccountUpdate) SetRoles(s []string) *AccountUpdate {
	au.mutation.SetRoles(s)
	return au
}

// ClearRoles clears the value of roles.
func (au *AccountUpdate) ClearRoles() *AccountUpdate {
	au.mutation.ClearRoles()
	return au
}

// SetTeams sets the teams field.
func (au *AccountUpdate) SetTeams(m map[string]string) *AccountUpdate {
	au.mutation.SetTeams(m)
	return au
}

// ClearTeams clears the value of teams.
func (au *AccountUpdate) ClearTeams() *AccountUpdate {
	au.mutation.ClearTeams()
	return au
}

// SetUpdatedAt sets the updated_at field.
func (au *AccountUpdate) SetUpdatedAt(t time.Time) *AccountUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetLastSigninAt sets the last_signin_at field.
func (au *AccountUpdate) SetLastSigninAt(t time.Time) *AccountUpdate {
	au.mutation.SetLastSigninAt(t)
	return au
}

// SetNillableLastSigninAt sets the last_signin_at field if the given value is not nil.
func (au *AccountUpdate) SetNillableLastSigninAt(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetLastSigninAt(*t)
	}
	return au
}

// ClearLastSigninAt clears the value of last_signin_at.
func (au *AccountUpdate) ClearLastSigninAt() *AccountUpdate {
	au.mutation.ClearLastSigninAt()
	return au
}

// SetWorkspaceID sets the workspace edge to Workspace by id.
func (au *AccountUpdate) SetWorkspaceID(id uuid.UUID) *AccountUpdate {
	au.mutation.SetWorkspaceID(id)
	return au
}

// SetNillableWorkspaceID sets the workspace edge to Workspace by id if the given value is not nil.
func (au *AccountUpdate) SetNillableWorkspaceID(id *uuid.UUID) *AccountUpdate {
	if id != nil {
		au = au.SetWorkspaceID(*id)
	}
	return au
}

// SetWorkspace sets the workspace edge to Workspace.
func (au *AccountUpdate) SetWorkspace(w *Workspace) *AccountUpdate {
	return au.SetWorkspaceID(w.ID)
}

// AddWorkspaceRoleIDs adds the workspace_roles edge to WorkspaceRole by ids.
func (au *AccountUpdate) AddWorkspaceRoleIDs(ids ...uuid.UUID) *AccountUpdate {
	au.mutation.AddWorkspaceRoleIDs(ids...)
	return au
}

// AddWorkspaceRoles adds the workspace_roles edges to WorkspaceRole.
func (au *AccountUpdate) AddWorkspaceRoles(w ...*WorkspaceRole) *AccountUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return au.AddWorkspaceRoleIDs(ids...)
}

// AddGroupRoleIDs adds the group_roles edge to GroupRole by ids.
func (au *AccountUpdate) AddGroupRoleIDs(ids ...uuid.UUID) *AccountUpdate {
	au.mutation.AddGroupRoleIDs(ids...)
	return au
}

// AddGroupRoles adds the group_roles edges to GroupRole.
func (au *AccountUpdate) AddGroupRoles(g ...*GroupRole) *AccountUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return au.AddGroupRoleIDs(ids...)
}

// AddAccountRoleIDs adds the account_roles edge to AccountRole by ids.
func (au *AccountUpdate) AddAccountRoleIDs(ids ...uuid.UUID) *AccountUpdate {
	au.mutation.AddAccountRoleIDs(ids...)
	return au
}

// AddAccountRoles adds the account_roles edges to AccountRole.
func (au *AccountUpdate) AddAccountRoles(a ...*AccountRole) *AccountUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddAccountRoleIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// ClearWorkspace clears the "workspace" edge to type Workspace.
func (au *AccountUpdate) ClearWorkspace() *AccountUpdate {
	au.mutation.ClearWorkspace()
	return au
}

// ClearWorkspaceRoles clears all "workspace_roles" edges to type WorkspaceRole.
func (au *AccountUpdate) ClearWorkspaceRoles() *AccountUpdate {
	au.mutation.ClearWorkspaceRoles()
	return au
}

// RemoveWorkspaceRoleIDs removes the workspace_roles edge to WorkspaceRole by ids.
func (au *AccountUpdate) RemoveWorkspaceRoleIDs(ids ...uuid.UUID) *AccountUpdate {
	au.mutation.RemoveWorkspaceRoleIDs(ids...)
	return au
}

// RemoveWorkspaceRoles removes workspace_roles edges to WorkspaceRole.
func (au *AccountUpdate) RemoveWorkspaceRoles(w ...*WorkspaceRole) *AccountUpdate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return au.RemoveWorkspaceRoleIDs(ids...)
}

// ClearGroupRoles clears all "group_roles" edges to type GroupRole.
func (au *AccountUpdate) ClearGroupRoles() *AccountUpdate {
	au.mutation.ClearGroupRoles()
	return au
}

// RemoveGroupRoleIDs removes the group_roles edge to GroupRole by ids.
func (au *AccountUpdate) RemoveGroupRoleIDs(ids ...uuid.UUID) *AccountUpdate {
	au.mutation.RemoveGroupRoleIDs(ids...)
	return au
}

// RemoveGroupRoles removes group_roles edges to GroupRole.
func (au *AccountUpdate) RemoveGroupRoles(g ...*GroupRole) *AccountUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return au.RemoveGroupRoleIDs(ids...)
}

// ClearAccountRoles clears all "account_roles" edges to type AccountRole.
func (au *AccountUpdate) ClearAccountRoles() *AccountUpdate {
	au.mutation.ClearAccountRoles()
	return au
}

// RemoveAccountRoleIDs removes the account_roles edge to AccountRole by ids.
func (au *AccountUpdate) RemoveAccountRoleIDs(ids ...uuid.UUID) *AccountUpdate {
	au.mutation.RemoveAccountRoleIDs(ids...)
	return au
}

// RemoveAccountRoles removes account_roles edges to AccountRole.
func (au *AccountUpdate) RemoveAccountRoles(a ...*AccountRole) *AccountUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemoveAccountRoleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	au.defaults()
	if len(au.hooks) == 0 {
		if err = au.check(); err != nil {
			return 0, err
		}
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = au.check(); err != nil {
				return 0, err
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AccountUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := account.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AccountUpdate) check() error {
	if v, ok := au.mutation.BillingID(); ok {
		if err := account.BillingIDValidator(v); err != nil {
			return &ValidationError{Name: "billing_id", err: fmt.Errorf("models: validator failed for field \"billing_id\": %w", err)}
		}
	}
	if v, ok := au.mutation.Provider(); ok {
		if err := account.ProviderValidator(v); err != nil {
			return &ValidationError{Name: "provider", err: fmt.Errorf("models: validator failed for field \"provider\": %w", err)}
		}
	}
	if v, ok := au.mutation.Email(); ok {
		if err := account.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf("models: validator failed for field \"email\": %w", err)}
		}
	}
	if v, ok := au.mutation.Password(); ok {
		if err := account.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("models: validator failed for field \"password\": %w", err)}
		}
	}
	if v, ok := au.mutation.APIKey(); ok {
		if err := account.APIKeyValidator(v); err != nil {
			return &ValidationError{Name: "api_key", err: fmt.Errorf("models: validator failed for field \"api_key\": %w", err)}
		}
	}
	if v, ok := au.mutation.ConfirmationToken(); ok {
		if err := account.ConfirmationTokenValidator(v); err != nil {
			return &ValidationError{Name: "confirmation_token", err: fmt.Errorf("models: validator failed for field \"confirmation_token\": %w", err)}
		}
	}
	if v, ok := au.mutation.RecoveryToken(); ok {
		if err := account.RecoveryTokenValidator(v); err != nil {
			return &ValidationError{Name: "recovery_token", err: fmt.Errorf("models: validator failed for field \"recovery_token\": %w", err)}
		}
	}
	if v, ok := au.mutation.Otp(); ok {
		if err := account.OtpValidator(v); err != nil {
			return &ValidationError{Name: "otp", err: fmt.Errorf("models: validator failed for field \"otp\": %w", err)}
		}
	}
	if v, ok := au.mutation.EmailChange(); ok {
		if err := account.EmailChangeValidator(v); err != nil {
			return &ValidationError{Name: "email_change", err: fmt.Errorf("models: validator failed for field \"email_change\": %w", err)}
		}
	}
	if v, ok := au.mutation.EmailChangeToken(); ok {
		if err := account.EmailChangeTokenValidator(v); err != nil {
			return &ValidationError{Name: "email_change_token", err: fmt.Errorf("models: validator failed for field \"email_change_token\": %w", err)}
		}
	}
	return nil
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: account.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.BillingID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldBillingID,
		})
	}
	if au.mutation.BillingIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldBillingID,
		})
	}
	if value, ok := au.mutation.Provider(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldProvider,
		})
	}
	if value, ok := au.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldEmail,
		})
	}
	if value, ok := au.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldPassword,
		})
	}
	if value, ok := au.mutation.APIKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldAPIKey,
		})
	}
	if au.mutation.APIKeyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldAPIKey,
		})
	}
	if value, ok := au.mutation.Confirmed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: account.FieldConfirmed,
		})
	}
	if au.mutation.ConfirmedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: account.FieldConfirmed,
		})
	}
	if value, ok := au.mutation.ConfirmationSentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldConfirmationSentAt,
		})
	}
	if au.mutation.ConfirmationSentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: account.FieldConfirmationSentAt,
		})
	}
	if value, ok := au.mutation.ConfirmationToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldConfirmationToken,
		})
	}
	if au.mutation.ConfirmationTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldConfirmationToken,
		})
	}
	if value, ok := au.mutation.RecoverySentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldRecoverySentAt,
		})
	}
	if au.mutation.RecoverySentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: account.FieldRecoverySentAt,
		})
	}
	if value, ok := au.mutation.RecoveryToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldRecoveryToken,
		})
	}
	if au.mutation.RecoveryTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldRecoveryToken,
		})
	}
	if value, ok := au.mutation.OtpSentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldOtpSentAt,
		})
	}
	if au.mutation.OtpSentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: account.FieldOtpSentAt,
		})
	}
	if value, ok := au.mutation.Otp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldOtp,
		})
	}
	if au.mutation.OtpCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldOtp,
		})
	}
	if value, ok := au.mutation.EmailChange(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldEmailChange,
		})
	}
	if au.mutation.EmailChangeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldEmailChange,
		})
	}
	if value, ok := au.mutation.EmailChangeSentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldEmailChangeSentAt,
		})
	}
	if au.mutation.EmailChangeSentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: account.FieldEmailChangeSentAt,
		})
	}
	if value, ok := au.mutation.EmailChangeToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldEmailChangeToken,
		})
	}
	if au.mutation.EmailChangeTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldEmailChangeToken,
		})
	}
	if value, ok := au.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: account.FieldMetadata,
		})
	}
	if value, ok := au.mutation.Roles(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: account.FieldRoles,
		})
	}
	if au.mutation.RolesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: account.FieldRoles,
		})
	}
	if value, ok := au.mutation.Teams(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: account.FieldTeams,
		})
	}
	if au.mutation.TeamsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: account.FieldTeams,
		})
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldUpdatedAt,
		})
	}
	if value, ok := au.mutation.LastSigninAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldLastSigninAt,
		})
	}
	if au.mutation.LastSigninAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: account.FieldLastSigninAt,
		})
	}
	if au.mutation.WorkspaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   account.WorkspaceTable,
			Columns: []string{account.WorkspaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workspace.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   account.WorkspaceTable,
			Columns: []string{account.WorkspaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workspace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.WorkspaceRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.WorkspaceRolesTable,
			Columns: []string{account.WorkspaceRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workspacerole.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedWorkspaceRolesIDs(); len(nodes) > 0 && !au.mutation.WorkspaceRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.WorkspaceRolesTable,
			Columns: []string{account.WorkspaceRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workspacerole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.WorkspaceRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.WorkspaceRolesTable,
			Columns: []string{account.WorkspaceRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workspacerole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.GroupRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.GroupRolesTable,
			Columns: []string{account.GroupRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: grouprole.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedGroupRolesIDs(); len(nodes) > 0 && !au.mutation.GroupRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.GroupRolesTable,
			Columns: []string{account.GroupRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: grouprole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.GroupRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.GroupRolesTable,
			Columns: []string{account.GroupRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: grouprole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.AccountRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.AccountRolesTable,
			Columns: []string{account.AccountRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: accountrole.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedAccountRolesIDs(); len(nodes) > 0 && !au.mutation.AccountRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.AccountRolesTable,
			Columns: []string{account.AccountRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: accountrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AccountRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.AccountRolesTable,
			Columns: []string{account.AccountRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: accountrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// SetBillingID sets the billing_id field.
func (auo *AccountUpdateOne) SetBillingID(s string) *AccountUpdateOne {
	auo.mutation.SetBillingID(s)
	return auo
}

// SetNillableBillingID sets the billing_id field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableBillingID(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetBillingID(*s)
	}
	return auo
}

// ClearBillingID clears the value of billing_id.
func (auo *AccountUpdateOne) ClearBillingID() *AccountUpdateOne {
	auo.mutation.ClearBillingID()
	return auo
}

// SetProvider sets the provider field.
func (auo *AccountUpdateOne) SetProvider(s string) *AccountUpdateOne {
	auo.mutation.SetProvider(s)
	return auo
}

// SetEmail sets the email field.
func (auo *AccountUpdateOne) SetEmail(s string) *AccountUpdateOne {
	auo.mutation.SetEmail(s)
	return auo
}

// SetPassword sets the password field.
func (auo *AccountUpdateOne) SetPassword(s string) *AccountUpdateOne {
	auo.mutation.SetPassword(s)
	return auo
}

// SetAPIKey sets the api_key field.
func (auo *AccountUpdateOne) SetAPIKey(s string) *AccountUpdateOne {
	auo.mutation.SetAPIKey(s)
	return auo
}

// SetNillableAPIKey sets the api_key field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableAPIKey(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetAPIKey(*s)
	}
	return auo
}

// ClearAPIKey clears the value of api_key.
func (auo *AccountUpdateOne) ClearAPIKey() *AccountUpdateOne {
	auo.mutation.ClearAPIKey()
	return auo
}

// SetConfirmed sets the confirmed field.
func (auo *AccountUpdateOne) SetConfirmed(b bool) *AccountUpdateOne {
	auo.mutation.SetConfirmed(b)
	return auo
}

// SetNillableConfirmed sets the confirmed field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableConfirmed(b *bool) *AccountUpdateOne {
	if b != nil {
		auo.SetConfirmed(*b)
	}
	return auo
}

// ClearConfirmed clears the value of confirmed.
func (auo *AccountUpdateOne) ClearConfirmed() *AccountUpdateOne {
	auo.mutation.ClearConfirmed()
	return auo
}

// SetConfirmationSentAt sets the confirmation_sent_at field.
func (auo *AccountUpdateOne) SetConfirmationSentAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetConfirmationSentAt(t)
	return auo
}

// SetNillableConfirmationSentAt sets the confirmation_sent_at field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableConfirmationSentAt(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetConfirmationSentAt(*t)
	}
	return auo
}

// ClearConfirmationSentAt clears the value of confirmation_sent_at.
func (auo *AccountUpdateOne) ClearConfirmationSentAt() *AccountUpdateOne {
	auo.mutation.ClearConfirmationSentAt()
	return auo
}

// SetConfirmationToken sets the confirmation_token field.
func (auo *AccountUpdateOne) SetConfirmationToken(s string) *AccountUpdateOne {
	auo.mutation.SetConfirmationToken(s)
	return auo
}

// SetNillableConfirmationToken sets the confirmation_token field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableConfirmationToken(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetConfirmationToken(*s)
	}
	return auo
}

// ClearConfirmationToken clears the value of confirmation_token.
func (auo *AccountUpdateOne) ClearConfirmationToken() *AccountUpdateOne {
	auo.mutation.ClearConfirmationToken()
	return auo
}

// SetRecoverySentAt sets the recovery_sent_at field.
func (auo *AccountUpdateOne) SetRecoverySentAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetRecoverySentAt(t)
	return auo
}

// SetNillableRecoverySentAt sets the recovery_sent_at field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableRecoverySentAt(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetRecoverySentAt(*t)
	}
	return auo
}

// ClearRecoverySentAt clears the value of recovery_sent_at.
func (auo *AccountUpdateOne) ClearRecoverySentAt() *AccountUpdateOne {
	auo.mutation.ClearRecoverySentAt()
	return auo
}

// SetRecoveryToken sets the recovery_token field.
func (auo *AccountUpdateOne) SetRecoveryToken(s string) *AccountUpdateOne {
	auo.mutation.SetRecoveryToken(s)
	return auo
}

// SetNillableRecoveryToken sets the recovery_token field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableRecoveryToken(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetRecoveryToken(*s)
	}
	return auo
}

// ClearRecoveryToken clears the value of recovery_token.
func (auo *AccountUpdateOne) ClearRecoveryToken() *AccountUpdateOne {
	auo.mutation.ClearRecoveryToken()
	return auo
}

// SetOtpSentAt sets the otp_sent_at field.
func (auo *AccountUpdateOne) SetOtpSentAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetOtpSentAt(t)
	return auo
}

// SetNillableOtpSentAt sets the otp_sent_at field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableOtpSentAt(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetOtpSentAt(*t)
	}
	return auo
}

// ClearOtpSentAt clears the value of otp_sent_at.
func (auo *AccountUpdateOne) ClearOtpSentAt() *AccountUpdateOne {
	auo.mutation.ClearOtpSentAt()
	return auo
}

// SetOtp sets the otp field.
func (auo *AccountUpdateOne) SetOtp(s string) *AccountUpdateOne {
	auo.mutation.SetOtp(s)
	return auo
}

// SetNillableOtp sets the otp field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableOtp(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetOtp(*s)
	}
	return auo
}

// ClearOtp clears the value of otp.
func (auo *AccountUpdateOne) ClearOtp() *AccountUpdateOne {
	auo.mutation.ClearOtp()
	return auo
}

// SetEmailChange sets the email_change field.
func (auo *AccountUpdateOne) SetEmailChange(s string) *AccountUpdateOne {
	auo.mutation.SetEmailChange(s)
	return auo
}

// SetNillableEmailChange sets the email_change field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableEmailChange(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetEmailChange(*s)
	}
	return auo
}

// ClearEmailChange clears the value of email_change.
func (auo *AccountUpdateOne) ClearEmailChange() *AccountUpdateOne {
	auo.mutation.ClearEmailChange()
	return auo
}

// SetEmailChangeSentAt sets the email_change_sent_at field.
func (auo *AccountUpdateOne) SetEmailChangeSentAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetEmailChangeSentAt(t)
	return auo
}

// SetNillableEmailChangeSentAt sets the email_change_sent_at field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableEmailChangeSentAt(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetEmailChangeSentAt(*t)
	}
	return auo
}

// ClearEmailChangeSentAt clears the value of email_change_sent_at.
func (auo *AccountUpdateOne) ClearEmailChangeSentAt() *AccountUpdateOne {
	auo.mutation.ClearEmailChangeSentAt()
	return auo
}

// SetEmailChangeToken sets the email_change_token field.
func (auo *AccountUpdateOne) SetEmailChangeToken(s string) *AccountUpdateOne {
	auo.mutation.SetEmailChangeToken(s)
	return auo
}

// SetNillableEmailChangeToken sets the email_change_token field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableEmailChangeToken(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetEmailChangeToken(*s)
	}
	return auo
}

// ClearEmailChangeToken clears the value of email_change_token.
func (auo *AccountUpdateOne) ClearEmailChangeToken() *AccountUpdateOne {
	auo.mutation.ClearEmailChangeToken()
	return auo
}

// SetMetadata sets the metadata field.
func (auo *AccountUpdateOne) SetMetadata(m map[string]interface{}) *AccountUpdateOne {
	auo.mutation.SetMetadata(m)
	return auo
}

// SetRoles sets the roles field.
func (auo *AccountUpdateOne) SetRoles(s []string) *AccountUpdateOne {
	auo.mutation.SetRoles(s)
	return auo
}

// ClearRoles clears the value of roles.
func (auo *AccountUpdateOne) ClearRoles() *AccountUpdateOne {
	auo.mutation.ClearRoles()
	return auo
}

// SetTeams sets the teams field.
func (auo *AccountUpdateOne) SetTeams(m map[string]string) *AccountUpdateOne {
	auo.mutation.SetTeams(m)
	return auo
}

// ClearTeams clears the value of teams.
func (auo *AccountUpdateOne) ClearTeams() *AccountUpdateOne {
	auo.mutation.ClearTeams()
	return auo
}

// SetUpdatedAt sets the updated_at field.
func (auo *AccountUpdateOne) SetUpdatedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetLastSigninAt sets the last_signin_at field.
func (auo *AccountUpdateOne) SetLastSigninAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetLastSigninAt(t)
	return auo
}

// SetNillableLastSigninAt sets the last_signin_at field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableLastSigninAt(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetLastSigninAt(*t)
	}
	return auo
}

// ClearLastSigninAt clears the value of last_signin_at.
func (auo *AccountUpdateOne) ClearLastSigninAt() *AccountUpdateOne {
	auo.mutation.ClearLastSigninAt()
	return auo
}

// SetWorkspaceID sets the workspace edge to Workspace by id.
func (auo *AccountUpdateOne) SetWorkspaceID(id uuid.UUID) *AccountUpdateOne {
	auo.mutation.SetWorkspaceID(id)
	return auo
}

// SetNillableWorkspaceID sets the workspace edge to Workspace by id if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableWorkspaceID(id *uuid.UUID) *AccountUpdateOne {
	if id != nil {
		auo = auo.SetWorkspaceID(*id)
	}
	return auo
}

// SetWorkspace sets the workspace edge to Workspace.
func (auo *AccountUpdateOne) SetWorkspace(w *Workspace) *AccountUpdateOne {
	return auo.SetWorkspaceID(w.ID)
}

// AddWorkspaceRoleIDs adds the workspace_roles edge to WorkspaceRole by ids.
func (auo *AccountUpdateOne) AddWorkspaceRoleIDs(ids ...uuid.UUID) *AccountUpdateOne {
	auo.mutation.AddWorkspaceRoleIDs(ids...)
	return auo
}

// AddWorkspaceRoles adds the workspace_roles edges to WorkspaceRole.
func (auo *AccountUpdateOne) AddWorkspaceRoles(w ...*WorkspaceRole) *AccountUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return auo.AddWorkspaceRoleIDs(ids...)
}

// AddGroupRoleIDs adds the group_roles edge to GroupRole by ids.
func (auo *AccountUpdateOne) AddGroupRoleIDs(ids ...uuid.UUID) *AccountUpdateOne {
	auo.mutation.AddGroupRoleIDs(ids...)
	return auo
}

// AddGroupRoles adds the group_roles edges to GroupRole.
func (auo *AccountUpdateOne) AddGroupRoles(g ...*GroupRole) *AccountUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return auo.AddGroupRoleIDs(ids...)
}

// AddAccountRoleIDs adds the account_roles edge to AccountRole by ids.
func (auo *AccountUpdateOne) AddAccountRoleIDs(ids ...uuid.UUID) *AccountUpdateOne {
	auo.mutation.AddAccountRoleIDs(ids...)
	return auo
}

// AddAccountRoles adds the account_roles edges to AccountRole.
func (auo *AccountUpdateOne) AddAccountRoles(a ...*AccountRole) *AccountUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddAccountRoleIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// ClearWorkspace clears the "workspace" edge to type Workspace.
func (auo *AccountUpdateOne) ClearWorkspace() *AccountUpdateOne {
	auo.mutation.ClearWorkspace()
	return auo
}

// ClearWorkspaceRoles clears all "workspace_roles" edges to type WorkspaceRole.
func (auo *AccountUpdateOne) ClearWorkspaceRoles() *AccountUpdateOne {
	auo.mutation.ClearWorkspaceRoles()
	return auo
}

// RemoveWorkspaceRoleIDs removes the workspace_roles edge to WorkspaceRole by ids.
func (auo *AccountUpdateOne) RemoveWorkspaceRoleIDs(ids ...uuid.UUID) *AccountUpdateOne {
	auo.mutation.RemoveWorkspaceRoleIDs(ids...)
	return auo
}

// RemoveWorkspaceRoles removes workspace_roles edges to WorkspaceRole.
func (auo *AccountUpdateOne) RemoveWorkspaceRoles(w ...*WorkspaceRole) *AccountUpdateOne {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return auo.RemoveWorkspaceRoleIDs(ids...)
}

// ClearGroupRoles clears all "group_roles" edges to type GroupRole.
func (auo *AccountUpdateOne) ClearGroupRoles() *AccountUpdateOne {
	auo.mutation.ClearGroupRoles()
	return auo
}

// RemoveGroupRoleIDs removes the group_roles edge to GroupRole by ids.
func (auo *AccountUpdateOne) RemoveGroupRoleIDs(ids ...uuid.UUID) *AccountUpdateOne {
	auo.mutation.RemoveGroupRoleIDs(ids...)
	return auo
}

// RemoveGroupRoles removes group_roles edges to GroupRole.
func (auo *AccountUpdateOne) RemoveGroupRoles(g ...*GroupRole) *AccountUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return auo.RemoveGroupRoleIDs(ids...)
}

// ClearAccountRoles clears all "account_roles" edges to type AccountRole.
func (auo *AccountUpdateOne) ClearAccountRoles() *AccountUpdateOne {
	auo.mutation.ClearAccountRoles()
	return auo
}

// RemoveAccountRoleIDs removes the account_roles edge to AccountRole by ids.
func (auo *AccountUpdateOne) RemoveAccountRoleIDs(ids ...uuid.UUID) *AccountUpdateOne {
	auo.mutation.RemoveAccountRoleIDs(ids...)
	return auo
}

// RemoveAccountRoles removes account_roles edges to AccountRole.
func (auo *AccountUpdateOne) RemoveAccountRoles(a ...*AccountRole) *AccountUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemoveAccountRoleIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	var (
		err  error
		node *Account
	)
	auo.defaults()
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auo.check(); err != nil {
				return nil, err
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AccountUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := account.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AccountUpdateOne) check() error {
	if v, ok := auo.mutation.BillingID(); ok {
		if err := account.BillingIDValidator(v); err != nil {
			return &ValidationError{Name: "billing_id", err: fmt.Errorf("models: validator failed for field \"billing_id\": %w", err)}
		}
	}
	if v, ok := auo.mutation.Provider(); ok {
		if err := account.ProviderValidator(v); err != nil {
			return &ValidationError{Name: "provider", err: fmt.Errorf("models: validator failed for field \"provider\": %w", err)}
		}
	}
	if v, ok := auo.mutation.Email(); ok {
		if err := account.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf("models: validator failed for field \"email\": %w", err)}
		}
	}
	if v, ok := auo.mutation.Password(); ok {
		if err := account.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("models: validator failed for field \"password\": %w", err)}
		}
	}
	if v, ok := auo.mutation.APIKey(); ok {
		if err := account.APIKeyValidator(v); err != nil {
			return &ValidationError{Name: "api_key", err: fmt.Errorf("models: validator failed for field \"api_key\": %w", err)}
		}
	}
	if v, ok := auo.mutation.ConfirmationToken(); ok {
		if err := account.ConfirmationTokenValidator(v); err != nil {
			return &ValidationError{Name: "confirmation_token", err: fmt.Errorf("models: validator failed for field \"confirmation_token\": %w", err)}
		}
	}
	if v, ok := auo.mutation.RecoveryToken(); ok {
		if err := account.RecoveryTokenValidator(v); err != nil {
			return &ValidationError{Name: "recovery_token", err: fmt.Errorf("models: validator failed for field \"recovery_token\": %w", err)}
		}
	}
	if v, ok := auo.mutation.Otp(); ok {
		if err := account.OtpValidator(v); err != nil {
			return &ValidationError{Name: "otp", err: fmt.Errorf("models: validator failed for field \"otp\": %w", err)}
		}
	}
	if v, ok := auo.mutation.EmailChange(); ok {
		if err := account.EmailChangeValidator(v); err != nil {
			return &ValidationError{Name: "email_change", err: fmt.Errorf("models: validator failed for field \"email_change\": %w", err)}
		}
	}
	if v, ok := auo.mutation.EmailChangeToken(); ok {
		if err := account.EmailChangeTokenValidator(v); err != nil {
			return &ValidationError{Name: "email_change_token", err: fmt.Errorf("models: validator failed for field \"email_change_token\": %w", err)}
		}
	}
	return nil
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (_node *Account, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: account.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Account.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := auo.mutation.BillingID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldBillingID,
		})
	}
	if auo.mutation.BillingIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldBillingID,
		})
	}
	if value, ok := auo.mutation.Provider(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldProvider,
		})
	}
	if value, ok := auo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldEmail,
		})
	}
	if value, ok := auo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldPassword,
		})
	}
	if value, ok := auo.mutation.APIKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldAPIKey,
		})
	}
	if auo.mutation.APIKeyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldAPIKey,
		})
	}
	if value, ok := auo.mutation.Confirmed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: account.FieldConfirmed,
		})
	}
	if auo.mutation.ConfirmedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: account.FieldConfirmed,
		})
	}
	if value, ok := auo.mutation.ConfirmationSentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldConfirmationSentAt,
		})
	}
	if auo.mutation.ConfirmationSentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: account.FieldConfirmationSentAt,
		})
	}
	if value, ok := auo.mutation.ConfirmationToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldConfirmationToken,
		})
	}
	if auo.mutation.ConfirmationTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldConfirmationToken,
		})
	}
	if value, ok := auo.mutation.RecoverySentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldRecoverySentAt,
		})
	}
	if auo.mutation.RecoverySentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: account.FieldRecoverySentAt,
		})
	}
	if value, ok := auo.mutation.RecoveryToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldRecoveryToken,
		})
	}
	if auo.mutation.RecoveryTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldRecoveryToken,
		})
	}
	if value, ok := auo.mutation.OtpSentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldOtpSentAt,
		})
	}
	if auo.mutation.OtpSentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: account.FieldOtpSentAt,
		})
	}
	if value, ok := auo.mutation.Otp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldOtp,
		})
	}
	if auo.mutation.OtpCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldOtp,
		})
	}
	if value, ok := auo.mutation.EmailChange(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldEmailChange,
		})
	}
	if auo.mutation.EmailChangeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldEmailChange,
		})
	}
	if value, ok := auo.mutation.EmailChangeSentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldEmailChangeSentAt,
		})
	}
	if auo.mutation.EmailChangeSentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: account.FieldEmailChangeSentAt,
		})
	}
	if value, ok := auo.mutation.EmailChangeToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldEmailChangeToken,
		})
	}
	if auo.mutation.EmailChangeTokenCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: account.FieldEmailChangeToken,
		})
	}
	if value, ok := auo.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: account.FieldMetadata,
		})
	}
	if value, ok := auo.mutation.Roles(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: account.FieldRoles,
		})
	}
	if auo.mutation.RolesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: account.FieldRoles,
		})
	}
	if value, ok := auo.mutation.Teams(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: account.FieldTeams,
		})
	}
	if auo.mutation.TeamsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: account.FieldTeams,
		})
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldUpdatedAt,
		})
	}
	if value, ok := auo.mutation.LastSigninAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldLastSigninAt,
		})
	}
	if auo.mutation.LastSigninAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: account.FieldLastSigninAt,
		})
	}
	if auo.mutation.WorkspaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   account.WorkspaceTable,
			Columns: []string{account.WorkspaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workspace.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.WorkspaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   account.WorkspaceTable,
			Columns: []string{account.WorkspaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workspace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.WorkspaceRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.WorkspaceRolesTable,
			Columns: []string{account.WorkspaceRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workspacerole.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedWorkspaceRolesIDs(); len(nodes) > 0 && !auo.mutation.WorkspaceRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.WorkspaceRolesTable,
			Columns: []string{account.WorkspaceRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workspacerole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.WorkspaceRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.WorkspaceRolesTable,
			Columns: []string{account.WorkspaceRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workspacerole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.GroupRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.GroupRolesTable,
			Columns: []string{account.GroupRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: grouprole.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedGroupRolesIDs(); len(nodes) > 0 && !auo.mutation.GroupRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.GroupRolesTable,
			Columns: []string{account.GroupRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: grouprole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.GroupRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.GroupRolesTable,
			Columns: []string{account.GroupRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: grouprole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.AccountRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.AccountRolesTable,
			Columns: []string{account.AccountRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: accountrole.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedAccountRolesIDs(); len(nodes) > 0 && !auo.mutation.AccountRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.AccountRolesTable,
			Columns: []string{account.AccountRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: accountrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AccountRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.AccountRolesTable,
			Columns: []string{account.AccountRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: accountrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Account{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
