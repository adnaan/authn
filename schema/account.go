package schema

import (
	"time"

	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/schema"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Annotations of the Account.
func (Account) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "accounts"},
	}
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("provider").NotEmpty(),
		field.String("email").Unique().NotEmpty(),
		field.String("password").NotEmpty().Sensitive().MinLen(8),
		field.Bool("locked").Default(false),

		field.Bool("confirmed").Default(false).Optional(),
		field.Time("confirmation_sent_at").Nillable().Optional(),
		field.String("confirmation_token").NotEmpty().Optional().Nillable().Unique(),

		field.Time("recovery_sent_at").Nillable().Optional(),
		field.String("recovery_token").NotEmpty().Optional().Nillable().Unique(),
		field.Time("otp_sent_at").Nillable().Optional(),
		field.String("otp").NotEmpty().Optional().Nillable().Unique(),

		field.String("email_change").NotEmpty().Optional(),
		field.Time("email_change_sent_at").Nillable().Optional(),
		field.String("email_change_token").NotEmpty().Optional().Nillable().Unique(),

		field.JSON("attributes", map[string]interface{}{}).Optional(),
		field.JSON("sensitive_attributes", map[string]string{}).Optional(),
		field.Bytes("attribute_bytes").Optional(),

		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("last_signin_at").Nillable().Optional(),
	}
}
