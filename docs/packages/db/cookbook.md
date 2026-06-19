---
title: "db Cookbook"
package: "db"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: knex</span><span class="api-badge import">github.com/sahilkhaire/gox/db</span></div>
# db Cookbook

Knex-style chainable SQL on sqlx.

```go
import "github.com/sahilkhaire/gox/db"

database, err := db.Open(ctx, "sqlite", ":memory:")
defer database.Close()

var user struct {
    ID   int    `db:"id"`
    Name string `db:"name"`
}
err = database.From("users").WhereEq("id", 1).First(ctx, &user)

err = database.Transaction(ctx, func(tx *db.Tx) error {
    _, err := tx.Insert(ctx, "users", map[string]any{"name": "Ada"})
    return err
})
```

Escape hatch: **database.SQL** exposes the underlying *sqlx.DB*.

← [Back to db overview](/packages/db/)
