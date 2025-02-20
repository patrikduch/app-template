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

// Keyauth is an object representing the database table.
type Keyauth struct {
	KeyID   int          `boil:"key_id" json:"key_id" toml:"key_id" yaml:"key_id"`
	Key     null.String  `boil:"key" json:"key,omitempty" toml:"key" yaml:"key,omitempty"`
	Expires null.Float64 `boil:"expires" json:"expires,omitempty" toml:"expires" yaml:"expires,omitempty"`

	R *keyauthR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L keyauthL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var KeyauthColumns = struct {
	KeyID   string
	Key     string
	Expires string
}{
	KeyID:   "key_id",
	Key:     "key",
	Expires: "expires",
}

var KeyauthTableColumns = struct {
	KeyID   string
	Key     string
	Expires string
}{
	KeyID:   "keyauth.key_id",
	Key:     "keyauth.key",
	Expires: "keyauth.expires",
}

// Generated where

var KeyauthWhere = struct {
	KeyID   whereHelperint
	Key     whereHelpernull_String
	Expires whereHelpernull_Float64
}{
	KeyID:   whereHelperint{field: "\"keyauth\".\"key_id\""},
	Key:     whereHelpernull_String{field: "\"keyauth\".\"key\""},
	Expires: whereHelpernull_Float64{field: "\"keyauth\".\"expires\""},
}

// KeyauthRels is where relationship names are stored.
var KeyauthRels = struct {
}{}

// keyauthR is where relationships are stored.
type keyauthR struct {
}

// NewStruct creates a new relationship struct
func (*keyauthR) NewStruct() *keyauthR {
	return &keyauthR{}
}

// keyauthL is where Load methods for each relationship are stored.
type keyauthL struct{}

var (
	keyauthAllColumns            = []string{"key_id", "key", "expires"}
	keyauthColumnsWithoutDefault = []string{"key_id"}
	keyauthColumnsWithDefault    = []string{"key", "expires"}
	keyauthPrimaryKeyColumns     = []string{"key_id"}
	keyauthGeneratedColumns      = []string{}
)

type (
	// KeyauthSlice is an alias for a slice of pointers to Keyauth.
	// This should almost always be used instead of []Keyauth.
	KeyauthSlice []*Keyauth
	// KeyauthHook is the signature for custom Keyauth hook methods
	KeyauthHook func(context.Context, boil.ContextExecutor, *Keyauth) error

	keyauthQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	keyauthType                 = reflect.TypeOf(&Keyauth{})
	keyauthMapping              = queries.MakeStructMapping(keyauthType)
	keyauthPrimaryKeyMapping, _ = queries.BindMapping(keyauthType, keyauthMapping, keyauthPrimaryKeyColumns)
	keyauthInsertCacheMut       sync.RWMutex
	keyauthInsertCache          = make(map[string]insertCache)
	keyauthUpdateCacheMut       sync.RWMutex
	keyauthUpdateCache          = make(map[string]updateCache)
	keyauthUpsertCacheMut       sync.RWMutex
	keyauthUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var keyauthAfterSelectHooks []KeyauthHook

var keyauthBeforeInsertHooks []KeyauthHook
var keyauthAfterInsertHooks []KeyauthHook

var keyauthBeforeUpdateHooks []KeyauthHook
var keyauthAfterUpdateHooks []KeyauthHook

var keyauthBeforeDeleteHooks []KeyauthHook
var keyauthAfterDeleteHooks []KeyauthHook

var keyauthBeforeUpsertHooks []KeyauthHook
var keyauthAfterUpsertHooks []KeyauthHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Keyauth) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range keyauthAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Keyauth) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range keyauthBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Keyauth) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range keyauthAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Keyauth) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range keyauthBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Keyauth) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range keyauthAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Keyauth) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range keyauthBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Keyauth) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range keyauthAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Keyauth) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range keyauthBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Keyauth) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range keyauthAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddKeyauthHook registers your hook function for all future operations.
func AddKeyauthHook(hookPoint boil.HookPoint, keyauthHook KeyauthHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		keyauthAfterSelectHooks = append(keyauthAfterSelectHooks, keyauthHook)
	case boil.BeforeInsertHook:
		keyauthBeforeInsertHooks = append(keyauthBeforeInsertHooks, keyauthHook)
	case boil.AfterInsertHook:
		keyauthAfterInsertHooks = append(keyauthAfterInsertHooks, keyauthHook)
	case boil.BeforeUpdateHook:
		keyauthBeforeUpdateHooks = append(keyauthBeforeUpdateHooks, keyauthHook)
	case boil.AfterUpdateHook:
		keyauthAfterUpdateHooks = append(keyauthAfterUpdateHooks, keyauthHook)
	case boil.BeforeDeleteHook:
		keyauthBeforeDeleteHooks = append(keyauthBeforeDeleteHooks, keyauthHook)
	case boil.AfterDeleteHook:
		keyauthAfterDeleteHooks = append(keyauthAfterDeleteHooks, keyauthHook)
	case boil.BeforeUpsertHook:
		keyauthBeforeUpsertHooks = append(keyauthBeforeUpsertHooks, keyauthHook)
	case boil.AfterUpsertHook:
		keyauthAfterUpsertHooks = append(keyauthAfterUpsertHooks, keyauthHook)
	}
}

// One returns a single keyauth record from the query.
func (q keyauthQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Keyauth, error) {
	o := &Keyauth{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for keyauth")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Keyauth records from the query.
func (q keyauthQuery) All(ctx context.Context, exec boil.ContextExecutor) (KeyauthSlice, error) {
	var o []*Keyauth

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Keyauth slice")
	}

	if len(keyauthAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Keyauth records in the query.
func (q keyauthQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count keyauth rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q keyauthQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if keyauth exists")
	}

	return count > 0, nil
}

// Keyauths retrieves all the records using an executor.
func Keyauths(mods ...qm.QueryMod) keyauthQuery {
	mods = append(mods, qm.From("\"keyauth\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"keyauth\".*"})
	}

	return keyauthQuery{q}
}

// FindKeyauth retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindKeyauth(ctx context.Context, exec boil.ContextExecutor, keyID int, selectCols ...string) (*Keyauth, error) {
	keyauthObj := &Keyauth{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"keyauth\" where \"key_id\"=$1", sel,
	)

	q := queries.Raw(query, keyID)

	err := q.Bind(ctx, exec, keyauthObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from keyauth")
	}

	if err = keyauthObj.doAfterSelectHooks(ctx, exec); err != nil {
		return keyauthObj, err
	}

	return keyauthObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Keyauth) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no keyauth provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(keyauthColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	keyauthInsertCacheMut.RLock()
	cache, cached := keyauthInsertCache[key]
	keyauthInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			keyauthAllColumns,
			keyauthColumnsWithDefault,
			keyauthColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(keyauthType, keyauthMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(keyauthType, keyauthMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"keyauth\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"keyauth\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into keyauth")
	}

	if !cached {
		keyauthInsertCacheMut.Lock()
		keyauthInsertCache[key] = cache
		keyauthInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Keyauth.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Keyauth) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	keyauthUpdateCacheMut.RLock()
	cache, cached := keyauthUpdateCache[key]
	keyauthUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			keyauthAllColumns,
			keyauthPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update keyauth, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"keyauth\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, keyauthPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(keyauthType, keyauthMapping, append(wl, keyauthPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update keyauth row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for keyauth")
	}

	if !cached {
		keyauthUpdateCacheMut.Lock()
		keyauthUpdateCache[key] = cache
		keyauthUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q keyauthQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for keyauth")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for keyauth")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o KeyauthSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), keyauthPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"keyauth\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, keyauthPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in keyauth slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all keyauth")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Keyauth) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no keyauth provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(keyauthColumnsWithDefault, o)

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

	keyauthUpsertCacheMut.RLock()
	cache, cached := keyauthUpsertCache[key]
	keyauthUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			keyauthAllColumns,
			keyauthColumnsWithDefault,
			keyauthColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			keyauthAllColumns,
			keyauthPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert keyauth, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(keyauthPrimaryKeyColumns))
			copy(conflict, keyauthPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"keyauth\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(keyauthType, keyauthMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(keyauthType, keyauthMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert keyauth")
	}

	if !cached {
		keyauthUpsertCacheMut.Lock()
		keyauthUpsertCache[key] = cache
		keyauthUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Keyauth record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Keyauth) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Keyauth provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), keyauthPrimaryKeyMapping)
	sql := "DELETE FROM \"keyauth\" WHERE \"key_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from keyauth")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for keyauth")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q keyauthQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no keyauthQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from keyauth")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for keyauth")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o KeyauthSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(keyauthBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), keyauthPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"keyauth\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, keyauthPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from keyauth slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for keyauth")
	}

	if len(keyauthAfterDeleteHooks) != 0 {
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
func (o *Keyauth) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindKeyauth(ctx, exec, o.KeyID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *KeyauthSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := KeyauthSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), keyauthPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"keyauth\".* FROM \"keyauth\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, keyauthPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in KeyauthSlice")
	}

	*o = slice

	return nil
}

// KeyauthExists checks if the Keyauth row exists.
func KeyauthExists(ctx context.Context, exec boil.ContextExecutor, keyID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"keyauth\" where \"key_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, keyID)
	}
	row := exec.QueryRowContext(ctx, sql, keyID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if keyauth exists")
	}

	return exists, nil
}
