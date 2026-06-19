# Knex → gox/db

Knex query builders map to chainable methods on `gox/db`.

## Connect

```go
import (
    "github.com/sahilkhaire/gox/db"
    _ "github.com/jackc/pgx/v5/stdlib" // or modernc.org/sqlite for SQLite
)

d, err := db.Open("pgx", os.Getenv("DATABASE_URL"))
```

## Queries

| Knex | gox/db |
|------|--------|
| `knex('users').select('*')` | `d.Table("users").Select()` |
| `.where('id', id)` | `.Where("id", "=", id)` |
| `.where({ id, active: true })` | `.WhereEq(map[string]any{...})` |
| `.orderBy('id', 'desc')` | `.OrderBy("id DESC")` |
| `.limit(10).offset(20)` | `.Limit(10).Offset(20)` |
| `.first()` | `.First(ctx, &row)` |
| `.then(rows => ...)` | `.Scan(ctx, &rows)` |

## Writes

| Knex | gox/db |
|------|--------|
| `knex('users').insert({...})` | `d.Insert(ctx, "users", map[string]any{...})` |
| `.update({...}).where({id})` | `d.Update(ctx, "users", set, where)` |
| `.del().where({id})` | `d.Delete(ctx, "users", where)` |
| `knex.transaction(async trx => ...)` | `d.Transaction(ctx, func(tx *db.Tx) error { ... })` |

See `examples/db-crud/main.go` for a SQLite walkthrough.
