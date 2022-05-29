// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/adnaan/authn/models/account"
	"github.com/adnaan/authn/models/predicate"
	"github.com/adnaan/authn/models/session"
	"github.com/google/uuid"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeAccount = "Account"
	TypeSession = "Session"
)

// AccountMutation represents an operation that mutates the Account nodes in the graph.
type AccountMutation struct {
	config
	op                   Op
	typ                  string
	id                   *uuid.UUID
	provider             *string
	email                *string
	password             *string
	locked               *bool
	confirmed            *bool
	confirmation_sent_at *time.Time
	confirmation_token   *string
	recovery_sent_at     *time.Time
	recovery_token       *string
	otp_sent_at          *time.Time
	otp                  *string
	email_change         *string
	email_change_sent_at *time.Time
	email_change_token   *string
	attributes           *map[string]interface{}
	sensitive_attributes *map[string]string
	attribute_bytes      *[]byte
	created_at           *time.Time
	updated_at           *time.Time
	last_signin_at       *time.Time
	clearedFields        map[string]struct{}
	done                 bool
	oldValue             func(context.Context) (*Account, error)
	predicates           []predicate.Account
}

var _ ent.Mutation = (*AccountMutation)(nil)

// accountOption allows management of the mutation configuration using functional options.
type accountOption func(*AccountMutation)

// newAccountMutation creates new mutation for the Account entity.
func newAccountMutation(c config, op Op, opts ...accountOption) *AccountMutation {
	m := &AccountMutation{
		config:        c,
		op:            op,
		typ:           TypeAccount,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAccountID sets the ID field of the mutation.
func withAccountID(id uuid.UUID) accountOption {
	return func(m *AccountMutation) {
		var (
			err   error
			once  sync.Once
			value *Account
		)
		m.oldValue = func(ctx context.Context) (*Account, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Account.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withAccount sets the old Account of the mutation.
func withAccount(node *Account) accountOption {
	return func(m *AccountMutation) {
		m.oldValue = func(context.Context) (*Account, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AccountMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AccountMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("models: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Account entities.
func (m *AccountMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *AccountMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *AccountMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Account.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetProvider sets the "provider" field.
func (m *AccountMutation) SetProvider(s string) {
	m.provider = &s
}

// Provider returns the value of the "provider" field in the mutation.
func (m *AccountMutation) Provider() (r string, exists bool) {
	v := m.provider
	if v == nil {
		return
	}
	return *v, true
}

// OldProvider returns the old "provider" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldProvider(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldProvider is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldProvider requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldProvider: %w", err)
	}
	return oldValue.Provider, nil
}

// ResetProvider resets all changes to the "provider" field.
func (m *AccountMutation) ResetProvider() {
	m.provider = nil
}

// SetEmail sets the "email" field.
func (m *AccountMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *AccountMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *AccountMutation) ResetEmail() {
	m.email = nil
}

// SetPassword sets the "password" field.
func (m *AccountMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *AccountMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *AccountMutation) ResetPassword() {
	m.password = nil
}

// SetLocked sets the "locked" field.
func (m *AccountMutation) SetLocked(b bool) {
	m.locked = &b
}

// Locked returns the value of the "locked" field in the mutation.
func (m *AccountMutation) Locked() (r bool, exists bool) {
	v := m.locked
	if v == nil {
		return
	}
	return *v, true
}

// OldLocked returns the old "locked" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldLocked(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLocked is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLocked requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLocked: %w", err)
	}
	return oldValue.Locked, nil
}

// ResetLocked resets all changes to the "locked" field.
func (m *AccountMutation) ResetLocked() {
	m.locked = nil
}

// SetConfirmed sets the "confirmed" field.
func (m *AccountMutation) SetConfirmed(b bool) {
	m.confirmed = &b
}

// Confirmed returns the value of the "confirmed" field in the mutation.
func (m *AccountMutation) Confirmed() (r bool, exists bool) {
	v := m.confirmed
	if v == nil {
		return
	}
	return *v, true
}

// OldConfirmed returns the old "confirmed" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldConfirmed(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldConfirmed is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldConfirmed requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldConfirmed: %w", err)
	}
	return oldValue.Confirmed, nil
}

// ClearConfirmed clears the value of the "confirmed" field.
func (m *AccountMutation) ClearConfirmed() {
	m.confirmed = nil
	m.clearedFields[account.FieldConfirmed] = struct{}{}
}

// ConfirmedCleared returns if the "confirmed" field was cleared in this mutation.
func (m *AccountMutation) ConfirmedCleared() bool {
	_, ok := m.clearedFields[account.FieldConfirmed]
	return ok
}

// ResetConfirmed resets all changes to the "confirmed" field.
func (m *AccountMutation) ResetConfirmed() {
	m.confirmed = nil
	delete(m.clearedFields, account.FieldConfirmed)
}

// SetConfirmationSentAt sets the "confirmation_sent_at" field.
func (m *AccountMutation) SetConfirmationSentAt(t time.Time) {
	m.confirmation_sent_at = &t
}

// ConfirmationSentAt returns the value of the "confirmation_sent_at" field in the mutation.
func (m *AccountMutation) ConfirmationSentAt() (r time.Time, exists bool) {
	v := m.confirmation_sent_at
	if v == nil {
		return
	}
	return *v, true
}

// OldConfirmationSentAt returns the old "confirmation_sent_at" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldConfirmationSentAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldConfirmationSentAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldConfirmationSentAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldConfirmationSentAt: %w", err)
	}
	return oldValue.ConfirmationSentAt, nil
}

// ClearConfirmationSentAt clears the value of the "confirmation_sent_at" field.
func (m *AccountMutation) ClearConfirmationSentAt() {
	m.confirmation_sent_at = nil
	m.clearedFields[account.FieldConfirmationSentAt] = struct{}{}
}

// ConfirmationSentAtCleared returns if the "confirmation_sent_at" field was cleared in this mutation.
func (m *AccountMutation) ConfirmationSentAtCleared() bool {
	_, ok := m.clearedFields[account.FieldConfirmationSentAt]
	return ok
}

// ResetConfirmationSentAt resets all changes to the "confirmation_sent_at" field.
func (m *AccountMutation) ResetConfirmationSentAt() {
	m.confirmation_sent_at = nil
	delete(m.clearedFields, account.FieldConfirmationSentAt)
}

// SetConfirmationToken sets the "confirmation_token" field.
func (m *AccountMutation) SetConfirmationToken(s string) {
	m.confirmation_token = &s
}

// ConfirmationToken returns the value of the "confirmation_token" field in the mutation.
func (m *AccountMutation) ConfirmationToken() (r string, exists bool) {
	v := m.confirmation_token
	if v == nil {
		return
	}
	return *v, true
}

// OldConfirmationToken returns the old "confirmation_token" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldConfirmationToken(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldConfirmationToken is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldConfirmationToken requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldConfirmationToken: %w", err)
	}
	return oldValue.ConfirmationToken, nil
}

// ClearConfirmationToken clears the value of the "confirmation_token" field.
func (m *AccountMutation) ClearConfirmationToken() {
	m.confirmation_token = nil
	m.clearedFields[account.FieldConfirmationToken] = struct{}{}
}

// ConfirmationTokenCleared returns if the "confirmation_token" field was cleared in this mutation.
func (m *AccountMutation) ConfirmationTokenCleared() bool {
	_, ok := m.clearedFields[account.FieldConfirmationToken]
	return ok
}

// ResetConfirmationToken resets all changes to the "confirmation_token" field.
func (m *AccountMutation) ResetConfirmationToken() {
	m.confirmation_token = nil
	delete(m.clearedFields, account.FieldConfirmationToken)
}

// SetRecoverySentAt sets the "recovery_sent_at" field.
func (m *AccountMutation) SetRecoverySentAt(t time.Time) {
	m.recovery_sent_at = &t
}

// RecoverySentAt returns the value of the "recovery_sent_at" field in the mutation.
func (m *AccountMutation) RecoverySentAt() (r time.Time, exists bool) {
	v := m.recovery_sent_at
	if v == nil {
		return
	}
	return *v, true
}

// OldRecoverySentAt returns the old "recovery_sent_at" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldRecoverySentAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRecoverySentAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRecoverySentAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRecoverySentAt: %w", err)
	}
	return oldValue.RecoverySentAt, nil
}

// ClearRecoverySentAt clears the value of the "recovery_sent_at" field.
func (m *AccountMutation) ClearRecoverySentAt() {
	m.recovery_sent_at = nil
	m.clearedFields[account.FieldRecoverySentAt] = struct{}{}
}

// RecoverySentAtCleared returns if the "recovery_sent_at" field was cleared in this mutation.
func (m *AccountMutation) RecoverySentAtCleared() bool {
	_, ok := m.clearedFields[account.FieldRecoverySentAt]
	return ok
}

// ResetRecoverySentAt resets all changes to the "recovery_sent_at" field.
func (m *AccountMutation) ResetRecoverySentAt() {
	m.recovery_sent_at = nil
	delete(m.clearedFields, account.FieldRecoverySentAt)
}

// SetRecoveryToken sets the "recovery_token" field.
func (m *AccountMutation) SetRecoveryToken(s string) {
	m.recovery_token = &s
}

// RecoveryToken returns the value of the "recovery_token" field in the mutation.
func (m *AccountMutation) RecoveryToken() (r string, exists bool) {
	v := m.recovery_token
	if v == nil {
		return
	}
	return *v, true
}

// OldRecoveryToken returns the old "recovery_token" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldRecoveryToken(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRecoveryToken is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRecoveryToken requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRecoveryToken: %w", err)
	}
	return oldValue.RecoveryToken, nil
}

// ClearRecoveryToken clears the value of the "recovery_token" field.
func (m *AccountMutation) ClearRecoveryToken() {
	m.recovery_token = nil
	m.clearedFields[account.FieldRecoveryToken] = struct{}{}
}

// RecoveryTokenCleared returns if the "recovery_token" field was cleared in this mutation.
func (m *AccountMutation) RecoveryTokenCleared() bool {
	_, ok := m.clearedFields[account.FieldRecoveryToken]
	return ok
}

// ResetRecoveryToken resets all changes to the "recovery_token" field.
func (m *AccountMutation) ResetRecoveryToken() {
	m.recovery_token = nil
	delete(m.clearedFields, account.FieldRecoveryToken)
}

// SetOtpSentAt sets the "otp_sent_at" field.
func (m *AccountMutation) SetOtpSentAt(t time.Time) {
	m.otp_sent_at = &t
}

// OtpSentAt returns the value of the "otp_sent_at" field in the mutation.
func (m *AccountMutation) OtpSentAt() (r time.Time, exists bool) {
	v := m.otp_sent_at
	if v == nil {
		return
	}
	return *v, true
}

// OldOtpSentAt returns the old "otp_sent_at" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldOtpSentAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldOtpSentAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldOtpSentAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOtpSentAt: %w", err)
	}
	return oldValue.OtpSentAt, nil
}

// ClearOtpSentAt clears the value of the "otp_sent_at" field.
func (m *AccountMutation) ClearOtpSentAt() {
	m.otp_sent_at = nil
	m.clearedFields[account.FieldOtpSentAt] = struct{}{}
}

// OtpSentAtCleared returns if the "otp_sent_at" field was cleared in this mutation.
func (m *AccountMutation) OtpSentAtCleared() bool {
	_, ok := m.clearedFields[account.FieldOtpSentAt]
	return ok
}

// ResetOtpSentAt resets all changes to the "otp_sent_at" field.
func (m *AccountMutation) ResetOtpSentAt() {
	m.otp_sent_at = nil
	delete(m.clearedFields, account.FieldOtpSentAt)
}

// SetOtp sets the "otp" field.
func (m *AccountMutation) SetOtp(s string) {
	m.otp = &s
}

// Otp returns the value of the "otp" field in the mutation.
func (m *AccountMutation) Otp() (r string, exists bool) {
	v := m.otp
	if v == nil {
		return
	}
	return *v, true
}

// OldOtp returns the old "otp" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldOtp(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldOtp is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldOtp requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOtp: %w", err)
	}
	return oldValue.Otp, nil
}

// ClearOtp clears the value of the "otp" field.
func (m *AccountMutation) ClearOtp() {
	m.otp = nil
	m.clearedFields[account.FieldOtp] = struct{}{}
}

// OtpCleared returns if the "otp" field was cleared in this mutation.
func (m *AccountMutation) OtpCleared() bool {
	_, ok := m.clearedFields[account.FieldOtp]
	return ok
}

// ResetOtp resets all changes to the "otp" field.
func (m *AccountMutation) ResetOtp() {
	m.otp = nil
	delete(m.clearedFields, account.FieldOtp)
}

// SetEmailChange sets the "email_change" field.
func (m *AccountMutation) SetEmailChange(s string) {
	m.email_change = &s
}

// EmailChange returns the value of the "email_change" field in the mutation.
func (m *AccountMutation) EmailChange() (r string, exists bool) {
	v := m.email_change
	if v == nil {
		return
	}
	return *v, true
}

// OldEmailChange returns the old "email_change" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldEmailChange(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmailChange is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmailChange requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmailChange: %w", err)
	}
	return oldValue.EmailChange, nil
}

// ClearEmailChange clears the value of the "email_change" field.
func (m *AccountMutation) ClearEmailChange() {
	m.email_change = nil
	m.clearedFields[account.FieldEmailChange] = struct{}{}
}

// EmailChangeCleared returns if the "email_change" field was cleared in this mutation.
func (m *AccountMutation) EmailChangeCleared() bool {
	_, ok := m.clearedFields[account.FieldEmailChange]
	return ok
}

// ResetEmailChange resets all changes to the "email_change" field.
func (m *AccountMutation) ResetEmailChange() {
	m.email_change = nil
	delete(m.clearedFields, account.FieldEmailChange)
}

// SetEmailChangeSentAt sets the "email_change_sent_at" field.
func (m *AccountMutation) SetEmailChangeSentAt(t time.Time) {
	m.email_change_sent_at = &t
}

// EmailChangeSentAt returns the value of the "email_change_sent_at" field in the mutation.
func (m *AccountMutation) EmailChangeSentAt() (r time.Time, exists bool) {
	v := m.email_change_sent_at
	if v == nil {
		return
	}
	return *v, true
}

// OldEmailChangeSentAt returns the old "email_change_sent_at" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldEmailChangeSentAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmailChangeSentAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmailChangeSentAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmailChangeSentAt: %w", err)
	}
	return oldValue.EmailChangeSentAt, nil
}

// ClearEmailChangeSentAt clears the value of the "email_change_sent_at" field.
func (m *AccountMutation) ClearEmailChangeSentAt() {
	m.email_change_sent_at = nil
	m.clearedFields[account.FieldEmailChangeSentAt] = struct{}{}
}

// EmailChangeSentAtCleared returns if the "email_change_sent_at" field was cleared in this mutation.
func (m *AccountMutation) EmailChangeSentAtCleared() bool {
	_, ok := m.clearedFields[account.FieldEmailChangeSentAt]
	return ok
}

// ResetEmailChangeSentAt resets all changes to the "email_change_sent_at" field.
func (m *AccountMutation) ResetEmailChangeSentAt() {
	m.email_change_sent_at = nil
	delete(m.clearedFields, account.FieldEmailChangeSentAt)
}

// SetEmailChangeToken sets the "email_change_token" field.
func (m *AccountMutation) SetEmailChangeToken(s string) {
	m.email_change_token = &s
}

// EmailChangeToken returns the value of the "email_change_token" field in the mutation.
func (m *AccountMutation) EmailChangeToken() (r string, exists bool) {
	v := m.email_change_token
	if v == nil {
		return
	}
	return *v, true
}

// OldEmailChangeToken returns the old "email_change_token" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldEmailChangeToken(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmailChangeToken is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmailChangeToken requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmailChangeToken: %w", err)
	}
	return oldValue.EmailChangeToken, nil
}

// ClearEmailChangeToken clears the value of the "email_change_token" field.
func (m *AccountMutation) ClearEmailChangeToken() {
	m.email_change_token = nil
	m.clearedFields[account.FieldEmailChangeToken] = struct{}{}
}

// EmailChangeTokenCleared returns if the "email_change_token" field was cleared in this mutation.
func (m *AccountMutation) EmailChangeTokenCleared() bool {
	_, ok := m.clearedFields[account.FieldEmailChangeToken]
	return ok
}

// ResetEmailChangeToken resets all changes to the "email_change_token" field.
func (m *AccountMutation) ResetEmailChangeToken() {
	m.email_change_token = nil
	delete(m.clearedFields, account.FieldEmailChangeToken)
}

// SetAttributes sets the "attributes" field.
func (m *AccountMutation) SetAttributes(value map[string]interface{}) {
	m.attributes = &value
}

// Attributes returns the value of the "attributes" field in the mutation.
func (m *AccountMutation) Attributes() (r map[string]interface{}, exists bool) {
	v := m.attributes
	if v == nil {
		return
	}
	return *v, true
}

// OldAttributes returns the old "attributes" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldAttributes(ctx context.Context) (v map[string]interface{}, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAttributes is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAttributes requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAttributes: %w", err)
	}
	return oldValue.Attributes, nil
}

// ClearAttributes clears the value of the "attributes" field.
func (m *AccountMutation) ClearAttributes() {
	m.attributes = nil
	m.clearedFields[account.FieldAttributes] = struct{}{}
}

// AttributesCleared returns if the "attributes" field was cleared in this mutation.
func (m *AccountMutation) AttributesCleared() bool {
	_, ok := m.clearedFields[account.FieldAttributes]
	return ok
}

// ResetAttributes resets all changes to the "attributes" field.
func (m *AccountMutation) ResetAttributes() {
	m.attributes = nil
	delete(m.clearedFields, account.FieldAttributes)
}

// SetSensitiveAttributes sets the "sensitive_attributes" field.
func (m *AccountMutation) SetSensitiveAttributes(value map[string]string) {
	m.sensitive_attributes = &value
}

// SensitiveAttributes returns the value of the "sensitive_attributes" field in the mutation.
func (m *AccountMutation) SensitiveAttributes() (r map[string]string, exists bool) {
	v := m.sensitive_attributes
	if v == nil {
		return
	}
	return *v, true
}

// OldSensitiveAttributes returns the old "sensitive_attributes" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldSensitiveAttributes(ctx context.Context) (v map[string]string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSensitiveAttributes is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSensitiveAttributes requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSensitiveAttributes: %w", err)
	}
	return oldValue.SensitiveAttributes, nil
}

// ClearSensitiveAttributes clears the value of the "sensitive_attributes" field.
func (m *AccountMutation) ClearSensitiveAttributes() {
	m.sensitive_attributes = nil
	m.clearedFields[account.FieldSensitiveAttributes] = struct{}{}
}

// SensitiveAttributesCleared returns if the "sensitive_attributes" field was cleared in this mutation.
func (m *AccountMutation) SensitiveAttributesCleared() bool {
	_, ok := m.clearedFields[account.FieldSensitiveAttributes]
	return ok
}

// ResetSensitiveAttributes resets all changes to the "sensitive_attributes" field.
func (m *AccountMutation) ResetSensitiveAttributes() {
	m.sensitive_attributes = nil
	delete(m.clearedFields, account.FieldSensitiveAttributes)
}

// SetAttributeBytes sets the "attribute_bytes" field.
func (m *AccountMutation) SetAttributeBytes(b []byte) {
	m.attribute_bytes = &b
}

// AttributeBytes returns the value of the "attribute_bytes" field in the mutation.
func (m *AccountMutation) AttributeBytes() (r []byte, exists bool) {
	v := m.attribute_bytes
	if v == nil {
		return
	}
	return *v, true
}

// OldAttributeBytes returns the old "attribute_bytes" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldAttributeBytes(ctx context.Context) (v []byte, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAttributeBytes is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAttributeBytes requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAttributeBytes: %w", err)
	}
	return oldValue.AttributeBytes, nil
}

// ClearAttributeBytes clears the value of the "attribute_bytes" field.
func (m *AccountMutation) ClearAttributeBytes() {
	m.attribute_bytes = nil
	m.clearedFields[account.FieldAttributeBytes] = struct{}{}
}

// AttributeBytesCleared returns if the "attribute_bytes" field was cleared in this mutation.
func (m *AccountMutation) AttributeBytesCleared() bool {
	_, ok := m.clearedFields[account.FieldAttributeBytes]
	return ok
}

// ResetAttributeBytes resets all changes to the "attribute_bytes" field.
func (m *AccountMutation) ResetAttributeBytes() {
	m.attribute_bytes = nil
	delete(m.clearedFields, account.FieldAttributeBytes)
}

// SetCreatedAt sets the "created_at" field.
func (m *AccountMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *AccountMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *AccountMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *AccountMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *AccountMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *AccountMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetLastSigninAt sets the "last_signin_at" field.
func (m *AccountMutation) SetLastSigninAt(t time.Time) {
	m.last_signin_at = &t
}

// LastSigninAt returns the value of the "last_signin_at" field in the mutation.
func (m *AccountMutation) LastSigninAt() (r time.Time, exists bool) {
	v := m.last_signin_at
	if v == nil {
		return
	}
	return *v, true
}

// OldLastSigninAt returns the old "last_signin_at" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldLastSigninAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLastSigninAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLastSigninAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastSigninAt: %w", err)
	}
	return oldValue.LastSigninAt, nil
}

// ClearLastSigninAt clears the value of the "last_signin_at" field.
func (m *AccountMutation) ClearLastSigninAt() {
	m.last_signin_at = nil
	m.clearedFields[account.FieldLastSigninAt] = struct{}{}
}

// LastSigninAtCleared returns if the "last_signin_at" field was cleared in this mutation.
func (m *AccountMutation) LastSigninAtCleared() bool {
	_, ok := m.clearedFields[account.FieldLastSigninAt]
	return ok
}

// ResetLastSigninAt resets all changes to the "last_signin_at" field.
func (m *AccountMutation) ResetLastSigninAt() {
	m.last_signin_at = nil
	delete(m.clearedFields, account.FieldLastSigninAt)
}

// Where appends a list predicates to the AccountMutation builder.
func (m *AccountMutation) Where(ps ...predicate.Account) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *AccountMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Account).
func (m *AccountMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *AccountMutation) Fields() []string {
	fields := make([]string, 0, 20)
	if m.provider != nil {
		fields = append(fields, account.FieldProvider)
	}
	if m.email != nil {
		fields = append(fields, account.FieldEmail)
	}
	if m.password != nil {
		fields = append(fields, account.FieldPassword)
	}
	if m.locked != nil {
		fields = append(fields, account.FieldLocked)
	}
	if m.confirmed != nil {
		fields = append(fields, account.FieldConfirmed)
	}
	if m.confirmation_sent_at != nil {
		fields = append(fields, account.FieldConfirmationSentAt)
	}
	if m.confirmation_token != nil {
		fields = append(fields, account.FieldConfirmationToken)
	}
	if m.recovery_sent_at != nil {
		fields = append(fields, account.FieldRecoverySentAt)
	}
	if m.recovery_token != nil {
		fields = append(fields, account.FieldRecoveryToken)
	}
	if m.otp_sent_at != nil {
		fields = append(fields, account.FieldOtpSentAt)
	}
	if m.otp != nil {
		fields = append(fields, account.FieldOtp)
	}
	if m.email_change != nil {
		fields = append(fields, account.FieldEmailChange)
	}
	if m.email_change_sent_at != nil {
		fields = append(fields, account.FieldEmailChangeSentAt)
	}
	if m.email_change_token != nil {
		fields = append(fields, account.FieldEmailChangeToken)
	}
	if m.attributes != nil {
		fields = append(fields, account.FieldAttributes)
	}
	if m.sensitive_attributes != nil {
		fields = append(fields, account.FieldSensitiveAttributes)
	}
	if m.attribute_bytes != nil {
		fields = append(fields, account.FieldAttributeBytes)
	}
	if m.created_at != nil {
		fields = append(fields, account.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, account.FieldUpdatedAt)
	}
	if m.last_signin_at != nil {
		fields = append(fields, account.FieldLastSigninAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *AccountMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case account.FieldProvider:
		return m.Provider()
	case account.FieldEmail:
		return m.Email()
	case account.FieldPassword:
		return m.Password()
	case account.FieldLocked:
		return m.Locked()
	case account.FieldConfirmed:
		return m.Confirmed()
	case account.FieldConfirmationSentAt:
		return m.ConfirmationSentAt()
	case account.FieldConfirmationToken:
		return m.ConfirmationToken()
	case account.FieldRecoverySentAt:
		return m.RecoverySentAt()
	case account.FieldRecoveryToken:
		return m.RecoveryToken()
	case account.FieldOtpSentAt:
		return m.OtpSentAt()
	case account.FieldOtp:
		return m.Otp()
	case account.FieldEmailChange:
		return m.EmailChange()
	case account.FieldEmailChangeSentAt:
		return m.EmailChangeSentAt()
	case account.FieldEmailChangeToken:
		return m.EmailChangeToken()
	case account.FieldAttributes:
		return m.Attributes()
	case account.FieldSensitiveAttributes:
		return m.SensitiveAttributes()
	case account.FieldAttributeBytes:
		return m.AttributeBytes()
	case account.FieldCreatedAt:
		return m.CreatedAt()
	case account.FieldUpdatedAt:
		return m.UpdatedAt()
	case account.FieldLastSigninAt:
		return m.LastSigninAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *AccountMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case account.FieldProvider:
		return m.OldProvider(ctx)
	case account.FieldEmail:
		return m.OldEmail(ctx)
	case account.FieldPassword:
		return m.OldPassword(ctx)
	case account.FieldLocked:
		return m.OldLocked(ctx)
	case account.FieldConfirmed:
		return m.OldConfirmed(ctx)
	case account.FieldConfirmationSentAt:
		return m.OldConfirmationSentAt(ctx)
	case account.FieldConfirmationToken:
		return m.OldConfirmationToken(ctx)
	case account.FieldRecoverySentAt:
		return m.OldRecoverySentAt(ctx)
	case account.FieldRecoveryToken:
		return m.OldRecoveryToken(ctx)
	case account.FieldOtpSentAt:
		return m.OldOtpSentAt(ctx)
	case account.FieldOtp:
		return m.OldOtp(ctx)
	case account.FieldEmailChange:
		return m.OldEmailChange(ctx)
	case account.FieldEmailChangeSentAt:
		return m.OldEmailChangeSentAt(ctx)
	case account.FieldEmailChangeToken:
		return m.OldEmailChangeToken(ctx)
	case account.FieldAttributes:
		return m.OldAttributes(ctx)
	case account.FieldSensitiveAttributes:
		return m.OldSensitiveAttributes(ctx)
	case account.FieldAttributeBytes:
		return m.OldAttributeBytes(ctx)
	case account.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case account.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case account.FieldLastSigninAt:
		return m.OldLastSigninAt(ctx)
	}
	return nil, fmt.Errorf("unknown Account field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AccountMutation) SetField(name string, value ent.Value) error {
	switch name {
	case account.FieldProvider:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetProvider(v)
		return nil
	case account.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case account.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case account.FieldLocked:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLocked(v)
		return nil
	case account.FieldConfirmed:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetConfirmed(v)
		return nil
	case account.FieldConfirmationSentAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetConfirmationSentAt(v)
		return nil
	case account.FieldConfirmationToken:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetConfirmationToken(v)
		return nil
	case account.FieldRecoverySentAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRecoverySentAt(v)
		return nil
	case account.FieldRecoveryToken:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRecoveryToken(v)
		return nil
	case account.FieldOtpSentAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOtpSentAt(v)
		return nil
	case account.FieldOtp:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOtp(v)
		return nil
	case account.FieldEmailChange:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmailChange(v)
		return nil
	case account.FieldEmailChangeSentAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmailChangeSentAt(v)
		return nil
	case account.FieldEmailChangeToken:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmailChangeToken(v)
		return nil
	case account.FieldAttributes:
		v, ok := value.(map[string]interface{})
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAttributes(v)
		return nil
	case account.FieldSensitiveAttributes:
		v, ok := value.(map[string]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSensitiveAttributes(v)
		return nil
	case account.FieldAttributeBytes:
		v, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAttributeBytes(v)
		return nil
	case account.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case account.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case account.FieldLastSigninAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastSigninAt(v)
		return nil
	}
	return fmt.Errorf("unknown Account field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *AccountMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *AccountMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AccountMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Account numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *AccountMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(account.FieldConfirmed) {
		fields = append(fields, account.FieldConfirmed)
	}
	if m.FieldCleared(account.FieldConfirmationSentAt) {
		fields = append(fields, account.FieldConfirmationSentAt)
	}
	if m.FieldCleared(account.FieldConfirmationToken) {
		fields = append(fields, account.FieldConfirmationToken)
	}
	if m.FieldCleared(account.FieldRecoverySentAt) {
		fields = append(fields, account.FieldRecoverySentAt)
	}
	if m.FieldCleared(account.FieldRecoveryToken) {
		fields = append(fields, account.FieldRecoveryToken)
	}
	if m.FieldCleared(account.FieldOtpSentAt) {
		fields = append(fields, account.FieldOtpSentAt)
	}
	if m.FieldCleared(account.FieldOtp) {
		fields = append(fields, account.FieldOtp)
	}
	if m.FieldCleared(account.FieldEmailChange) {
		fields = append(fields, account.FieldEmailChange)
	}
	if m.FieldCleared(account.FieldEmailChangeSentAt) {
		fields = append(fields, account.FieldEmailChangeSentAt)
	}
	if m.FieldCleared(account.FieldEmailChangeToken) {
		fields = append(fields, account.FieldEmailChangeToken)
	}
	if m.FieldCleared(account.FieldAttributes) {
		fields = append(fields, account.FieldAttributes)
	}
	if m.FieldCleared(account.FieldSensitiveAttributes) {
		fields = append(fields, account.FieldSensitiveAttributes)
	}
	if m.FieldCleared(account.FieldAttributeBytes) {
		fields = append(fields, account.FieldAttributeBytes)
	}
	if m.FieldCleared(account.FieldLastSigninAt) {
		fields = append(fields, account.FieldLastSigninAt)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *AccountMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *AccountMutation) ClearField(name string) error {
	switch name {
	case account.FieldConfirmed:
		m.ClearConfirmed()
		return nil
	case account.FieldConfirmationSentAt:
		m.ClearConfirmationSentAt()
		return nil
	case account.FieldConfirmationToken:
		m.ClearConfirmationToken()
		return nil
	case account.FieldRecoverySentAt:
		m.ClearRecoverySentAt()
		return nil
	case account.FieldRecoveryToken:
		m.ClearRecoveryToken()
		return nil
	case account.FieldOtpSentAt:
		m.ClearOtpSentAt()
		return nil
	case account.FieldOtp:
		m.ClearOtp()
		return nil
	case account.FieldEmailChange:
		m.ClearEmailChange()
		return nil
	case account.FieldEmailChangeSentAt:
		m.ClearEmailChangeSentAt()
		return nil
	case account.FieldEmailChangeToken:
		m.ClearEmailChangeToken()
		return nil
	case account.FieldAttributes:
		m.ClearAttributes()
		return nil
	case account.FieldSensitiveAttributes:
		m.ClearSensitiveAttributes()
		return nil
	case account.FieldAttributeBytes:
		m.ClearAttributeBytes()
		return nil
	case account.FieldLastSigninAt:
		m.ClearLastSigninAt()
		return nil
	}
	return fmt.Errorf("unknown Account nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *AccountMutation) ResetField(name string) error {
	switch name {
	case account.FieldProvider:
		m.ResetProvider()
		return nil
	case account.FieldEmail:
		m.ResetEmail()
		return nil
	case account.FieldPassword:
		m.ResetPassword()
		return nil
	case account.FieldLocked:
		m.ResetLocked()
		return nil
	case account.FieldConfirmed:
		m.ResetConfirmed()
		return nil
	case account.FieldConfirmationSentAt:
		m.ResetConfirmationSentAt()
		return nil
	case account.FieldConfirmationToken:
		m.ResetConfirmationToken()
		return nil
	case account.FieldRecoverySentAt:
		m.ResetRecoverySentAt()
		return nil
	case account.FieldRecoveryToken:
		m.ResetRecoveryToken()
		return nil
	case account.FieldOtpSentAt:
		m.ResetOtpSentAt()
		return nil
	case account.FieldOtp:
		m.ResetOtp()
		return nil
	case account.FieldEmailChange:
		m.ResetEmailChange()
		return nil
	case account.FieldEmailChangeSentAt:
		m.ResetEmailChangeSentAt()
		return nil
	case account.FieldEmailChangeToken:
		m.ResetEmailChangeToken()
		return nil
	case account.FieldAttributes:
		m.ResetAttributes()
		return nil
	case account.FieldSensitiveAttributes:
		m.ResetSensitiveAttributes()
		return nil
	case account.FieldAttributeBytes:
		m.ResetAttributeBytes()
		return nil
	case account.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case account.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case account.FieldLastSigninAt:
		m.ResetLastSigninAt()
		return nil
	}
	return fmt.Errorf("unknown Account field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *AccountMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *AccountMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *AccountMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *AccountMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *AccountMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *AccountMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *AccountMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Account unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *AccountMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Account edge %s", name)
}

// SessionMutation represents an operation that mutates the Session nodes in the graph.
type SessionMutation struct {
	config
	op            Op
	typ           string
	id            *string
	data          *string
	created_at    *time.Time
	updated_at    *time.Time
	expires_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Session, error)
	predicates    []predicate.Session
}

var _ ent.Mutation = (*SessionMutation)(nil)

// sessionOption allows management of the mutation configuration using functional options.
type sessionOption func(*SessionMutation)

// newSessionMutation creates new mutation for the Session entity.
func newSessionMutation(c config, op Op, opts ...sessionOption) *SessionMutation {
	m := &SessionMutation{
		config:        c,
		op:            op,
		typ:           TypeSession,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withSessionID sets the ID field of the mutation.
func withSessionID(id string) sessionOption {
	return func(m *SessionMutation) {
		var (
			err   error
			once  sync.Once
			value *Session
		)
		m.oldValue = func(ctx context.Context) (*Session, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Session.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withSession sets the old Session of the mutation.
func withSession(node *Session) sessionOption {
	return func(m *SessionMutation) {
		m.oldValue = func(context.Context) (*Session, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m SessionMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m SessionMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("models: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Session entities.
func (m *SessionMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *SessionMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *SessionMutation) IDs(ctx context.Context) ([]string, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []string{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Session.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetData sets the "data" field.
func (m *SessionMutation) SetData(s string) {
	m.data = &s
}

// Data returns the value of the "data" field in the mutation.
func (m *SessionMutation) Data() (r string, exists bool) {
	v := m.data
	if v == nil {
		return
	}
	return *v, true
}

// OldData returns the old "data" field's value of the Session entity.
// If the Session object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SessionMutation) OldData(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldData is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldData requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldData: %w", err)
	}
	return oldValue.Data, nil
}

// ResetData resets all changes to the "data" field.
func (m *SessionMutation) ResetData() {
	m.data = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *SessionMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *SessionMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Session entity.
// If the Session object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SessionMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *SessionMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *SessionMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *SessionMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Session entity.
// If the Session object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SessionMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *SessionMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetExpiresAt sets the "expires_at" field.
func (m *SessionMutation) SetExpiresAt(t time.Time) {
	m.expires_at = &t
}

// ExpiresAt returns the value of the "expires_at" field in the mutation.
func (m *SessionMutation) ExpiresAt() (r time.Time, exists bool) {
	v := m.expires_at
	if v == nil {
		return
	}
	return *v, true
}

// OldExpiresAt returns the old "expires_at" field's value of the Session entity.
// If the Session object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SessionMutation) OldExpiresAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldExpiresAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldExpiresAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldExpiresAt: %w", err)
	}
	return oldValue.ExpiresAt, nil
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (m *SessionMutation) ClearExpiresAt() {
	m.expires_at = nil
	m.clearedFields[session.FieldExpiresAt] = struct{}{}
}

// ExpiresAtCleared returns if the "expires_at" field was cleared in this mutation.
func (m *SessionMutation) ExpiresAtCleared() bool {
	_, ok := m.clearedFields[session.FieldExpiresAt]
	return ok
}

// ResetExpiresAt resets all changes to the "expires_at" field.
func (m *SessionMutation) ResetExpiresAt() {
	m.expires_at = nil
	delete(m.clearedFields, session.FieldExpiresAt)
}

// Where appends a list predicates to the SessionMutation builder.
func (m *SessionMutation) Where(ps ...predicate.Session) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *SessionMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Session).
func (m *SessionMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *SessionMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.data != nil {
		fields = append(fields, session.FieldData)
	}
	if m.created_at != nil {
		fields = append(fields, session.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, session.FieldUpdatedAt)
	}
	if m.expires_at != nil {
		fields = append(fields, session.FieldExpiresAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *SessionMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case session.FieldData:
		return m.Data()
	case session.FieldCreatedAt:
		return m.CreatedAt()
	case session.FieldUpdatedAt:
		return m.UpdatedAt()
	case session.FieldExpiresAt:
		return m.ExpiresAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *SessionMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case session.FieldData:
		return m.OldData(ctx)
	case session.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case session.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case session.FieldExpiresAt:
		return m.OldExpiresAt(ctx)
	}
	return nil, fmt.Errorf("unknown Session field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SessionMutation) SetField(name string, value ent.Value) error {
	switch name {
	case session.FieldData:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetData(v)
		return nil
	case session.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case session.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case session.FieldExpiresAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetExpiresAt(v)
		return nil
	}
	return fmt.Errorf("unknown Session field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *SessionMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *SessionMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SessionMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Session numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *SessionMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(session.FieldExpiresAt) {
		fields = append(fields, session.FieldExpiresAt)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *SessionMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *SessionMutation) ClearField(name string) error {
	switch name {
	case session.FieldExpiresAt:
		m.ClearExpiresAt()
		return nil
	}
	return fmt.Errorf("unknown Session nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *SessionMutation) ResetField(name string) error {
	switch name {
	case session.FieldData:
		m.ResetData()
		return nil
	case session.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case session.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case session.FieldExpiresAt:
		m.ResetExpiresAt()
		return nil
	}
	return fmt.Errorf("unknown Session field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *SessionMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *SessionMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *SessionMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *SessionMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *SessionMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *SessionMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *SessionMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Session unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *SessionMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Session edge %s", name)
}
