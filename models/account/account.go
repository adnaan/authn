// Code generated (@generated) by entc, DO NOT EDIT.

package account

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProvider holds the string denoting the provider field in the database.
	FieldProvider = "provider"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldLocked holds the string denoting the locked field in the database.
	FieldLocked = "locked"
	// FieldConfirmed holds the string denoting the confirmed field in the database.
	FieldConfirmed = "confirmed"
	// FieldConfirmationSentAt holds the string denoting the confirmation_sent_at field in the database.
	FieldConfirmationSentAt = "confirmation_sent_at"
	// FieldConfirmationToken holds the string denoting the confirmation_token field in the database.
	FieldConfirmationToken = "confirmation_token"
	// FieldRecoverySentAt holds the string denoting the recovery_sent_at field in the database.
	FieldRecoverySentAt = "recovery_sent_at"
	// FieldRecoveryToken holds the string denoting the recovery_token field in the database.
	FieldRecoveryToken = "recovery_token"
	// FieldOtpSentAt holds the string denoting the otp_sent_at field in the database.
	FieldOtpSentAt = "otp_sent_at"
	// FieldOtp holds the string denoting the otp field in the database.
	FieldOtp = "otp"
	// FieldEmailChange holds the string denoting the email_change field in the database.
	FieldEmailChange = "email_change"
	// FieldEmailChangeSentAt holds the string denoting the email_change_sent_at field in the database.
	FieldEmailChangeSentAt = "email_change_sent_at"
	// FieldEmailChangeToken holds the string denoting the email_change_token field in the database.
	FieldEmailChangeToken = "email_change_token"
	// FieldAttributes holds the string denoting the attributes field in the database.
	FieldAttributes = "attributes"
	// FieldSensitiveAttributes holds the string denoting the sensitive_attributes field in the database.
	FieldSensitiveAttributes = "sensitive_attributes"
	// FieldAttributeBytes holds the string denoting the attribute_bytes field in the database.
	FieldAttributeBytes = "attribute_bytes"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldLastSigninAt holds the string denoting the last_signin_at field in the database.
	FieldLastSigninAt = "last_signin_at"
	// Table holds the table name of the account in the database.
	Table = "accounts"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldProvider,
	FieldEmail,
	FieldPassword,
	FieldLocked,
	FieldConfirmed,
	FieldConfirmationSentAt,
	FieldConfirmationToken,
	FieldRecoverySentAt,
	FieldRecoveryToken,
	FieldOtpSentAt,
	FieldOtp,
	FieldEmailChange,
	FieldEmailChangeSentAt,
	FieldEmailChangeToken,
	FieldAttributes,
	FieldSensitiveAttributes,
	FieldAttributeBytes,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldLastSigninAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// ProviderValidator is a validator for the "provider" field. It is called by the builders before save.
	ProviderValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultLocked holds the default value on creation for the "locked" field.
	DefaultLocked bool
	// DefaultConfirmed holds the default value on creation for the "confirmed" field.
	DefaultConfirmed bool
	// ConfirmationTokenValidator is a validator for the "confirmation_token" field. It is called by the builders before save.
	ConfirmationTokenValidator func(string) error
	// RecoveryTokenValidator is a validator for the "recovery_token" field. It is called by the builders before save.
	RecoveryTokenValidator func(string) error
	// OtpValidator is a validator for the "otp" field. It is called by the builders before save.
	OtpValidator func(string) error
	// EmailChangeValidator is a validator for the "email_change" field. It is called by the builders before save.
	EmailChangeValidator func(string) error
	// EmailChangeTokenValidator is a validator for the "email_change_token" field. It is called by the builders before save.
	EmailChangeTokenValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
