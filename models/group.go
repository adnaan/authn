// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/adnaan/authzen/models/group"
	"github.com/adnaan/authzen/models/workspace"
	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// Group is the model entity for the Group schema.
type Group struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupQuery when eager-loading is set.
	Edges            GroupEdges `json:"edges"`
	workspace_groups *uuid.UUID
}

// GroupEdges holds the relations/edges for other nodes in the graph.
type GroupEdges struct {
	// GroupRoles holds the value of the group_roles edge.
	GroupRoles []*GroupRole
	// Workspaces holds the value of the workspaces edge.
	Workspaces *Workspace
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// GroupRolesOrErr returns the GroupRoles value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) GroupRolesOrErr() ([]*GroupRole, error) {
	if e.loadedTypes[0] {
		return e.GroupRoles, nil
	}
	return nil, &NotLoadedError{edge: "group_roles"}
}

// WorkspacesOrErr returns the Workspaces value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupEdges) WorkspacesOrErr() (*Workspace, error) {
	if e.loadedTypes[1] {
		if e.Workspaces == nil {
			// The edge workspaces was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: workspace.Label}
		}
		return e.Workspaces, nil
	}
	return nil, &NotLoadedError{edge: "workspaces"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Group) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&sql.NullString{}, // name
		&sql.NullString{}, // description
		&[]byte{},         // metadata
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Group) fkValues() []interface{} {
	return []interface{}{
		&uuid.UUID{}, // workspace_groups
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Group fields.
func (gr *Group) assignValues(values ...interface{}) error {
	if m, n := len(values), len(group.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		gr.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		gr.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field description", values[1])
	} else if value.Valid {
		gr.Description = value.String
	}

	if value, ok := values[2].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field metadata", values[2])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &gr.Metadata); err != nil {
			return fmt.Errorf("unmarshal field metadata: %v", err)
		}
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[3])
	} else if value.Valid {
		gr.CreatedAt = value.Time
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[4])
	} else if value.Valid {
		gr.UpdatedAt = value.Time
	}
	values = values[5:]
	if len(values) == len(group.ForeignKeys) {
		if value, ok := values[0].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field workspace_groups", values[0])
		} else if value != nil {
			gr.workspace_groups = value
		}
	}
	return nil
}

// QueryGroupRoles queries the group_roles edge of the Group.
func (gr *Group) QueryGroupRoles() *GroupRoleQuery {
	return (&GroupClient{config: gr.config}).QueryGroupRoles(gr)
}

// QueryWorkspaces queries the workspaces edge of the Group.
func (gr *Group) QueryWorkspaces() *WorkspaceQuery {
	return (&GroupClient{config: gr.config}).QueryWorkspaces(gr)
}

// Update returns a builder for updating this Group.
// Note that, you need to call Group.Unwrap() before calling this method, if this Group
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *Group) Update() *GroupUpdateOne {
	return (&GroupClient{config: gr.config}).UpdateOne(gr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (gr *Group) Unwrap() *Group {
	tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("models: Group is not a transactional entity")
	}
	gr.config.driver = tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *Group) String() string {
	var builder strings.Builder
	builder.WriteString("Group(")
	builder.WriteString(fmt.Sprintf("id=%v", gr.ID))
	builder.WriteString(", name=")
	builder.WriteString(gr.Name)
	builder.WriteString(", description=")
	builder.WriteString(gr.Description)
	builder.WriteString(", metadata=")
	builder.WriteString(fmt.Sprintf("%v", gr.Metadata))
	builder.WriteString(", created_at=")
	builder.WriteString(gr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(gr.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Groups is a parsable slice of Group.
type Groups []*Group

func (gr Groups) config(cfg config) {
	for _i := range gr {
		gr[_i].config = cfg
	}
}
