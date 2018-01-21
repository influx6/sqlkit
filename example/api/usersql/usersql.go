// Package usersql provides a auto-generated package which contains a sql CRUD API for the specific User struct in package api.
//
//
package usersql

import (
	"context"

	"fmt"

	"errors"

	dsql "database/sql"

	"github.com/influx6/faux/db"

	"github.com/influx6/faux/metrics"

	"github.com/influx6/faux/db/sql"

	"github.com/influx6/faux/db/sql/tables"

	"github.com/gokit/sqlkit/example/api"
)

// errors ...
var (
	ErrNotFound = errors.New("record not found")
)

// mapFields defines a type for a map that exposes a Fields() method.
type mapFields map[string]interface{}

// Fields returns the map itself and provides a method to match the sql.TableField interface.
func (m mapFields) Fields() (map[string]interface{}, error) {
	return m, nil
}

// Validation defines an interface which expose a method to validate a giving type.
type Validation interface {
	Validate() error
}

// UserFields defines an interface which exposes method to return a map of all
// attributes associated with the defined structure as decided by the structure.
type UserFields interface {
	Fields() (map[string]interface{}, error)
}

// UserConsumer defines an interface which accepts a map of data which will be consumed
// into the giving implementing structure as decided by the structure.
type UserConsumer interface {
	Consume(map[string]interface{}) error
}

// UserDB defines a structure which provide DB CRUD operations
// using sql as the underline db.
type UserDB struct {
	col     string
	sx      sql.DB
	dx      *sql.SQL
	metrics metrics.Metrics
	table   db.TableIdentity
}

// New returns a new instance of UserDB.
func New(table string, m metrics.Metrics, sx sql.DB, tm ...tables.TableMigration) *UserDB {
	dx := sql.New(m, sx, tm...)

	return &UserDB{
		sx:      sx,
		dx:      dx,
		col:     table,
		metrics: m,
		table:   db.TableName{Name: table},
	}
}

// Delete attempts to remove the record from the db using the provided publicID.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given api.User struct.
func (mdb *UserDB) Delete(ctx context.Context, publicID string) error {
	defer mdb.metrics.CollectMetrics("UserDB.Delete")

	if isContextExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to delete record"), metrics.With("publicID", publicID), metrics.With("table", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	if err := mdb.dx.Delete(mdb.table, "public_id", publicID); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to delete record"), metrics.With("table", mdb.col), metrics.With("publicID", publicID), metrics.With("error", err.Error()))
		if err == dsql.ErrNoRows {
			return ErrNotFound
		}
		return err
	}

	mdb.metrics.Emit(metrics.Info("Deleted record"), metrics.With("table", mdb.col), metrics.With("publicID", publicID))

	return nil
}

// Create attempts to add the record into the db using the provided instance of the
// api.User.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given User struct.
func (mdb *UserDB) Create(ctx context.Context, elem api.User) error {
	defer mdb.metrics.CollectMetrics("UserDB.Create")

	if isContextExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to create record"), metrics.With("publicID", elem.PublicID), metrics.With("table", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	if validator, ok := interface{}(elem).(Validation); ok {
		if err := validator.Validate(); err != nil {
			mdb.metrics.Emit(metrics.Errorf("Failed to create record: record failed validation"), metrics.With("publicID", elem.PublicID), metrics.With("table", mdb.col), metrics.With("error", err.Error()))
			return err
		}
	}

	content := mapFields(map[string]interface{}{

		"name": elem.Name,

		"public_id": elem.PublicID,
	})
	if err := mdb.dx.Save(mdb.table, content); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to create User record"), metrics.With("table", mdb.col), metrics.With("query", content), metrics.With("error", err.Error()))
		return err
	}

	mdb.metrics.Emit(metrics.Info("Create record"), metrics.With("table", mdb.col), metrics.With("query", content))

	return nil
}

// GetAllByOrder attempts to retrieve all elements from db using provided order and orderby
// values.
func (mdb *UserDB) GetAllByOrder(ctx context.Context, order string, orderby string) ([]api.User, error) {
	defer mdb.metrics.CollectMetrics("UserDB.GetAllByOrder")
	res, _, err := mdb.GetAll(ctx, order, orderby, -1, -1)
	return res, err
}

// GetAll retrieves all records from the db and returns a slice of api.User type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given User struct.
func (mdb *UserDB) GetAll(ctx context.Context, order string, orderby string, page int, responsePerPage int) ([]api.User, int, error) {
	defer mdb.metrics.CollectMetrics("UserDB.GetAll")

	if isContextExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With("table", mdb.col), metrics.With("error", err.Error()))

		return nil, -1, err
	}

	var total int
	var ritems []api.User

	panic("User must implement UserConsumer")

	return ritems, total, nil
}

// GetByField retrieves a record from the db using the field key and value,
// returns the api.User type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given User struct.
func (mdb *UserDB) GetByField(ctx context.Context, key string, value interface{}) (api.User, error) {
	defer mdb.metrics.CollectMetrics("UserDB.GetByField")

	if isContextExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With(key, value), metrics.With("table", mdb.col), metrics.With("error", err.Error()))

		return api.User{}, err
	}

	var elem api.User

	panic("User must implement UserConsumer")

	return elem, nil
}

// Get retrieves a record from the db using the publicID and returns the api.User type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given User struct.
func (mdb *UserDB) Get(ctx context.Context, publicID string) (api.User, error) {
	defer mdb.metrics.CollectMetrics("UserDB.Get")

	if isContextExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to retrieve record"), metrics.With("publicID", publicID), metrics.With("table", mdb.col), metrics.With("error", err.Error()))
		return api.User{}, err
	}

	var elem api.User

	panic("User must implement UserConsumer")

	return elem, nil
}

// Update uses a record from the db using the publicID and returns the api.User type.
// Records using this DB must have a public id value, expressed either by a bson or json tag
// on the given User struct.
func (mdb *UserDB) Update(ctx context.Context, publicID string, elem api.User) error {
	defer mdb.metrics.CollectMetrics("UserDB.Update")

	if isContextExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to finish, context has expired"), metrics.With("table", mdb.col), metrics.With("public_id", publicID), metrics.With("error", err.Error()))
		return err
	}

	if validator, ok := interface{}(elem).(Validation); ok {
		if err := validator.Validate(); err != nil {
			mdb.metrics.Emit(metrics.Errorf("Failed to finish, record failed validation"), metrics.With("table", mdb.col), metrics.With("public_id", publicID), metrics.With("error", err.Error()))
			return err
		}
	}

	data := mapFields(map[string]interface{}{

		"name": elem.Name,

		"public_id": elem.PublicID,
	})
	if err := mdb.dx.Update(mdb.table, data, "public_id", publicID); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to update User record"), metrics.With("query", elem), metrics.With("table", mdb.col), metrics.With("public_id", publicID), metrics.With("error", err.Error()))
		if err == dsql.ErrNoRows {
			return ErrNotFound
		}
		return err
	}

	mdb.metrics.Emit(metrics.Info("Update record"), metrics.With("table", mdb.col), metrics.With("public_id", publicID), metrics.With("query", elem))

	return nil
}

// Exec provides a function which allows the execution of a custom function against the table.
func (mdb *UserDB) Exec(ctx context.Context, fx func(*sql.SQL, sql.DB) error) error {
	defer mdb.metrics.CollectMetrics("UserDB.Exec")

	if isContextExpired(ctx) {
		err := fmt.Errorf("Context has expired")
		mdb.metrics.Emit(metrics.Errorf("Failed to execute operation"), metrics.With("table", mdb.col), metrics.With("error", err.Error()))
		return err
	}

	if err := fx(mdb.dx, mdb.sx); err != nil {
		mdb.metrics.Emit(metrics.Errorf("Failed to execute operation"), metrics.With("table", mdb.col), metrics.With("error", err.Error()))
		if err == dsql.ErrNoRows {
			return ErrNotFound
		}
		return err
	}

	mdb.metrics.Emit(metrics.Info("Operation executed"), metrics.With("table", mdb.col))

	return nil
}

func isContextExpired(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
