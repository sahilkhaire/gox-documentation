---
title: "Query"
package: "db"
import: "github.com/sahilkhaire/gox/db"
gox-doc-version: "10"
---

<SymbolHeader pkg="db" title="Query" node="knex" import-path="github.com/sahilkhaire/gox/db" />
## Overview

Query is a chainable SELECT builder.

## Signature

<div class="signature-block">

```go
type Query struct {
	// contains filtered or unexported fields
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

_ = db.Query
```

:::

← [Back to db package overview](/packages/db/)
