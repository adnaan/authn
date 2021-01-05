// Code generated (@generated) by entc, DO NOT EDIT.

package permission

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the permission type in the database.
	Label = "permission"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRoleID holds the string denoting the role_id field in the database.
	FieldRoleID = "role_id"
	// FieldAction holds the string denoting the action field in the database.
	FieldAction = "action"
	// FieldTarget holds the string denoting the target field in the database.
	FieldTarget = "target"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"

	// Table holds the table name of the permission in the database.
	Table = "permissions"
)

// Columns holds all SQL columns for permission fields.
var Columns = []string{
	FieldID,
	FieldRoleID,
	FieldAction,
	FieldTarget,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// ActionValidator is a validator for the "action" field. It is called by the builders before save.
	ActionValidator func(string) error
	// TargetValidator is a validator for the "target" field. It is called by the builders before save.
	TargetValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the id field.
	DefaultID func() uuid.UUID
)
