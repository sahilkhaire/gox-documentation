---
title: "DB.Open"
package: "db"
import: "github.com/sahilkhaire/gox/db"
gox-doc-version: "11"
---

<SymbolHeader pkg="db" title="DB.Open" node="knex" import-path="github.com/sahilkhaire/gox/db" />
## Overview

Open connects using driver name and DSN.

Part of the **`db`** package — Node.js analog: *knex*.

Method on **`DB`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func Open(driver, dsn string) (*DB, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const db = knex({ client: 'pg', connection: process.env.DATABASE_URL });
```

```go [Standard Go]
sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
```

```go [gox]
import "github.com/sahilkhaire/gox/db"

database, err := db.Open(ctx, "postgres", os.Getenv("DATABASE_URL"))
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/db"

database, err := db.Open(ctx, "postgres", os.Getenv("DATABASE_URL"))
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/db/db-must-open">DB.MustOpen</a>
</div>

← [Back to db package overview](/packages/db/)
