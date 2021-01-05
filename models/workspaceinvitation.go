// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/adnaan/authzen/models/workspaceinvitation"
	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// WorkspaceInvitation is the model entity for the WorkspaceInvitation schema.
type WorkspaceInvitation struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// WorkspaceID holds the value of the "workspace_id" field.
	WorkspaceID uuid.UUID `json:"workspace_id,omitempty"`
	// Role holds the value of the "role" field.
	Role string `json:"role,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WorkspaceInvitation) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&uuid.UUID{},      // workspace_id
		&sql.NullString{}, // role
		&sql.NullString{}, // email
		&sql.NullTime{},   // created_at
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WorkspaceInvitation fields.
func (wi *WorkspaceInvitation) assignValues(values ...interface{}) error {
	if m, n := len(values), len(workspaceinvitation.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		wi.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field workspace_id", values[0])
	} else if value != nil {
		wi.WorkspaceID = *value
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field role", values[1])
	} else if value.Valid {
		wi.Role = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[2])
	} else if value.Valid {
		wi.Email = value.String
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[3])
	} else if value.Valid {
		wi.CreatedAt = value.Time
	}
	return nil
}

// Update returns a builder for updating this WorkspaceInvitation.
// Note that, you need to call WorkspaceInvitation.Unwrap() before calling this method, if this WorkspaceInvitation
// was returned from a transaction, and the transaction was committed or rolled back.
func (wi *WorkspaceInvitation) Update() *WorkspaceInvitationUpdateOne {
	return (&WorkspaceInvitationClient{config: wi.config}).UpdateOne(wi)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (wi *WorkspaceInvitation) Unwrap() *WorkspaceInvitation {
	tx, ok := wi.config.driver.(*txDriver)
	if !ok {
		panic("models: WorkspaceInvitation is not a transactional entity")
	}
	wi.config.driver = tx.drv
	return wi
}

// String implements the fmt.Stringer.
func (wi *WorkspaceInvitation) String() string {
	var builder strings.Builder
	builder.WriteString("WorkspaceInvitation(")
	builder.WriteString(fmt.Sprintf("id=%v", wi.ID))
	builder.WriteString(", workspace_id=")
	builder.WriteString(fmt.Sprintf("%v", wi.WorkspaceID))
	builder.WriteString(", role=")
	builder.WriteString(wi.Role)
	builder.WriteString(", email=")
	builder.WriteString(wi.Email)
	builder.WriteString(", created_at=")
	builder.WriteString(wi.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// WorkspaceInvitations is a parsable slice of WorkspaceInvitation.
type WorkspaceInvitations []*WorkspaceInvitation

func (wi WorkspaceInvitations) config(cfg config) {
	for _i := range wi {
		wi[_i].config = cfg
	}
}
