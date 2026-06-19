---
title: "Tx"
package: "db"
import: "github.com/sahilkhaire/gox/db"
gox-doc-version: "10"
---

<SymbolHeader pkg="db" title="Tx" node="knex" import-path="github.com/sahilkhaire/gox/db" />
## Overview

Tx wraps a sqlx transaction.

## Signature

<div class="signature-block">

```go
type Tx struct {
	SQL *sqlx.Tx
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

_ = db.Tx
```

:::

← [Back to db package overview](/packages/db/)
