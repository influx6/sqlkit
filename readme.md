SqlKit
--------
Sqlkit implements a code generator which automatically generates go packages for sql db implementations for annotated structs.

## Install

```
go get -u github.com/gokit/sqlkit
```

## Usage

Running the following commands instantly generates all necessary files and packages for giving code gen.

```go
> sqlkit generate
```

## How It works

You annotate any giving struct with `@sqlapi` which marks giving struct has a target for code generation. 

It also respects `@associates` annotation which gives extra information to the generator for the following data:

1. What struct to be used as representing a new struct type.
2. What struct contain information for representing the updated struct type.

Sample below:

```go
// User is a type defining the given user related fields for a given.
//@sqlapi
//@associates(@sqlapi, New, NewUser)
//@associates(@sqlapi, Update, UpdateUser)
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