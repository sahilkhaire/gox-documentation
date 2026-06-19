---
title: "Values"
package: "maputil"
import: "github.com/sahilkhaire/gox/maputil"
node: "Object.values(obj)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: Object.values(obj)</span><span class="api-badge import">github.com/sahilkhaire/gox/maputil</span></div>
# Values

## Overview

Maps the Node.js pattern `Object.values(obj)` to gox `maputil.Values(obj)`.

**Node.js equivalent:** `Object.values(obj)`

## Signature

```go
func Values[K comparable, V any](m map[K]V) []V
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
Object.values(obj)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/maputil"

maputil.Values(obj)
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
