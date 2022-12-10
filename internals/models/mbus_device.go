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

// MbusDevice is an object representing the database table.
type MbusDevice struct {
	DeviceID        int         `boil:"device_id" json:"device_id" toml:"device_id" yaml:"device_id"`
	BridgeID        int         `boil:"bridge_id" json:"bridge_id" toml:"bridge_id" yaml:"bridge_id"`
	Manufacturer    null.String `boil:"manufacturer" json:"manufacturer,omitempty" toml:"manufacturer" yaml:"manufacturer,omitempty"`
	Model           null.String `boil:"model" json:"model,omitempty" toml:"model" yaml:"model,omitempty"`
	Address         null.Int16  `boil:"address" json:"address,omitempty" toml:"address" yaml:"address,omitempty"`
	SecAddress      null.String `boil:"sec_address" json:"sec_address,omitempty" toml:"sec_address" yaml:"sec_address,omitempty"`
	Enable          bool        `boil:"enable" json:"enable" toml:"enable" yaml:"enable"`
	Raster          null.String `boil:"raster" json:"raster,omitempty" toml:"raster" yaml:"raster,omitempty"`
	MaxFail         null.Int    `boil:"max_fail" json:"max_fail,omitempty" toml:"max_fail" yaml:"max_fail,omitempty"`
	MaxRetry        null.Int    `boil:"max_retry" json:"max_retry,omitempty" toml:"max_retry" yaml:"max_retry,omitempty"`
	SendNke         null.Bool   `boil:"send_nke" json:"send_nke,omitempty" toml:"send_nke" yaml:"send_nke,omitempty"`
	AppResetSubcode null.Int16  `boil:"app_reset_subcode" json:"app_reset_subcode,omitempty" toml:"app_reset_subcode" yaml:"app_reset_subcode,omitempty"`
	MultiFrames     null.Int16  `boil:"multi_frames" json:"multi_frames,omitempty" toml:"multi_frames" yaml:"multi_frames,omitempty"`

	R *mbusDeviceR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L mbusDeviceL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MbusDeviceColumns = struct {
	DeviceID        string
	BridgeID        string
	Manufacturer    string
	Model           string
	Address         string
	SecAddress      string
	Enable          string
	Raster          string
	MaxFail         string
	MaxRetry        string
	SendNke         string
	AppResetSubcode string
	MultiFrames     string
}{
	DeviceID:        "device_id",
	BridgeID:        "bridge_id",
	Manufacturer:    "manufacturer",
	Model:           "model",
	Address:         "address",
	SecAddress:      "sec_address",
	Enable:          "enable",
	Raster:          "raster",
	MaxFail:         "max_fail",
	MaxRetry:        "max_retry",
	SendNke:         "send_nke",
	AppResetSubcode: "app_reset_subcode",
	MultiFrames:     "multi_frames",
}

var MbusDeviceTableColumns = struct {
	DeviceID        string
	BridgeID        string
	Manufacturer    string
	Model           string
	Address         string
	SecAddress      string
	Enable          string
	Raster          string
	MaxFail         string
	MaxRetry        string
	SendNke         string
	AppResetSubcode string
	MultiFrames     string
}{
	DeviceID:        "mbus_device.device_id",
	BridgeID:        "mbus_device.bridge_id",
	Manufacturer:    "mbus_device.manufacturer",
	Model:           "mbus_device.model",
	Address:         "mbus_device.address",
	SecAddress:      "mbus_device.sec_address",
	Enable:          "mbus_device.enable",
	Raster:          "mbus_device.raster",
	MaxFail:         "mbus_device.max_fail",
	MaxRetry:        "mbus_device.max_retry",
	SendNke:         "mbus_device.send_nke",
	AppResetSubcode: "mbus_device.app_reset_subcode",
	MultiFrames:     "mbus_device.multi_frames",
}

// Generated where

var MbusDeviceWhere = struct {
	DeviceID        whereHelperint
	BridgeID        whereHelperint
	Manufacturer    whereHelpernull_String
	Model           whereHelpernull_String
	Address         whereHelpernull_Int16
	SecAddress      whereHelpernull_String
	Enable          whereHelperbool
	Raster          whereHelpernull_String
	MaxFail         whereHelpernull_Int
	MaxRetry        whereHelpernull_Int
	SendNke         whereHelpernull_Bool
	AppResetSubcode whereHelpernull_Int16
	MultiFrames     whereHelpernull_Int16
}{
	DeviceID:        whereHelperint{field: "\"mbus_device\".\"device_id\""},
	BridgeID:        whereHelperint{field: "\"mbus_device\".\"bridge_id\""},
	Manufacturer:    whereHelpernull_String{field: "\"mbus_device\".\"manufacturer\""},
	Model:           whereHelpernull_String{field: "\"mbus_device\".\"model\""},
	Address:         whereHelpernull_Int16{field: "\"mbus_device\".\"address\""},
	SecAddress:      whereHelpernull_String{field: "\"mbus_device\".\"sec_address\""},
	Enable:          whereHelperbool{field: "\"mbus_device\".\"enable\""},
	Raster:          whereHelpernull_String{field: "\"mbus_device\".\"raster\""},
	MaxFail:         whereHelpernull_Int{field: "\"mbus_device\".\"max_fail\""},
	MaxRetry:        whereHelpernull_Int{field: "\"mbus_device\".\"max_retry\""},
	SendNke:         whereHelpernull_Bool{field: "\"mbus_device\".\"send_nke\""},
	AppResetSubcode: whereHelpernull_Int16{field: "\"mbus_device\".\"app_reset_subcode\""},
	MultiFrames:     whereHelpernull_Int16{field: "\"mbus_device\".\"multi_frames\""},
}

// MbusDeviceRels is where relationship names are stored.
var MbusDeviceRels = struct {
}{}

// mbusDeviceR is where relationships are stored.
type mbusDeviceR struct {
}

// NewStruct creates a new relationship struct
func (*mbusDeviceR) NewStruct() *mbusDeviceR {
	return &mbusDeviceR{}
}

// mbusDeviceL is where Load methods for each relationship are stored.
type mbusDeviceL struct{}

var (
	mbusDeviceAllColumns            = []string{"device_id", "bridge_id", "manufacturer", "model", "address", "sec_address", "enable", "raster", "max_fail", "max_retry", "send_nke", "app_reset_subcode", "multi_frames"}
	mbusDeviceColumnsWithoutDefault = []string{"device_id", "bridge_id"}
	mbusDeviceColumnsWithDefault    = []string{"manufacturer", "model", "address", "sec_address", "enable", "raster", "max_fail", "max_retry", "send_nke", "app_reset_subcode", "multi_frames"}
	mbusDevicePrimaryKeyColumns     = []string{"device_id"}
	mbusDeviceGeneratedColumns      = []string{}
)

type (
	// MbusDeviceSlice is an alias for a slice of pointers to MbusDevice.
	// This should almost always be used instead of []MbusDevice.
	MbusDeviceSlice []*MbusDevice
	// MbusDeviceHook is the signature for custom MbusDevice hook methods
	MbusDeviceHook func(context.Context, boil.ContextExecutor, *MbusDevice) error

	mbusDeviceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	mbusDeviceType                 = reflect.TypeOf(&MbusDevice{})
	mbusDeviceMapping              = queries.MakeStructMapping(mbusDeviceType)
	mbusDevicePrimaryKeyMapping, _ = queries.BindMapping(mbusDeviceType, mbusDeviceMapping, mbusDevicePrimaryKeyColumns)
	mbusDeviceInsertCacheMut       sync.RWMutex
	mbusDeviceInsertCache          = make(map[string]insertCache)
	mbusDeviceUpdateCacheMut       sync.RWMutex
	mbusDeviceUpdateCache          = make(map[string]updateCache)
	mbusDeviceUpsertCacheMut       sync.RWMutex
	mbusDeviceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var mbusDeviceAfterSelectHooks []MbusDeviceHook

var mbusDeviceBeforeInsertHooks []MbusDeviceHook
var mbusDeviceAfterInsertHooks []MbusDeviceHook

var mbusDeviceBeforeUpdateHooks []MbusDeviceHook
var mbusDeviceAfterUpdateHooks []MbusDeviceHook

var mbusDeviceBeforeDeleteHooks []MbusDeviceHook
var mbusDeviceAfterDeleteHooks []MbusDeviceHook

var mbusDeviceBeforeUpsertHooks []MbusDeviceHook
var mbusDeviceAfterUpsertHooks []MbusDeviceHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *MbusDevice) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mbusDeviceAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *MbusDevice) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mbusDeviceBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *MbusDevice) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mbusDeviceAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *MbusDevice) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mbusDeviceBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *MbusDevice) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mbusDeviceAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *MbusDevice) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mbusDeviceBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *MbusDevice) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mbusDeviceAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *MbusDevice) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mbusDeviceBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *MbusDevice) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mbusDeviceAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMbusDeviceHook registers your hook function for all future operations.
func AddMbusDeviceHook(hookPoint boil.HookPoint, mbusDeviceHook MbusDeviceHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		mbusDeviceAfterSelectHooks = append(mbusDeviceAfterSelectHooks, mbusDeviceHook)
	case boil.BeforeInsertHook:
		mbusDeviceBeforeInsertHooks = append(mbusDeviceBeforeInsertHooks, mbusDeviceHook)
	case boil.AfterInsertHook:
		mbusDeviceAfterInsertHooks = append(mbusDeviceAfterInsertHooks, mbusDeviceHook)
	case boil.BeforeUpdateHook:
		mbusDeviceBeforeUpdateHooks = append(mbusDeviceBeforeUpdateHooks, mbusDeviceHook)
	case boil.AfterUpdateHook:
		mbusDeviceAfterUpdateHooks = append(mbusDeviceAfterUpdateHooks, mbusDeviceHook)
	case boil.BeforeDeleteHook:
		mbusDeviceBeforeDeleteHooks = append(mbusDeviceBeforeDeleteHooks, mbusDeviceHook)
	case boil.AfterDeleteHook:
		mbusDeviceAfterDeleteHooks = append(mbusDeviceAfterDeleteHooks, mbusDeviceHook)
	case boil.BeforeUpsertHook:
		mbusDeviceBeforeUpsertHooks = append(mbusDeviceBeforeUpsertHooks, mbusDeviceHook)
	case boil.AfterUpsertHook:
		mbusDeviceAfterUpsertHooks = append(mbusDeviceAfterUpsertHooks, mbusDeviceHook)
	}
}

// One returns a single mbusDevice record from the query.
func (q mbusDeviceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MbusDevice, error) {
	o := &MbusDevice{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for mbus_device")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all MbusDevice records from the query.
func (q mbusDeviceQuery) All(ctx context.Context, exec boil.ContextExecutor) (MbusDeviceSlice, error) {
	var o []*MbusDevice

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to MbusDevice slice")
	}

	if len(mbusDeviceAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all MbusDevice records in the query.
func (q mbusDeviceQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count mbus_device rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q mbusDeviceQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if mbus_device exists")
	}

	return count > 0, nil
}

// MbusDevices retrieves all the records using an executor.
func MbusDevices(mods ...qm.QueryMod) mbusDeviceQuery {
	mods = append(mods, qm.From("\"mbus_device\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"mbus_device\".*"})
	}

	return mbusDeviceQuery{q}
}

// FindMbusDevice retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMbusDevice(ctx context.Context, exec boil.ContextExecutor, deviceID int, selectCols ...string) (*MbusDevice, error) {
	mbusDeviceObj := &MbusDevice{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"mbus_device\" where \"device_id\"=$1", sel,
	)

	q := queries.Raw(query, deviceID)

	err := q.Bind(ctx, exec, mbusDeviceObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from mbus_device")
	}

	if err = mbusDeviceObj.doAfterSelectHooks(ctx, exec); err != nil {
		return mbusDeviceObj, err
	}

	return mbusDeviceObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MbusDevice) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no mbus_device provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(mbusDeviceColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	mbusDeviceInsertCacheMut.RLock()
	cache, cached := mbusDeviceInsertCache[key]
	mbusDeviceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			mbusDeviceAllColumns,
			mbusDeviceColumnsWithDefault,
			mbusDeviceColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(mbusDeviceType, mbusDeviceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(mbusDeviceType, mbusDeviceMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"mbus_device\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"mbus_device\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into mbus_device")
	}

	if !cached {
		mbusDeviceInsertCacheMut.Lock()
		mbusDeviceInsertCache[key] = cache
		mbusDeviceInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the MbusDevice.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MbusDevice) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	mbusDeviceUpdateCacheMut.RLock()
	cache, cached := mbusDeviceUpdateCache[key]
	mbusDeviceUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			mbusDeviceAllColumns,
			mbusDevicePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update mbus_device, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"mbus_device\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, mbusDevicePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(mbusDeviceType, mbusDeviceMapping, append(wl, mbusDevicePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update mbus_device row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for mbus_device")
	}

	if !cached {
		mbusDeviceUpdateCacheMut.Lock()
		mbusDeviceUpdateCache[key] = cache
		mbusDeviceUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q mbusDeviceQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for mbus_device")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for mbus_device")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MbusDeviceSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mbusDevicePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"mbus_device\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, mbusDevicePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in mbusDevice slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all mbusDevice")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MbusDevice) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no mbus_device provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(mbusDeviceColumnsWithDefault, o)

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

	mbusDeviceUpsertCacheMut.RLock()
	cache, cached := mbusDeviceUpsertCache[key]
	mbusDeviceUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			mbusDeviceAllColumns,
			mbusDeviceColumnsWithDefault,
			mbusDeviceColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			mbusDeviceAllColumns,
			mbusDevicePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert mbus_device, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(mbusDevicePrimaryKeyColumns))
			copy(conflict, mbusDevicePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"mbus_device\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(mbusDeviceType, mbusDeviceMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(mbusDeviceType, mbusDeviceMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert mbus_device")
	}

	if !cached {
		mbusDeviceUpsertCacheMut.Lock()
		mbusDeviceUpsertCache[key] = cache
		mbusDeviceUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single MbusDevice record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MbusDevice) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MbusDevice provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), mbusDevicePrimaryKeyMapping)
	sql := "DELETE FROM \"mbus_device\" WHERE \"device_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from mbus_device")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for mbus_device")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q mbusDeviceQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no mbusDeviceQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from mbus_device")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for mbus_device")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MbusDeviceSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(mbusDeviceBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mbusDevicePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"mbus_device\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, mbusDevicePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from mbusDevice slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for mbus_device")
	}

	if len(mbusDeviceAfterDeleteHooks) != 0 {
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
func (o *MbusDevice) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMbusDevice(ctx, exec, o.DeviceID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MbusDeviceSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MbusDeviceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mbusDevicePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"mbus_device\".* FROM \"mbus_device\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, mbusDevicePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MbusDeviceSlice")
	}

	*o = slice

	return nil
}

// MbusDeviceExists checks if the MbusDevice row exists.
func MbusDeviceExists(ctx context.Context, exec boil.ContextExecutor, deviceID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"mbus_device\" where \"device_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, deviceID)
	}
	row := exec.QueryRowContext(ctx, sql, deviceID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if mbus_device exists")
	}

	return exists, nil
}
