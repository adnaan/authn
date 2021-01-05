// Code generated (@generated) by entc, DO NOT EDIT.

package group

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldAttributes holds the string denoting the attributes field in the database.
	FieldAttributes = "attributes"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"

	// EdgeGroupRoles holds the string denoting the group_roles edge name in mutations.
	EdgeGroupRoles = "group_roles"
	// EdgeWorkspaces holds the string denoting the workspaces edge name in mutations.
	EdgeWorkspaces = "workspaces"

	// Table holds the table name of the group in the database.
	Table = "groups"
	// GroupRolesTable is the table the holds the group_roles relation/edge.
	GroupRolesTable = "group_roles"
	// GroupRolesInverseTable is the table name for the GroupRole entity.
	// It exists in this package in order to avoid circular dependency with the "grouprole" package.
	GroupRolesInverseTable = "group_roles"
	// GroupRolesColumn is the table column denoting the group_roles relation/edge.
	GroupRolesColumn = "group_group_roles"
	// WorkspacesTable is the table the holds the workspaces relation/edge.
	WorkspacesTable = "groups"
	// WorkspacesInverseTable is the table name for the Workspace entity.
	// It exists in this package in order to avoid circular dependency with the "workspace" package.
	WorkspacesInverseTable = "workspaces"
	// WorkspacesColumn is the table column denoting the workspaces relation/edge.
	WorkspacesColumn = "workspace_groups"
)

// Columns holds all SQL columns for group fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldAttributes,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Group type.
var ForeignKeys = []string{
	"workspace_groups",
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
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the id field.
	DefaultID func() uuid.UUID
)
