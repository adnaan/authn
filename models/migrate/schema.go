// Code generated (@generated) by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "provider", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "locked", Type: field.TypeBool, Default: false},
		{Name: "confirmed", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "confirmation_sent_at", Type: field.TypeTime, Nullable: true},
		{Name: "confirmation_token", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "recovery_sent_at", Type: field.TypeTime, Nullable: true},
		{Name: "recovery_token", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "otp_sent_at", Type: field.TypeTime, Nullable: true},
		{Name: "otp", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "email_change", Type: field.TypeString, Nullable: true},
		{Name: "email_change_sent_at", Type: field.TypeTime, Nullable: true},
		{Name: "email_change_token", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "attributes", Type: field.TypeJSON, Nullable: true},
		{Name: "sensitive_attributes", Type: field.TypeJSON, Nullable: true},
		{Name: "attribute_bytes", Type: field.TypeBytes, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "last_signin_at", Type: field.TypeTime, Nullable: true},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
	}
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "data", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "expires_at", Type: field.TypeTime, Nullable: true},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:       "sessions",
		Columns:    SessionsColumns,
		PrimaryKey: []*schema.Column{SessionsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		SessionsTable,
	}
)

func init() {
	AccountsTable.Annotation = &entsql.Annotation{
		Table: "accounts",
	}
	SessionsTable.Annotation = &entsql.Annotation{
		Table: "sessions",
	}
}
