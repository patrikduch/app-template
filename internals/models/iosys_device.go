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

// IosysDevice is an object representing the database table.
type IosysDevice struct {
	DeviceID    int         `boil:"device_id" json:"device_id" toml:"device_id" yaml:"device_id"`
	BridgeID    int         `boil:"bridge_id" json:"bridge_id" toml:"bridge_id" yaml:"bridge_id"`
	Enable      bool        `boil:"enable" json:"enable" toml:"enable" yaml:"enable"`
	Port        null.Int    `boil:"port" json:"port,omitempty" toml:"port" yaml:"port,omitempty"`
	Certificate null.String `boil:"certificate" json:"certificate,omitempty" toml:"certificate" yaml:"certificate,omitempty"`
	Key         null.String `boil:"key" json:"key,omitempty" toml:"key" yaml:"key,omitempty"`
	Timeout     null.Int16  `boil:"timeout" json:"timeout,omitempty" toml:"timeout" yaml:"timeout,omitempty"`
	Reconnect   null.Int16  `boil:"reconnect" json:"reconnect,omitempty" toml:"reconnect" yaml:"reconnect,omitempty"`

	R *iosysDeviceR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L iosysDeviceL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var IosysDeviceColumns = struct {
	DeviceID    string
	BridgeID    string
	Enable      string
	Port        string
	Certificate string
	Key         string
	Timeout     string
	Reconnect   string
}{
	DeviceID:    "device_id",
	BridgeID:    "bridge_id",
	Enable:      "enable",
	Port:        "port",
	Certificate: "certificate",
	Key:         "key",
	Timeout:     "timeout",
	Reconnect:   "reconnect",
}

var IosysDeviceTableColumns = struct {
	DeviceID    string
	BridgeID    string
	Enable      string
	Port        string
	Certificate string
	Key         string
	Timeout     string
	Reconnect   string
}{
	DeviceID:    "iosys_device.device_id",
	BridgeID:    "iosys_device.bridge_id",
	Enable:      "iosys_device.enable",
	Port:        "iosys_device.port",
	Certificate: "iosys_device.certificate",
	Key:         "iosys_device.key",
	Timeout:     "iosys_device.timeout",
	Reconnect:   "iosys_device.reconnect",
}

// Generated where

var IosysDeviceWhere = struct {
	DeviceID    whereHelperint
	BridgeID    whereHelperint
	Enable      whereHelperbool
	Port        whereHelpernull_Int
	Certificate whereHelpernull_String
	Key         whereHelpernull_String
	Timeout     whereHelpernull_Int16
	Reconnect   whereHelpernull_Int16
}{
	DeviceID:    whereHelperint{field: "\"iosys_device\".\"device_id\""},
	BridgeID:    whereHelperint{field: "\"iosys_device\".\"bridge_id\""},
	Enable:      whereHelperbool{field: "\"iosys_device\".\"enable\""},
	Port:        whereHelpernull_Int{field: "\"iosys_device\".\"port\""},
	Certificate: whereHelpernull_String{field: "\"iosys_device\".\"certificate\""},
	Key:         whereHelpernull_String{field: "\"iosys_device\".\"key\""},
	Timeout:     whereHelpernull_Int16{field: "\"iosys_device\".\"timeout\""},
	Reconnect:   whereHelpernull_Int16{field: "\"iosys_device\".\"reconnect\""},
}

// IosysDeviceRels is where relationship names are stored.
var IosysDeviceRels = struct {
}{}

// iosysDeviceR is where relationships are stored.
type iosysDeviceR struct {
}

// NewStruct creates a new relationship struct
func (*iosysDeviceR) NewStruct() *iosysDeviceR {
	return &iosysDeviceR{}
}

// iosysDeviceL is where Load methods for each relationship are stored.
type iosysDeviceL struct{}

var (
	iosysDeviceAllColumns            = []string{"device_id", "bridge_id", "enable", "port", "certificate", "key", "timeout", "reconnect"}
	iosysDeviceColumnsWithoutDefault = []string{"device_id", "bridge_id"}
	iosysDeviceColumnsWithDefault    = []string{"enable", "port", "certificate", "key", "timeout", "reconnect"}
	iosysDevicePrimaryKeyColumns     = []string{"device_id"}
	iosysDeviceGeneratedColumns      = []string{}
)

type (
	// IosysDeviceSlice is an alias for a slice of pointers to IosysDevice.
	// This should almost always be used instead of []IosysDevice.
	IosysDeviceSlice []*IosysDevice
	// IosysDeviceHook is the signature for custom IosysDevice hook methods
	IosysDeviceHook func(context.Context, boil.ContextExecutor, *IosysDevice) error

	iosysDeviceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	iosysDeviceType                 = reflect.TypeOf(&IosysDevice{})
	iosysDeviceMapping              = queries.MakeStructMapping(iosysDeviceType)
	iosysDevicePrimaryKeyMapping, _ = queries.BindMapping(iosysDeviceType, iosysDeviceMapping, iosysDevicePrimaryKeyColumns)
	iosysDeviceInsertCacheMut       sync.RWMutex
	iosysDeviceInsertCache          = make(map[string]insertCache)
	iosysDeviceUpdateCacheMut       sync.RWMutex
	iosysDeviceUpdateCache          = make(map[string]updateCache)
	iosysDeviceUpsertCacheMut       sync.RWMutex
	iosysDeviceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var iosysDeviceAfterSelectHooks []IosysDeviceHook

var iosysDeviceBeforeInsertHooks []IosysDeviceHook
var iosysDeviceAfterInsertHooks []IosysDeviceHook

var iosysDeviceBeforeUpdateHooks []IosysDeviceHook
var iosysDeviceAfterUpdateHooks []IosysDeviceHook

var iosysDeviceBeforeDeleteHooks []IosysDeviceHook
var iosysDeviceAfterDeleteHooks []IosysDeviceHook

var iosysDeviceBeforeUpsertHooks []IosysDeviceHook
var iosysDeviceAfterUpsertHooks []IosysDeviceHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *IosysDevice) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range iosysDeviceAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *IosysDevice) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range iosysDeviceBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *IosysDevice) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range iosysDeviceAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *IosysDevice) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range iosysDeviceBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *IosysDevice) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range iosysDeviceAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *IosysDevice) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range iosysDeviceBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *IosysDevice) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range iosysDeviceAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *IosysDevice) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range iosysDeviceBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *IosysDevice) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range iosysDeviceAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddIosysDeviceHook registers your hook function for all future operations.
func AddIosysDeviceHook(hookPoint boil.HookPoint, iosysDeviceHook IosysDeviceHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		iosysDeviceAfterSelectHooks = append(iosysDeviceAfterSelectHooks, iosysDeviceHook)
	case boil.BeforeInsertHook:
		iosysDeviceBeforeInsertHooks = append(iosysDeviceBeforeInsertHooks, iosysDeviceHook)
	case boil.AfterInsertHook:
		iosysDeviceAfterInsertHooks = append(iosysDeviceAfterInsertHooks, iosysDeviceHook)
	case boil.BeforeUpdateHook:
		iosysDeviceBeforeUpdateHooks = append(iosysDeviceBeforeUpdateHooks, iosysDeviceHook)
	case boil.AfterUpdateHook:
		iosysDeviceAfterUpdateHooks = append(iosysDeviceAfterUpdateHooks, iosysDeviceHook)
	case boil.BeforeDeleteHook:
		iosysDeviceBeforeDeleteHooks = append(iosysDeviceBeforeDeleteHooks, iosysDeviceHook)
	case boil.AfterDeleteHook:
		iosysDeviceAfterDeleteHooks = append(iosysDeviceAfterDeleteHooks, iosysDeviceHook)
	case boil.BeforeUpsertHook:
		iosysDeviceBeforeUpsertHooks = append(iosysDeviceBeforeUpsertHooks, iosysDeviceHook)
	case boil.AfterUpsertHook:
		iosysDeviceAfterUpsertHooks = append(iosysDeviceAfterUpsertHooks, iosysDeviceHook)
	}
}

// One returns a single iosysDevice record from the query.
func (q iosysDeviceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*IosysDevice, error) {
	o := &IosysDevice{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for iosys_device")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all IosysDevice records from the query.
func (q iosysDeviceQuery) All(ctx context.Context, exec boil.ContextExecutor) (IosysDeviceSlice, error) {
	var o []*IosysDevice

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to IosysDevice slice")
	}

	if len(iosysDeviceAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all IosysDevice records in the query.
func (q iosysDeviceQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count iosys_device rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q iosysDeviceQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if iosys_device exists")
	}

	return count > 0, nil
}

// IosysDevices retrieves all the records using an executor.
func IosysDevices(mods ...qm.QueryMod) iosysDeviceQuery {
	mods = append(mods, qm.From("\"iosys_device\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"iosys_device\".*"})
	}

	return iosysDeviceQuery{q}
}

// FindIosysDevice retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindIosysDevice(ctx context.Context, exec boil.ContextExecutor, deviceID int, selectCols ...string) (*IosysDevice, error) {
	iosysDeviceObj := &IosysDevice{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"iosys_device\" where \"device_id\"=$1", sel,
	)

	q := queries.Raw(query, deviceID)

	err := q.Bind(ctx, exec, iosysDeviceObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from iosys_device")
	}

	if err = iosysDeviceObj.doAfterSelectHooks(ctx, exec); err != nil {
		return iosysDeviceObj, err
	}

	return iosysDeviceObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *IosysDevice) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no iosys_device provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(iosysDeviceColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	iosysDeviceInsertCacheMut.RLock()
	cache, cached := iosysDeviceInsertCache[key]
	iosysDeviceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			iosysDeviceAllColumns,
			iosysDeviceColumnsWithDefault,
			iosysDeviceColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(iosysDeviceType, iosysDeviceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(iosysDeviceType, iosysDeviceMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"iosys_device\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"iosys_device\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into iosys_device")
	}

	if !cached {
		iosysDeviceInsertCacheMut.Lock()
		iosysDeviceInsertCache[key] = cache
		iosysDeviceInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the IosysDevice.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *IosysDevice) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	iosysDeviceUpdateCacheMut.RLock()
	cache, cached := iosysDeviceUpdateCache[key]
	iosysDeviceUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			iosysDeviceAllColumns,
			iosysDevicePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update iosys_device, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"iosys_device\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, iosysDevicePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(iosysDeviceType, iosysDeviceMapping, append(wl, iosysDevicePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update iosys_device row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for iosys_device")
	}

	if !cached {
		iosysDeviceUpdateCacheMut.Lock()
		iosysDeviceUpdateCache[key] = cache
		iosysDeviceUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q iosysDeviceQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for iosys_device")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for iosys_device")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o IosysDeviceSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), iosysDevicePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"iosys_device\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, iosysDevicePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in iosysDevice slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all iosysDevice")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *IosysDevice) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no iosys_device provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(iosysDeviceColumnsWithDefault, o)

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

	iosysDeviceUpsertCacheMut.RLock()
	cache, cached := iosysDeviceUpsertCache[key]
	iosysDeviceUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			iosysDeviceAllColumns,
			iosysDeviceColumnsWithDefault,
			iosysDeviceColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			iosysDeviceAllColumns,
			iosysDevicePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert iosys_device, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(iosysDevicePrimaryKeyColumns))
			copy(conflict, iosysDevicePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"iosys_device\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(iosysDeviceType, iosysDeviceMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(iosysDeviceType, iosysDeviceMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert iosys_device")
	}

	if !cached {
		iosysDeviceUpsertCacheMut.Lock()
		iosysDeviceUpsertCache[key] = cache
		iosysDeviceUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single IosysDevice record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *IosysDevice) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no IosysDevice provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), iosysDevicePrimaryKeyMapping)
	sql := "DELETE FROM \"iosys_device\" WHERE \"device_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from iosys_device")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for iosys_device")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q iosysDeviceQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no iosysDeviceQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from iosys_device")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for iosys_device")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o IosysDeviceSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(iosysDeviceBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), iosysDevicePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"iosys_device\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, iosysDevicePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from iosysDevice slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for iosys_device")
	}

	if len(iosysDeviceAfterDeleteHooks) != 0 {
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
func (o *IosysDevice) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindIosysDevice(ctx, exec, o.DeviceID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *IosysDeviceSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := IosysDeviceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), iosysDevicePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"iosys_device\".* FROM \"iosys_device\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, iosysDevicePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in IosysDeviceSlice")
	}

	*o = slice

	return nil
}

// IosysDeviceExists checks if the IosysDevice row exists.
func IosysDeviceExists(ctx context.Context, exec boil.ContextExecutor, deviceID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"iosys_device\" where \"device_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, deviceID)
	}
	row := exec.QueryRowContext(ctx, sql, deviceID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if iosys_device exists")
	}

	return exists, nil
}
