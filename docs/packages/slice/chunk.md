---
title: "Chunk"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "_.chunk(arr, n)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: _.chunk(arr, n)</span><span class="api-badge import">github.com/sahilkhaire/gox/slice</span></div>
# Chunk

## Overview

Maps the Node.js pattern `_.chunk(arr, n)` to gox `slice.Chunk(arr, n)`. Part of the slice package — Node.js analog: lodash / Array.*.

**Node.js equivalent:** `_.chunk(arr, n)`

## Signature

```go
func Chunk[T any](in []T, size int) [][]T
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
_.chunk(arr, n)
```

```go [Standard Go]
// Manual loop or Go 1.21+ slices package
for i, v := range items {
    // transform v
}
```

```go [gox]
import "github.com/sahilkhaire/gox/slice"

slice.Chunk(arr, n)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Contains](/packages/slice/contains)
- [Every](/packages/slice/every)
- [Filter](/packages/slice/filter)

← [Back to slice package overview](/packages/slice/)
