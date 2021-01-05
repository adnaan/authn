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

// Workspace is the model entity for the Workspace schema.
type Workspace struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Plan holds the value of the "plan" field.
	Plan string `json:"plan,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkspaceQuery when eager-loading is set.
	Edges             WorkspaceEdges `json:"edges"`
	account_workspace *uuid.UUID
}

// WorkspaceEdges holds the relations/edges for other nodes in the graph.
type WorkspaceEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Account
	// WorkspaceRoles holds the value of the workspace_roles edge.
	WorkspaceRoles []*WorkspaceRole
	// Groups holds the value of the groups edge.
	Groups []*Group
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkspaceEdges) OwnerOrErr() (*Account, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: account.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// WorkspaceRolesOrErr returns the WorkspaceRoles value or an error if the edge
// was not loaded in eager-loading.
func (e WorkspaceEdges) WorkspaceRolesOrErr() ([]*WorkspaceRole, error) {
	if e.loadedTypes[1] {
		return e.WorkspaceRoles, nil
	}
	return nil, &NotLoadedError{edge: "workspace_roles"}
}

// GroupsOrErr returns the Groups value or an error if the edge
// was not loaded in eager-loading.
func (e WorkspaceEdges) GroupsOrErr() ([]*Group, error) {
	if e.loadedTypes[2] {
		return e.Groups, nil
	}
	return nil, &NotLoadedError{edge: "groups"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Workspace) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&sql.NullString{}, // name
		&sql.NullString{}, // plan
		&sql.NullString{}, // description
		&[]byte{},         // metadata
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Workspace) fkValues() []interface{} {
	return []interface{}{
		&uuid.UUID{}, // account_workspace
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Workspace fields.
func (w *Workspace) assignValues(values ...interface{}) error {
	if m, n := len(values), len(workspace.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		w.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		w.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field plan", values[1])
	} else if value.Valid {
		w.Plan = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field description", values[2])
	} else if value.Valid {
		w.Description = value.String
	}

	if value, ok := values[3].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field metadata", values[3])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &w.Metadata); err != nil {
			return fmt.Errorf("unmarshal field metadata: %v", err)
		}
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[4])
	} else if value.Valid {
		w.CreatedAt = value.Time
	}
	if value, ok := values[5].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[5])
	} else if value.Valid {
		w.UpdatedAt = value.Time
	}
	values = values[6:]
	if len(values) == len(workspace.ForeignKeys) {
		if value, ok := values[0].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field account_workspace", values[0])
		} else if value != nil {
			w.account_workspace = value
		}
	}
	return nil
}

// QueryOwner queries the owner edge of the Workspace.
func (w *Workspace) QueryOwner() *AccountQuery {
	return (&WorkspaceClient{config: w.config}).QueryOwner(w)
}

// QueryWorkspaceRoles queries the workspace_roles edge of the Workspace.
func (w *Workspace) QueryWorkspaceRoles() *WorkspaceRoleQuery {
	return (&WorkspaceClient{config: w.config}).QueryWorkspaceRoles(w)
}

// QueryGroups queries the groups edge of the Workspace.
func (w *Workspace) QueryGroups() *GroupQuery {
	return (&WorkspaceClient{config: w.config}).QueryGroups(w)
}

// Update returns a builder for updating this Workspace.
// Note that, you need to call Workspace.Unwrap() before calling this method, if this Workspace
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Workspace) Update() *WorkspaceUpdateOne {
	return (&WorkspaceClient{config: w.config}).UpdateOne(w)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (w *Workspace) Unwrap() *Workspace {
	tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("models: Workspace is not a transactional entity")
	}
	w.config.driver = tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Workspace) String() string {
	var builder strings.Builder
	builder.WriteString("Workspace(")
	builder.WriteString(fmt.Sprintf("id=%v", w.ID))
	builder.WriteString(", name=")
	builder.WriteString(w.Name)
	builder.WriteString(", plan=")
	builder.WriteString(w.Plan)
	builder.WriteString(", description=")
	builder.WriteString(w.Description)
	builder.WriteString(", metadata=")
	builder.WriteString(fmt.Sprintf("%v", w.Metadata))
	builder.WriteString(", created_at=")
	builder.WriteString(w.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(w.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Workspaces is a parsable slice of Workspace.
type Workspaces []*Workspace

func (w Workspaces) config(cfg config) {
	for _i := range w {
		w[_i].config = cfg
	}
}
