package schema

import (
	"time"

	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/schema"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// WorkspaceInvitation holds the schema definition for the WorkspaceInvitation entity.
type WorkspaceInvitation struct {
	ent.Schema
}

// Annotations of the WorkspaceInvitation.
func (WorkspaceInvitation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "workspace_invitations"},
	}
}

// Fields of the WorkspaceInvitation.
func (WorkspaceInvitation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("workspace_id", uuid.UUID{}),
		field.String("role").NotEmpty(),
		field.String("email").NotEmpty(),
		field.Time("created_at").Immutable().Default(time.Now),
	}
}

// Edges of the WorkspaceInvitation.
func (WorkspaceInvitation) Edges() []ent.Edge {
	return []ent.Edge{}
}
