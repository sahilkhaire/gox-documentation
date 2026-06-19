---
title: "DB.Open"
package: "db"
import: "github.com/sahilkhaire/gox/db"
gox-doc-version: "10"
---

<SymbolHeader pkg="db" title="DB.Open" node="knex" import-path="github.com/sahilkhaire/gox/db" />
## Overview

Open connects using driver name and DSN.

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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/db/db-must-open">DB.MustOpen</a>
</div>

← [Back to db package overview](/packages/db/)
