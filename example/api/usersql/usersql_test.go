package usersql_test

import (
	"os"

	"time"

	"testing"

	"context"

	"github.com/influx6/faux/tests"

	"github.com/influx6/faux/db/sql"

	"github.com/influx6/faux/metrics"

	"github.com/influx6/faux/metrics/custom"

	_ "github.com/go-sql-driver/mysql"

	_ "github.com/lib/pq"

	_ "github.com/mattn/go-sqlite3"

	sqldb "github.com/gokit/sqlkit/example/api/usersql"

	fixtures "github.com/gokit/sqlkit/example/api/usersql/fixtures"
)

var (
	config = sql.Config{
		DBName:       os.Getenv("API_SQL_TEST_DB"),
		User:         os.Getenv("API_SQL_TEST_USER"),
		DBIP:         os.Getenv("API_SQL_TEST_ADDR"),
		DBPort:       os.Getenv("API_SQL_TEST_PORT"),
		DBDriver:     os.Getenv("API_SQL_TEST_Driver"),
		UserPassword: os.Getenv("API_SQL_TEST_PASSWORD"),
	}

	testCol = "user_test"
)

// TestGetUser validates the retrieval of a User
// record from a sqldb.
func TestGetUser(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}

	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")

	_, err = api.Get(ctx, elem.PublicID)
	if err != nil {
		tests.Failed("Successfully retrieved stored record for User from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved stored record for User from db.")
}

// TestGetAllUser validates the retrieval of all User
// record from a sqldb.
func TestGetAllUser(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")

	records, err := api.GetAllByOrder(ctx, "asc", "public_id")
	if err != nil {
		tests.Failed("Successfully retrieved all records for User from db: %+q.", err)
	}
	tests.Passed("Successfully retrieved all records for User from db.")

	if len(records) == 0 {
		tests.Failed("Successfully retrieved atleast 1 record for User from db.")
	}
	tests.Passed("Successfully retrieved atleast 1 record for User from db.")
}

// TestUserCreate validates the creation of a User
// record with a sqldb.
func TestUserCreate(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")
}

// TestUserUpdate validates the update of a User
// record with a sqldb.
func TestUserUpdate(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	defer api.Delete(ctx, elem.PublicID)

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")

	elem2, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	elem2.PublicID = elem.PublicID

	if err := api.Update(ctx, elem2.PublicID, elem2); err != nil {
		tests.Failed("Successfully updated record for User into db: %+q.", err)
	}
	tests.Passed("Successfully updated record for User into db.")
}

// TestUserDelete validates the removal of a User
// record from a sqldb.
func TestUserDelete(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}
	api := sqldb.New(testCol, events, sql.NewDB(config, events))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	if err := api.Create(ctx, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")

	if err := api.Delete(ctx, elem.PublicID); err != nil {
		tests.Failed("Successfully removed record for User into db: %+q.", err)
	}
	tests.Passed("Successfully removed record for User into db.")

	if _, err = api.Get(ctx, elem.PublicID); err == nil {
		tests.Failed("Successfully failed to get deleted record for User into db.")
	}
	tests.Passed("Successfully failed to get deleted record for User into db.")
}
