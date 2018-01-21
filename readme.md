SqlKit
--------
[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/sqlkit)](https://goreportcard.com/report/github.com/gokit/sqlkit)

Sqlkit implements a code generator which automatically generates go packages for sql db implementations for annotated structs.

Due to the nature of the sql, the code gen mainly caters for flag structures, that is where all struct fields are not 
referenced types like list and maps. If you have such, it would be more beneficial to write custom code that handle the 
intricacies of such structs.

## Install

```
go get -u github.com/gokit/sqlkit
```

## Examples

See [Examples](./example) for all usage of provided annotations for code generation with associated results.

## Annotations

- `@sql`

Generate SQL package-level methods for interacting with database.


- `sqlapi`

Generate API package for interacting with CRUD operations for structs.

## Usage

Running the following commands instantly generates all necessary files and packages for giving code gen.

```go
> sqlkit generate
```

## How It works

- Generate sql package with package-level functions


You annotate any target package with `@sql` which marks giving package has a target for code generation. 

```go
// @sql
package users
```


- Generate API with sql as underline db

You annotate any giving struct with `@sqlapi` which marks giving struct has a target for code generation. 

*All struct must have a `PublicID` field.*

Sample below:

```go
// User is a type defining the given user related fields for a given.
//@sqlapi
type User struct {
	Username      string    `json:"username"`
	PublicID      string    `json:"public_id"`
	PrivateID     string    `json:"private_id"`
	Hash          string    `json:"hash"`
	TwoFactorAuth bool      `json:"two_factor_auth"`
	Created       time.Time `json:"created_at"`
	Updated       time.Time `json:"updated_at"`
}
```

SqlKit expects structs to match specific interfaces, which allows it to function as seperate from the declared struct has much has possible, because the interfaces allow to perform serialization and deserialization.

Each interface is included with all generated files.

Sample interfaces to be implemented for `User` struct:

```go
type UserFields  interface {
	Fields() map[string]interface{}
}

type UserConsumer interface {
	Consume(map[string]interface{}) error
}
```

## Customization

If you wish to use a custome name prefix for the config environment variables names generating in the test, then setting 
a attribute of `ENVName` on the attribute will generate a config in the ff format:

```go
// @sql(ENVName => BOB)
```

```go
    config := mdb.Config{
        Mode: mgo.Monotonic,
        DB: os.Getenv("{{.ENVName}}_SQL_TEST_DB"),
        Host: os.Getenv("{{.ENVName}}_SQL_TEST_HOST"),
        User: os.Getenv("{{.ENVName}}_SQL_TEST_USER"),
        AuthDB: os.Getenv("{{.ENVName}}_SQL_TEST_AUTHDB"),
        Password: os.Getenv("{{.ENVName}}_SQL_TEST_PASSWORD"),
    }
```

Will result in:

```go
    config := mdb.Config{
        Mode: mgo.Monotonic,
        DB: os.Getenv("BOB_SQL_TEST_DB"),
        Host: os.Getenv("BOB_SQL_TEST_HOST"),
        User: os.Getenv("BOB_SQL_TEST_USER"),
        AuthDB: os.Getenv("BOB_SQL_TEST_AUTHDB"),
        Password: os.Getenv("BOB_SQL_TEST_PASSWORD"),
    }
```
