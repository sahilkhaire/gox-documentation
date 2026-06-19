---
title: "Contains"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "arr.includes(x)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: arr.includes(x)</span><span class="api-badge import">github.com/sahilkhaire/gox/slice</span></div>
# Contains

## Overview

Maps the Node.js pattern `arr.includes(x)` to gox `slice.Contains(arr, x)`. Part of the slice package — Node.js analog: lodash / Array.*.

**Node.js equivalent:** `arr.includes(x)`

## Signature

```go
func Contains[T comparable](in []T, v T) bool
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
arr.includes(x)
```

```go [Standard Go]
// Manual loop or Go 1.21+ slices package
for i, v := range items {
    // transform v
}
```

```go [gox]
import "github.com/sahilkhaire/gox/slice"

slice.Contains(arr, x)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Chunk](/packages/slice/chunk)
- [Every](/packages/slice/every)
- [Filter](/packages/slice/filter)

← [Back to slice package overview](/packages/slice/)
