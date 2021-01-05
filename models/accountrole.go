// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/adnaan/authzen/models/account"
	"github.com/adnaan/authzen/models/accountrole"
	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// AccountRole is the model entity for the AccountRole schema.
type AccountRole struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// AccountID holds the value of the "account_id" field.
	AccountID uuid.UUID `json:"account_id,omitempty"`
	// Attributes holds the value of the "attributes" field.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountRoleQuery when eager-loading is set.
	Edges                 AccountRoleEdges `json:"edges"`
	account_account_roles *uuid.UUID
}

// AccountRoleEdges holds the relations/edges for other nodes in the graph.
type AccountRoleEdges struct {
	// Accounts holds the value of the accounts edge.
	Accounts *Account
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AccountsOrErr returns the Accounts value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountRoleEdges) AccountsOrErr() (*Account, error) {
	if e.loadedTypes[0] {
		if e.Accounts == nil {
			// The edge accounts was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: account.Label}
		}
		return e.Accounts, nil
	}
	return nil, &NotLoadedError{edge: "accounts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AccountRole) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&sql.NullString{}, // name
		&uuid.UUID{},      // account_id
		&[]byte{},         // attributes
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*AccountRole) fkValues() []interface{} {
	return []interface{}{
		&uuid.UUID{}, // account_account_roles
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AccountRole fields.
func (ar *AccountRole) assignValues(values ...interface{}) error {
	if m, n := len(values), len(accountrole.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		ar.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		ar.Name = value.String
	}
	if value, ok := values[1].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field account_id", values[1])
	} else if value != nil {
		ar.AccountID = *value
	}

	if value, ok := values[2].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field attributes", values[2])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &ar.Attributes); err != nil {
			return fmt.Errorf("unmarshal field attributes: %v", err)
		}
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[3])
	} else if value.Valid {
		ar.CreatedAt = value.Time
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[4])
	} else if value.Valid {
		ar.UpdatedAt = value.Time
	}
	values = values[5:]
	if len(values) == len(accountrole.ForeignKeys) {
		if value, ok := values[0].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field account_account_roles", values[0])
		} else if value != nil {
			ar.account_account_roles = value
		}
	}
	return nil
}

// QueryAccounts queries the accounts edge of the AccountRole.
func (ar *AccountRole) QueryAccounts() *AccountQuery {
	return (&AccountRoleClient{config: ar.config}).QueryAccounts(ar)
}

// Update returns a builder for updating this AccountRole.
// Note that, you need to call AccountRole.Unwrap() before calling this method, if this AccountRole
// was returned from a transaction, and the transaction was committed or rolled back.
func (ar *AccountRole) Update() *AccountRoleUpdateOne {
	return (&AccountRoleClient{config: ar.config}).UpdateOne(ar)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (ar *AccountRole) Unwrap() *AccountRole {
	tx, ok := ar.config.driver.(*txDriver)
	if !ok {
		panic("models: AccountRole is not a transactional entity")
	}
	ar.config.driver = tx.drv
	return ar
}

// String implements the fmt.Stringer.
func (ar *AccountRole) String() string {
	var builder strings.Builder
	builder.WriteString("AccountRole(")
	builder.WriteString(fmt.Sprintf("id=%v", ar.ID))
	builder.WriteString(", name=")
	builder.WriteString(ar.Name)
	builder.WriteString(", account_id=")
	builder.WriteString(fmt.Sprintf("%v", ar.AccountID))
	builder.WriteString(", attributes=")
	builder.WriteString(fmt.Sprintf("%v", ar.Attributes))
	builder.WriteString(", created_at=")
	builder.WriteString(ar.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(ar.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// AccountRoles is a parsable slice of AccountRole.
type AccountRoles []*AccountRole

func (ar AccountRoles) config(cfg config) {
	for _i := range ar {
		ar[_i].config = cfg
	}
}
