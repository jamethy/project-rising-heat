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

// Weather is an object representing the database table.
type Weather struct {
	ID            int          `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt     time.Time    `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	Timestamp     time.Time    `boil:"timestamp" json:"timestamp" toml:"timestamp" yaml:"timestamp"`
	Provider      null.String  `boil:"provider" json:"provider,omitempty" toml:"provider" yaml:"provider,omitempty"`
	Temperature   null.Float32 `boil:"temperature" json:"temperature,omitempty" toml:"temperature" yaml:"temperature,omitempty"`
	FeelsLike     null.Float32 `boil:"feels_like" json:"feels_like,omitempty" toml:"feels_like" yaml:"feels_like,omitempty"`
	Pressure      null.Float32 `boil:"pressure" json:"pressure,omitempty" toml:"pressure" yaml:"pressure,omitempty"`
	Humidity      null.Float32 `boil:"humidity" json:"humidity,omitempty" toml:"humidity" yaml:"humidity,omitempty"`
	WindSpeed     null.Float32 `boil:"wind_speed" json:"wind_speed,omitempty" toml:"wind_speed" yaml:"wind_speed,omitempty"`
	WindDirection null.Float32 `boil:"wind_direction" json:"wind_direction,omitempty" toml:"wind_direction" yaml:"wind_direction,omitempty"`
	Clouds        null.Float32 `boil:"clouds" json:"clouds,omitempty" toml:"clouds" yaml:"clouds,omitempty"`

	R *weatherR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L weatherL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var WeatherColumns = struct {
	ID            string
	CreatedAt     string
	Timestamp     string
	Provider      string
	Temperature   string
	FeelsLike     string
	Pressure      string
	Humidity      string
	WindSpeed     string
	WindDirection string
	Clouds        string
}{
	ID:            "id",
	CreatedAt:     "created_at",
	Timestamp:     "timestamp",
	Provider:      "provider",
	Temperature:   "temperature",
	FeelsLike:     "feels_like",
	Pressure:      "pressure",
	Humidity:      "humidity",
	WindSpeed:     "wind_speed",
	WindDirection: "wind_direction",
	Clouds:        "clouds",
}

// Generated where

var WeatherWhere = struct {
	ID            whereHelperint
	CreatedAt     whereHelpertime_Time
	Timestamp     whereHelpertime_Time
	Provider      whereHelpernull_String
	Temperature   whereHelpernull_Float32
	FeelsLike     whereHelpernull_Float32
	Pressure      whereHelpernull_Float32
	Humidity      whereHelpernull_Float32
	WindSpeed     whereHelpernull_Float32
	WindDirection whereHelpernull_Float32
	Clouds        whereHelpernull_Float32
}{
	ID:            whereHelperint{field: `id`},
	CreatedAt:     whereHelpertime_Time{field: `created_at`},
	Timestamp:     whereHelpertime_Time{field: `timestamp`},
	Provider:      whereHelpernull_String{field: `provider`},
	Temperature:   whereHelpernull_Float32{field: `temperature`},
	FeelsLike:     whereHelpernull_Float32{field: `feels_like`},
	Pressure:      whereHelpernull_Float32{field: `pressure`},
	Humidity:      whereHelpernull_Float32{field: `humidity`},
	WindSpeed:     whereHelpernull_Float32{field: `wind_speed`},
	WindDirection: whereHelpernull_Float32{field: `wind_direction`},
	Clouds:        whereHelpernull_Float32{field: `clouds`},
}

// WeatherRels is where relationship names are stored.
var WeatherRels = struct {
}{}

// weatherR is where relationships are stored.
type weatherR struct {
}

// NewStruct creates a new relationship struct
func (*weatherR) NewStruct() *weatherR {
	return &weatherR{}
}

// weatherL is where Load methods for each relationship are stored.
type weatherL struct{}

var (
	weatherColumns               = []string{"id", "created_at", "timestamp", "provider", "temperature", "feels_like", "pressure", "humidity", "wind_speed", "wind_direction", "clouds"}
	weatherColumnsWithoutDefault = []string{"timestamp", "provider", "temperature", "feels_like", "pressure", "humidity", "wind_speed", "wind_direction", "clouds"}
	weatherColumnsWithDefault    = []string{"id", "created_at"}
	weatherPrimaryKeyColumns     = []string{"id"}
)

type (
	// WeatherSlice is an alias for a slice of pointers to Weather.
	// This should generally be used opposed to []Weather.
	WeatherSlice []*Weather
	// WeatherHook is the signature for custom Weather hook methods
	WeatherHook func(context.Context, boil.ContextExecutor, *Weather) error

	weatherQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	weatherType                 = reflect.TypeOf(&Weather{})
	weatherMapping              = queries.MakeStructMapping(weatherType)
	weatherPrimaryKeyMapping, _ = queries.BindMapping(weatherType, weatherMapping, weatherPrimaryKeyColumns)
	weatherInsertCacheMut       sync.RWMutex
	weatherInsertCache          = make(map[string]insertCache)
	weatherUpdateCacheMut       sync.RWMutex
	weatherUpdateCache          = make(map[string]updateCache)
	weatherUpsertCacheMut       sync.RWMutex
	weatherUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var weatherBeforeInsertHooks []WeatherHook
var weatherBeforeUpdateHooks []WeatherHook
var weatherBeforeDeleteHooks []WeatherHook
var weatherBeforeUpsertHooks []WeatherHook

var weatherAfterInsertHooks []WeatherHook
var weatherAfterSelectHooks []WeatherHook
var weatherAfterUpdateHooks []WeatherHook
var weatherAfterDeleteHooks []WeatherHook
var weatherAfterUpsertHooks []WeatherHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Weather) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range weatherBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Weather) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range weatherBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Weather) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range weatherBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Weather) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range weatherBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Weather) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range weatherAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Weather) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range weatherAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Weather) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range weatherAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Weather) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range weatherAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Weather) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range weatherAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddWeatherHook registers your hook function for all future operations.
func AddWeatherHook(hookPoint boil.HookPoint, weatherHook WeatherHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		weatherBeforeInsertHooks = append(weatherBeforeInsertHooks, weatherHook)
	case boil.BeforeUpdateHook:
		weatherBeforeUpdateHooks = append(weatherBeforeUpdateHooks, weatherHook)
	case boil.BeforeDeleteHook:
		weatherBeforeDeleteHooks = append(weatherBeforeDeleteHooks, weatherHook)
	case boil.BeforeUpsertHook:
		weatherBeforeUpsertHooks = append(weatherBeforeUpsertHooks, weatherHook)
	case boil.AfterInsertHook:
		weatherAfterInsertHooks = append(weatherAfterInsertHooks, weatherHook)
	case boil.AfterSelectHook:
		weatherAfterSelectHooks = append(weatherAfterSelectHooks, weatherHook)
	case boil.AfterUpdateHook:
		weatherAfterUpdateHooks = append(weatherAfterUpdateHooks, weatherHook)
	case boil.AfterDeleteHook:
		weatherAfterDeleteHooks = append(weatherAfterDeleteHooks, weatherHook)
	case boil.AfterUpsertHook:
		weatherAfterUpsertHooks = append(weatherAfterUpsertHooks, weatherHook)
	}
}

// One returns a single weather record from the query.
func (q weatherQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Weather, error) {
	o := &Weather{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db: failed to execute a one query for weather")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Weather records from the query.
func (q weatherQuery) All(ctx context.Context, exec boil.ContextExecutor) (WeatherSlice, error) {
	var o []*Weather

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "db: failed to assign all query results to Weather slice")
	}

	if len(weatherAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Weather records in the query.
func (q weatherQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to count weather rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q weatherQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "db: failed to check if weather exists")
	}

	return count > 0, nil
}

// Weathers retrieves all the records using an executor.
func Weathers(mods ...qm.QueryMod) weatherQuery {
	mods = append(mods, qm.From("\"prh\".\"weather\""))
	return weatherQuery{NewQuery(mods...)}
}

// FindWeather retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindWeather(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Weather, error) {
	weatherObj := &Weather{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"prh\".\"weather\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, weatherObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db: unable to select from weather")
	}

	return weatherObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Weather) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("db: no weather provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(weatherColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	weatherInsertCacheMut.RLock()
	cache, cached := weatherInsertCache[key]
	weatherInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			weatherColumns,
			weatherColumnsWithDefault,
			weatherColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(weatherType, weatherMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(weatherType, weatherMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"prh\".\"weather\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"prh\".\"weather\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "db: unable to insert into weather")
	}

	if !cached {
		weatherInsertCacheMut.Lock()
		weatherInsertCache[key] = cache
		weatherInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Weather.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Weather) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	weatherUpdateCacheMut.RLock()
	cache, cached := weatherUpdateCache[key]
	weatherUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			weatherColumns,
			weatherPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("db: unable to update weather, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"prh\".\"weather\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, weatherPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(weatherType, weatherMapping, append(wl, weatherPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "db: unable to update weather row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by update for weather")
	}

	if !cached {
		weatherUpdateCacheMut.Lock()
		weatherUpdateCache[key] = cache
		weatherUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q weatherQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update all for weather")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to retrieve rows affected for weather")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o WeatherSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), weatherPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"prh\".\"weather\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, weatherPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update all in weather slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to retrieve rows affected all in update all weather")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Weather) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("db: no weather provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(weatherColumnsWithDefault, o)

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

	weatherUpsertCacheMut.RLock()
	cache, cached := weatherUpsertCache[key]
	weatherUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			weatherColumns,
			weatherColumnsWithDefault,
			weatherColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			weatherColumns,
			weatherPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("db: unable to upsert weather, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(weatherPrimaryKeyColumns))
			copy(conflict, weatherPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"prh\".\"weather\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(weatherType, weatherMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(weatherType, weatherMapping, ret)
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
		return errors.Wrap(err, "db: unable to upsert weather")
	}

	if !cached {
		weatherUpsertCacheMut.Lock()
		weatherUpsertCache[key] = cache
		weatherUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Weather record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Weather) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("db: no Weather provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), weatherPrimaryKeyMapping)
	sql := "DELETE FROM \"prh\".\"weather\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete from weather")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by delete for weather")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q weatherQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("db: no weatherQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete all from weather")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by deleteall for weather")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o WeatherSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("db: no Weather slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(weatherBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), weatherPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"prh\".\"weather\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, weatherPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete all from weather slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by deleteall for weather")
	}

	if len(weatherAfterDeleteHooks) != 0 {
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
func (o *Weather) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindWeather(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *WeatherSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := WeatherSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), weatherPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"prh\".\"weather\".* FROM \"prh\".\"weather\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, weatherPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "db: unable to reload all in WeatherSlice")
	}

	*o = slice

	return nil
}

// WeatherExists checks if the Weather row exists.
func WeatherExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"prh\".\"weather\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "db: unable to check if weather exists")
	}

	return exists, nil
}