// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ashwnacharya/working-without-mocks/adapters/persistence/sqlite/ent/ingredient"
	"github.com/ashwnacharya/working-without-mocks/adapters/persistence/sqlite/ent/predicate"
)

// IngredientQuery is the builder for querying Ingredient entities.
type IngredientQuery struct {
	config
	ctx        *QueryContext
	order      []ingredient.OrderOption
	inters     []Interceptor
	predicates []predicate.Ingredient
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the IngredientQuery builder.
func (iq *IngredientQuery) Where(ps ...predicate.Ingredient) *IngredientQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit the number of records to be returned by this query.
func (iq *IngredientQuery) Limit(limit int) *IngredientQuery {
	iq.ctx.Limit = &limit
	return iq
}

// Offset to start from.
func (iq *IngredientQuery) Offset(offset int) *IngredientQuery {
	iq.ctx.Offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *IngredientQuery) Unique(unique bool) *IngredientQuery {
	iq.ctx.Unique = &unique
	return iq
}

// Order specifies how the records should be ordered.
func (iq *IngredientQuery) Order(o ...ingredient.OrderOption) *IngredientQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// First returns the first Ingredient entity from the query.
// Returns a *NotFoundError when no Ingredient was found.
func (iq *IngredientQuery) First(ctx context.Context) (*Ingredient, error) {
	nodes, err := iq.Limit(1).All(setContextOp(ctx, iq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ingredient.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *IngredientQuery) FirstX(ctx context.Context) *Ingredient {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Ingredient ID from the query.
// Returns a *NotFoundError when no Ingredient ID was found.
func (iq *IngredientQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(1).IDs(setContextOp(ctx, iq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ingredient.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *IngredientQuery) FirstIDX(ctx context.Context) int {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Ingredient entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Ingredient entity is found.
// Returns a *NotFoundError when no Ingredient entities are found.
func (iq *IngredientQuery) Only(ctx context.Context) (*Ingredient, error) {
	nodes, err := iq.Limit(2).All(setContextOp(ctx, iq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ingredient.Label}
	default:
		return nil, &NotSingularError{ingredient.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *IngredientQuery) OnlyX(ctx context.Context) *Ingredient {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Ingredient ID in the query.
// Returns a *NotSingularError when more than one Ingredient ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *IngredientQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(2).IDs(setContextOp(ctx, iq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ingredient.Label}
	default:
		err = &NotSingularError{ingredient.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *IngredientQuery) OnlyIDX(ctx context.Context) int {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Ingredients.
func (iq *IngredientQuery) All(ctx context.Context) ([]*Ingredient, error) {
	ctx = setContextOp(ctx, iq.ctx, "All")
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Ingredient, *IngredientQuery]()
	return withInterceptors[[]*Ingredient](ctx, iq, qr, iq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iq *IngredientQuery) AllX(ctx context.Context) []*Ingredient {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Ingredient IDs.
func (iq *IngredientQuery) IDs(ctx context.Context) (ids []int, err error) {
	if iq.ctx.Unique == nil && iq.path != nil {
		iq.Unique(true)
	}
	ctx = setContextOp(ctx, iq.ctx, "IDs")
	if err = iq.Select(ingredient.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *IngredientQuery) IDsX(ctx context.Context) []int {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *IngredientQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iq.ctx, "Count")
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iq, querierCount[*IngredientQuery](), iq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iq *IngredientQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *IngredientQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iq.ctx, "Exist")
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *IngredientQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the IngredientQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *IngredientQuery) Clone() *IngredientQuery {
	if iq == nil {
		return nil
	}
	return &IngredientQuery{
		config:     iq.config,
		ctx:        iq.ctx.Clone(),
		order:      append([]ingredient.OrderOption{}, iq.order...),
		inters:     append([]Interceptor{}, iq.inters...),
		predicates: append([]predicate.Ingredient{}, iq.predicates...),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Ingredient.Query().
//		GroupBy(ingredient.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (iq *IngredientQuery) GroupBy(field string, fields ...string) *IngredientGroupBy {
	iq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &IngredientGroupBy{build: iq}
	grbuild.flds = &iq.ctx.Fields
	grbuild.label = ingredient.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Ingredient.Query().
//		Select(ingredient.FieldName).
//		Scan(ctx, &v)
func (iq *IngredientQuery) Select(fields ...string) *IngredientSelect {
	iq.ctx.Fields = append(iq.ctx.Fields, fields...)
	sbuild := &IngredientSelect{IngredientQuery: iq}
	sbuild.label = ingredient.Label
	sbuild.flds, sbuild.scan = &iq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a IngredientSelect configured with the given aggregations.
func (iq *IngredientQuery) Aggregate(fns ...AggregateFunc) *IngredientSelect {
	return iq.Select().Aggregate(fns...)
}

func (iq *IngredientQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iq); err != nil {
				return err
			}
		}
	}
	for _, f := range iq.ctx.Fields {
		if !ingredient.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *IngredientQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Ingredient, error) {
	var (
		nodes = []*Ingredient{}
		_spec = iq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Ingredient).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Ingredient{config: iq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (iq *IngredientQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.ctx.Fields
	if len(iq.ctx.Fields) > 0 {
		_spec.Unique = iq.ctx.Unique != nil && *iq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *IngredientQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(ingredient.Table, ingredient.Columns, sqlgraph.NewFieldSpec(ingredient.FieldID, field.TypeInt))
	_spec.From = iq.sql
	if unique := iq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iq.path != nil {
		_spec.Unique = true
	}
	if fields := iq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ingredient.FieldID)
		for i := range fields {
			if fields[i] != ingredient.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *IngredientQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(ingredient.Table)
	columns := iq.ctx.Fields
	if len(columns) == 0 {
		columns = ingredient.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.ctx.Unique != nil && *iq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// IngredientGroupBy is the group-by builder for Ingredient entities.
type IngredientGroupBy struct {
	selector
	build *IngredientQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *IngredientGroupBy) Aggregate(fns ...AggregateFunc) *IngredientGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the selector query and scans the result into the given value.
func (igb *IngredientGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, igb.build.ctx, "GroupBy")
	if err := igb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IngredientQuery, *IngredientGroupBy](ctx, igb.build, igb, igb.build.inters, v)
}

func (igb *IngredientGroupBy) sqlScan(ctx context.Context, root *IngredientQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*igb.flds)+len(igb.fns))
		for _, f := range *igb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*igb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// IngredientSelect is the builder for selecting fields of Ingredient entities.
type IngredientSelect struct {
	*IngredientQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (is *IngredientSelect) Aggregate(fns ...AggregateFunc) *IngredientSelect {
	is.fns = append(is.fns, fns...)
	return is
}

// Scan applies the selector query and scans the result into the given value.
func (is *IngredientSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, is.ctx, "Select")
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IngredientQuery, *IngredientSelect](ctx, is.IngredientQuery, is, is.inters, v)
}

func (is *IngredientSelect) sqlScan(ctx context.Context, root *IngredientQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(is.fns))
	for _, fn := range is.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*is.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
