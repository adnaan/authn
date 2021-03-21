// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/adnaan/authn/models/account"
	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// Account is the model entity for the Account schema.
type Account struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Provider holds the value of the "provider" field.
	Provider string `json:"provider,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// Locked holds the value of the "locked" field.
	Locked bool `json:"locked,omitempty"`
	// Confirmed holds the value of the "confirmed" field.
	Confirmed bool `json:"confirmed,omitempty"`
	// ConfirmationSentAt holds the value of the "confirmation_sent_at" field.
	ConfirmationSentAt *time.Time `json:"confirmation_sent_at,omitempty"`
	// ConfirmationToken holds the value of the "confirmation_token" field.
	ConfirmationToken *string `json:"confirmation_token,omitempty"`
	// RecoverySentAt holds the value of the "recovery_sent_at" field.
	RecoverySentAt *time.Time `json:"recovery_sent_at,omitempty"`
	// RecoveryToken holds the value of the "recovery_token" field.
	RecoveryToken *string `json:"recovery_token,omitempty"`
	// OtpSentAt holds the value of the "otp_sent_at" field.
	OtpSentAt *time.Time `json:"otp_sent_at,omitempty"`
	// Otp holds the value of the "otp" field.
	Otp *string `json:"otp,omitempty"`
	// EmailChange holds the value of the "email_change" field.
	EmailChange string `json:"email_change,omitempty"`
	// EmailChangeSentAt holds the value of the "email_change_sent_at" field.
	EmailChangeSentAt *time.Time `json:"email_change_sent_at,omitempty"`
	// EmailChangeToken holds the value of the "email_change_token" field.
	EmailChangeToken *string `json:"email_change_token,omitempty"`
	// Attributes holds the value of the "attributes" field.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// SensitiveAttributes holds the value of the "sensitive_attributes" field.
	SensitiveAttributes map[string]string `json:"sensitive_attributes,omitempty"`
	// AttributeBytes holds the value of the "attribute_bytes" field.
	AttributeBytes []byte `json:"attribute_bytes,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// LastSigninAt holds the value of the "last_signin_at" field.
	LastSigninAt *time.Time `json:"last_signin_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Account) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&sql.NullString{}, // provider
		&sql.NullString{}, // email
		&sql.NullString{}, // password
		&sql.NullBool{},   // locked
		&sql.NullBool{},   // confirmed
		&sql.NullTime{},   // confirmation_sent_at
		&sql.NullString{}, // confirmation_token
		&sql.NullTime{},   // recovery_sent_at
		&sql.NullString{}, // recovery_token
		&sql.NullTime{},   // otp_sent_at
		&sql.NullString{}, // otp
		&sql.NullString{}, // email_change
		&sql.NullTime{},   // email_change_sent_at
		&sql.NullString{}, // email_change_token
		&[]byte{},         // attributes
		&[]byte{},         // sensitive_attributes
		&[]byte{},         // attribute_bytes
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
		&sql.NullTime{},   // last_signin_at
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Account fields.
func (a *Account) assignValues(values ...interface{}) error {
	if m, n := len(values), len(account.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		a.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field provider", values[0])
	} else if value.Valid {
		a.Provider = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[1])
	} else if value.Valid {
		a.Email = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field password", values[2])
	} else if value.Valid {
		a.Password = value.String
	}
	if value, ok := values[3].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field locked", values[3])
	} else if value.Valid {
		a.Locked = value.Bool
	}
	if value, ok := values[4].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field confirmed", values[4])
	} else if value.Valid {
		a.Confirmed = value.Bool
	}
	if value, ok := values[5].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field confirmation_sent_at", values[5])
	} else if value.Valid {
		a.ConfirmationSentAt = new(time.Time)
		*a.ConfirmationSentAt = value.Time
	}
	if value, ok := values[6].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field confirmation_token", values[6])
	} else if value.Valid {
		a.ConfirmationToken = new(string)
		*a.ConfirmationToken = value.String
	}
	if value, ok := values[7].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field recovery_sent_at", values[7])
	} else if value.Valid {
		a.RecoverySentAt = new(time.Time)
		*a.RecoverySentAt = value.Time
	}
	if value, ok := values[8].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field recovery_token", values[8])
	} else if value.Valid {
		a.RecoveryToken = new(string)
		*a.RecoveryToken = value.String
	}
	if value, ok := values[9].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field otp_sent_at", values[9])
	} else if value.Valid {
		a.OtpSentAt = new(time.Time)
		*a.OtpSentAt = value.Time
	}
	if value, ok := values[10].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field otp", values[10])
	} else if value.Valid {
		a.Otp = new(string)
		*a.Otp = value.String
	}
	if value, ok := values[11].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email_change", values[11])
	} else if value.Valid {
		a.EmailChange = value.String
	}
	if value, ok := values[12].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field email_change_sent_at", values[12])
	} else if value.Valid {
		a.EmailChangeSentAt = new(time.Time)
		*a.EmailChangeSentAt = value.Time
	}
	if value, ok := values[13].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email_change_token", values[13])
	} else if value.Valid {
		a.EmailChangeToken = new(string)
		*a.EmailChangeToken = value.String
	}

	if value, ok := values[14].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field attributes", values[14])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &a.Attributes); err != nil {
			return fmt.Errorf("unmarshal field attributes: %v", err)
		}
	}

	if value, ok := values[15].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field sensitive_attributes", values[15])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &a.SensitiveAttributes); err != nil {
			return fmt.Errorf("unmarshal field sensitive_attributes: %v", err)
		}
	}
	if value, ok := values[16].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field attribute_bytes", values[16])
	} else if value != nil {
		a.AttributeBytes = *value
	}
	if value, ok := values[17].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[17])
	} else if value.Valid {
		a.CreatedAt = value.Time
	}
	if value, ok := values[18].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[18])
	} else if value.Valid {
		a.UpdatedAt = value.Time
	}
	if value, ok := values[19].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field last_signin_at", values[19])
	} else if value.Valid {
		a.LastSigninAt = new(time.Time)
		*a.LastSigninAt = value.Time
	}
	return nil
}

// Update returns a builder for updating this Account.
// Note that, you need to call Account.Unwrap() before calling this method, if this Account
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Account) Update() *AccountUpdateOne {
	return (&AccountClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (a *Account) Unwrap() *Account {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("models: Account is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Account) String() string {
	var builder strings.Builder
	builder.WriteString("Account(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", provider=")
	builder.WriteString(a.Provider)
	builder.WriteString(", email=")
	builder.WriteString(a.Email)
	builder.WriteString(", password=<sensitive>")
	builder.WriteString(", locked=")
	builder.WriteString(fmt.Sprintf("%v", a.Locked))
	builder.WriteString(", confirmed=")
	builder.WriteString(fmt.Sprintf("%v", a.Confirmed))
	if v := a.ConfirmationSentAt; v != nil {
		builder.WriteString(", confirmation_sent_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := a.ConfirmationToken; v != nil {
		builder.WriteString(", confirmation_token=")
		builder.WriteString(*v)
	}
	if v := a.RecoverySentAt; v != nil {
		builder.WriteString(", recovery_sent_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := a.RecoveryToken; v != nil {
		builder.WriteString(", recovery_token=")
		builder.WriteString(*v)
	}
	if v := a.OtpSentAt; v != nil {
		builder.WriteString(", otp_sent_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := a.Otp; v != nil {
		builder.WriteString(", otp=")
		builder.WriteString(*v)
	}
	builder.WriteString(", email_change=")
	builder.WriteString(a.EmailChange)
	if v := a.EmailChangeSentAt; v != nil {
		builder.WriteString(", email_change_sent_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := a.EmailChangeToken; v != nil {
		builder.WriteString(", email_change_token=")
		builder.WriteString(*v)
	}
	builder.WriteString(", attributes=")
	builder.WriteString(fmt.Sprintf("%v", a.Attributes))
	builder.WriteString(", sensitive_attributes=")
	builder.WriteString(fmt.Sprintf("%v", a.SensitiveAttributes))
	builder.WriteString(", attribute_bytes=")
	builder.WriteString(fmt.Sprintf("%v", a.AttributeBytes))
	builder.WriteString(", created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	if v := a.LastSigninAt; v != nil {
		builder.WriteString(", last_signin_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Accounts is a parsable slice of Account.
type Accounts []*Account

func (a Accounts) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
