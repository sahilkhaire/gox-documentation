---
title: "In"
package: "mongo"
import: "github.com/sahilkhaire/gox/mongo"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: mongoose</span><span class="api-badge import">github.com/sahilkhaire/gox/mongo</span></div>
# In

## Overview

In matches field in values.

## Signature

```go
func In(field string, values ...any) bson.M
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
import "github.com/sahilkhaire/gox/mongo"

// mongo
_ = mongo.In(/* args */)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Eq](/packages/mongo/eq)
- [Gt](/packages/mongo/gt)
- [Set](/packages/mongo/set)

← [Back to mongo package overview](/packages/mongo/)
