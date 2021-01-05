// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/adnaan/authzen/models/account"
	"github.com/adnaan/authzen/models/workspace"
	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// Account is the model entity for the Account schema.
type Account struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// BillingID holds the value of the "billing_id" field.
	BillingID string `json:"billing_id,omitempty"`
	// Provider holds the value of the "provider" field.
	Provider string `json:"provider,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// APIKey holds the value of the "api_key" field.
	APIKey string `json:"-"`
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
	// Roles holds the value of the "roles" field.
	Roles []string `json:"roles,omitempty"`
	// Teams holds the value of the "teams" field.
	Teams map[string]string `json:"teams,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// LastSigninAt holds the value of the "last_signin_at" field.
	LastSigninAt *time.Time `json:"last_signin_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountQuery when eager-loading is set.
	Edges AccountEdges `json:"edges"`
}

// AccountEdges holds the relations/edges for other nodes in the graph.
type AccountEdges struct {
	// Workspace holds the value of the workspace edge.
	Workspace *Workspace
	// WorkspaceRoles holds the value of the workspace_roles edge.
	WorkspaceRoles []*WorkspaceRole
	// GroupRoles holds the value of the group_roles edge.
	GroupRoles []*GroupRole
	// AccountRoles holds the value of the account_roles edge.
	AccountRoles []*AccountRole
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// WorkspaceOrErr returns the Workspace value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountEdges) WorkspaceOrErr() (*Workspace, error) {
	if e.loadedTypes[0] {
		if e.Workspace == nil {
			// The edge workspace was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: workspace.Label}
		}
		return e.Workspace, nil
	}
	return nil, &NotLoadedError{edge: "workspace"}
}

// WorkspaceRolesOrErr returns the WorkspaceRoles value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) WorkspaceRolesOrErr() ([]*WorkspaceRole, error) {
	if e.loadedTypes[1] {
		return e.WorkspaceRoles, nil
	}
	return nil, &NotLoadedError{edge: "workspace_roles"}
}

// GroupRolesOrErr returns the GroupRoles value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) GroupRolesOrErr() ([]*GroupRole, error) {
	if e.loadedTypes[2] {
		return e.GroupRoles, nil
	}
	return nil, &NotLoadedError{edge: "group_roles"}
}

// AccountRolesOrErr returns the AccountRoles value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) AccountRolesOrErr() ([]*AccountRole, error) {
	if e.loadedTypes[3] {
		return e.AccountRoles, nil
	}
	return nil, &NotLoadedError{edge: "account_roles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Account) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&sql.NullString{}, // billing_id
		&sql.NullString{}, // provider
		&sql.NullString{}, // email
		&sql.NullString{}, // password
		&sql.NullString{}, // api_key
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
		&[]byte{},         // roles
		&[]byte{},         // teams
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
		return fmt.Errorf("unexpected type %T for field billing_id", values[0])
	} else if value.Valid {
		a.BillingID = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field provider", values[1])
	} else if value.Valid {
		a.Provider = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[2])
	} else if value.Valid {
		a.Email = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field password", values[3])
	} else if value.Valid {
		a.Password = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field api_key", values[4])
	} else if value.Valid {
		a.APIKey = value.String
	}
	if value, ok := values[5].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field confirmed", values[5])
	} else if value.Valid {
		a.Confirmed = value.Bool
	}
	if value, ok := values[6].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field confirmation_sent_at", values[6])
	} else if value.Valid {
		a.ConfirmationSentAt = new(time.Time)
		*a.ConfirmationSentAt = value.Time
	}
	if value, ok := values[7].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field confirmation_token", values[7])
	} else if value.Valid {
		a.ConfirmationToken = new(string)
		*a.ConfirmationToken = value.String
	}
	if value, ok := values[8].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field recovery_sent_at", values[8])
	} else if value.Valid {
		a.RecoverySentAt = new(time.Time)
		*a.RecoverySentAt = value.Time
	}
	if value, ok := values[9].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field recovery_token", values[9])
	} else if value.Valid {
		a.RecoveryToken = new(string)
		*a.RecoveryToken = value.String
	}
	if value, ok := values[10].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field otp_sent_at", values[10])
	} else if value.Valid {
		a.OtpSentAt = new(time.Time)
		*a.OtpSentAt = value.Time
	}
	if value, ok := values[11].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field otp", values[11])
	} else if value.Valid {
		a.Otp = new(string)
		*a.Otp = value.String
	}
	if value, ok := values[12].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email_change", values[12])
	} else if value.Valid {
		a.EmailChange = value.String
	}
	if value, ok := values[13].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field email_change_sent_at", values[13])
	} else if value.Valid {
		a.EmailChangeSentAt = new(time.Time)
		*a.EmailChangeSentAt = value.Time
	}
	if value, ok := values[14].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email_change_token", values[14])
	} else if value.Valid {
		a.EmailChangeToken = new(string)
		*a.EmailChangeToken = value.String
	}

	if value, ok := values[15].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field attributes", values[15])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &a.Attributes); err != nil {
			return fmt.Errorf("unmarshal field attributes: %v", err)
		}
	}

	if value, ok := values[16].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field roles", values[16])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &a.Roles); err != nil {
			return fmt.Errorf("unmarshal field roles: %v", err)
		}
	}

	if value, ok := values[17].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field teams", values[17])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &a.Teams); err != nil {
			return fmt.Errorf("unmarshal field teams: %v", err)
		}
	}
	if value, ok := values[18].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[18])
	} else if value.Valid {
		a.CreatedAt = value.Time
	}
	if value, ok := values[19].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[19])
	} else if value.Valid {
		a.UpdatedAt = value.Time
	}
	if value, ok := values[20].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field last_signin_at", values[20])
	} else if value.Valid {
		a.LastSigninAt = new(time.Time)
		*a.LastSigninAt = value.Time
	}
	return nil
}

// QueryWorkspace queries the workspace edge of the Account.
func (a *Account) QueryWorkspace() *WorkspaceQuery {
	return (&AccountClient{config: a.config}).QueryWorkspace(a)
}

// QueryWorkspaceRoles queries the workspace_roles edge of the Account.
func (a *Account) QueryWorkspaceRoles() *WorkspaceRoleQuery {
	return (&AccountClient{config: a.config}).QueryWorkspaceRoles(a)
}

// QueryGroupRoles queries the group_roles edge of the Account.
func (a *Account) QueryGroupRoles() *GroupRoleQuery {
	return (&AccountClient{config: a.config}).QueryGroupRoles(a)
}

// QueryAccountRoles queries the account_roles edge of the Account.
func (a *Account) QueryAccountRoles() *AccountRoleQuery {
	return (&AccountClient{config: a.config}).QueryAccountRoles(a)
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
	builder.WriteString(", billing_id=")
	builder.WriteString(a.BillingID)
	builder.WriteString(", provider=")
	builder.WriteString(a.Provider)
	builder.WriteString(", email=")
	builder.WriteString(a.Email)
	builder.WriteString(", password=<sensitive>")
	builder.WriteString(", api_key=<sensitive>")
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
	builder.WriteString(", roles=")
	builder.WriteString(fmt.Sprintf("%v", a.Roles))
	builder.WriteString(", teams=")
	builder.WriteString(fmt.Sprintf("%v", a.Teams))
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
