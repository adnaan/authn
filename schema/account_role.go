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

// AccountRole holds the schema definition for the AccountRole entity.
type AccountRole struct {
	ent.Schema
}

// Annotations of the AccountRole.
func (AccountRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "account_roles"},
	}
}

// Fields of the AccountRole.
func (AccountRole) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty(),
		field.UUID("account_id", uuid.UUID{}),
		field.JSON("attributes", map[string]interface{}{}).Optional(),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the AccountRole.
func (AccountRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("accounts", Account.Type).Ref("account_roles").Required().Unique(),
	}
}
