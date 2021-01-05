// Code generated (@generated) by entc, DO NOT EDIT.

package workspace

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the workspace type in the database.
	Label = "workspace"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPlan holds the string denoting the plan field in the database.
	FieldPlan = "plan"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldAttributes holds the string denoting the attributes field in the database.
	FieldAttributes = "attributes"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeWorkspaceRoles holds the string denoting the workspace_roles edge name in mutations.
	EdgeWorkspaceRoles = "workspace_roles"
	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"

	// Table holds the table name of the workspace in the database.
	Table = "workspaces"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "workspaces"
	// OwnerInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	OwnerInverseTable = "accounts"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "account_workspace"
	// WorkspaceRolesTable is the table the holds the workspace_roles relation/edge.
	WorkspaceRolesTable = "workspace_roles"
	// WorkspaceRolesInverseTable is the table name for the WorkspaceRole entity.
	// It exists in this package in order to avoid circular dependency with the "workspacerole" package.
	WorkspaceRolesInverseTable = "workspace_roles"
	// WorkspaceRolesColumn is the table column denoting the workspace_roles relation/edge.
	WorkspaceRolesColumn = "workspace_workspace_roles"
	// GroupsTable is the table the holds the groups relation/edge.
	GroupsTable = "groups"
	// GroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupsInverseTable = "groups"
	// GroupsColumn is the table column denoting the groups relation/edge.
	GroupsColumn = "workspace_groups"
)

// Columns holds all SQL columns for workspace fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPlan,
	FieldDescription,
	FieldAttributes,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Workspace type.
var ForeignKeys = []string{
	"account_workspace",
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
	// PlanValidator is a validator for the "plan" field. It is called by the builders before save.
	PlanValidator func(string) error
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
