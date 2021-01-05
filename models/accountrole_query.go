// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/adnaan/authzen/models/account"
	"github.com/adnaan/authzen/models/accountrole"
	"github.com/adnaan/authzen/models/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// AccountRoleQuery is the builder for querying AccountRole entities.
type AccountRoleQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	predicates []predicate.AccountRole
	// eager-loading edges.
	withAccounts *AccountQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (arq *AccountRoleQuery) Where(ps ...predicate.AccountRole) *AccountRoleQuery {
	arq.predicates = append(arq.predicates, ps...)
	return arq
}

// Limit adds a limit step to the query.
func (arq *AccountRoleQuery) Limit(limit int) *AccountRoleQuery {
	arq.limit = &limit
	return arq
}

// Offset adds an offset step to the query.
func (arq *AccountRoleQuery) Offset(offset int) *AccountRoleQuery {
	arq.offset = &offset
	return arq
}

// Order adds an order step to the query.
func (arq *AccountRoleQuery) Order(o ...OrderFunc) *AccountRoleQuery {
	arq.order = append(arq.order, o...)
	return arq
}

// QueryAccounts chains the current query on the accounts edge.
func (arq *AccountRoleQuery) QueryAccounts() *AccountQuery {
	query := &AccountQuery{config: arq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := arq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := arq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(accountrole.Table, accountrole.FieldID, selector),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, accountrole.AccountsTable, accountrole.AccountsColumn),
		)
		fromU = sqlgraph.SetNeighbors(arq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AccountRole entity in the query. Returns *NotFoundError when no accountrole was found.
func (arq *AccountRoleQuery) First(ctx context.Context) (*AccountRole, error) {
	nodes, err := arq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{accountrole.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (arq *AccountRoleQuery) FirstX(ctx context.Context) *AccountRole {
	node, err := arq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AccountRole id in the query. Returns *NotFoundError when no id was found.
func (arq *AccountRoleQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = arq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{accountrole.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (arq *AccountRoleQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := arq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only AccountRole entity in the query, returns an error if not exactly one entity was returned.
func (arq *AccountRoleQuery) Only(ctx context.Context) (*AccountRole, error) {
	nodes, err := arq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{accountrole.Label}
	default:
		return nil, &NotSingularError{accountrole.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (arq *AccountRoleQuery) OnlyX(ctx context.Context) *AccountRole {
	node, err := arq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only AccountRole id in the query, returns an error if not exactly one id was returned.
func (arq *AccountRoleQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = arq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{accountrole.Label}
	default:
		err = &NotSingularError{accountrole.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (arq *AccountRoleQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := arq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AccountRoles.
func (arq *AccountRoleQuery) All(ctx context.Context) ([]*AccountRole, error) {
	if err := arq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return arq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (arq *AccountRoleQuery) AllX(ctx context.Context) []*AccountRole {
	nodes, err := arq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AccountRole ids.
func (arq *AccountRoleQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := arq.Select(accountrole.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (arq *AccountRoleQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := arq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (arq *AccountRoleQuery) Count(ctx context.Context) (int, error) {
	if err := arq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return arq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (arq *AccountRoleQuery) CountX(ctx context.Context) int {
	count, err := arq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (arq *AccountRoleQuery) Exist(ctx context.Context) (bool, error) {
	if err := arq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return arq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (arq *AccountRoleQuery) ExistX(ctx context.Context) bool {
	exist, err := arq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (arq *AccountRoleQuery) Clone() *AccountRoleQuery {
	if arq == nil {
		return nil
	}
	return &AccountRoleQuery{
		config:       arq.config,
		limit:        arq.limit,
		offset:       arq.offset,
		order:        append([]OrderFunc{}, arq.order...),
		predicates:   append([]predicate.AccountRole{}, arq.predicates...),
		withAccounts: arq.withAccounts.Clone(),
		// clone intermediate query.
		sql:  arq.sql.Clone(),
		path: arq.path,
	}
}

//  WithAccounts tells the query-builder to eager-loads the nodes that are connected to
// the "accounts" edge. The optional arguments used to configure the query builder of the edge.
func (arq *AccountRoleQuery) WithAccounts(opts ...func(*AccountQuery)) *AccountRoleQuery {
	query := &AccountQuery{config: arq.config}
	for _, opt := range opts {
		opt(query)
	}
	arq.withAccounts = query
	return arq
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
//	client.AccountRole.Query().
//		GroupBy(accountrole.FieldName).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (arq *AccountRoleQuery) GroupBy(field string, fields ...string) *AccountRoleGroupBy {
	group := &AccountRoleGroupBy{config: arq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := arq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return arq.sqlQuery(), nil
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
//	client.AccountRole.Query().
//		Select(accountrole.FieldName).
//		Scan(ctx, &v)
//
func (arq *AccountRoleQuery) Select(field string, fields ...string) *AccountRoleSelect {
	selector := &AccountRoleSelect{config: arq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := arq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return arq.sqlQuery(), nil
	}
	return selector
}

func (arq *AccountRoleQuery) prepareQuery(ctx context.Context) error {
	if arq.path != nil {
		prev, err := arq.path(ctx)
		if err != nil {
			return err
		}
		arq.sql = prev
	}
	return nil
}

func (arq *AccountRoleQuery) sqlAll(ctx context.Context) ([]*AccountRole, error) {
	var (
		nodes       = []*AccountRole{}
		withFKs     = arq.withFKs
		_spec       = arq.querySpec()
		loadedTypes = [1]bool{
			arq.withAccounts != nil,
		}
	)
	if arq.withAccounts != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, accountrole.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &AccountRole{config: arq.config}
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
	if err := sqlgraph.QueryNodes(ctx, arq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := arq.withAccounts; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*AccountRole)
		for i := range nodes {
			if fk := nodes[i].account_account_roles; fk != nil {
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
				return nil, fmt.Errorf(`unexpected foreign-key "account_account_roles" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Accounts = n
			}
		}
	}

	return nodes, nil
}

func (arq *AccountRoleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := arq.querySpec()
	return sqlgraph.CountNodes(ctx, arq.driver, _spec)
}

func (arq *AccountRoleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := arq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %v", err)
	}
	return n > 0, nil
}

func (arq *AccountRoleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   accountrole.Table,
			Columns: accountrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: accountrole.FieldID,
			},
		},
		From:   arq.sql,
		Unique: true,
	}
	if ps := arq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := arq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := arq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := arq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, accountrole.ValidColumn)
			}
		}
	}
	return _spec
}

func (arq *AccountRoleQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(arq.driver.Dialect())
	t1 := builder.Table(accountrole.Table)
	selector := builder.Select(t1.Columns(accountrole.Columns...)...).From(t1)
	if arq.sql != nil {
		selector = arq.sql
		selector.Select(selector.Columns(accountrole.Columns...)...)
	}
	for _, p := range arq.predicates {
		p(selector)
	}
	for _, p := range arq.order {
		p(selector, accountrole.ValidColumn)
	}
	if offset := arq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := arq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AccountRoleGroupBy is the builder for group-by AccountRole entities.
type AccountRoleGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (argb *AccountRoleGroupBy) Aggregate(fns ...AggregateFunc) *AccountRoleGroupBy {
	argb.fns = append(argb.fns, fns...)
	return argb
}

// Scan applies the group-by query and scan the result into the given value.
func (argb *AccountRoleGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := argb.path(ctx)
	if err != nil {
		return err
	}
	argb.sql = query
	return argb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (argb *AccountRoleGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := argb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (argb *AccountRoleGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(argb.fields) > 1 {
		return nil, errors.New("models: AccountRoleGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := argb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (argb *AccountRoleGroupBy) StringsX(ctx context.Context) []string {
	v, err := argb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (argb *AccountRoleGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = argb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{accountrole.Label}
	default:
		err = fmt.Errorf("models: AccountRoleGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (argb *AccountRoleGroupBy) StringX(ctx context.Context) string {
	v, err := argb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (argb *AccountRoleGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(argb.fields) > 1 {
		return nil, errors.New("models: AccountRoleGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := argb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (argb *AccountRoleGroupBy) IntsX(ctx context.Context) []int {
	v, err := argb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (argb *AccountRoleGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = argb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{accountrole.Label}
	default:
		err = fmt.Errorf("models: AccountRoleGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (argb *AccountRoleGroupBy) IntX(ctx context.Context) int {
	v, err := argb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (argb *AccountRoleGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(argb.fields) > 1 {
		return nil, errors.New("models: AccountRoleGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := argb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (argb *AccountRoleGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := argb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (argb *AccountRoleGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = argb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{accountrole.Label}
	default:
		err = fmt.Errorf("models: AccountRoleGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (argb *AccountRoleGroupBy) Float64X(ctx context.Context) float64 {
	v, err := argb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (argb *AccountRoleGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(argb.fields) > 1 {
		return nil, errors.New("models: AccountRoleGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := argb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (argb *AccountRoleGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := argb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (argb *AccountRoleGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = argb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{accountrole.Label}
	default:
		err = fmt.Errorf("models: AccountRoleGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (argb *AccountRoleGroupBy) BoolX(ctx context.Context) bool {
	v, err := argb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (argb *AccountRoleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range argb.fields {
		if !accountrole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := argb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := argb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (argb *AccountRoleGroupBy) sqlQuery() *sql.Selector {
	selector := argb.sql
	columns := make([]string, 0, len(argb.fields)+len(argb.fns))
	columns = append(columns, argb.fields...)
	for _, fn := range argb.fns {
		columns = append(columns, fn(selector, accountrole.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(argb.fields...)
}

// AccountRoleSelect is the builder for select fields of AccountRole entities.
type AccountRoleSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ars *AccountRoleSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ars.path(ctx)
	if err != nil {
		return err
	}
	ars.sql = query
	return ars.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ars *AccountRoleSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ars.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ars *AccountRoleSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ars.fields) > 1 {
		return nil, errors.New("models: AccountRoleSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ars.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ars *AccountRoleSelect) StringsX(ctx context.Context) []string {
	v, err := ars.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ars *AccountRoleSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ars.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{accountrole.Label}
	default:
		err = fmt.Errorf("models: AccountRoleSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ars *AccountRoleSelect) StringX(ctx context.Context) string {
	v, err := ars.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ars *AccountRoleSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ars.fields) > 1 {
		return nil, errors.New("models: AccountRoleSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ars.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ars *AccountRoleSelect) IntsX(ctx context.Context) []int {
	v, err := ars.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ars *AccountRoleSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ars.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{accountrole.Label}
	default:
		err = fmt.Errorf("models: AccountRoleSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ars *AccountRoleSelect) IntX(ctx context.Context) int {
	v, err := ars.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ars *AccountRoleSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ars.fields) > 1 {
		return nil, errors.New("models: AccountRoleSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ars.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ars *AccountRoleSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ars.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ars *AccountRoleSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ars.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{accountrole.Label}
	default:
		err = fmt.Errorf("models: AccountRoleSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ars *AccountRoleSelect) Float64X(ctx context.Context) float64 {
	v, err := ars.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ars *AccountRoleSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ars.fields) > 1 {
		return nil, errors.New("models: AccountRoleSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ars.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ars *AccountRoleSelect) BoolsX(ctx context.Context) []bool {
	v, err := ars.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ars *AccountRoleSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ars.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{accountrole.Label}
	default:
		err = fmt.Errorf("models: AccountRoleSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ars *AccountRoleSelect) BoolX(ctx context.Context) bool {
	v, err := ars.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ars *AccountRoleSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ars.fields {
		if !accountrole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := ars.sqlQuery().Query()
	if err := ars.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ars *AccountRoleSelect) sqlQuery() sql.Querier {
	selector := ars.sql
	selector.Select(selector.Columns(ars.fields...)...)
	return selector
}
