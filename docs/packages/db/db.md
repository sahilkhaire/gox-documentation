---
title: "DB"
package: "db"
import: "github.com/sahilkhaire/gox/db"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: knex</span><span class="api-badge import">github.com/sahilkhaire/gox/db</span></div>
# DB

## Overview

DB wraps *sqlx.DB with a chainable query API.

## Signature

```go
type DB struct {
	SQL *sqlx.DB
}
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/db"

_ = db.DB
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

← [Back to db package overview](/packages/db/)
