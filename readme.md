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

Generate SQL package-level methods for interacting with the database.

- `@sql_methods`

Generate SQL package-level methods for CRUD interacting with the database.


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


- Generate package-level methods for CRUD with sql as underline db

You annotate any giving struct with `@sql_methods` which marks giving struct has a target for code generation. 

*All struct must have a `PublicID` field.*

Sample below:

```go
// User is a type defining the given user related fields for a given.
//@sql_methods
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

## Interfaces

Sqlkit will generate specific interfaces where each allows the internal functions validate struct validity and are able to get and consume maps containing record details.


```go
type Validate interface{
	Validate() error
}
```

These are required and must be added to target struct as 
they are the means the generated code uses to get data to be saved and sets internal struct's fields:

```go
type UserFields  interface {
	Fields() map[string]interface{}
}

type UserConsumer interface {
	Consume(map[string]interface{}) error
}
```

