// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Thermostat is an object representing the database table.
type Thermostat struct {
	ID           int          `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt    time.Time    `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	Timestamp    time.Time    `boil:"timestamp" json:"timestamp" toml:"timestamp" yaml:"timestamp"`
	Provider     null.String  `boil:"provider" json:"provider,omitempty" toml:"provider" yaml:"provider,omitempty"`
	ThermostatID null.String  `boil:"thermostat_id" json:"thermostat_id,omitempty" toml:"thermostat_id" yaml:"thermostat_id,omitempty"`
	TargetCool   null.Float32 `boil:"target_cool" json:"target_cool,omitempty" toml:"target_cool" yaml:"target_cool,omitempty"`
	TargetHeat   null.Float32 `boil:"target_heat" json:"target_heat,omitempty" toml:"target_heat" yaml:"target_heat,omitempty"`
	ActualTemp   null.Float32 `boil:"actual_temp" json:"actual_temp,omitempty" toml:"actual_temp" yaml:"actual_temp,omitempty"`
	Humidity     null.Float32 `boil:"humidity" json:"humidity,omitempty" toml:"humidity" yaml:"humidity,omitempty"`
	IsHeating    null.Bool    `boil:"is_heating" json:"is_heating,omitempty" toml:"is_heating" yaml:"is_heating,omitempty"`
	IsCooling    null.Bool    `boil:"is_cooling" json:"is_cooling,omitempty" toml:"is_cooling" yaml:"is_cooling,omitempty"`

	R *thermostatR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L thermostatL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ThermostatColumns = struct {
	ID           string
	CreatedAt    string
	Timestamp    string
	Provider     string
	ThermostatID string
	TargetCool   string
	TargetHeat   string
	ActualTemp   string
	Humidity     string
	IsHeating    string
	IsCooling    string
}{
	ID:           "id",
	CreatedAt:    "created_at",
	Timestamp:    "timestamp",
	Provider:     "provider",
	ThermostatID: "thermostat_id",
	TargetCool:   "target_cool",
	TargetHeat:   "target_heat",
	ActualTemp:   "actual_temp",
	Humidity:     "humidity",
	IsHeating:    "is_heating",
	IsCooling:    "is_cooling",
}

// Generated where

type whereHelpernull_Bool struct{ field string }

func (w whereHelpernull_Bool) EQ(x null.Bool) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Bool) NEQ(x null.Bool) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Bool) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Bool) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Bool) LT(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Bool) LTE(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Bool) GT(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Bool) GTE(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var ThermostatWhere = struct {
	ID           whereHelperint
	CreatedAt    whereHelpertime_Time
	Timestamp    whereHelpertime_Time
	Provider     whereHelpernull_String
	ThermostatID whereHelpernull_String
	TargetCool   whereHelpernull_Float32
	TargetHeat   whereHelpernull_Float32
	ActualTemp   whereHelpernull_Float32
	Humidity     whereHelpernull_Float32
	IsHeating    whereHelpernull_Bool
	IsCooling    whereHelpernull_Bool
}{
	ID:           whereHelperint{field: `id`},
	CreatedAt:    whereHelpertime_Time{field: `created_at`},
	Timestamp:    whereHelpertime_Time{field: `timestamp`},
	Provider:     whereHelpernull_String{field: `provider`},
	ThermostatID: whereHelpernull_String{field: `thermostat_id`},
	TargetCool:   whereHelpernull_Float32{field: `target_cool`},
	TargetHeat:   whereHelpernull_Float32{field: `target_heat`},
	ActualTemp:   whereHelpernull_Float32{field: `actual_temp`},
	Humidity:     whereHelpernull_Float32{field: `humidity`},
	IsHeating:    whereHelpernull_Bool{field: `is_heating`},
	IsCooling:    whereHelpernull_Bool{field: `is_cooling`},
}

// ThermostatRels is where relationship names are stored.
var ThermostatRels = struct {
}{}

// thermostatR is where relationships are stored.
type thermostatR struct {
}

// NewStruct creates a new relationship struct
func (*thermostatR) NewStruct() *thermostatR {
	return &thermostatR{}
}

// thermostatL is where Load methods for each relationship are stored.
type thermostatL struct{}

var (
	thermostatColumns               = []string{"id", "created_at", "timestamp", "provider", "thermostat_id", "target_cool", "target_heat", "actual_temp", "humidity", "is_heating", "is_cooling"}
	thermostatColumnsWithoutDefault = []string{"timestamp", "provider", "thermostat_id", "target_cool", "target_heat", "actual_temp", "humidity", "is_heating", "is_cooling"}
	thermostatColumnsWithDefault    = []string{"id", "created_at"}
	thermostatPrimaryKeyColumns     = []string{"id"}
)

type (
	// ThermostatSlice is an alias for a slice of pointers to Thermostat.
	// This should generally be used opposed to []Thermostat.
	ThermostatSlice []*Thermostat
	// ThermostatHook is the signature for custom Thermostat hook methods
	ThermostatHook func(context.Context, boil.ContextExecutor, *Thermostat) error

	thermostatQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	thermostatType                 = reflect.TypeOf(&Thermostat{})
	thermostatMapping              = queries.MakeStructMapping(thermostatType)
	thermostatPrimaryKeyMapping, _ = queries.BindMapping(thermostatType, thermostatMapping, thermostatPrimaryKeyColumns)
	thermostatInsertCacheMut       sync.RWMutex
	thermostatInsertCache          = make(map[string]insertCache)
	thermostatUpdateCacheMut       sync.RWMutex
	thermostatUpdateCache          = make(map[string]updateCache)
	thermostatUpsertCacheMut       sync.RWMutex
	thermostatUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var thermostatBeforeInsertHooks []ThermostatHook
var thermostatBeforeUpdateHooks []ThermostatHook
var thermostatBeforeDeleteHooks []ThermostatHook
var thermostatBeforeUpsertHooks []ThermostatHook

var thermostatAfterInsertHooks []ThermostatHook
var thermostatAfterSelectHooks []ThermostatHook
var thermostatAfterUpdateHooks []ThermostatHook
var thermostatAfterDeleteHooks []ThermostatHook
var thermostatAfterUpsertHooks []ThermostatHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Thermostat) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thermostatBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Thermostat) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thermostatBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Thermostat) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thermostatBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Thermostat) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thermostatBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Thermostat) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thermostatAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Thermostat) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thermostatAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Thermostat) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thermostatAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Thermostat) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thermostatAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Thermostat) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range thermostatAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddThermostatHook registers your hook function for all future operations.
func AddThermostatHook(hookPoint boil.HookPoint, thermostatHook ThermostatHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		thermostatBeforeInsertHooks = append(thermostatBeforeInsertHooks, thermostatHook)
	case boil.BeforeUpdateHook:
		thermostatBeforeUpdateHooks = append(thermostatBeforeUpdateHooks, thermostatHook)
	case boil.BeforeDeleteHook:
		thermostatBeforeDeleteHooks = append(thermostatBeforeDeleteHooks, thermostatHook)
	case boil.BeforeUpsertHook:
		thermostatBeforeUpsertHooks = append(thermostatBeforeUpsertHooks, thermostatHook)
	case boil.AfterInsertHook:
		thermostatAfterInsertHooks = append(thermostatAfterInsertHooks, thermostatHook)
	case boil.AfterSelectHook:
		thermostatAfterSelectHooks = append(thermostatAfterSelectHooks, thermostatHook)
	case boil.AfterUpdateHook:
		thermostatAfterUpdateHooks = append(thermostatAfterUpdateHooks, thermostatHook)
	case boil.AfterDeleteHook:
		thermostatAfterDeleteHooks = append(thermostatAfterDeleteHooks, thermostatHook)
	case boil.AfterUpsertHook:
		thermostatAfterUpsertHooks = append(thermostatAfterUpsertHooks, thermostatHook)
	}
}

// One returns a single thermostat record from the query.
func (q thermostatQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Thermostat, error) {
	o := &Thermostat{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db: failed to execute a one query for thermostat")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Thermostat records from the query.
func (q thermostatQuery) All(ctx context.Context, exec boil.ContextExecutor) (ThermostatSlice, error) {
	var o []*Thermostat

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "db: failed to assign all query results to Thermostat slice")
	}

	if len(thermostatAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Thermostat records in the query.
func (q thermostatQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to count thermostat rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q thermostatQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "db: failed to check if thermostat exists")
	}

	return count > 0, nil
}

// Thermostats retrieves all the records using an executor.
func Thermostats(mods ...qm.QueryMod) thermostatQuery {
	mods = append(mods, qm.From("\"prh\".\"thermostat\""))
	return thermostatQuery{NewQuery(mods...)}
}

// FindThermostat retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindThermostat(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Thermostat, error) {
	thermostatObj := &Thermostat{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"prh\".\"thermostat\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, thermostatObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db: unable to select from thermostat")
	}

	return thermostatObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Thermostat) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("db: no thermostat provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(thermostatColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	thermostatInsertCacheMut.RLock()
	cache, cached := thermostatInsertCache[key]
	thermostatInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			thermostatColumns,
			thermostatColumnsWithDefault,
			thermostatColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(thermostatType, thermostatMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(thermostatType, thermostatMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"prh\".\"thermostat\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"prh\".\"thermostat\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "db: unable to insert into thermostat")
	}

	if !cached {
		thermostatInsertCacheMut.Lock()
		thermostatInsertCache[key] = cache
		thermostatInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Thermostat.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Thermostat) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	thermostatUpdateCacheMut.RLock()
	cache, cached := thermostatUpdateCache[key]
	thermostatUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			thermostatColumns,
			thermostatPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("db: unable to update thermostat, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"prh\".\"thermostat\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, thermostatPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(thermostatType, thermostatMapping, append(wl, thermostatPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update thermostat row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by update for thermostat")
	}

	if !cached {
		thermostatUpdateCacheMut.Lock()
		thermostatUpdateCache[key] = cache
		thermostatUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q thermostatQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update all for thermostat")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to retrieve rows affected for thermostat")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ThermostatSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("db: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thermostatPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"prh\".\"thermostat\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, thermostatPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update all in thermostat slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to retrieve rows affected all in update all thermostat")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Thermostat) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("db: no thermostat provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(thermostatColumnsWithDefault, o)

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

	thermostatUpsertCacheMut.RLock()
	cache, cached := thermostatUpsertCache[key]
	thermostatUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			thermostatColumns,
			thermostatColumnsWithDefault,
			thermostatColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			thermostatColumns,
			thermostatPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("db: unable to upsert thermostat, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(thermostatPrimaryKeyColumns))
			copy(conflict, thermostatPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"prh\".\"thermostat\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(thermostatType, thermostatMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(thermostatType, thermostatMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "db: unable to upsert thermostat")
	}

	if !cached {
		thermostatUpsertCacheMut.Lock()
		thermostatUpsertCache[key] = cache
		thermostatUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Thermostat record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Thermostat) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("db: no Thermostat provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), thermostatPrimaryKeyMapping)
	sql := "DELETE FROM \"prh\".\"thermostat\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete from thermostat")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by delete for thermostat")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q thermostatQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("db: no thermostatQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete all from thermostat")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by deleteall for thermostat")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ThermostatSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("db: no Thermostat slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(thermostatBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thermostatPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"prh\".\"thermostat\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, thermostatPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete all from thermostat slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by deleteall for thermostat")
	}

	if len(thermostatAfterDeleteHooks) != 0 {
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
func (o *Thermostat) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindThermostat(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ThermostatSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ThermostatSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), thermostatPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"prh\".\"thermostat\".* FROM \"prh\".\"thermostat\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, thermostatPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "db: unable to reload all in ThermostatSlice")
	}

	*o = slice

	return nil
}

// ThermostatExists checks if the Thermostat row exists.
func ThermostatExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"prh\".\"thermostat\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "db: unable to check if thermostat exists")
	}

	return exists, nil
}
