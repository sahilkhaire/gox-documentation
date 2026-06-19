---
title: "DB"
package: "db"
import: "github.com/sahilkhaire/gox/db"
gox-doc-version: "10"
---

<SymbolHeader pkg="db" title="DB" node="knex" import-path="github.com/sahilkhaire/gox/db" />
## Overview

DB wraps *sqlx.DB with a chainable query API.

## Signature

<div class="signature-block">

```go
type DB struct {
	SQL *sqlx.DB
}
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
db, err := sqlx.Connect("postgres", dsn)
db.GetContext(ctx, &row, query, args...)
```

```go [gox]
import "github.com/sahilkhaire/gox/db"

_ = db.DB
```

:::

← [Back to db package overview](/packages/db/)
