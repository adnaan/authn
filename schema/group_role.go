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

// GroupRole holds the schema definition for the GroupRole entity.
type GroupRole struct {
	ent.Schema
}

// Annotations of the GroupRole.
func (GroupRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "group_roles"},
	}
}

// Fields of the GroupRole.
func (GroupRole) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty(),
		field.UUID("group_id", uuid.UUID{}),
		field.UUID("account_id", uuid.UUID{}),
		field.JSON("attributes", map[string]interface{}{}).Optional(),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the GroupRole.
func (GroupRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("groups", Group.Type).Ref("group_roles").Required().Unique(),
		edge.From("accounts", Account.Type).Ref("group_roles").Required().Unique(),
	}
}
