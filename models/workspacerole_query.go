// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/adnaan/authzen/models/account"
	"github.com/adnaan/authzen/models/predicate"
	"github.com/adnaan/authzen/models/workspace"
	"github.com/adnaan/authzen/models/workspacerole"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// WorkspaceRoleQuery is the builder for querying WorkspaceRole entities.
type WorkspaceRoleQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	predicates []predicate.WorkspaceRole
	// eager-loading edges.
	withWorkspaces *WorkspaceQuery
	withAccounts   *AccountQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (wrq *WorkspaceRoleQuery) Where(ps ...predicate.WorkspaceRole) *WorkspaceRoleQuery {
	wrq.predicates = append(wrq.predicates, ps...)
	return wrq
}

// Limit adds a limit step to the query.
func (wrq *WorkspaceRoleQuery) Limit(limit int) *WorkspaceRoleQuery {
	wrq.limit = &limit
	return wrq
}

// Offset adds an offset step to the query.
func (wrq *WorkspaceRoleQuery) Offset(offset int) *WorkspaceRoleQuery {
	wrq.offset = &offset
	return wrq
}

// Order adds an order step to the query.
func (wrq *WorkspaceRoleQuery) Order(o ...OrderFunc) *WorkspaceRoleQuery {
	wrq.order = append(wrq.order, o...)
	return wrq
}

// QueryWorkspaces chains the current query on the workspaces edge.
func (wrq *WorkspaceRoleQuery) QueryWorkspaces() *WorkspaceQuery {
	query := &WorkspaceQuery{config: wrq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wrq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workspacerole.Table, workspacerole.FieldID, selector),
			sqlgraph.To(workspace.Table, workspace.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workspacerole.WorkspacesTable, workspacerole.WorkspacesColumn),
		)
		fromU = sqlgraph.SetNeighbors(wrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAccounts chains the current query on the accounts edge.
func (wrq *WorkspaceRoleQuery) QueryAccounts() *AccountQuery {
	query := &AccountQuery{config: wrq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wrq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workspacerole.Table, workspacerole.FieldID, selector),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workspacerole.AccountsTable, workspacerole.AccountsColumn),
		)
		fromU = sqlgraph.SetNeighbors(wrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WorkspaceRole entity in the query. Returns *NotFoundError when no workspacerole was found.
func (wrq *WorkspaceRoleQuery) First(ctx context.Context) (*WorkspaceRole, error) {
	nodes, err := wrq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{workspacerole.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wrq *WorkspaceRoleQuery) FirstX(ctx context.Context) *WorkspaceRole {
	node, err := wrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WorkspaceRole id in the query. Returns *NotFoundError when no id was found.
func (wrq *WorkspaceRoleQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wrq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{workspacerole.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wrq *WorkspaceRoleQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := wrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only WorkspaceRole entity in the query, returns an error if not exactly one entity was returned.
func (wrq *WorkspaceRoleQuery) Only(ctx context.Context) (*WorkspaceRole, error) {
	nodes, err := wrq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{workspacerole.Label}
	default:
		return nil, &NotSingularError{workspacerole.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wrq *WorkspaceRoleQuery) OnlyX(ctx context.Context) *WorkspaceRole {
	node, err := wrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only WorkspaceRole id in the query, returns an error if not exactly one id was returned.
func (wrq *WorkspaceRoleQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wrq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{workspacerole.Label}
	default:
		err = &NotSingularError{workspacerole.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wrq *WorkspaceRoleQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := wrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WorkspaceRoles.
func (wrq *WorkspaceRoleQuery) All(ctx context.Context) ([]*WorkspaceRole, error) {
	if err := wrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wrq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wrq *WorkspaceRoleQuery) AllX(ctx context.Context) []*WorkspaceRole {
	nodes, err := wrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WorkspaceRole ids.
func (wrq *WorkspaceRoleQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := wrq.Select(workspacerole.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wrq *WorkspaceRoleQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := wrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wrq *WorkspaceRoleQuery) Count(ctx context.Context) (int, error) {
	if err := wrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wrq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wrq *WorkspaceRoleQuery) CountX(ctx context.Context) int {
	count, err := wrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wrq *WorkspaceRoleQuery) Exist(ctx context.Context) (bool, error) {
	if err := wrq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wrq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wrq *WorkspaceRoleQuery) ExistX(ctx context.Context) bool {
	exist, err := wrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wrq *WorkspaceRoleQuery) Clone() *WorkspaceRoleQuery {
	if wrq == nil {
		return nil
	}
	return &WorkspaceRoleQuery{
		config:         wrq.config,
		limit:          wrq.limit,
		offset:         wrq.offset,
		order:          append([]OrderFunc{}, wrq.order...),
		predicates:     append([]predicate.WorkspaceRole{}, wrq.predicates...),
		withWorkspaces: wrq.withWorkspaces.Clone(),
		withAccounts:   wrq.withAccounts.Clone(),
		// clone intermediate query.
		sql:  wrq.sql.Clone(),
		path: wrq.path,
	}
}

//  WithWorkspaces tells the query-builder to eager-loads the nodes that are connected to
// the "workspaces" edge. The optional arguments used to configure the query builder of the edge.
func (wrq *WorkspaceRoleQuery) WithWorkspaces(opts ...func(*WorkspaceQuery)) *WorkspaceRoleQuery {
	query := &WorkspaceQuery{config: wrq.config}
	for _, opt := range opts {
		opt(query)
	}
	wrq.withWorkspaces = query
	return wrq
}

//  WithAccounts tells the query-builder to eager-loads the nodes that are connected to
// the "accounts" edge. The optional arguments used to configure the query builder of the edge.
func (wrq *WorkspaceRoleQuery) WithAccounts(opts ...func(*AccountQuery)) *WorkspaceRoleQuery {
	query := &AccountQuery{config: wrq.config}
	for _, opt := range opts {
		opt(query)
	}
	wrq.withAccounts = query
	return wrq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WorkspaceRole.Query().
//		GroupBy(workspacerole.FieldName).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (wrq *WorkspaceRoleQuery) GroupBy(field string, fields ...string) *WorkspaceRoleGroupBy {
	group := &WorkspaceRoleGroupBy{config: wrq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wrq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.WorkspaceRole.Query().
//		Select(workspacerole.FieldName).
//		Scan(ctx, &v)
//
func (wrq *WorkspaceRoleQuery) Select(field string, fields ...string) *WorkspaceRoleSelect {
	selector := &WorkspaceRoleSelect{config: wrq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wrq.sqlQuery(), nil
	}
	return selector
}

func (wrq *WorkspaceRoleQuery) prepareQuery(ctx context.Context) error {
	if wrq.path != nil {
		prev, err := wrq.path(ctx)
		if err != nil {
			return err
		}
		wrq.sql = prev
	}
	return nil
}

func (wrq *WorkspaceRoleQuery) sqlAll(ctx context.Context) ([]*WorkspaceRole, error) {
	var (
		nodes       = []*WorkspaceRole{}
		withFKs     = wrq.withFKs
		_spec       = wrq.querySpec()
		loadedTypes = [2]bool{
			wrq.withWorkspaces != nil,
			wrq.withAccounts != nil,
		}
	)
	if wrq.withWorkspaces != nil || wrq.withAccounts != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, workspacerole.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &WorkspaceRole{config: wrq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("models: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, wrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := wrq.withWorkspaces; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*WorkspaceRole)
		for i := range nodes {
			if fk := nodes[i].workspace_workspace_roles; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(workspace.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "workspace_workspace_roles" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Workspaces = n
			}
		}
	}

	if query := wrq.withAccounts; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*WorkspaceRole)
		for i := range nodes {
			if fk := nodes[i].account_workspace_roles; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(account.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "account_workspace_roles" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Accounts = n
			}
		}
	}

	return nodes, nil
}

func (wrq *WorkspaceRoleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wrq.querySpec()
	return sqlgraph.CountNodes(ctx, wrq.driver, _spec)
}

func (wrq *WorkspaceRoleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wrq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %v", err)
	}
	return n > 0, nil
}

func (wrq *WorkspaceRoleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workspacerole.Table,
			Columns: workspacerole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: workspacerole.FieldID,
			},
		},
		From:   wrq.sql,
		Unique: true,
	}
	if ps := wrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wrq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wrq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, workspacerole.ValidColumn)
			}
		}
	}
	return _spec
}

func (wrq *WorkspaceRoleQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(wrq.driver.Dialect())
	t1 := builder.Table(workspacerole.Table)
	selector := builder.Select(t1.Columns(workspacerole.Columns...)...).From(t1)
	if wrq.sql != nil {
		selector = wrq.sql
		selector.Select(selector.Columns(workspacerole.Columns...)...)
	}
	for _, p := range wrq.predicates {
		p(selector)
	}
	for _, p := range wrq.order {
		p(selector, workspacerole.ValidColumn)
	}
	if offset := wrq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wrq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WorkspaceRoleGroupBy is the builder for group-by WorkspaceRole entities.
type WorkspaceRoleGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wrgb *WorkspaceRoleGroupBy) Aggregate(fns ...AggregateFunc) *WorkspaceRoleGroupBy {
	wrgb.fns = append(wrgb.fns, fns...)
	return wrgb
}

// Scan applies the group-by query and scan the result into the given value.
func (wrgb *WorkspaceRoleGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := wrgb.path(ctx)
	if err != nil {
		return err
	}
	wrgb.sql = query
	return wrgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wrgb *WorkspaceRoleGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := wrgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (wrgb *WorkspaceRoleGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(wrgb.fields) > 1 {
		return nil, errors.New("models: WorkspaceRoleGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := wrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wrgb *WorkspaceRoleGroupBy) StringsX(ctx context.Context) []string {
	v, err := wrgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (wrgb *WorkspaceRoleGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wrgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspacerole.Label}
	default:
		err = fmt.Errorf("models: WorkspaceRoleGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wrgb *WorkspaceRoleGroupBy) StringX(ctx context.Context) string {
	v, err := wrgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (wrgb *WorkspaceRoleGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(wrgb.fields) > 1 {
		return nil, errors.New("models: WorkspaceRoleGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := wrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wrgb *WorkspaceRoleGroupBy) IntsX(ctx context.Context) []int {
	v, err := wrgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (wrgb *WorkspaceRoleGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wrgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspacerole.Label}
	default:
		err = fmt.Errorf("models: WorkspaceRoleGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wrgb *WorkspaceRoleGroupBy) IntX(ctx context.Context) int {
	v, err := wrgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (wrgb *WorkspaceRoleGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(wrgb.fields) > 1 {
		return nil, errors.New("models: WorkspaceRoleGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := wrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wrgb *WorkspaceRoleGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := wrgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (wrgb *WorkspaceRoleGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wrgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspacerole.Label}
	default:
		err = fmt.Errorf("models: WorkspaceRoleGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wrgb *WorkspaceRoleGroupBy) Float64X(ctx context.Context) float64 {
	v, err := wrgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (wrgb *WorkspaceRoleGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(wrgb.fields) > 1 {
		return nil, errors.New("models: WorkspaceRoleGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := wrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wrgb *WorkspaceRoleGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := wrgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (wrgb *WorkspaceRoleGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wrgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspacerole.Label}
	default:
		err = fmt.Errorf("models: WorkspaceRoleGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wrgb *WorkspaceRoleGroupBy) BoolX(ctx context.Context) bool {
	v, err := wrgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wrgb *WorkspaceRoleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range wrgb.fields {
		if !workspacerole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wrgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wrgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wrgb *WorkspaceRoleGroupBy) sqlQuery() *sql.Selector {
	selector := wrgb.sql
	columns := make([]string, 0, len(wrgb.fields)+len(wrgb.fns))
	columns = append(columns, wrgb.fields...)
	for _, fn := range wrgb.fns {
		columns = append(columns, fn(selector, workspacerole.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(wrgb.fields...)
}

// WorkspaceRoleSelect is the builder for select fields of WorkspaceRole entities.
type WorkspaceRoleSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (wrs *WorkspaceRoleSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := wrs.path(ctx)
	if err != nil {
		return err
	}
	wrs.sql = query
	return wrs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wrs *WorkspaceRoleSelect) ScanX(ctx context.Context, v interface{}) {
	if err := wrs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (wrs *WorkspaceRoleSelect) Strings(ctx context.Context) ([]string, error) {
	if len(wrs.fields) > 1 {
		return nil, errors.New("models: WorkspaceRoleSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := wrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wrs *WorkspaceRoleSelect) StringsX(ctx context.Context) []string {
	v, err := wrs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (wrs *WorkspaceRoleSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wrs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspacerole.Label}
	default:
		err = fmt.Errorf("models: WorkspaceRoleSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wrs *WorkspaceRoleSelect) StringX(ctx context.Context) string {
	v, err := wrs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (wrs *WorkspaceRoleSelect) Ints(ctx context.Context) ([]int, error) {
	if len(wrs.fields) > 1 {
		return nil, errors.New("models: WorkspaceRoleSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := wrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wrs *WorkspaceRoleSelect) IntsX(ctx context.Context) []int {
	v, err := wrs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (wrs *WorkspaceRoleSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wrs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspacerole.Label}
	default:
		err = fmt.Errorf("models: WorkspaceRoleSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wrs *WorkspaceRoleSelect) IntX(ctx context.Context) int {
	v, err := wrs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (wrs *WorkspaceRoleSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(wrs.fields) > 1 {
		return nil, errors.New("models: WorkspaceRoleSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := wrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wrs *WorkspaceRoleSelect) Float64sX(ctx context.Context) []float64 {
	v, err := wrs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (wrs *WorkspaceRoleSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wrs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspacerole.Label}
	default:
		err = fmt.Errorf("models: WorkspaceRoleSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wrs *WorkspaceRoleSelect) Float64X(ctx context.Context) float64 {
	v, err := wrs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (wrs *WorkspaceRoleSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(wrs.fields) > 1 {
		return nil, errors.New("models: WorkspaceRoleSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := wrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wrs *WorkspaceRoleSelect) BoolsX(ctx context.Context) []bool {
	v, err := wrs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (wrs *WorkspaceRoleSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wrs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspacerole.Label}
	default:
		err = fmt.Errorf("models: WorkspaceRoleSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wrs *WorkspaceRoleSelect) BoolX(ctx context.Context) bool {
	v, err := wrs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wrs *WorkspaceRoleSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range wrs.fields {
		if !workspacerole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := wrs.sqlQuery().Query()
	if err := wrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wrs *WorkspaceRoleSelect) sqlQuery() sql.Querier {
	selector := wrs.sql
	selector.Select(selector.Columns(wrs.fields...)...)
	return selector
}
