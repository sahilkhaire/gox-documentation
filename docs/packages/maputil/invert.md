---
title: "Invert"
package: "maputil"
import: "github.com/sahilkhaire/gox/maputil"
node: "_.invert(obj)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: _.invert(obj)</span><span class="api-badge import">github.com/sahilkhaire/gox/maputil</span></div>
# Invert

## Overview

Maps the Node.js pattern `_.invert(obj)` to gox `maputil.Invert(obj)`.

**Node.js equivalent:** `_.invert(obj)`

## Signature

```go
func Invert[K, V comparable](m map[K]V) map[V]K
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
_.invert(obj)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/maputil"

maputil.Invert(obj)
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
- [Keys](/packages/maputil/keys)

← [Back to maputil package overview](/packages/maputil/)
