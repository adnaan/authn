// Code generated (@generated) by entc, DO NOT EDIT.

package workspacerole

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the workspacerole type in the database.
	Label = "workspace_role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldWorkspaceID holds the string denoting the workspace_id field in the database.
	FieldWorkspaceID = "workspace_id"
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"

	// EdgeWorkspaces holds the string denoting the workspaces edge name in mutations.
	EdgeWorkspaces = "workspaces"
	// EdgeAccounts holds the string denoting the accounts edge name in mutations.
	EdgeAccounts = "accounts"

	// Table holds the table name of the workspacerole in the database.
	Table = "workspace_roles"
	// WorkspacesTable is the table the holds the workspaces relation/edge.
	WorkspacesTable = "workspace_roles"
	// WorkspacesInverseTable is the table name for the Workspace entity.
	// It exists in this package in order to avoid circular dependency with the "workspace" package.
	WorkspacesInverseTable = "workspaces"
	// WorkspacesColumn is the table column denoting the workspaces relation/edge.
	WorkspacesColumn = "workspace_workspace_roles"
	// AccountsTable is the table the holds the accounts relation/edge.
	AccountsTable = "workspace_roles"
	// AccountsInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AccountsInverseTable = "accounts"
	// AccountsColumn is the table column denoting the accounts relation/edge.
	AccountsColumn = "account_workspace_roles"
)

// Columns holds all SQL columns for workspacerole fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldWorkspaceID,
	FieldAccountID,
	FieldMetadata,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the WorkspaceRole type.
var ForeignKeys = []string{
	"account_workspace_roles",
	"workspace_workspace_roles",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the id field.
	DefaultID func() uuid.UUID
)
