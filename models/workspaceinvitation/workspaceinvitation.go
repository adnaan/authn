// Code generated (@generated) by entc, DO NOT EDIT.

package workspaceinvitation

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the workspaceinvitation type in the database.
	Label = "workspace_invitation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldWorkspaceID holds the string denoting the workspace_id field in the database.
	FieldWorkspaceID = "workspace_id"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"

	// Table holds the table name of the workspaceinvitation in the database.
	Table = "workspace_invitations"
)

// Columns holds all SQL columns for workspaceinvitation fields.
var Columns = []string{
	FieldID,
	FieldWorkspaceID,
	FieldRole,
	FieldEmail,
	FieldCreatedAt,
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
	// RoleValidator is a validator for the "role" field. It is called by the builders before save.
	RoleValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the id field.
	DefaultID func() uuid.UUID
)
