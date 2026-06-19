---
title: "SortBy"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "_.sortBy(arr, fn)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: _.sortBy(arr, fn)</span><span class="api-badge import">github.com/sahilkhaire/gox/slice</span></div>
# SortBy

## Overview

Maps the Node.js pattern `_.sortBy(arr, fn)` to gox `slice.SortBy(arr, fn)`. Part of the slice package — Node.js analog: lodash / Array.*.

**Node.js equivalent:** `_.sortBy(arr, fn)`

## Signature

```go
func SortBy[T any, K cmp.Ordered](in []T, fn func(T) K) []T
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
_.sortBy(arr, fn)
```

```go [Standard Go]
// Manual loop or Go 1.21+ slices package
for i, v := range items {
    // transform v
}
```

```go [gox]
import "github.com/sahilkhaire/gox/slice"

slice.SortBy(arr, fn)
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
- [Contains](/packages/slice/contains)
- [Every](/packages/slice/every)

← [Back to slice package overview](/packages/slice/)
