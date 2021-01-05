// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/adnaan/authzen/models/account"
	"github.com/adnaan/authzen/models/group"
	"github.com/adnaan/authzen/models/grouprole"
	"github.com/adnaan/authzen/models/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// GroupRoleQuery is the builder for querying GroupRole entities.
type GroupRoleQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	predicates []predicate.GroupRole
	// eager-loading edges.
	withGroups   *GroupQuery
	withAccounts *AccountQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (grq *GroupRoleQuery) Where(ps ...predicate.GroupRole) *GroupRoleQuery {
	grq.predicates = append(grq.predicates, ps...)
	return grq
}

// Limit adds a limit step to the query.
func (grq *GroupRoleQuery) Limit(limit int) *GroupRoleQuery {
	grq.limit = &limit
	return grq
}

// Offset adds an offset step to the query.
func (grq *GroupRoleQuery) Offset(offset int) *GroupRoleQuery {
	grq.offset = &offset
	return grq
}

// Order adds an order step to the query.
func (grq *GroupRoleQuery) Order(o ...OrderFunc) *GroupRoleQuery {
	grq.order = append(grq.order, o...)
	return grq
}

// QueryGroups chains the current query on the groups edge.
func (grq *GroupRoleQuery) QueryGroups() *GroupQuery {
	query := &GroupQuery{config: grq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := grq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := grq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(grouprole.Table, grouprole.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, grouprole.GroupsTable, grouprole.GroupsColumn),
		)
		fromU = sqlgraph.SetNeighbors(grq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAccounts chains the current query on the accounts edge.
func (grq *GroupRoleQuery) QueryAccounts() *AccountQuery {
	query := &AccountQuery{config: grq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := grq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := grq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(grouprole.Table, grouprole.FieldID, selector),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, grouprole.AccountsTable, grouprole.AccountsColumn),
		)
		fromU = sqlgraph.SetNeighbors(grq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GroupRole entity in the query. Returns *NotFoundError when no grouprole was found.
func (grq *GroupRoleQuery) First(ctx context.Context) (*GroupRole, error) {
	nodes, err := grq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{grouprole.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (grq *GroupRoleQuery) FirstX(ctx context.Context) *GroupRole {
	node, err := grq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GroupRole id in the query. Returns *NotFoundError when no id was found.
func (grq *GroupRoleQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = grq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{grouprole.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (grq *GroupRoleQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := grq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only GroupRole entity in the query, returns an error if not exactly one entity was returned.
func (grq *GroupRoleQuery) Only(ctx context.Context) (*GroupRole, error) {
	nodes, err := grq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{grouprole.Label}
	default:
		return nil, &NotSingularError{grouprole.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (grq *GroupRoleQuery) OnlyX(ctx context.Context) *GroupRole {
	node, err := grq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only GroupRole id in the query, returns an error if not exactly one id was returned.
func (grq *GroupRoleQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = grq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{grouprole.Label}
	default:
		err = &NotSingularError{grouprole.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (grq *GroupRoleQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := grq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GroupRoles.
func (grq *GroupRoleQuery) All(ctx context.Context) ([]*GroupRole, error) {
	if err := grq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return grq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (grq *GroupRoleQuery) AllX(ctx context.Context) []*GroupRole {
	nodes, err := grq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GroupRole ids.
func (grq *GroupRoleQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := grq.Select(grouprole.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (grq *GroupRoleQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := grq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (grq *GroupRoleQuery) Count(ctx context.Context) (int, error) {
	if err := grq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return grq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (grq *GroupRoleQuery) CountX(ctx context.Context) int {
	count, err := grq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (grq *GroupRoleQuery) Exist(ctx context.Context) (bool, error) {
	if err := grq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return grq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (grq *GroupRoleQuery) ExistX(ctx context.Context) bool {
	exist, err := grq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (grq *GroupRoleQuery) Clone() *GroupRoleQuery {
	if grq == nil {
		return nil
	}
	return &GroupRoleQuery{
		config:       grq.config,
		limit:        grq.limit,
		offset:       grq.offset,
		order:        append([]OrderFunc{}, grq.order...),
		predicates:   append([]predicate.GroupRole{}, grq.predicates...),
		withGroups:   grq.withGroups.Clone(),
		withAccounts: grq.withAccounts.Clone(),
		// clone intermediate query.
		sql:  grq.sql.Clone(),
		path: grq.path,
	}
}

//  WithGroups tells the query-builder to eager-loads the nodes that are connected to
// the "groups" edge. The optional arguments used to configure the query builder of the edge.
func (grq *GroupRoleQuery) WithGroups(opts ...func(*GroupQuery)) *GroupRoleQuery {
	query := &GroupQuery{config: grq.config}
	for _, opt := range opts {
		opt(query)
	}
	grq.withGroups = query
	return grq
}

//  WithAccounts tells the query-builder to eager-loads the nodes that are connected to
// the "accounts" edge. The optional arguments used to configure the query builder of the edge.
func (grq *GroupRoleQuery) WithAccounts(opts ...func(*AccountQuery)) *GroupRoleQuery {
	query := &AccountQuery{config: grq.config}
	for _, opt := range opts {
		opt(query)
	}
	grq.withAccounts = query
	return grq
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
//	client.GroupRole.Query().
//		GroupBy(grouprole.FieldName).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (grq *GroupRoleQuery) GroupBy(field string, fields ...string) *GroupRoleGroupBy {
	group := &GroupRoleGroupBy{config: grq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := grq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return grq.sqlQuery(), nil
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
//	client.GroupRole.Query().
//		Select(grouprole.FieldName).
//		Scan(ctx, &v)
//
func (grq *GroupRoleQuery) Select(field string, fields ...string) *GroupRoleSelect {
	selector := &GroupRoleSelect{config: grq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := grq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return grq.sqlQuery(), nil
	}
	return selector
}

func (grq *GroupRoleQuery) prepareQuery(ctx context.Context) error {
	if grq.path != nil {
		prev, err := grq.path(ctx)
		if err != nil {
			return err
		}
		grq.sql = prev
	}
	return nil
}

func (grq *GroupRoleQuery) sqlAll(ctx context.Context) ([]*GroupRole, error) {
	var (
		nodes       = []*GroupRole{}
		withFKs     = grq.withFKs
		_spec       = grq.querySpec()
		loadedTypes = [2]bool{
			grq.withGroups != nil,
			grq.withAccounts != nil,
		}
	)
	if grq.withGroups != nil || grq.withAccounts != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, grouprole.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &GroupRole{config: grq.config}
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
	if err := sqlgraph.QueryNodes(ctx, grq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := grq.withGroups; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*GroupRole)
		for i := range nodes {
			if fk := nodes[i].group_group_roles; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(group.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "group_group_roles" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Groups = n
			}
		}
	}

	if query := grq.withAccounts; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*GroupRole)
		for i := range nodes {
			if fk := nodes[i].account_group_roles; fk != nil {
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
				return nil, fmt.Errorf(`unexpected foreign-key "account_group_roles" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Accounts = n
			}
		}
	}

	return nodes, nil
}

func (grq *GroupRoleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := grq.querySpec()
	return sqlgraph.CountNodes(ctx, grq.driver, _spec)
}

func (grq *GroupRoleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := grq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %v", err)
	}
	return n > 0, nil
}

func (grq *GroupRoleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   grouprole.Table,
			Columns: grouprole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: grouprole.FieldID,
			},
		},
		From:   grq.sql,
		Unique: true,
	}
	if ps := grq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := grq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := grq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := grq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, grouprole.ValidColumn)
			}
		}
	}
	return _spec
}

func (grq *GroupRoleQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(grq.driver.Dialect())
	t1 := builder.Table(grouprole.Table)
	selector := builder.Select(t1.Columns(grouprole.Columns...)...).From(t1)
	if grq.sql != nil {
		selector = grq.sql
		selector.Select(selector.Columns(grouprole.Columns...)...)
	}
	for _, p := range grq.predicates {
		p(selector)
	}
	for _, p := range grq.order {
		p(selector, grouprole.ValidColumn)
	}
	if offset := grq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := grq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GroupRoleGroupBy is the builder for group-by GroupRole entities.
type GroupRoleGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (grgb *GroupRoleGroupBy) Aggregate(fns ...AggregateFunc) *GroupRoleGroupBy {
	grgb.fns = append(grgb.fns, fns...)
	return grgb
}

// Scan applies the group-by query and scan the result into the given value.
func (grgb *GroupRoleGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := grgb.path(ctx)
	if err != nil {
		return err
	}
	grgb.sql = query
	return grgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (grgb *GroupRoleGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := grgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (grgb *GroupRoleGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(grgb.fields) > 1 {
		return nil, errors.New("models: GroupRoleGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := grgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (grgb *GroupRoleGroupBy) StringsX(ctx context.Context) []string {
	v, err := grgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (grgb *GroupRoleGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = grgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grouprole.Label}
	default:
		err = fmt.Errorf("models: GroupRoleGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (grgb *GroupRoleGroupBy) StringX(ctx context.Context) string {
	v, err := grgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (grgb *GroupRoleGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(grgb.fields) > 1 {
		return nil, errors.New("models: GroupRoleGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := grgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (grgb *GroupRoleGroupBy) IntsX(ctx context.Context) []int {
	v, err := grgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (grgb *GroupRoleGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = grgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grouprole.Label}
	default:
		err = fmt.Errorf("models: GroupRoleGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (grgb *GroupRoleGroupBy) IntX(ctx context.Context) int {
	v, err := grgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (grgb *GroupRoleGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(grgb.fields) > 1 {
		return nil, errors.New("models: GroupRoleGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := grgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (grgb *GroupRoleGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := grgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (grgb *GroupRoleGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = grgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grouprole.Label}
	default:
		err = fmt.Errorf("models: GroupRoleGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (grgb *GroupRoleGroupBy) Float64X(ctx context.Context) float64 {
	v, err := grgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (grgb *GroupRoleGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(grgb.fields) > 1 {
		return nil, errors.New("models: GroupRoleGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := grgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (grgb *GroupRoleGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := grgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (grgb *GroupRoleGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = grgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grouprole.Label}
	default:
		err = fmt.Errorf("models: GroupRoleGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (grgb *GroupRoleGroupBy) BoolX(ctx context.Context) bool {
	v, err := grgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (grgb *GroupRoleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range grgb.fields {
		if !grouprole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := grgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := grgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (grgb *GroupRoleGroupBy) sqlQuery() *sql.Selector {
	selector := grgb.sql
	columns := make([]string, 0, len(grgb.fields)+len(grgb.fns))
	columns = append(columns, grgb.fields...)
	for _, fn := range grgb.fns {
		columns = append(columns, fn(selector, grouprole.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(grgb.fields...)
}

// GroupRoleSelect is the builder for select fields of GroupRole entities.
type GroupRoleSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (grs *GroupRoleSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := grs.path(ctx)
	if err != nil {
		return err
	}
	grs.sql = query
	return grs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (grs *GroupRoleSelect) ScanX(ctx context.Context, v interface{}) {
	if err := grs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (grs *GroupRoleSelect) Strings(ctx context.Context) ([]string, error) {
	if len(grs.fields) > 1 {
		return nil, errors.New("models: GroupRoleSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := grs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (grs *GroupRoleSelect) StringsX(ctx context.Context) []string {
	v, err := grs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (grs *GroupRoleSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = grs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grouprole.Label}
	default:
		err = fmt.Errorf("models: GroupRoleSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (grs *GroupRoleSelect) StringX(ctx context.Context) string {
	v, err := grs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (grs *GroupRoleSelect) Ints(ctx context.Context) ([]int, error) {
	if len(grs.fields) > 1 {
		return nil, errors.New("models: GroupRoleSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := grs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (grs *GroupRoleSelect) IntsX(ctx context.Context) []int {
	v, err := grs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (grs *GroupRoleSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = grs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grouprole.Label}
	default:
		err = fmt.Errorf("models: GroupRoleSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (grs *GroupRoleSelect) IntX(ctx context.Context) int {
	v, err := grs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (grs *GroupRoleSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(grs.fields) > 1 {
		return nil, errors.New("models: GroupRoleSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := grs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (grs *GroupRoleSelect) Float64sX(ctx context.Context) []float64 {
	v, err := grs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (grs *GroupRoleSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = grs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grouprole.Label}
	default:
		err = fmt.Errorf("models: GroupRoleSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (grs *GroupRoleSelect) Float64X(ctx context.Context) float64 {
	v, err := grs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (grs *GroupRoleSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(grs.fields) > 1 {
		return nil, errors.New("models: GroupRoleSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := grs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (grs *GroupRoleSelect) BoolsX(ctx context.Context) []bool {
	v, err := grs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (grs *GroupRoleSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = grs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grouprole.Label}
	default:
		err = fmt.Errorf("models: GroupRoleSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (grs *GroupRoleSelect) BoolX(ctx context.Context) bool {
	v, err := grs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (grs *GroupRoleSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range grs.fields {
		if !grouprole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := grs.sqlQuery().Query()
	if err := grs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (grs *GroupRoleSelect) sqlQuery() sql.Querier {
	selector := grs.sql
	selector.Select(selector.Columns(grs.fields...)...)
	return selector
}
