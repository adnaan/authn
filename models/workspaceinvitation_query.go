// Code generated (@generated) by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/adnaan/authzen/models/predicate"
	"github.com/adnaan/authzen/models/workspaceinvitation"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// WorkspaceInvitationQuery is the builder for querying WorkspaceInvitation entities.
type WorkspaceInvitationQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	predicates []predicate.WorkspaceInvitation
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (wiq *WorkspaceInvitationQuery) Where(ps ...predicate.WorkspaceInvitation) *WorkspaceInvitationQuery {
	wiq.predicates = append(wiq.predicates, ps...)
	return wiq
}

// Limit adds a limit step to the query.
func (wiq *WorkspaceInvitationQuery) Limit(limit int) *WorkspaceInvitationQuery {
	wiq.limit = &limit
	return wiq
}

// Offset adds an offset step to the query.
func (wiq *WorkspaceInvitationQuery) Offset(offset int) *WorkspaceInvitationQuery {
	wiq.offset = &offset
	return wiq
}

// Order adds an order step to the query.
func (wiq *WorkspaceInvitationQuery) Order(o ...OrderFunc) *WorkspaceInvitationQuery {
	wiq.order = append(wiq.order, o...)
	return wiq
}

// First returns the first WorkspaceInvitation entity in the query. Returns *NotFoundError when no workspaceinvitation was found.
func (wiq *WorkspaceInvitationQuery) First(ctx context.Context) (*WorkspaceInvitation, error) {
	nodes, err := wiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{workspaceinvitation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wiq *WorkspaceInvitationQuery) FirstX(ctx context.Context) *WorkspaceInvitation {
	node, err := wiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WorkspaceInvitation id in the query. Returns *NotFoundError when no id was found.
func (wiq *WorkspaceInvitationQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{workspaceinvitation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wiq *WorkspaceInvitationQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := wiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only WorkspaceInvitation entity in the query, returns an error if not exactly one entity was returned.
func (wiq *WorkspaceInvitationQuery) Only(ctx context.Context) (*WorkspaceInvitation, error) {
	nodes, err := wiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{workspaceinvitation.Label}
	default:
		return nil, &NotSingularError{workspaceinvitation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wiq *WorkspaceInvitationQuery) OnlyX(ctx context.Context) *WorkspaceInvitation {
	node, err := wiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only WorkspaceInvitation id in the query, returns an error if not exactly one id was returned.
func (wiq *WorkspaceInvitationQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wiq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{workspaceinvitation.Label}
	default:
		err = &NotSingularError{workspaceinvitation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wiq *WorkspaceInvitationQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := wiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WorkspaceInvitations.
func (wiq *WorkspaceInvitationQuery) All(ctx context.Context) ([]*WorkspaceInvitation, error) {
	if err := wiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wiq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wiq *WorkspaceInvitationQuery) AllX(ctx context.Context) []*WorkspaceInvitation {
	nodes, err := wiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WorkspaceInvitation ids.
func (wiq *WorkspaceInvitationQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := wiq.Select(workspaceinvitation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wiq *WorkspaceInvitationQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := wiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wiq *WorkspaceInvitationQuery) Count(ctx context.Context) (int, error) {
	if err := wiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wiq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wiq *WorkspaceInvitationQuery) CountX(ctx context.Context) int {
	count, err := wiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wiq *WorkspaceInvitationQuery) Exist(ctx context.Context) (bool, error) {
	if err := wiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wiq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wiq *WorkspaceInvitationQuery) ExistX(ctx context.Context) bool {
	exist, err := wiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wiq *WorkspaceInvitationQuery) Clone() *WorkspaceInvitationQuery {
	if wiq == nil {
		return nil
	}
	return &WorkspaceInvitationQuery{
		config:     wiq.config,
		limit:      wiq.limit,
		offset:     wiq.offset,
		order:      append([]OrderFunc{}, wiq.order...),
		predicates: append([]predicate.WorkspaceInvitation{}, wiq.predicates...),
		// clone intermediate query.
		sql:  wiq.sql.Clone(),
		path: wiq.path,
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		WorkspaceID uuid.UUID `json:"workspace_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WorkspaceInvitation.Query().
//		GroupBy(workspaceinvitation.FieldWorkspaceID).
//		Aggregate(models.Count()).
//		Scan(ctx, &v)
//
func (wiq *WorkspaceInvitationQuery) GroupBy(field string, fields ...string) *WorkspaceInvitationGroupBy {
	group := &WorkspaceInvitationGroupBy{config: wiq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wiq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		WorkspaceID uuid.UUID `json:"workspace_id,omitempty"`
//	}
//
//	client.WorkspaceInvitation.Query().
//		Select(workspaceinvitation.FieldWorkspaceID).
//		Scan(ctx, &v)
//
func (wiq *WorkspaceInvitationQuery) Select(field string, fields ...string) *WorkspaceInvitationSelect {
	selector := &WorkspaceInvitationSelect{config: wiq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wiq.sqlQuery(), nil
	}
	return selector
}

func (wiq *WorkspaceInvitationQuery) prepareQuery(ctx context.Context) error {
	if wiq.path != nil {
		prev, err := wiq.path(ctx)
		if err != nil {
			return err
		}
		wiq.sql = prev
	}
	return nil
}

func (wiq *WorkspaceInvitationQuery) sqlAll(ctx context.Context) ([]*WorkspaceInvitation, error) {
	var (
		nodes = []*WorkspaceInvitation{}
		_spec = wiq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &WorkspaceInvitation{config: wiq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("models: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, wiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (wiq *WorkspaceInvitationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wiq.querySpec()
	return sqlgraph.CountNodes(ctx, wiq.driver, _spec)
}

func (wiq *WorkspaceInvitationQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wiq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("models: check existence: %v", err)
	}
	return n > 0, nil
}

func (wiq *WorkspaceInvitationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workspaceinvitation.Table,
			Columns: workspaceinvitation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: workspaceinvitation.FieldID,
			},
		},
		From:   wiq.sql,
		Unique: true,
	}
	if ps := wiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wiq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wiq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, workspaceinvitation.ValidColumn)
			}
		}
	}
	return _spec
}

func (wiq *WorkspaceInvitationQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(wiq.driver.Dialect())
	t1 := builder.Table(workspaceinvitation.Table)
	selector := builder.Select(t1.Columns(workspaceinvitation.Columns...)...).From(t1)
	if wiq.sql != nil {
		selector = wiq.sql
		selector.Select(selector.Columns(workspaceinvitation.Columns...)...)
	}
	for _, p := range wiq.predicates {
		p(selector)
	}
	for _, p := range wiq.order {
		p(selector, workspaceinvitation.ValidColumn)
	}
	if offset := wiq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wiq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WorkspaceInvitationGroupBy is the builder for group-by WorkspaceInvitation entities.
type WorkspaceInvitationGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wigb *WorkspaceInvitationGroupBy) Aggregate(fns ...AggregateFunc) *WorkspaceInvitationGroupBy {
	wigb.fns = append(wigb.fns, fns...)
	return wigb
}

// Scan applies the group-by query and scan the result into the given value.
func (wigb *WorkspaceInvitationGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := wigb.path(ctx)
	if err != nil {
		return err
	}
	wigb.sql = query
	return wigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wigb *WorkspaceInvitationGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := wigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (wigb *WorkspaceInvitationGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(wigb.fields) > 1 {
		return nil, errors.New("models: WorkspaceInvitationGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := wigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wigb *WorkspaceInvitationGroupBy) StringsX(ctx context.Context) []string {
	v, err := wigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (wigb *WorkspaceInvitationGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspaceinvitation.Label}
	default:
		err = fmt.Errorf("models: WorkspaceInvitationGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wigb *WorkspaceInvitationGroupBy) StringX(ctx context.Context) string {
	v, err := wigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (wigb *WorkspaceInvitationGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(wigb.fields) > 1 {
		return nil, errors.New("models: WorkspaceInvitationGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := wigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wigb *WorkspaceInvitationGroupBy) IntsX(ctx context.Context) []int {
	v, err := wigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (wigb *WorkspaceInvitationGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspaceinvitation.Label}
	default:
		err = fmt.Errorf("models: WorkspaceInvitationGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wigb *WorkspaceInvitationGroupBy) IntX(ctx context.Context) int {
	v, err := wigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (wigb *WorkspaceInvitationGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(wigb.fields) > 1 {
		return nil, errors.New("models: WorkspaceInvitationGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := wigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wigb *WorkspaceInvitationGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := wigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (wigb *WorkspaceInvitationGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspaceinvitation.Label}
	default:
		err = fmt.Errorf("models: WorkspaceInvitationGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wigb *WorkspaceInvitationGroupBy) Float64X(ctx context.Context) float64 {
	v, err := wigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (wigb *WorkspaceInvitationGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(wigb.fields) > 1 {
		return nil, errors.New("models: WorkspaceInvitationGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := wigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wigb *WorkspaceInvitationGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := wigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (wigb *WorkspaceInvitationGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspaceinvitation.Label}
	default:
		err = fmt.Errorf("models: WorkspaceInvitationGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wigb *WorkspaceInvitationGroupBy) BoolX(ctx context.Context) bool {
	v, err := wigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wigb *WorkspaceInvitationGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range wigb.fields {
		if !workspaceinvitation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wigb *WorkspaceInvitationGroupBy) sqlQuery() *sql.Selector {
	selector := wigb.sql
	columns := make([]string, 0, len(wigb.fields)+len(wigb.fns))
	columns = append(columns, wigb.fields...)
	for _, fn := range wigb.fns {
		columns = append(columns, fn(selector, workspaceinvitation.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(wigb.fields...)
}

// WorkspaceInvitationSelect is the builder for select fields of WorkspaceInvitation entities.
type WorkspaceInvitationSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (wis *WorkspaceInvitationSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := wis.path(ctx)
	if err != nil {
		return err
	}
	wis.sql = query
	return wis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wis *WorkspaceInvitationSelect) ScanX(ctx context.Context, v interface{}) {
	if err := wis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (wis *WorkspaceInvitationSelect) Strings(ctx context.Context) ([]string, error) {
	if len(wis.fields) > 1 {
		return nil, errors.New("models: WorkspaceInvitationSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := wis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wis *WorkspaceInvitationSelect) StringsX(ctx context.Context) []string {
	v, err := wis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (wis *WorkspaceInvitationSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspaceinvitation.Label}
	default:
		err = fmt.Errorf("models: WorkspaceInvitationSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wis *WorkspaceInvitationSelect) StringX(ctx context.Context) string {
	v, err := wis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (wis *WorkspaceInvitationSelect) Ints(ctx context.Context) ([]int, error) {
	if len(wis.fields) > 1 {
		return nil, errors.New("models: WorkspaceInvitationSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := wis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wis *WorkspaceInvitationSelect) IntsX(ctx context.Context) []int {
	v, err := wis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (wis *WorkspaceInvitationSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspaceinvitation.Label}
	default:
		err = fmt.Errorf("models: WorkspaceInvitationSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wis *WorkspaceInvitationSelect) IntX(ctx context.Context) int {
	v, err := wis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (wis *WorkspaceInvitationSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(wis.fields) > 1 {
		return nil, errors.New("models: WorkspaceInvitationSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := wis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wis *WorkspaceInvitationSelect) Float64sX(ctx context.Context) []float64 {
	v, err := wis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (wis *WorkspaceInvitationSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspaceinvitation.Label}
	default:
		err = fmt.Errorf("models: WorkspaceInvitationSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wis *WorkspaceInvitationSelect) Float64X(ctx context.Context) float64 {
	v, err := wis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (wis *WorkspaceInvitationSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(wis.fields) > 1 {
		return nil, errors.New("models: WorkspaceInvitationSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := wis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wis *WorkspaceInvitationSelect) BoolsX(ctx context.Context) []bool {
	v, err := wis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (wis *WorkspaceInvitationSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workspaceinvitation.Label}
	default:
		err = fmt.Errorf("models: WorkspaceInvitationSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wis *WorkspaceInvitationSelect) BoolX(ctx context.Context) bool {
	v, err := wis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wis *WorkspaceInvitationSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range wis.fields {
		if !workspaceinvitation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := wis.sqlQuery().Query()
	if err := wis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wis *WorkspaceInvitationSelect) sqlQuery() sql.Querier {
	selector := wis.sql
	selector.Select(selector.Columns(wis.fields...)...)
	return selector
}
