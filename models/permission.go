// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/adnaan/authzen/models/permission"
	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// Permission is the model entity for the Permission schema.
type Permission struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// RoleID holds the value of the "role_id" field.
	RoleID uuid.UUID `json:"role_id,omitempty"`
	// Action holds the value of the "action" field.
	Action string `json:"action,omitempty"`
	// Target holds the value of the "target" field.
	Target string `json:"target,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Permission) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&uuid.UUID{},      // role_id
		&sql.NullString{}, // action
		&sql.NullString{}, // target
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Permission fields.
func (pe *Permission) assignValues(values ...interface{}) error {
	if m, n := len(values), len(permission.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		pe.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field role_id", values[0])
	} else if value != nil {
		pe.RoleID = *value
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field action", values[1])
	} else if value.Valid {
		pe.Action = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field target", values[2])
	} else if value.Valid {
		pe.Target = value.String
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[3])
	} else if value.Valid {
		pe.CreatedAt = value.Time
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[4])
	} else if value.Valid {
		pe.UpdatedAt = value.Time
	}
	return nil
}

// Update returns a builder for updating this Permission.
// Note that, you need to call Permission.Unwrap() before calling this method, if this Permission
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Permission) Update() *PermissionUpdateOne {
	return (&PermissionClient{config: pe.config}).UpdateOne(pe)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pe *Permission) Unwrap() *Permission {
	tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("models: Permission is not a transactional entity")
	}
	pe.config.driver = tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Permission) String() string {
	var builder strings.Builder
	builder.WriteString("Permission(")
	builder.WriteString(fmt.Sprintf("id=%v", pe.ID))
	builder.WriteString(", role_id=")
	builder.WriteString(fmt.Sprintf("%v", pe.RoleID))
	builder.WriteString(", action=")
	builder.WriteString(pe.Action)
	builder.WriteString(", target=")
	builder.WriteString(pe.Target)
	builder.WriteString(", created_at=")
	builder.WriteString(pe.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(pe.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Permissions is a parsable slice of Permission.
type Permissions []*Permission

func (pe Permissions) config(cfg config) {
	for _i := range pe {
		pe[_i].config = cfg
	}
}
