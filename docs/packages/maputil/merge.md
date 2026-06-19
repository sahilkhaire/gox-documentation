---
title: "Merge"
package: "maputil"
import: "github.com/sahilkhaire/gox/maputil"
node: "_.merge(a, b)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: _.merge(a, b)</span><span class="api-badge import">github.com/sahilkhaire/gox/maputil</span></div>
# Merge

## Overview

Maps the Node.js pattern `_.merge(a, b)` to gox `maputil.Merge(a, b)`.

**Node.js equivalent:** `_.merge(a, b)`

## Signature

```go
func Merge[K comparable, V any](dst map[K]V, sources ...map[K]V) map[K]V
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
_.merge(a, b)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/maputil"

maputil.Merge(a, b)
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
