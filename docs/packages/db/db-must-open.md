---
title: "DB.MustOpen"
package: "db"
import: "github.com/sahilkhaire/gox/db"
gox-doc-version: "10"
---

<SymbolHeader pkg="db" title="DB.MustOpen" node="knex" import-path="github.com/sahilkhaire/gox/db" />
## Overview

MustOpen connects or panics.

## Signature

<div class="signature-block">

```go
func MustOpen(driver, dsn string) *DB
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

var v DB
v.MustOpen(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/db/db-open">DB.Open</a>
</div>

← [Back to db package overview](/packages/db/)
