---
title: "Gt"
package: "mongo"
import: "github.com/sahilkhaire/gox/mongo"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: mongoose</span><span class="api-badge import">github.com/sahilkhaire/gox/mongo</span></div>
# Gt

## Overview

Gt matches field &gt; value.

## Signature

```go
func Gt(field string, value any) bson.M
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
_ = mongo.Gt(/* args */)
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
- [In](/packages/mongo/in)
- [Set](/packages/mongo/set)

← [Back to mongo package overview](/packages/mongo/)
