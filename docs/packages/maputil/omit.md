---
title: "Omit"
package: "maputil"
import: "github.com/sahilkhaire/gox/maputil"
node: "_.omit(obj, keys)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: _.omit(obj, keys)</span><span class="api-badge import">github.com/sahilkhaire/gox/maputil</span></div>
# Omit

## Overview

Maps the Node.js pattern `_.omit(obj, keys)` to gox `maputil.Omit(obj, keys...)`.

**Node.js equivalent:** `_.omit(obj, keys)`

## Signature

```go
func Omit[K comparable, V any](m map[K]V, keys ...K) map[K]V
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
_.omit(obj, keys)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/maputil"

maputil.Omit(obj, keys...)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Clone](/packages/maputil/clone)
- [Get](/packages/maputil/get)
- [Invert](/packages/maputil/invert)

← [Back to maputil package overview](/packages/maputil/)
