// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

import (
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

// Event is an object representing the database table.
type Event struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	UserID    int64     `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *eventR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L eventL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EventColumns = struct {
	ID        string
	Name      string
	UserID    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Name:      "name",
	UserID:    "user_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var EventTableColumns = struct {
	ID        string
	Name      string
	UserID    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "events.id",
	Name:      "events.name",
	UserID:    "events.user_id",
	CreatedAt: "events.created_at",
	UpdatedAt: "events.updated_at",
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

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var EventWhere = struct {
	ID        whereHelperint64
	Name      whereHelperstring
	UserID    whereHelperint64
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpernull_Time
}{
	ID:        whereHelperint64{field: "`events`.`id`"},
	Name:      whereHelperstring{field: "`events`.`name`"},
	UserID:    whereHelperint64{field: "`events`.`user_id`"},
	CreatedAt: whereHelpertime_Time{field: "`events`.`created_at`"},
	UpdatedAt: whereHelpernull_Time{field: "`events`.`updated_at`"},
}

// EventRels is where relationship names are stored.
var EventRels = struct {
}{}

// eventR is where relationships are stored.
type eventR struct {
}

// NewStruct creates a new relationship struct
func (*eventR) NewStruct() *eventR {
	return &eventR{}
}

// eventL is where Load methods for each relationship are stored.
type eventL struct{}

var (
	eventAllColumns            = []string{"id", "name", "user_id", "created_at", "updated_at"}
	eventColumnsWithoutDefault = []string{"name", "user_id", "created_at", "updated_at"}
	eventColumnsWithDefault    = []string{"id"}
	eventPrimaryKeyColumns     = []string{"id"}
)

type (
	// EventSlice is an alias for a slice of pointers to Event.
	// This should almost always be used instead of []Event.
	EventSlice []*Event
	// EventHook is the signature for custom Event hook methods
	EventHook func(boil.Executor, *Event) error

	eventQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	eventType                 = reflect.TypeOf(&Event{})
	eventMapping              = queries.MakeStructMapping(eventType)
	eventPrimaryKeyMapping, _ = queries.BindMapping(eventType, eventMapping, eventPrimaryKeyColumns)
	eventInsertCacheMut       sync.RWMutex
	eventInsertCache          = make(map[string]insertCache)
	eventUpdateCacheMut       sync.RWMutex
	eventUpdateCache          = make(map[string]updateCache)
	eventUpsertCacheMut       sync.RWMutex
	eventUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var eventBeforeInsertHooks []EventHook
var eventBeforeUpdateHooks []EventHook
var eventBeforeDeleteHooks []EventHook
var eventBeforeUpsertHooks []EventHook

var eventAfterInsertHooks []EventHook
var eventAfterSelectHooks []EventHook
var eventAfterUpdateHooks []EventHook
var eventAfterDeleteHooks []EventHook
var eventAfterUpsertHooks []EventHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Event) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range eventBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Event) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range eventBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Event) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range eventBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Event) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range eventBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Event) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range eventAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Event) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range eventAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Event) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range eventAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Event) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range eventAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Event) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range eventAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEventHook registers your hook function for all future operations.
func AddEventHook(hookPoint boil.HookPoint, eventHook EventHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		eventBeforeInsertHooks = append(eventBeforeInsertHooks, eventHook)
	case boil.BeforeUpdateHook:
		eventBeforeUpdateHooks = append(eventBeforeUpdateHooks, eventHook)
	case boil.BeforeDeleteHook:
		eventBeforeDeleteHooks = append(eventBeforeDeleteHooks, eventHook)
	case boil.BeforeUpsertHook:
		eventBeforeUpsertHooks = append(eventBeforeUpsertHooks, eventHook)
	case boil.AfterInsertHook:
		eventAfterInsertHooks = append(eventAfterInsertHooks, eventHook)
	case boil.AfterSelectHook:
		eventAfterSelectHooks = append(eventAfterSelectHooks, eventHook)
	case boil.AfterUpdateHook:
		eventAfterUpdateHooks = append(eventAfterUpdateHooks, eventHook)
	case boil.AfterDeleteHook:
		eventAfterDeleteHooks = append(eventAfterDeleteHooks, eventHook)
	case boil.AfterUpsertHook:
		eventAfterUpsertHooks = append(eventAfterUpsertHooks, eventHook)
	}
}

// OneG returns a single event record from the query using the global executor.
func (q eventQuery) OneG() (*Event, error) {
	return q.One(boil.GetDB())
}

// OneGP returns a single event record from the query using the global executor, and panics on error.
func (q eventQuery) OneGP() *Event {
	o, err := q.One(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// OneP returns a single event record from the query, and panics on error.
func (q eventQuery) OneP(exec boil.Executor) *Event {
	o, err := q.One(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single event record from the query.
func (q eventQuery) One(exec boil.Executor) (*Event, error) {
	o := &Event{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db: failed to execute a one query for events")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Event records from the query using the global executor.
func (q eventQuery) AllG() (EventSlice, error) {
	return q.All(boil.GetDB())
}

// AllGP returns all Event records from the query using the global executor, and panics on error.
func (q eventQuery) AllGP() EventSlice {
	o, err := q.All(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// AllP returns all Event records from the query, and panics on error.
func (q eventQuery) AllP(exec boil.Executor) EventSlice {
	o, err := q.All(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Event records from the query.
func (q eventQuery) All(exec boil.Executor) (EventSlice, error) {
	var o []*Event

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "db: failed to assign all query results to Event slice")
	}

	if len(eventAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Event records in the query, and panics on error.
func (q eventQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// CountGP returns the count of all Event records in the query using the global executor, and panics on error.
func (q eventQuery) CountGP() int64 {
	c, err := q.Count(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// CountP returns the count of all Event records in the query, and panics on error.
func (q eventQuery) CountP(exec boil.Executor) int64 {
	c, err := q.Count(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Event records in the query.
func (q eventQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to count events rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q eventQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// ExistsGP checks if the row exists in the table using the global executor, and panics on error.
func (q eventQuery) ExistsGP() bool {
	e, err := q.Exists(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// ExistsP checks if the row exists in the table, and panics on error.
func (q eventQuery) ExistsP(exec boil.Executor) bool {
	e, err := q.Exists(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q eventQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "db: failed to check if events exists")
	}

	return count > 0, nil
}

// Events retrieves all the records using an executor.
func Events(mods ...qm.QueryMod) eventQuery {
	mods = append(mods, qm.From("`events`"))
	return eventQuery{NewQuery(mods...)}
}

// FindEventG retrieves a single record by ID.
func FindEventG(iD int64, selectCols ...string) (*Event, error) {
	return FindEvent(boil.GetDB(), iD, selectCols...)
}

// FindEventP retrieves a single record by ID with an executor, and panics on error.
func FindEventP(exec boil.Executor, iD int64, selectCols ...string) *Event {
	retobj, err := FindEvent(exec, iD, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindEventGP retrieves a single record by ID, and panics on error.
func FindEventGP(iD int64, selectCols ...string) *Event {
	retobj, err := FindEvent(boil.GetDB(), iD, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindEvent retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEvent(exec boil.Executor, iD int64, selectCols ...string) (*Event, error) {
	eventObj := &Event{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `events` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, eventObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db: unable to select from events")
	}

	if err = eventObj.doAfterSelectHooks(exec); err != nil {
		return eventObj, err
	}

	return eventObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Event) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Event) InsertP(exec boil.Executor, columns boil.Columns) {
	if err := o.Insert(exec, columns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Event) InsertGP(columns boil.Columns) {
	if err := o.Insert(boil.GetDB(), columns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Event) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("db: no events provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if queries.MustTime(o.UpdatedAt).IsZero() {
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(eventColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	eventInsertCacheMut.RLock()
	cache, cached := eventInsertCache[key]
	eventInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			eventAllColumns,
			eventColumnsWithDefault,
			eventColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(eventType, eventMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(eventType, eventMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `events` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `events` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `events` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, eventPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "db: unable to insert into events")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == eventMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}
	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "db: unable to populate default values for events")
	}

CacheNoHooks:
	if !cached {
		eventInsertCacheMut.Lock()
		eventInsertCache[key] = cache
		eventInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Event record using the global executor.
// See Update for more documentation.
func (o *Event) UpdateG(columns boil.Columns) (int64, error) {
	return o.Update(boil.GetDB(), columns)
}

// UpdateP uses an executor to update the Event, and panics on error.
// See Update for more documentation.
func (o *Event) UpdateP(exec boil.Executor, columns boil.Columns) int64 {
	rowsAff, err := o.Update(exec, columns)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// UpdateGP a single Event record using the global executor. Panics on error.
// See Update for more documentation.
func (o *Event) UpdateGP(columns boil.Columns) int64 {
	rowsAff, err := o.Update(boil.GetDB(), columns)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// Update uses an executor to update the Event.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Event) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	queries.SetScanner(&o.UpdatedAt, currTime)

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	eventUpdateCacheMut.RLock()
	cache, cached := eventUpdateCache[key]
	eventUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			eventAllColumns,
			eventPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("db: unable to update events, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `events` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, eventPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(eventType, eventMapping, append(wl, eventPrimaryKeyColumns...))
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
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update events row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by update for events")
	}

	if !cached {
		eventUpdateCacheMut.Lock()
		eventUpdateCache[key] = cache
		eventUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q eventQuery) UpdateAllP(exec boil.Executor, cols M) int64 {
	rowsAff, err := q.UpdateAll(exec, cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// UpdateAllG updates all rows with the specified column values.
func (q eventQuery) UpdateAllG(cols M) (int64, error) {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q eventQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update all for events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to retrieve rows affected for events")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o EventSlice) UpdateAllG(cols M) (int64, error) {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o EventSlice) UpdateAllGP(cols M) int64 {
	rowsAff, err := o.UpdateAll(boil.GetDB(), cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o EventSlice) UpdateAllP(exec boil.Executor, cols M) int64 {
	rowsAff, err := o.UpdateAll(exec, cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EventSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `events` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, eventPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update all in event slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to retrieve rows affected all in update all event")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Event) UpsertG(updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateColumns, insertColumns)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Event) UpsertGP(updateColumns, insertColumns boil.Columns) {
	if err := o.Upsert(boil.GetDB(), updateColumns, insertColumns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Event) UpsertP(exec boil.Executor, updateColumns, insertColumns boil.Columns) {
	if err := o.Upsert(exec, updateColumns, insertColumns); err != nil {
		panic(boil.WrapErr(err))
	}
}

var mySQLEventUniqueColumns = []string{
	"id",
	"user_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Event) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("db: no events provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	queries.SetScanner(&o.UpdatedAt, currTime)

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(eventColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLEventUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	eventUpsertCacheMut.RLock()
	cache, cached := eventUpsertCache[key]
	eventUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			eventAllColumns,
			eventColumnsWithDefault,
			eventColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			eventAllColumns,
			eventPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("db: unable to upsert events, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`events`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `events` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(eventType, eventMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(eventType, eventMapping, ret)
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
	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "db: unable to upsert for events")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == eventMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(eventType, eventMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "db: unable to retrieve unique values for events")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}
	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "db: unable to populate default values for events")
	}

CacheNoHooks:
	if !cached {
		eventUpsertCacheMut.Lock()
		eventUpsertCache[key] = cache
		eventUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteG deletes a single Event record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Event) DeleteG() (int64, error) {
	return o.Delete(boil.GetDB())
}

// DeleteP deletes a single Event record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Event) DeleteP(exec boil.Executor) int64 {
	rowsAff, err := o.Delete(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// DeleteGP deletes a single Event record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Event) DeleteGP() int64 {
	rowsAff, err := o.Delete(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// Delete deletes a single Event record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Event) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("db: no Event provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), eventPrimaryKeyMapping)
	sql := "DELETE FROM `events` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete from events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by delete for events")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q eventQuery) DeleteAllG() (int64, error) {
	return q.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows, and panics on error.
func (q eventQuery) DeleteAllP(exec boil.Executor) int64 {
	rowsAff, err := q.DeleteAll(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// DeleteAll deletes all matching rows.
func (q eventQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("db: no eventQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete all from events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by deleteall for events")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o EventSlice) DeleteAllG() (int64, error) {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o EventSlice) DeleteAllP(exec boil.Executor) int64 {
	rowsAff, err := o.DeleteAll(exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o EventSlice) DeleteAllGP() int64 {
	rowsAff, err := o.DeleteAll(boil.GetDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EventSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(eventBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `events` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, eventPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete all from event slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by deleteall for events")
	}

	if len(eventAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Event) ReloadG() error {
	if o == nil {
		return errors.New("db: no Event provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Event) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Event) ReloadGP() {
	if err := o.Reload(boil.GetDB()); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Event) Reload(exec boil.Executor) error {
	ret, err := FindEvent(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EventSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("db: empty EventSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *EventSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *EventSlice) ReloadAllGP() {
	if err := o.ReloadAll(boil.GetDB()); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EventSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EventSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `events`.* FROM `events` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, eventPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "db: unable to reload all in EventSlice")
	}

	*o = slice

	return nil
}

// EventExistsG checks if the Event row exists.
func EventExistsG(iD int64) (bool, error) {
	return EventExists(boil.GetDB(), iD)
}

// EventExistsP checks if the Event row exists. Panics on error.
func EventExistsP(exec boil.Executor, iD int64) bool {
	e, err := EventExists(exec, iD)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// EventExistsGP checks if the Event row exists. Panics on error.
func EventExistsGP(iD int64) bool {
	e, err := EventExists(boil.GetDB(), iD)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// EventExists checks if the Event row exists.
func EventExists(exec boil.Executor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `events` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "db: unable to check if events exists")
	}

	return exists, nil
}
