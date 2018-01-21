User SQLDB API
===================================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/sqlkit/example/api/usersql)](https://goreportcard.com/report/github.com/gokit/sqlkit/example/api/usersql)

User SQLDB API is a auto-generated CRUD implementation for the `User` in package `github.com/gokit/sqlkit/example/api`.

The following method exists for custom operations:

## Exec

```go
Exec(ctx context.Context, fx func(*sql.SQL, sql.DB) error) error
```

The following methods exists in the generated API as pertaining to CRUD:

## Count

```go
Count(ctx context.Context) (int, error)
```

## Create

```go
Create(ctx context.Context, elem api.User) error
```

## GetByField

```go
GetByField(ctx context.Context, key string, value string)  (api.User,  error)
```

## Get

```go
Get(ctx context.Context, publicID string) (api.User, error)
```

## Get All

```go
GetAll(ctx context.Context) ([]api.User, error)
```

## Update

```go
Update(ctx context.Context, publicID string, elem api.User) error
```

## Delete

```go
Delete(ctx context.Context, publicID string) error
```
