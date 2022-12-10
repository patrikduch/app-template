// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Heap is an object representing the database table.
type Heap struct {
	AssetID           int         `boil:"asset_id" json:"asset_id" toml:"asset_id" yaml:"asset_id"`
	Subtype           string      `boil:"subtype" json:"subtype" toml:"subtype" yaml:"subtype"`
	His               bool        `boil:"his" json:"his" toml:"his" yaml:"his"`
	TS                time.Time   `boil:"ts" json:"ts" toml:"ts" yaml:"ts"`
	Data              null.JSON   `boil:"data" json:"data,omitempty" toml:"data" yaml:"data,omitempty"`
	Valid             null.Bool   `boil:"valid" json:"valid,omitempty" toml:"valid" yaml:"valid,omitempty"`
	AllowedInactivity null.String `boil:"allowed_inactivity" json:"allowed_inactivity,omitempty" toml:"allowed_inactivity" yaml:"allowed_inactivity,omitempty"`
	UpdateCNT         int64       `boil:"update_cnt" json:"update_cnt" toml:"update_cnt" yaml:"update_cnt"`
	UpdateCNTResetTS  time.Time   `boil:"update_cnt_reset_ts" json:"update_cnt_reset_ts" toml:"update_cnt_reset_ts" yaml:"update_cnt_reset_ts"`

	R *heapR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L heapL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var HeapColumns = struct {
	AssetID           string
	Subtype           string
	His               string
	TS                string
	Data              string
	Valid             string
	AllowedInactivity string
	UpdateCNT         string
	UpdateCNTResetTS  string
}{
	AssetID:           "asset_id",
	Subtype:           "subtype",
	His:               "his",
	TS:                "ts",
	Data:              "data",
	Valid:             "valid",
	AllowedInactivity: "allowed_inactivity",
	UpdateCNT:         "update_cnt",
	UpdateCNTResetTS:  "update_cnt_reset_ts",
}

var HeapTableColumns = struct {
	AssetID           string
	Subtype           string
	His               string
	TS                string
	Data              string
	Valid             string
	AllowedInactivity string
	UpdateCNT         string
	UpdateCNTResetTS  string
}{
	AssetID:           "heap.asset_id",
	Subtype:           "heap.subtype",
	His:               "heap.his",
	TS:                "heap.ts",
	Data:              "heap.data",
	Valid:             "heap.valid",
	AllowedInactivity: "heap.allowed_inactivity",
	UpdateCNT:         "heap.update_cnt",
	UpdateCNTResetTS:  "heap.update_cnt_reset_ts",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var HeapWhere = struct {
	AssetID           whereHelperint
	Subtype           whereHelperstring
	His               whereHelperbool
	TS                whereHelpertime_Time
	Data              whereHelpernull_JSON
	Valid             whereHelpernull_Bool
	AllowedInactivity whereHelpernull_String
	UpdateCNT         whereHelperint64
	UpdateCNTResetTS  whereHelpertime_Time
}{
	AssetID:           whereHelperint{field: "\"heap\".\"asset_id\""},
	Subtype:           whereHelperstring{field: "\"heap\".\"subtype\""},
	His:               whereHelperbool{field: "\"heap\".\"his\""},
	TS:                whereHelpertime_Time{field: "\"heap\".\"ts\""},
	Data:              whereHelpernull_JSON{field: "\"heap\".\"data\""},
	Valid:             whereHelpernull_Bool{field: "\"heap\".\"valid\""},
	AllowedInactivity: whereHelpernull_String{field: "\"heap\".\"allowed_inactivity\""},
	UpdateCNT:         whereHelperint64{field: "\"heap\".\"update_cnt\""},
	UpdateCNTResetTS:  whereHelpertime_Time{field: "\"heap\".\"update_cnt_reset_ts\""},
}

// HeapRels is where relationship names are stored.
var HeapRels = struct {
}{}

// heapR is where relationships are stored.
type heapR struct {
}

// NewStruct creates a new relationship struct
func (*heapR) NewStruct() *heapR {
	return &heapR{}
}

// heapL is where Load methods for each relationship are stored.
type heapL struct{}

var (
	heapAllColumns            = []string{"asset_id", "subtype", "his", "ts", "data", "valid", "allowed_inactivity", "update_cnt", "update_cnt_reset_ts"}
	heapColumnsWithoutDefault = []string{"asset_id", "subtype"}
	heapColumnsWithDefault    = []string{"his", "ts", "data", "valid", "allowed_inactivity", "update_cnt", "update_cnt_reset_ts"}
	heapPrimaryKeyColumns     = []string{"asset_id", "subtype"}
	heapGeneratedColumns      = []string{}
)

type (
	// HeapSlice is an alias for a slice of pointers to Heap.
	// This should almost always be used instead of []Heap.
	HeapSlice []*Heap
	// HeapHook is the signature for custom Heap hook methods
	HeapHook func(context.Context, boil.ContextExecutor, *Heap) error

	heapQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	heapType                 = reflect.TypeOf(&Heap{})
	heapMapping              = queries.MakeStructMapping(heapType)
	heapPrimaryKeyMapping, _ = queries.BindMapping(heapType, heapMapping, heapPrimaryKeyColumns)
	heapInsertCacheMut       sync.RWMutex
	heapInsertCache          = make(map[string]insertCache)
	heapUpdateCacheMut       sync.RWMutex
	heapUpdateCache          = make(map[string]updateCache)
	heapUpsertCacheMut       sync.RWMutex
	heapUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var heapAfterSelectHooks []HeapHook

var heapBeforeInsertHooks []HeapHook
var heapAfterInsertHooks []HeapHook

var heapBeforeUpdateHooks []HeapHook
var heapAfterUpdateHooks []HeapHook

var heapBeforeDeleteHooks []HeapHook
var heapAfterDeleteHooks []HeapHook

var heapBeforeUpsertHooks []HeapHook
var heapAfterUpsertHooks []HeapHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Heap) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range heapAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Heap) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range heapBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Heap) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range heapAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Heap) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range heapBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Heap) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range heapAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Heap) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range heapBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Heap) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range heapAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Heap) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range heapBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Heap) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range heapAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddHeapHook registers your hook function for all future operations.
func AddHeapHook(hookPoint boil.HookPoint, heapHook HeapHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		heapAfterSelectHooks = append(heapAfterSelectHooks, heapHook)
	case boil.BeforeInsertHook:
		heapBeforeInsertHooks = append(heapBeforeInsertHooks, heapHook)
	case boil.AfterInsertHook:
		heapAfterInsertHooks = append(heapAfterInsertHooks, heapHook)
	case boil.BeforeUpdateHook:
		heapBeforeUpdateHooks = append(heapBeforeUpdateHooks, heapHook)
	case boil.AfterUpdateHook:
		heapAfterUpdateHooks = append(heapAfterUpdateHooks, heapHook)
	case boil.BeforeDeleteHook:
		heapBeforeDeleteHooks = append(heapBeforeDeleteHooks, heapHook)
	case boil.AfterDeleteHook:
		heapAfterDeleteHooks = append(heapAfterDeleteHooks, heapHook)
	case boil.BeforeUpsertHook:
		heapBeforeUpsertHooks = append(heapBeforeUpsertHooks, heapHook)
	case boil.AfterUpsertHook:
		heapAfterUpsertHooks = append(heapAfterUpsertHooks, heapHook)
	}
}

// One returns a single heap record from the query.
func (q heapQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Heap, error) {
	o := &Heap{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for heap")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Heap records from the query.
func (q heapQuery) All(ctx context.Context, exec boil.ContextExecutor) (HeapSlice, error) {
	var o []*Heap

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Heap slice")
	}

	if len(heapAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Heap records in the query.
func (q heapQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count heap rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q heapQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if heap exists")
	}

	return count > 0, nil
}

// Heaps retrieves all the records using an executor.
func Heaps(mods ...qm.QueryMod) heapQuery {
	mods = append(mods, qm.From("\"heap\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"heap\".*"})
	}

	return heapQuery{q}
}

// FindHeap retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindHeap(ctx context.Context, exec boil.ContextExecutor, assetID int, subtype string, selectCols ...string) (*Heap, error) {
	heapObj := &Heap{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"heap\" where \"asset_id\"=$1 AND \"subtype\"=$2", sel,
	)

	q := queries.Raw(query, assetID, subtype)

	err := q.Bind(ctx, exec, heapObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from heap")
	}

	if err = heapObj.doAfterSelectHooks(ctx, exec); err != nil {
		return heapObj, err
	}

	return heapObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Heap) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no heap provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(heapColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	heapInsertCacheMut.RLock()
	cache, cached := heapInsertCache[key]
	heapInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			heapAllColumns,
			heapColumnsWithDefault,
			heapColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(heapType, heapMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(heapType, heapMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"heap\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"heap\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into heap")
	}

	if !cached {
		heapInsertCacheMut.Lock()
		heapInsertCache[key] = cache
		heapInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Heap.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Heap) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	heapUpdateCacheMut.RLock()
	cache, cached := heapUpdateCache[key]
	heapUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			heapAllColumns,
			heapPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update heap, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"heap\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, heapPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(heapType, heapMapping, append(wl, heapPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update heap row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for heap")
	}

	if !cached {
		heapUpdateCacheMut.Lock()
		heapUpdateCache[key] = cache
		heapUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q heapQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for heap")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for heap")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o HeapSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), heapPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"heap\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, heapPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in heap slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all heap")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Heap) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no heap provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(heapColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	heapUpsertCacheMut.RLock()
	cache, cached := heapUpsertCache[key]
	heapUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			heapAllColumns,
			heapColumnsWithDefault,
			heapColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			heapAllColumns,
			heapPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert heap, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(heapPrimaryKeyColumns))
			copy(conflict, heapPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"heap\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(heapType, heapMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(heapType, heapMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert heap")
	}

	if !cached {
		heapUpsertCacheMut.Lock()
		heapUpsertCache[key] = cache
		heapUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Heap record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Heap) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Heap provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), heapPrimaryKeyMapping)
	sql := "DELETE FROM \"heap\" WHERE \"asset_id\"=$1 AND \"subtype\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from heap")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for heap")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q heapQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no heapQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from heap")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for heap")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o HeapSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(heapBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), heapPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"heap\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, heapPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from heap slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for heap")
	}

	if len(heapAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Heap) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindHeap(ctx, exec, o.AssetID, o.Subtype)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *HeapSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := HeapSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), heapPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"heap\".* FROM \"heap\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, heapPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in HeapSlice")
	}

	*o = slice

	return nil
}

// HeapExists checks if the Heap row exists.
func HeapExists(ctx context.Context, exec boil.ContextExecutor, assetID int, subtype string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"heap\" where \"asset_id\"=$1 AND \"subtype\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, assetID, subtype)
	}
	row := exec.QueryRowContext(ctx, sql, assetID, subtype)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if heap exists")
	}

	return exists, nil
}
