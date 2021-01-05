// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"fmt"
	"log"

	"github.com/adnaan/authzen/models/migrate"
	"github.com/google/uuid"

	"github.com/adnaan/authzen/models/account"
	"github.com/adnaan/authzen/models/accountrole"
	"github.com/adnaan/authzen/models/group"
	"github.com/adnaan/authzen/models/grouprole"
	"github.com/adnaan/authzen/models/permission"
	"github.com/adnaan/authzen/models/session"
	"github.com/adnaan/authzen/models/workspace"
	"github.com/adnaan/authzen/models/workspaceinvitation"
	"github.com/adnaan/authzen/models/workspacerole"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Account is the client for interacting with the Account builders.
	Account *AccountClient
	// AccountRole is the client for interacting with the AccountRole builders.
	AccountRole *AccountRoleClient
	// Group is the client for interacting with the Group builders.
	Group *GroupClient
	// GroupRole is the client for interacting with the GroupRole builders.
	GroupRole *GroupRoleClient
	// Permission is the client for interacting with the Permission builders.
	Permission *PermissionClient
	// Session is the client for interacting with the Session builders.
	Session *SessionClient
	// Workspace is the client for interacting with the Workspace builders.
	Workspace *WorkspaceClient
	// WorkspaceInvitation is the client for interacting with the WorkspaceInvitation builders.
	WorkspaceInvitation *WorkspaceInvitationClient
	// WorkspaceRole is the client for interacting with the WorkspaceRole builders.
	WorkspaceRole *WorkspaceRoleClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Account = NewAccountClient(c.config)
	c.AccountRole = NewAccountRoleClient(c.config)
	c.Group = NewGroupClient(c.config)
	c.GroupRole = NewGroupRoleClient(c.config)
	c.Permission = NewPermissionClient(c.config)
	c.Session = NewSessionClient(c.config)
	c.Workspace = NewWorkspaceClient(c.config)
	c.WorkspaceInvitation = NewWorkspaceInvitationClient(c.config)
	c.WorkspaceRole = NewWorkspaceRoleClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("models: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("models: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:                 ctx,
		config:              cfg,
		Account:             NewAccountClient(cfg),
		AccountRole:         NewAccountRoleClient(cfg),
		Group:               NewGroupClient(cfg),
		GroupRole:           NewGroupRoleClient(cfg),
		Permission:          NewPermissionClient(cfg),
		Session:             NewSessionClient(cfg),
		Workspace:           NewWorkspaceClient(cfg),
		WorkspaceInvitation: NewWorkspaceInvitationClient(cfg),
		WorkspaceRole:       NewWorkspaceRoleClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:              cfg,
		Account:             NewAccountClient(cfg),
		AccountRole:         NewAccountRoleClient(cfg),
		Group:               NewGroupClient(cfg),
		GroupRole:           NewGroupRoleClient(cfg),
		Permission:          NewPermissionClient(cfg),
		Session:             NewSessionClient(cfg),
		Workspace:           NewWorkspaceClient(cfg),
		WorkspaceInvitation: NewWorkspaceInvitationClient(cfg),
		WorkspaceRole:       NewWorkspaceRoleClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Account.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Account.Use(hooks...)
	c.AccountRole.Use(hooks...)
	c.Group.Use(hooks...)
	c.GroupRole.Use(hooks...)
	c.Permission.Use(hooks...)
	c.Session.Use(hooks...)
	c.Workspace.Use(hooks...)
	c.WorkspaceInvitation.Use(hooks...)
	c.WorkspaceRole.Use(hooks...)
}

// AccountClient is a client for the Account schema.
type AccountClient struct {
	config
}

// NewAccountClient returns a client for the Account from the given config.
func NewAccountClient(c config) *AccountClient {
	return &AccountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `account.Hooks(f(g(h())))`.
func (c *AccountClient) Use(hooks ...Hook) {
	c.hooks.Account = append(c.hooks.Account, hooks...)
}

// Create returns a create builder for Account.
func (c *AccountClient) Create() *AccountCreate {
	mutation := newAccountMutation(c.config, OpCreate)
	return &AccountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Account entities.
func (c *AccountClient) CreateBulk(builders ...*AccountCreate) *AccountCreateBulk {
	return &AccountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Account.
func (c *AccountClient) Update() *AccountUpdate {
	mutation := newAccountMutation(c.config, OpUpdate)
	return &AccountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountClient) UpdateOne(a *Account) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccount(a))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountClient) UpdateOneID(id uuid.UUID) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccountID(id))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Account.
func (c *AccountClient) Delete() *AccountDelete {
	mutation := newAccountMutation(c.config, OpDelete)
	return &AccountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AccountClient) DeleteOne(a *Account) *AccountDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AccountClient) DeleteOneID(id uuid.UUID) *AccountDeleteOne {
	builder := c.Delete().Where(account.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AccountDeleteOne{builder}
}

// Query returns a query builder for Account.
func (c *AccountClient) Query() *AccountQuery {
	return &AccountQuery{config: c.config}
}

// Get returns a Account entity by its id.
func (c *AccountClient) Get(ctx context.Context, id uuid.UUID) (*Account, error) {
	return c.Query().Where(account.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountClient) GetX(ctx context.Context, id uuid.UUID) *Account {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryWorkspace queries the workspace edge of a Account.
func (c *AccountClient) QueryWorkspace(a *Account) *WorkspaceQuery {
	query := &WorkspaceQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, id),
			sqlgraph.To(workspace.Table, workspace.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, account.WorkspaceTable, account.WorkspaceColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryWorkspaceRoles queries the workspace_roles edge of a Account.
func (c *AccountClient) QueryWorkspaceRoles(a *Account) *WorkspaceRoleQuery {
	query := &WorkspaceRoleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, id),
			sqlgraph.To(workspacerole.Table, workspacerole.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, account.WorkspaceRolesTable, account.WorkspaceRolesColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGroupRoles queries the group_roles edge of a Account.
func (c *AccountClient) QueryGroupRoles(a *Account) *GroupRoleQuery {
	query := &GroupRoleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, id),
			sqlgraph.To(grouprole.Table, grouprole.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, account.GroupRolesTable, account.GroupRolesColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAccountRoles queries the account_roles edge of a Account.
func (c *AccountClient) QueryAccountRoles(a *Account) *AccountRoleQuery {
	query := &AccountRoleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, id),
			sqlgraph.To(accountrole.Table, accountrole.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, account.AccountRolesTable, account.AccountRolesColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AccountClient) Hooks() []Hook {
	return c.hooks.Account
}

// AccountRoleClient is a client for the AccountRole schema.
type AccountRoleClient struct {
	config
}

// NewAccountRoleClient returns a client for the AccountRole from the given config.
func NewAccountRoleClient(c config) *AccountRoleClient {
	return &AccountRoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `accountrole.Hooks(f(g(h())))`.
func (c *AccountRoleClient) Use(hooks ...Hook) {
	c.hooks.AccountRole = append(c.hooks.AccountRole, hooks...)
}

// Create returns a create builder for AccountRole.
func (c *AccountRoleClient) Create() *AccountRoleCreate {
	mutation := newAccountRoleMutation(c.config, OpCreate)
	return &AccountRoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AccountRole entities.
func (c *AccountRoleClient) CreateBulk(builders ...*AccountRoleCreate) *AccountRoleCreateBulk {
	return &AccountRoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AccountRole.
func (c *AccountRoleClient) Update() *AccountRoleUpdate {
	mutation := newAccountRoleMutation(c.config, OpUpdate)
	return &AccountRoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountRoleClient) UpdateOne(ar *AccountRole) *AccountRoleUpdateOne {
	mutation := newAccountRoleMutation(c.config, OpUpdateOne, withAccountRole(ar))
	return &AccountRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountRoleClient) UpdateOneID(id uuid.UUID) *AccountRoleUpdateOne {
	mutation := newAccountRoleMutation(c.config, OpUpdateOne, withAccountRoleID(id))
	return &AccountRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AccountRole.
func (c *AccountRoleClient) Delete() *AccountRoleDelete {
	mutation := newAccountRoleMutation(c.config, OpDelete)
	return &AccountRoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AccountRoleClient) DeleteOne(ar *AccountRole) *AccountRoleDeleteOne {
	return c.DeleteOneID(ar.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AccountRoleClient) DeleteOneID(id uuid.UUID) *AccountRoleDeleteOne {
	builder := c.Delete().Where(accountrole.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AccountRoleDeleteOne{builder}
}

// Query returns a query builder for AccountRole.
func (c *AccountRoleClient) Query() *AccountRoleQuery {
	return &AccountRoleQuery{config: c.config}
}

// Get returns a AccountRole entity by its id.
func (c *AccountRoleClient) Get(ctx context.Context, id uuid.UUID) (*AccountRole, error) {
	return c.Query().Where(accountrole.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountRoleClient) GetX(ctx context.Context, id uuid.UUID) *AccountRole {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAccounts queries the accounts edge of a AccountRole.
func (c *AccountRoleClient) QueryAccounts(ar *AccountRole) *AccountQuery {
	query := &AccountQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ar.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(accountrole.Table, accountrole.FieldID, id),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, accountrole.AccountsTable, accountrole.AccountsColumn),
		)
		fromV = sqlgraph.Neighbors(ar.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AccountRoleClient) Hooks() []Hook {
	return c.hooks.AccountRole
}

// GroupClient is a client for the Group schema.
type GroupClient struct {
	config
}

// NewGroupClient returns a client for the Group from the given config.
func NewGroupClient(c config) *GroupClient {
	return &GroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `group.Hooks(f(g(h())))`.
func (c *GroupClient) Use(hooks ...Hook) {
	c.hooks.Group = append(c.hooks.Group, hooks...)
}

// Create returns a create builder for Group.
func (c *GroupClient) Create() *GroupCreate {
	mutation := newGroupMutation(c.config, OpCreate)
	return &GroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Group entities.
func (c *GroupClient) CreateBulk(builders ...*GroupCreate) *GroupCreateBulk {
	return &GroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Group.
func (c *GroupClient) Update() *GroupUpdate {
	mutation := newGroupMutation(c.config, OpUpdate)
	return &GroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupClient) UpdateOne(gr *Group) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroup(gr))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupClient) UpdateOneID(id uuid.UUID) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroupID(id))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Group.
func (c *GroupClient) Delete() *GroupDelete {
	mutation := newGroupMutation(c.config, OpDelete)
	return &GroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GroupClient) DeleteOne(gr *Group) *GroupDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GroupClient) DeleteOneID(id uuid.UUID) *GroupDeleteOne {
	builder := c.Delete().Where(group.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GroupDeleteOne{builder}
}

// Query returns a query builder for Group.
func (c *GroupClient) Query() *GroupQuery {
	return &GroupQuery{config: c.config}
}

// Get returns a Group entity by its id.
func (c *GroupClient) Get(ctx context.Context, id uuid.UUID) (*Group, error) {
	return c.Query().Where(group.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupClient) GetX(ctx context.Context, id uuid.UUID) *Group {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryGroupRoles queries the group_roles edge of a Group.
func (c *GroupClient) QueryGroupRoles(gr *Group) *GroupRoleQuery {
	query := &GroupRoleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(grouprole.Table, grouprole.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, group.GroupRolesTable, group.GroupRolesColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryWorkspaces queries the workspaces edge of a Group.
func (c *GroupClient) QueryWorkspaces(gr *Group) *WorkspaceQuery {
	query := &WorkspaceQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(workspace.Table, workspace.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, group.WorkspacesTable, group.WorkspacesColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GroupClient) Hooks() []Hook {
	return c.hooks.Group
}

// GroupRoleClient is a client for the GroupRole schema.
type GroupRoleClient struct {
	config
}

// NewGroupRoleClient returns a client for the GroupRole from the given config.
func NewGroupRoleClient(c config) *GroupRoleClient {
	return &GroupRoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `grouprole.Hooks(f(g(h())))`.
func (c *GroupRoleClient) Use(hooks ...Hook) {
	c.hooks.GroupRole = append(c.hooks.GroupRole, hooks...)
}

// Create returns a create builder for GroupRole.
func (c *GroupRoleClient) Create() *GroupRoleCreate {
	mutation := newGroupRoleMutation(c.config, OpCreate)
	return &GroupRoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of GroupRole entities.
func (c *GroupRoleClient) CreateBulk(builders ...*GroupRoleCreate) *GroupRoleCreateBulk {
	return &GroupRoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for GroupRole.
func (c *GroupRoleClient) Update() *GroupRoleUpdate {
	mutation := newGroupRoleMutation(c.config, OpUpdate)
	return &GroupRoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupRoleClient) UpdateOne(gr *GroupRole) *GroupRoleUpdateOne {
	mutation := newGroupRoleMutation(c.config, OpUpdateOne, withGroupRole(gr))
	return &GroupRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupRoleClient) UpdateOneID(id uuid.UUID) *GroupRoleUpdateOne {
	mutation := newGroupRoleMutation(c.config, OpUpdateOne, withGroupRoleID(id))
	return &GroupRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for GroupRole.
func (c *GroupRoleClient) Delete() *GroupRoleDelete {
	mutation := newGroupRoleMutation(c.config, OpDelete)
	return &GroupRoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GroupRoleClient) DeleteOne(gr *GroupRole) *GroupRoleDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GroupRoleClient) DeleteOneID(id uuid.UUID) *GroupRoleDeleteOne {
	builder := c.Delete().Where(grouprole.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GroupRoleDeleteOne{builder}
}

// Query returns a query builder for GroupRole.
func (c *GroupRoleClient) Query() *GroupRoleQuery {
	return &GroupRoleQuery{config: c.config}
}

// Get returns a GroupRole entity by its id.
func (c *GroupRoleClient) Get(ctx context.Context, id uuid.UUID) (*GroupRole, error) {
	return c.Query().Where(grouprole.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupRoleClient) GetX(ctx context.Context, id uuid.UUID) *GroupRole {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryGroups queries the groups edge of a GroupRole.
func (c *GroupRoleClient) QueryGroups(gr *GroupRole) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(grouprole.Table, grouprole.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, grouprole.GroupsTable, grouprole.GroupsColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAccounts queries the accounts edge of a GroupRole.
func (c *GroupRoleClient) QueryAccounts(gr *GroupRole) *AccountQuery {
	query := &AccountQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(grouprole.Table, grouprole.FieldID, id),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, grouprole.AccountsTable, grouprole.AccountsColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GroupRoleClient) Hooks() []Hook {
	return c.hooks.GroupRole
}

// PermissionClient is a client for the Permission schema.
type PermissionClient struct {
	config
}

// NewPermissionClient returns a client for the Permission from the given config.
func NewPermissionClient(c config) *PermissionClient {
	return &PermissionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `permission.Hooks(f(g(h())))`.
func (c *PermissionClient) Use(hooks ...Hook) {
	c.hooks.Permission = append(c.hooks.Permission, hooks...)
}

// Create returns a create builder for Permission.
func (c *PermissionClient) Create() *PermissionCreate {
	mutation := newPermissionMutation(c.config, OpCreate)
	return &PermissionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Permission entities.
func (c *PermissionClient) CreateBulk(builders ...*PermissionCreate) *PermissionCreateBulk {
	return &PermissionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Permission.
func (c *PermissionClient) Update() *PermissionUpdate {
	mutation := newPermissionMutation(c.config, OpUpdate)
	return &PermissionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PermissionClient) UpdateOne(pe *Permission) *PermissionUpdateOne {
	mutation := newPermissionMutation(c.config, OpUpdateOne, withPermission(pe))
	return &PermissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PermissionClient) UpdateOneID(id uuid.UUID) *PermissionUpdateOne {
	mutation := newPermissionMutation(c.config, OpUpdateOne, withPermissionID(id))
	return &PermissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Permission.
func (c *PermissionClient) Delete() *PermissionDelete {
	mutation := newPermissionMutation(c.config, OpDelete)
	return &PermissionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PermissionClient) DeleteOne(pe *Permission) *PermissionDeleteOne {
	return c.DeleteOneID(pe.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PermissionClient) DeleteOneID(id uuid.UUID) *PermissionDeleteOne {
	builder := c.Delete().Where(permission.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PermissionDeleteOne{builder}
}

// Query returns a query builder for Permission.
func (c *PermissionClient) Query() *PermissionQuery {
	return &PermissionQuery{config: c.config}
}

// Get returns a Permission entity by its id.
func (c *PermissionClient) Get(ctx context.Context, id uuid.UUID) (*Permission, error) {
	return c.Query().Where(permission.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PermissionClient) GetX(ctx context.Context, id uuid.UUID) *Permission {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PermissionClient) Hooks() []Hook {
	return c.hooks.Permission
}

// SessionClient is a client for the Session schema.
type SessionClient struct {
	config
}

// NewSessionClient returns a client for the Session from the given config.
func NewSessionClient(c config) *SessionClient {
	return &SessionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `session.Hooks(f(g(h())))`.
func (c *SessionClient) Use(hooks ...Hook) {
	c.hooks.Session = append(c.hooks.Session, hooks...)
}

// Create returns a create builder for Session.
func (c *SessionClient) Create() *SessionCreate {
	mutation := newSessionMutation(c.config, OpCreate)
	return &SessionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Session entities.
func (c *SessionClient) CreateBulk(builders ...*SessionCreate) *SessionCreateBulk {
	return &SessionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Session.
func (c *SessionClient) Update() *SessionUpdate {
	mutation := newSessionMutation(c.config, OpUpdate)
	return &SessionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SessionClient) UpdateOne(s *Session) *SessionUpdateOne {
	mutation := newSessionMutation(c.config, OpUpdateOne, withSession(s))
	return &SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SessionClient) UpdateOneID(id string) *SessionUpdateOne {
	mutation := newSessionMutation(c.config, OpUpdateOne, withSessionID(id))
	return &SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Session.
func (c *SessionClient) Delete() *SessionDelete {
	mutation := newSessionMutation(c.config, OpDelete)
	return &SessionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SessionClient) DeleteOne(s *Session) *SessionDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SessionClient) DeleteOneID(id string) *SessionDeleteOne {
	builder := c.Delete().Where(session.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SessionDeleteOne{builder}
}

// Query returns a query builder for Session.
func (c *SessionClient) Query() *SessionQuery {
	return &SessionQuery{config: c.config}
}

// Get returns a Session entity by its id.
func (c *SessionClient) Get(ctx context.Context, id string) (*Session, error) {
	return c.Query().Where(session.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SessionClient) GetX(ctx context.Context, id string) *Session {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SessionClient) Hooks() []Hook {
	return c.hooks.Session
}

// WorkspaceClient is a client for the Workspace schema.
type WorkspaceClient struct {
	config
}

// NewWorkspaceClient returns a client for the Workspace from the given config.
func NewWorkspaceClient(c config) *WorkspaceClient {
	return &WorkspaceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `workspace.Hooks(f(g(h())))`.
func (c *WorkspaceClient) Use(hooks ...Hook) {
	c.hooks.Workspace = append(c.hooks.Workspace, hooks...)
}

// Create returns a create builder for Workspace.
func (c *WorkspaceClient) Create() *WorkspaceCreate {
	mutation := newWorkspaceMutation(c.config, OpCreate)
	return &WorkspaceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Workspace entities.
func (c *WorkspaceClient) CreateBulk(builders ...*WorkspaceCreate) *WorkspaceCreateBulk {
	return &WorkspaceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Workspace.
func (c *WorkspaceClient) Update() *WorkspaceUpdate {
	mutation := newWorkspaceMutation(c.config, OpUpdate)
	return &WorkspaceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WorkspaceClient) UpdateOne(w *Workspace) *WorkspaceUpdateOne {
	mutation := newWorkspaceMutation(c.config, OpUpdateOne, withWorkspace(w))
	return &WorkspaceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WorkspaceClient) UpdateOneID(id uuid.UUID) *WorkspaceUpdateOne {
	mutation := newWorkspaceMutation(c.config, OpUpdateOne, withWorkspaceID(id))
	return &WorkspaceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Workspace.
func (c *WorkspaceClient) Delete() *WorkspaceDelete {
	mutation := newWorkspaceMutation(c.config, OpDelete)
	return &WorkspaceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *WorkspaceClient) DeleteOne(w *Workspace) *WorkspaceDeleteOne {
	return c.DeleteOneID(w.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *WorkspaceClient) DeleteOneID(id uuid.UUID) *WorkspaceDeleteOne {
	builder := c.Delete().Where(workspace.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WorkspaceDeleteOne{builder}
}

// Query returns a query builder for Workspace.
func (c *WorkspaceClient) Query() *WorkspaceQuery {
	return &WorkspaceQuery{config: c.config}
}

// Get returns a Workspace entity by its id.
func (c *WorkspaceClient) Get(ctx context.Context, id uuid.UUID) (*Workspace, error) {
	return c.Query().Where(workspace.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WorkspaceClient) GetX(ctx context.Context, id uuid.UUID) *Workspace {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Workspace.
func (c *WorkspaceClient) QueryOwner(w *Workspace) *AccountQuery {
	query := &AccountQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(workspace.Table, workspace.FieldID, id),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, workspace.OwnerTable, workspace.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryWorkspaceRoles queries the workspace_roles edge of a Workspace.
func (c *WorkspaceClient) QueryWorkspaceRoles(w *Workspace) *WorkspaceRoleQuery {
	query := &WorkspaceRoleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(workspace.Table, workspace.FieldID, id),
			sqlgraph.To(workspacerole.Table, workspacerole.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, workspace.WorkspaceRolesTable, workspace.WorkspaceRolesColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGroups queries the groups edge of a Workspace.
func (c *WorkspaceClient) QueryGroups(w *Workspace) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(workspace.Table, workspace.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, workspace.GroupsTable, workspace.GroupsColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *WorkspaceClient) Hooks() []Hook {
	return c.hooks.Workspace
}

// WorkspaceInvitationClient is a client for the WorkspaceInvitation schema.
type WorkspaceInvitationClient struct {
	config
}

// NewWorkspaceInvitationClient returns a client for the WorkspaceInvitation from the given config.
func NewWorkspaceInvitationClient(c config) *WorkspaceInvitationClient {
	return &WorkspaceInvitationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `workspaceinvitation.Hooks(f(g(h())))`.
func (c *WorkspaceInvitationClient) Use(hooks ...Hook) {
	c.hooks.WorkspaceInvitation = append(c.hooks.WorkspaceInvitation, hooks...)
}

// Create returns a create builder for WorkspaceInvitation.
func (c *WorkspaceInvitationClient) Create() *WorkspaceInvitationCreate {
	mutation := newWorkspaceInvitationMutation(c.config, OpCreate)
	return &WorkspaceInvitationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of WorkspaceInvitation entities.
func (c *WorkspaceInvitationClient) CreateBulk(builders ...*WorkspaceInvitationCreate) *WorkspaceInvitationCreateBulk {
	return &WorkspaceInvitationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for WorkspaceInvitation.
func (c *WorkspaceInvitationClient) Update() *WorkspaceInvitationUpdate {
	mutation := newWorkspaceInvitationMutation(c.config, OpUpdate)
	return &WorkspaceInvitationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WorkspaceInvitationClient) UpdateOne(wi *WorkspaceInvitation) *WorkspaceInvitationUpdateOne {
	mutation := newWorkspaceInvitationMutation(c.config, OpUpdateOne, withWorkspaceInvitation(wi))
	return &WorkspaceInvitationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WorkspaceInvitationClient) UpdateOneID(id uuid.UUID) *WorkspaceInvitationUpdateOne {
	mutation := newWorkspaceInvitationMutation(c.config, OpUpdateOne, withWorkspaceInvitationID(id))
	return &WorkspaceInvitationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for WorkspaceInvitation.
func (c *WorkspaceInvitationClient) Delete() *WorkspaceInvitationDelete {
	mutation := newWorkspaceInvitationMutation(c.config, OpDelete)
	return &WorkspaceInvitationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *WorkspaceInvitationClient) DeleteOne(wi *WorkspaceInvitation) *WorkspaceInvitationDeleteOne {
	return c.DeleteOneID(wi.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *WorkspaceInvitationClient) DeleteOneID(id uuid.UUID) *WorkspaceInvitationDeleteOne {
	builder := c.Delete().Where(workspaceinvitation.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WorkspaceInvitationDeleteOne{builder}
}

// Query returns a query builder for WorkspaceInvitation.
func (c *WorkspaceInvitationClient) Query() *WorkspaceInvitationQuery {
	return &WorkspaceInvitationQuery{config: c.config}
}

// Get returns a WorkspaceInvitation entity by its id.
func (c *WorkspaceInvitationClient) Get(ctx context.Context, id uuid.UUID) (*WorkspaceInvitation, error) {
	return c.Query().Where(workspaceinvitation.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WorkspaceInvitationClient) GetX(ctx context.Context, id uuid.UUID) *WorkspaceInvitation {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *WorkspaceInvitationClient) Hooks() []Hook {
	return c.hooks.WorkspaceInvitation
}

// WorkspaceRoleClient is a client for the WorkspaceRole schema.
type WorkspaceRoleClient struct {
	config
}

// NewWorkspaceRoleClient returns a client for the WorkspaceRole from the given config.
func NewWorkspaceRoleClient(c config) *WorkspaceRoleClient {
	return &WorkspaceRoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `workspacerole.Hooks(f(g(h())))`.
func (c *WorkspaceRoleClient) Use(hooks ...Hook) {
	c.hooks.WorkspaceRole = append(c.hooks.WorkspaceRole, hooks...)
}

// Create returns a create builder for WorkspaceRole.
func (c *WorkspaceRoleClient) Create() *WorkspaceRoleCreate {
	mutation := newWorkspaceRoleMutation(c.config, OpCreate)
	return &WorkspaceRoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of WorkspaceRole entities.
func (c *WorkspaceRoleClient) CreateBulk(builders ...*WorkspaceRoleCreate) *WorkspaceRoleCreateBulk {
	return &WorkspaceRoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for WorkspaceRole.
func (c *WorkspaceRoleClient) Update() *WorkspaceRoleUpdate {
	mutation := newWorkspaceRoleMutation(c.config, OpUpdate)
	return &WorkspaceRoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WorkspaceRoleClient) UpdateOne(wr *WorkspaceRole) *WorkspaceRoleUpdateOne {
	mutation := newWorkspaceRoleMutation(c.config, OpUpdateOne, withWorkspaceRole(wr))
	return &WorkspaceRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WorkspaceRoleClient) UpdateOneID(id uuid.UUID) *WorkspaceRoleUpdateOne {
	mutation := newWorkspaceRoleMutation(c.config, OpUpdateOne, withWorkspaceRoleID(id))
	return &WorkspaceRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for WorkspaceRole.
func (c *WorkspaceRoleClient) Delete() *WorkspaceRoleDelete {
	mutation := newWorkspaceRoleMutation(c.config, OpDelete)
	return &WorkspaceRoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *WorkspaceRoleClient) DeleteOne(wr *WorkspaceRole) *WorkspaceRoleDeleteOne {
	return c.DeleteOneID(wr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *WorkspaceRoleClient) DeleteOneID(id uuid.UUID) *WorkspaceRoleDeleteOne {
	builder := c.Delete().Where(workspacerole.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WorkspaceRoleDeleteOne{builder}
}

// Query returns a query builder for WorkspaceRole.
func (c *WorkspaceRoleClient) Query() *WorkspaceRoleQuery {
	return &WorkspaceRoleQuery{config: c.config}
}

// Get returns a WorkspaceRole entity by its id.
func (c *WorkspaceRoleClient) Get(ctx context.Context, id uuid.UUID) (*WorkspaceRole, error) {
	return c.Query().Where(workspacerole.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WorkspaceRoleClient) GetX(ctx context.Context, id uuid.UUID) *WorkspaceRole {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryWorkspaces queries the workspaces edge of a WorkspaceRole.
func (c *WorkspaceRoleClient) QueryWorkspaces(wr *WorkspaceRole) *WorkspaceQuery {
	query := &WorkspaceQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := wr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(workspacerole.Table, workspacerole.FieldID, id),
			sqlgraph.To(workspace.Table, workspace.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workspacerole.WorkspacesTable, workspacerole.WorkspacesColumn),
		)
		fromV = sqlgraph.Neighbors(wr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAccounts queries the accounts edge of a WorkspaceRole.
func (c *WorkspaceRoleClient) QueryAccounts(wr *WorkspaceRole) *AccountQuery {
	query := &AccountQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := wr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(workspacerole.Table, workspacerole.FieldID, id),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workspacerole.AccountsTable, workspacerole.AccountsColumn),
		)
		fromV = sqlgraph.Neighbors(wr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *WorkspaceRoleClient) Hooks() []Hook {
	return c.hooks.WorkspaceRole
}
