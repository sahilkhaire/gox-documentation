---
title: "DB.MustOpen"
package: "db"
import: "github.com/sahilkhaire/gox/db"
gox-doc-version: "11"
---

<SymbolHeader pkg="db" title="DB.MustOpen" node="knex" import-path="github.com/sahilkhaire/gox/db" />
## Overview

MustOpen connects or panics.

Part of the **`db`** package — Node.js analog: *knex*.

Method on **`DB`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func MustOpen(driver, dsn string) *DB
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

var v DB
v.MustOpen(/* args */)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/db"

var v DB
v.MustOpen(/* args */)
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/db/db-open">DB.Open</a>
</div>

← [Back to db package overview](/packages/db/)
