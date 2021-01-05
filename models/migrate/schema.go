// Code generated (@generated) by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "billing_id", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "provider", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "api_key", Type: field.TypeString, Nullable: true},
		{Name: "confirmed", Type: field.TypeBool, Nullable: true},
		{Name: "confirmation_sent_at", Type: field.TypeTime, Nullable: true},
		{Name: "confirmation_token", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "recovery_sent_at", Type: field.TypeTime, Nullable: true},
		{Name: "recovery_token", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "otp_sent_at", Type: field.TypeTime, Nullable: true},
		{Name: "otp", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "email_change", Type: field.TypeString, Nullable: true},
		{Name: "email_change_sent_at", Type: field.TypeTime, Nullable: true},
		{Name: "email_change_token", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "metadata", Type: field.TypeJSON},
		{Name: "roles", Type: field.TypeJSON, Nullable: true},
		{Name: "teams", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "last_signin_at", Type: field.TypeTime, Nullable: true},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:        "accounts",
		Columns:     AccountsColumns,
		PrimaryKey:  []*schema.Column{AccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Annotation:  &entsql.Annotation{Table: "accounts"},
	}
	// AccountRolesColumns holds the columns for the "account_roles" table.
	AccountRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "account_id", Type: field.TypeUUID},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "account_account_roles", Type: field.TypeUUID, Nullable: true},
	}
	// AccountRolesTable holds the schema information for the "account_roles" table.
	AccountRolesTable = &schema.Table{
		Name:       "account_roles",
		Columns:    AccountRolesColumns,
		PrimaryKey: []*schema.Column{AccountRolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "account_roles_accounts_account_roles",
				Columns: []*schema.Column{AccountRolesColumns[6]},

				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Annotation: &entsql.Annotation{Table: "account_roles"},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "workspace_groups", Type: field.TypeUUID, Nullable: true},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "groups_workspaces_groups",
				Columns: []*schema.Column{GroupsColumns[6]},

				RefColumns: []*schema.Column{WorkspacesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Annotation: &entsql.Annotation{Table: "groups"},
	}
	// GroupRolesColumns holds the columns for the "group_roles" table.
	GroupRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "group_id", Type: field.TypeUUID},
		{Name: "account_id", Type: field.TypeUUID},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "account_group_roles", Type: field.TypeUUID, Nullable: true},
		{Name: "group_group_roles", Type: field.TypeUUID, Nullable: true},
	}
	// GroupRolesTable holds the schema information for the "group_roles" table.
	GroupRolesTable = &schema.Table{
		Name:       "group_roles",
		Columns:    GroupRolesColumns,
		PrimaryKey: []*schema.Column{GroupRolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "group_roles_accounts_group_roles",
				Columns: []*schema.Column{GroupRolesColumns[7]},

				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "group_roles_groups_group_roles",
				Columns: []*schema.Column{GroupRolesColumns[8]},

				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Annotation: &entsql.Annotation{Table: "group_roles"},
	}
	// PermissionsColumns holds the columns for the "permissions" table.
	PermissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "role_id", Type: field.TypeUUID},
		{Name: "action", Type: field.TypeString},
		{Name: "target", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// PermissionsTable holds the schema information for the "permissions" table.
	PermissionsTable = &schema.Table{
		Name:        "permissions",
		Columns:     PermissionsColumns,
		PrimaryKey:  []*schema.Column{PermissionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Annotation:  &entsql.Annotation{Table: "permissions"},
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
		Name:        "sessions",
		Columns:     SessionsColumns,
		PrimaryKey:  []*schema.Column{SessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Annotation:  &entsql.Annotation{Table: "sessions"},
	}
	// WorkspacesColumns holds the columns for the "workspaces" table.
	WorkspacesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "plan", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "account_workspace", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// WorkspacesTable holds the schema information for the "workspaces" table.
	WorkspacesTable = &schema.Table{
		Name:       "workspaces",
		Columns:    WorkspacesColumns,
		PrimaryKey: []*schema.Column{WorkspacesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "workspaces_accounts_workspace",
				Columns: []*schema.Column{WorkspacesColumns[7]},

				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Annotation: &entsql.Annotation{Table: "workspaces"},
	}
	// WorkspaceInvitationsColumns holds the columns for the "workspace_invitations" table.
	WorkspaceInvitationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "workspace_id", Type: field.TypeUUID},
		{Name: "role", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
	}
	// WorkspaceInvitationsTable holds the schema information for the "workspace_invitations" table.
	WorkspaceInvitationsTable = &schema.Table{
		Name:        "workspace_invitations",
		Columns:     WorkspaceInvitationsColumns,
		PrimaryKey:  []*schema.Column{WorkspaceInvitationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Annotation:  &entsql.Annotation{Table: "workspace_invitations"},
	}
	// WorkspaceRolesColumns holds the columns for the "workspace_roles" table.
	WorkspaceRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "workspace_id", Type: field.TypeUUID},
		{Name: "account_id", Type: field.TypeUUID},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "account_workspace_roles", Type: field.TypeUUID, Nullable: true},
		{Name: "workspace_workspace_roles", Type: field.TypeUUID, Nullable: true},
	}
	// WorkspaceRolesTable holds the schema information for the "workspace_roles" table.
	WorkspaceRolesTable = &schema.Table{
		Name:       "workspace_roles",
		Columns:    WorkspaceRolesColumns,
		PrimaryKey: []*schema.Column{WorkspaceRolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "workspace_roles_accounts_workspace_roles",
				Columns: []*schema.Column{WorkspaceRolesColumns[7]},

				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "workspace_roles_workspaces_workspace_roles",
				Columns: []*schema.Column{WorkspaceRolesColumns[8]},

				RefColumns: []*schema.Column{WorkspacesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Annotation: &entsql.Annotation{Table: "workspace_roles"},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		AccountRolesTable,
		GroupsTable,
		GroupRolesTable,
		PermissionsTable,
		SessionsTable,
		WorkspacesTable,
		WorkspaceInvitationsTable,
		WorkspaceRolesTable,
	}
)

func init() {
	AccountRolesTable.ForeignKeys[0].RefTable = AccountsTable
	GroupsTable.ForeignKeys[0].RefTable = WorkspacesTable
	GroupRolesTable.ForeignKeys[0].RefTable = AccountsTable
	GroupRolesTable.ForeignKeys[1].RefTable = GroupsTable
	WorkspacesTable.ForeignKeys[0].RefTable = AccountsTable
	WorkspaceRolesTable.ForeignKeys[0].RefTable = AccountsTable
	WorkspaceRolesTable.ForeignKeys[1].RefTable = WorkspacesTable
}
