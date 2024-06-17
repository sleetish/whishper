// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/pluja/anysub/ent/predicate"
	"github.com/pluja/anysub/ent/transcription"
	"github.com/pluja/anysub/ent/translation"
	"github.com/pluja/anysub/ent/user"
)

// TranscriptionQuery is the builder for querying Transcription entities.
type TranscriptionQuery struct {
	config
	ctx              *QueryContext
	order            []transcription.OrderOption
	inters           []Interceptor
	predicates       []predicate.Transcription
	withTranslations *TranslationQuery
	withUser         *UserQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TranscriptionQuery builder.
func (tq *TranscriptionQuery) Where(ps ...predicate.Transcription) *TranscriptionQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TranscriptionQuery) Limit(limit int) *TranscriptionQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TranscriptionQuery) Offset(offset int) *TranscriptionQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TranscriptionQuery) Unique(unique bool) *TranscriptionQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TranscriptionQuery) Order(o ...transcription.OrderOption) *TranscriptionQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryTranslations chains the current query on the "translations" edge.
func (tq *TranscriptionQuery) QueryTranslations() *TranslationQuery {
	query := (&TranslationClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transcription.Table, transcription.FieldID, selector),
			sqlgraph.To(translation.Table, translation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, transcription.TranslationsTable, transcription.TranslationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (tq *TranscriptionQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transcription.Table, transcription.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, transcription.UserTable, transcription.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Transcription entity from the query.
// Returns a *NotFoundError when no Transcription was found.
func (tq *TranscriptionQuery) First(ctx context.Context) (*Transcription, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{transcription.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TranscriptionQuery) FirstX(ctx context.Context) *Transcription {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Transcription ID from the query.
// Returns a *NotFoundError when no Transcription ID was found.
func (tq *TranscriptionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{transcription.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TranscriptionQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Transcription entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Transcription entity is found.
// Returns a *NotFoundError when no Transcription entities are found.
func (tq *TranscriptionQuery) Only(ctx context.Context) (*Transcription, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{transcription.Label}
	default:
		return nil, &NotSingularError{transcription.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TranscriptionQuery) OnlyX(ctx context.Context) *Transcription {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Transcription ID in the query.
// Returns a *NotSingularError when more than one Transcription ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TranscriptionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{transcription.Label}
	default:
		err = &NotSingularError{transcription.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TranscriptionQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Transcriptions.
func (tq *TranscriptionQuery) All(ctx context.Context) ([]*Transcription, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Transcription, *TranscriptionQuery]()
	return withInterceptors[[]*Transcription](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TranscriptionQuery) AllX(ctx context.Context) []*Transcription {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Transcription IDs.
func (tq *TranscriptionQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err = tq.Select(transcription.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TranscriptionQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TranscriptionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TranscriptionQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TranscriptionQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TranscriptionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, "Exist")
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TranscriptionQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TranscriptionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TranscriptionQuery) Clone() *TranscriptionQuery {
	if tq == nil {
		return nil
	}
	return &TranscriptionQuery{
		config:           tq.config,
		ctx:              tq.ctx.Clone(),
		order:            append([]transcription.OrderOption{}, tq.order...),
		inters:           append([]Interceptor{}, tq.inters...),
		predicates:       append([]predicate.Transcription{}, tq.predicates...),
		withTranslations: tq.withTranslations.Clone(),
		withUser:         tq.withUser.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithTranslations tells the query-builder to eager-load the nodes that are connected to
// the "translations" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TranscriptionQuery) WithTranslations(opts ...func(*TranslationQuery)) *TranscriptionQuery {
	query := (&TranslationClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withTranslations = query
	return tq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TranscriptionQuery) WithUser(opts ...func(*UserQuery)) *TranscriptionQuery {
	query := (&UserClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withUser = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Status string `json:"status"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Transcription.Query().
//		GroupBy(transcription.FieldStatus).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TranscriptionQuery) GroupBy(field string, fields ...string) *TranscriptionGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TranscriptionGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = transcription.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Status string `json:"status"`
//	}
//
//	client.Transcription.Query().
//		Select(transcription.FieldStatus).
//		Scan(ctx, &v)
func (tq *TranscriptionQuery) Select(fields ...string) *TranscriptionSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TranscriptionSelect{TranscriptionQuery: tq}
	sbuild.label = transcription.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TranscriptionSelect configured with the given aggregations.
func (tq *TranscriptionQuery) Aggregate(fns ...AggregateFunc) *TranscriptionSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TranscriptionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !transcription.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TranscriptionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Transcription, error) {
	var (
		nodes       = []*Transcription{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [2]bool{
			tq.withTranslations != nil,
			tq.withUser != nil,
		}
	)
	if tq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, transcription.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Transcription).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Transcription{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withTranslations; query != nil {
		if err := tq.loadTranslations(ctx, query, nodes,
			func(n *Transcription) { n.Edges.Translations = []*Translation{} },
			func(n *Transcription, e *Translation) { n.Edges.Translations = append(n.Edges.Translations, e) }); err != nil {
			return nil, err
		}
	}
	if query := tq.withUser; query != nil {
		if err := tq.loadUser(ctx, query, nodes, nil,
			func(n *Transcription, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TranscriptionQuery) loadTranslations(ctx context.Context, query *TranslationQuery, nodes []*Transcription, init func(*Transcription), assign func(*Transcription, *Translation)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Transcription)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Translation(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(transcription.TranslationsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.transcription_translations
		if fk == nil {
			return fmt.Errorf(`foreign-key "transcription_translations" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "transcription_translations" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (tq *TranscriptionQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Transcription, init func(*Transcription), assign func(*Transcription, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Transcription)
	for i := range nodes {
		if nodes[i].user_transcriptions == nil {
			continue
		}
		fk := *nodes[i].user_transcriptions
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_transcriptions" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (tq *TranscriptionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TranscriptionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(transcription.Table, transcription.Columns, sqlgraph.NewFieldSpec(transcription.FieldID, field.TypeInt))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transcription.FieldID)
		for i := range fields {
			if fields[i] != transcription.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TranscriptionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(transcription.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = transcription.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TranscriptionGroupBy is the group-by builder for Transcription entities.
type TranscriptionGroupBy struct {
	selector
	build *TranscriptionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TranscriptionGroupBy) Aggregate(fns ...AggregateFunc) *TranscriptionGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TranscriptionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TranscriptionQuery, *TranscriptionGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TranscriptionGroupBy) sqlScan(ctx context.Context, root *TranscriptionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TranscriptionSelect is the builder for selecting fields of Transcription entities.
type TranscriptionSelect struct {
	*TranscriptionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TranscriptionSelect) Aggregate(fns ...AggregateFunc) *TranscriptionSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TranscriptionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TranscriptionQuery, *TranscriptionSelect](ctx, ts.TranscriptionQuery, ts, ts.inters, v)
}

func (ts *TranscriptionSelect) sqlScan(ctx context.Context, root *TranscriptionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
