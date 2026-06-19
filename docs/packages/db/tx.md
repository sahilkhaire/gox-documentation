---
title: "Tx"
package: "db"
import: "github.com/sahilkhaire/gox/db"
gox-doc-version: "11"
---

<SymbolHeader pkg="db" title="Tx" node="knex" import-path="github.com/sahilkhaire/gox/db" />
## Overview

Tx wraps a sqlx transaction.

Part of the **`db`** package — Node.js analog: *knex*.

`Tx` is a type exported by gox. Methods on this type are documented separately.

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
// Typical knex pattern in Node.js
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

## Example

```go
import "github.com/sahilkhaire/gox/db"

_ = db.Tx
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to db package overview](/packages/db/)
