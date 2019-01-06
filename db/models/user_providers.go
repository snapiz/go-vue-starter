// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
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
	"github.com/volatiletech/sqlboiler/strmangle"
)

// UserProvider is an object representing the database table.
type UserProvider struct {
	UserID     int       `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Provider   string    `boil:"provider" json:"provider" toml:"provider" yaml:"provider"`
	ProviderID string    `boil:"provider_id" json:"provider_id" toml:"provider_id" yaml:"provider_id"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *userProviderR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userProviderL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserProviderColumns = struct {
	UserID     string
	Provider   string
	ProviderID string
	CreatedAt  string
	UpdatedAt  string
}{
	UserID:     "user_id",
	Provider:   "provider",
	ProviderID: "provider_id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// UserProviderRels is where relationship names are stored.
var UserProviderRels = struct {
	User string
}{
	User: "User",
}

// userProviderR is where relationships are stored.
type userProviderR struct {
	User *User
}

// NewStruct creates a new relationship struct
func (*userProviderR) NewStruct() *userProviderR {
	return &userProviderR{}
}

// userProviderL is where Load methods for each relationship are stored.
type userProviderL struct{}

var (
	userProviderColumns               = []string{"user_id", "provider", "provider_id", "created_at", "updated_at"}
	userProviderColumnsWithoutDefault = []string{"user_id", "provider", "provider_id", "updated_at"}
	userProviderColumnsWithDefault    = []string{"created_at"}
	userProviderPrimaryKeyColumns     = []string{"provider", "provider_id"}
)

type (
	// UserProviderSlice is an alias for a slice of pointers to UserProvider.
	// This should generally be used opposed to []UserProvider.
	UserProviderSlice []*UserProvider
	// UserProviderHook is the signature for custom UserProvider hook methods
	UserProviderHook func(boil.Executor, *UserProvider) error

	userProviderQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userProviderType                 = reflect.TypeOf(&UserProvider{})
	userProviderMapping              = queries.MakeStructMapping(userProviderType)
	userProviderPrimaryKeyMapping, _ = queries.BindMapping(userProviderType, userProviderMapping, userProviderPrimaryKeyColumns)
	userProviderInsertCacheMut       sync.RWMutex
	userProviderInsertCache          = make(map[string]insertCache)
	userProviderUpdateCacheMut       sync.RWMutex
	userProviderUpdateCache          = make(map[string]updateCache)
	userProviderUpsertCacheMut       sync.RWMutex
	userProviderUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var userProviderBeforeInsertHooks []UserProviderHook
var userProviderBeforeUpdateHooks []UserProviderHook
var userProviderBeforeDeleteHooks []UserProviderHook
var userProviderBeforeUpsertHooks []UserProviderHook

var userProviderAfterInsertHooks []UserProviderHook
var userProviderAfterSelectHooks []UserProviderHook
var userProviderAfterUpdateHooks []UserProviderHook
var userProviderAfterDeleteHooks []UserProviderHook
var userProviderAfterUpsertHooks []UserProviderHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserProvider) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userProviderBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserProvider) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range userProviderBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserProvider) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range userProviderBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserProvider) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userProviderBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserProvider) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userProviderAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserProvider) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range userProviderAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserProvider) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range userProviderAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserProvider) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range userProviderAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserProvider) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range userProviderAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserProviderHook registers your hook function for all future operations.
func AddUserProviderHook(hookPoint boil.HookPoint, userProviderHook UserProviderHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		userProviderBeforeInsertHooks = append(userProviderBeforeInsertHooks, userProviderHook)
	case boil.BeforeUpdateHook:
		userProviderBeforeUpdateHooks = append(userProviderBeforeUpdateHooks, userProviderHook)
	case boil.BeforeDeleteHook:
		userProviderBeforeDeleteHooks = append(userProviderBeforeDeleteHooks, userProviderHook)
	case boil.BeforeUpsertHook:
		userProviderBeforeUpsertHooks = append(userProviderBeforeUpsertHooks, userProviderHook)
	case boil.AfterInsertHook:
		userProviderAfterInsertHooks = append(userProviderAfterInsertHooks, userProviderHook)
	case boil.AfterSelectHook:
		userProviderAfterSelectHooks = append(userProviderAfterSelectHooks, userProviderHook)
	case boil.AfterUpdateHook:
		userProviderAfterUpdateHooks = append(userProviderAfterUpdateHooks, userProviderHook)
	case boil.AfterDeleteHook:
		userProviderAfterDeleteHooks = append(userProviderAfterDeleteHooks, userProviderHook)
	case boil.AfterUpsertHook:
		userProviderAfterUpsertHooks = append(userProviderAfterUpsertHooks, userProviderHook)
	}
}

// OneG returns a single userProvider record from the query using the global executor.
func (q userProviderQuery) OneG() (*UserProvider, error) {
	return q.One(boil.GetDB())
}

// One returns a single userProvider record from the query.
func (q userProviderQuery) One(exec boil.Executor) (*UserProvider, error) {
	o := &UserProvider{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for user_providers")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all UserProvider records from the query using the global executor.
func (q userProviderQuery) AllG() (UserProviderSlice, error) {
	return q.All(boil.GetDB())
}

// All returns all UserProvider records from the query.
func (q userProviderQuery) All(exec boil.Executor) (UserProviderSlice, error) {
	var o []*UserProvider

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UserProvider slice")
	}

	if len(userProviderAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all UserProvider records in the query, and panics on error.
func (q userProviderQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all UserProvider records in the query.
func (q userProviderQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count user_providers rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q userProviderQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q userProviderQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if user_providers exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *UserProvider) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userProviderL) LoadUser(e boil.Executor, singular bool, maybeUserProvider interface{}, mods queries.Applicator) error {
	var slice []*UserProvider
	var object *UserProvider

	if singular {
		object = maybeUserProvider.(*UserProvider)
	} else {
		slice = *maybeUserProvider.(*[]*UserProvider)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userProviderR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userProviderR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	query := NewQuery(qm.From(`users`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userProviderAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.UserProviders = append(foreign.R.UserProviders, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UserProviders = append(foreign.R.UserProviders, local)
				break
			}
		}
	}

	return nil
}

// SetUserG of the userProvider to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserProviders.
// Uses the global database handle.
func (o *UserProvider) SetUserG(insert bool, related *User) error {
	return o.SetUser(boil.GetDB(), insert, related)
}

// SetUser of the userProvider to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserProviders.
func (o *UserProvider) SetUser(exec boil.Executor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_providers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, userProviderPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.Provider, o.ProviderID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &userProviderR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			UserProviders: UserProviderSlice{o},
		}
	} else {
		related.R.UserProviders = append(related.R.UserProviders, o)
	}

	return nil
}

// UserProviders retrieves all the records using an executor.
func UserProviders(mods ...qm.QueryMod) userProviderQuery {
	mods = append(mods, qm.From("\"user_providers\""))
	return userProviderQuery{NewQuery(mods...)}
}

// FindUserProviderG retrieves a single record by ID.
func FindUserProviderG(provider string, providerID string, selectCols ...string) (*UserProvider, error) {
	return FindUserProvider(boil.GetDB(), provider, providerID, selectCols...)
}

// FindUserProvider retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserProvider(exec boil.Executor, provider string, providerID string, selectCols ...string) (*UserProvider, error) {
	userProviderObj := &UserProvider{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user_providers\" where \"provider\"=$1 AND \"provider_id\"=$2", sel,
	)

	q := queries.Raw(query, provider, providerID)

	err := q.Bind(nil, exec, userProviderObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from user_providers")
	}

	return userProviderObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *UserProvider) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserProvider) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_providers provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userProviderColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userProviderInsertCacheMut.RLock()
	cache, cached := userProviderInsertCache[key]
	userProviderInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userProviderColumns,
			userProviderColumnsWithDefault,
			userProviderColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userProviderType, userProviderMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userProviderType, userProviderMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"user_providers\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"user_providers\" %sDEFAULT VALUES%s"
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
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into user_providers")
	}

	if !cached {
		userProviderInsertCacheMut.Lock()
		userProviderInsertCache[key] = cache
		userProviderInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single UserProvider record using the global executor.
// See Update for more documentation.
func (o *UserProvider) UpdateG(columns boil.Columns) (int64, error) {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the UserProvider.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserProvider) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userProviderUpdateCacheMut.RLock()
	cache, cached := userProviderUpdateCache[key]
	userProviderUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userProviderColumns,
			userProviderPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update user_providers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user_providers\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userProviderPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userProviderType, userProviderMapping, append(wl, userProviderPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update user_providers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for user_providers")
	}

	if !cached {
		userProviderUpdateCacheMut.Lock()
		userProviderUpdateCache[key] = cache
		userProviderUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q userProviderQuery) UpdateAllG(cols M) (int64, error) {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q userProviderQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for user_providers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for user_providers")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o UserProviderSlice) UpdateAllG(cols M) (int64, error) {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserProviderSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userProviderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"user_providers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userProviderPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in userProvider slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all userProvider")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *UserProvider) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserProvider) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_providers provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userProviderColumnsWithDefault, o)

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

	userProviderUpsertCacheMut.RLock()
	cache, cached := userProviderUpsertCache[key]
	userProviderUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userProviderColumns,
			userProviderColumnsWithDefault,
			userProviderColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			userProviderColumns,
			userProviderPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert user_providers, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(userProviderPrimaryKeyColumns))
			copy(conflict, userProviderPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"user_providers\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(userProviderType, userProviderMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userProviderType, userProviderMapping, ret)
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
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert user_providers")
	}

	if !cached {
		userProviderUpsertCacheMut.Lock()
		userProviderUpsertCache[key] = cache
		userProviderUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteG deletes a single UserProvider record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *UserProvider) DeleteG() (int64, error) {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single UserProvider record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserProvider) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UserProvider provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userProviderPrimaryKeyMapping)
	sql := "DELETE FROM \"user_providers\" WHERE \"provider\"=$1 AND \"provider_id\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from user_providers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for user_providers")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userProviderQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userProviderQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from user_providers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_providers")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o UserProviderSlice) DeleteAllG() (int64, error) {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserProviderSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UserProvider slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(userProviderBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userProviderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"user_providers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userProviderPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from userProvider slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_providers")
	}

	if len(userProviderAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *UserProvider) ReloadG() error {
	if o == nil {
		return errors.New("models: no UserProvider provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *UserProvider) Reload(exec boil.Executor) error {
	ret, err := FindUserProvider(exec, o.Provider, o.ProviderID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserProviderSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty UserProviderSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserProviderSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserProviderSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userProviderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"user_providers\".* FROM \"user_providers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userProviderPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserProviderSlice")
	}

	*o = slice

	return nil
}

// UserProviderExistsG checks if the UserProvider row exists.
func UserProviderExistsG(provider string, providerID string) (bool, error) {
	return UserProviderExists(boil.GetDB(), provider, providerID)
}

// UserProviderExists checks if the UserProvider row exists.
func UserProviderExists(exec boil.Executor, provider string, providerID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"user_providers\" where \"provider\"=$1 AND \"provider_id\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, provider, providerID)
	}

	row := exec.QueryRow(sql, provider, providerID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if user_providers exists")
	}

	return exists, nil
}
