---
title: "DB.Open"
package: "db"
import: "github.com/sahilkhaire/gox/db"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: knex</span><span class="api-badge import">github.com/sahilkhaire/gox/db</span></div>
# DB.Open

## Overview

Open connects using driver name and DSN.

## Signature

```go
func Open(driver, dsn string) (*DB, error)
```

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

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [DB.MustOpen](/packages/db/db-must-open)

← [Back to db package overview](/packages/db/)
