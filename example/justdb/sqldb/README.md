SQLDB API
===================================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/sqlkit/example/justdb/sqldb)](https://goreportcard.com/report/github.com/gokit/sqlkit/example/justdb/sqldb)

SQLDB API is a auto-generated base implementation for executing sql queries through a db instance.

The following method exists for custom operations:

## Exec

```go
Exec(ctx context.Context, fx func(*sql.SQL, sql.DB) error) error
```

