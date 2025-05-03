# Golang Migrate

## Install CLI tool

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Create migration file

```bash
migrate create -ext sql -dir internal/database/migrations -seq create_users_table
```

## Run migrations

```bash
migrate -path internal/database/migrations -database postgres://postgres:password@localhost:5432/postgres?sslmode=disable up
```