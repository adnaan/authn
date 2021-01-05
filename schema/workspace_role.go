package schema

import (
	"time"

	"github.com/facebook/ent/schema/edge"

	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/schema"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// WorkspaceRole holds the schema definition for the WorkspaceRole entity.
type WorkspaceRole struct {
	ent.Schema
}

// Annotations of the WorkspaceRole.
func (WorkspaceRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "workspace_roles"},
	}
}

// Fields of the WorkspaceRole.
func (WorkspaceRole) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty(),
		field.UUID("workspace_id", uuid.UUID{}),
		field.UUID("account_id", uuid.UUID{}),
		field.JSON("metadata", map[string]interface{}{}).Optional(),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the WorkspaceRole.
func (WorkspaceRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("workspaces", Workspace.Type).Ref("workspace_roles").Required().Unique(),
		edge.From("accounts", Account.Type).Ref("workspace_roles").Required().Unique(),
	}
}
