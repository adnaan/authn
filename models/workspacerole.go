// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/adnaan/authzen/models/account"
	"github.com/adnaan/authzen/models/workspace"
	"github.com/adnaan/authzen/models/workspacerole"
	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// WorkspaceRole is the model entity for the WorkspaceRole schema.
type WorkspaceRole struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// WorkspaceID holds the value of the "workspace_id" field.
	WorkspaceID uuid.UUID `json:"workspace_id,omitempty"`
	// AccountID holds the value of the "account_id" field.
	AccountID uuid.UUID `json:"account_id,omitempty"`
	// Attributes holds the value of the "attributes" field.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkspaceRoleQuery when eager-loading is set.
	Edges                     WorkspaceRoleEdges `json:"edges"`
	account_workspace_roles   *uuid.UUID
	workspace_workspace_roles *uuid.UUID
}

// WorkspaceRoleEdges holds the relations/edges for other nodes in the graph.
type WorkspaceRoleEdges struct {
	// Workspaces holds the value of the workspaces edge.
	Workspaces *Workspace
	// Accounts holds the value of the accounts edge.
	Accounts *Account
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// WorkspacesOrErr returns the Workspaces value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkspaceRoleEdges) WorkspacesOrErr() (*Workspace, error) {
	if e.loadedTypes[0] {
		if e.Workspaces == nil {
			// The edge workspaces was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: workspace.Label}
		}
		return e.Workspaces, nil
	}
	return nil, &NotLoadedError{edge: "workspaces"}
}

// AccountsOrErr returns the Accounts value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkspaceRoleEdges) AccountsOrErr() (*Account, error) {
	if e.loadedTypes[1] {
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
func (*WorkspaceRole) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&sql.NullString{}, // name
		&uuid.UUID{},      // workspace_id
		&uuid.UUID{},      // account_id
		&[]byte{},         // attributes
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*WorkspaceRole) fkValues() []interface{} {
	return []interface{}{
		&uuid.UUID{}, // account_workspace_roles
		&uuid.UUID{}, // workspace_workspace_roles
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WorkspaceRole fields.
func (wr *WorkspaceRole) assignValues(values ...interface{}) error {
	if m, n := len(values), len(workspacerole.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		wr.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		wr.Name = value.String
	}
	if value, ok := values[1].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field workspace_id", values[1])
	} else if value != nil {
		wr.WorkspaceID = *value
	}
	if value, ok := values[2].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field account_id", values[2])
	} else if value != nil {
		wr.AccountID = *value
	}

	if value, ok := values[3].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field attributes", values[3])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &wr.Attributes); err != nil {
			return fmt.Errorf("unmarshal field attributes: %v", err)
		}
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[4])
	} else if value.Valid {
		wr.CreatedAt = value.Time
	}
	if value, ok := values[5].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[5])
	} else if value.Valid {
		wr.UpdatedAt = value.Time
	}
	values = values[6:]
	if len(values) == len(workspacerole.ForeignKeys) {
		if value, ok := values[0].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field account_workspace_roles", values[0])
		} else if value != nil {
			wr.account_workspace_roles = value
		}
		if value, ok := values[1].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field workspace_workspace_roles", values[1])
		} else if value != nil {
			wr.workspace_workspace_roles = value
		}
	}
	return nil
}

// QueryWorkspaces queries the workspaces edge of the WorkspaceRole.
func (wr *WorkspaceRole) QueryWorkspaces() *WorkspaceQuery {
	return (&WorkspaceRoleClient{config: wr.config}).QueryWorkspaces(wr)
}

// QueryAccounts queries the accounts edge of the WorkspaceRole.
func (wr *WorkspaceRole) QueryAccounts() *AccountQuery {
	return (&WorkspaceRoleClient{config: wr.config}).QueryAccounts(wr)
}

// Update returns a builder for updating this WorkspaceRole.
// Note that, you need to call WorkspaceRole.Unwrap() before calling this method, if this WorkspaceRole
// was returned from a transaction, and the transaction was committed or rolled back.
func (wr *WorkspaceRole) Update() *WorkspaceRoleUpdateOne {
	return (&WorkspaceRoleClient{config: wr.config}).UpdateOne(wr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (wr *WorkspaceRole) Unwrap() *WorkspaceRole {
	tx, ok := wr.config.driver.(*txDriver)
	if !ok {
		panic("models: WorkspaceRole is not a transactional entity")
	}
	wr.config.driver = tx.drv
	return wr
}

// String implements the fmt.Stringer.
func (wr *WorkspaceRole) String() string {
	var builder strings.Builder
	builder.WriteString("WorkspaceRole(")
	builder.WriteString(fmt.Sprintf("id=%v", wr.ID))
	builder.WriteString(", name=")
	builder.WriteString(wr.Name)
	builder.WriteString(", workspace_id=")
	builder.WriteString(fmt.Sprintf("%v", wr.WorkspaceID))
	builder.WriteString(", account_id=")
	builder.WriteString(fmt.Sprintf("%v", wr.AccountID))
	builder.WriteString(", attributes=")
	builder.WriteString(fmt.Sprintf("%v", wr.Attributes))
	builder.WriteString(", created_at=")
	builder.WriteString(wr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(wr.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// WorkspaceRoles is a parsable slice of WorkspaceRole.
type WorkspaceRoles []*WorkspaceRole

func (wr WorkspaceRoles) config(cfg config) {
	for _i := range wr {
		wr[_i].config = cfg
	}
}
