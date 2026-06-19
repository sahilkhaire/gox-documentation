---
title: "Clone"
package: "maputil"
import: "github.com/sahilkhaire/gox/maputil"
node: "_.cloneDeep(obj)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: _.cloneDeep(obj)</span><span class="api-badge import">github.com/sahilkhaire/gox/maputil</span></div>
# Clone

## Overview

Maps the Node.js pattern `_.cloneDeep(obj)` to gox `maputil.Clone(obj)`.

**Node.js equivalent:** `_.cloneDeep(obj)`

## Signature

```go
func Clone[K comparable, V any](m map[K]V) map[K]V
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
_.cloneDeep(obj)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/maputil"

maputil.Clone(obj)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Get](/packages/maputil/get)
- [Invert](/packages/maputil/invert)
- [Keys](/packages/maputil/keys)

← [Back to maputil package overview](/packages/maputil/)
