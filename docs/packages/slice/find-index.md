---
title: "FindIndex"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "arr.findIndex(fn)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: arr.findIndex(fn)</span><span class="api-badge import">github.com/sahilkhaire/gox/slice</span></div>
# FindIndex

## Overview

Maps the Node.js pattern `arr.findIndex(fn)` to gox `slice.FindIndex(arr, fn)`. Part of the slice package — Node.js analog: lodash / Array.*.

**Node.js equivalent:** `arr.findIndex(fn)`

## Signature

```go
func FindIndex[T any](in []T, fn func(T) bool) int
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
arr.findIndex(fn)
```

```go [Standard Go]
// Manual loop or Go 1.21+ slices package
for i, v := range items {
    // transform v
}
```

```go [gox]
import "github.com/sahilkhaire/gox/slice"

slice.FindIndex(arr, fn)
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
