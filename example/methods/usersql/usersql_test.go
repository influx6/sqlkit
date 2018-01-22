package usersql_test

import (
	"os"

	"time"

	"testing"

	"context"

	"github.com/influx6/faux/tests"

	"github.com/influx6/faux/metrics"

	"github.com/influx6/faux/metrics/custom"

	_ "github.com/go-sql-driver/mysql"

	_ "github.com/mattn/go-sqlite3"

	sqldb "github.com/gokit/sqlkit/example/methods/usersql"

	fixtures "github.com/gokit/sqlkit/example/methods/usersql/fixtures"
)

var (
	config = sqldb.Config{
		DB:       os.Getenv("SQL_TEST_DB"),
		User:     os.Getenv("SQL_TEST_USER"),
		Host:     os.Getenv("SQL_TEST_ADDR"),
		Port:     os.Getenv("SQL_TEST_PORT"),
		Driver:   os.Getenv("SQL_TEST_Driver"),
		Password: os.Getenv("SQL_TEST_PASSWORD"),
	}

	db      = sqldb.NewDB(config)
	testCol = "user_test"
)

// TestGetUser validates the retrieval of a User
// record from a sqldb.
func TestGetUser(t *testing.T) {
	events := metrics.New()
	if testing.Verbose() {
		events = metrics.New(custom.StackDisplay(os.Stdout))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	defer sqldb.Delete(ctx, events, db, testCol, elem.PublicID)

	if err := sqldb.Create(ctx, events, db, testCol, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")

	_, err = sqldb.Get(ctx, events, db, testCol, elem.PublicID)
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	defer sqldb.Delete(ctx, events, db, testCol, elem.PublicID)

	if err := sqldb.Create(ctx, events, db, testCol, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")

	records, err := sqldb.GetAll(ctx, events, db, testCol)
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	defer sqldb.Delete(ctx, events, db, testCol, elem.PublicID)

	if err := sqldb.Create(ctx, events, db, testCol, elem); err != nil {
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	defer sqldb.Delete(ctx, events, db, testCol, elem.PublicID)

	if err := sqldb.Create(ctx, events, db, testCol, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")

	elem2, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	elem2.PublicID = elem.PublicID

	if err := sqldb.Update(ctx, events, db, testCol, elem2.PublicID, elem2); err != nil {
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	elem, err := fixtures.LoadUserJSON(fixtures.UserJSON)
	if err != nil {
		tests.Failed("Successfully loaded JSON for User record: %+q.", err)
	}
	tests.Passed("Successfully loaded JSON for User record")

	if err := sqldb.Create(ctx, events, db, testCol, elem); err != nil {
		tests.Failed("Successfully added record for User into db: %+q.", err)
	}
	tests.Passed("Successfully added record for User into db.")

	if err := sqldb.Delete(ctx, events, db, testCol, elem.PublicID); err != nil {
		tests.Failed("Successfully removed record for User into db: %+q.", err)
	}
	tests.Passed("Successfully removed record for User into db.")

	if _, err = sqldb.Get(ctx, events, db, testCol, elem.PublicID); err == nil {
		tests.Failed("Successfully failed to get deleted record for User into db.")
	}
	tests.Passed("Successfully failed to get deleted record for User into db.")
}
