---
title: "Some"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "arr.some(fn)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: arr.some(fn)</span><span class="api-badge import">github.com/sahilkhaire/gox/slice</span></div>
# Some

## Overview

Maps the Node.js pattern `arr.some(fn)` to gox `slice.Some(arr, fn)`. Part of the slice package — Node.js analog: lodash / Array.*.

**Node.js equivalent:** `arr.some(fn)`

## Signature

```go
func Some[T any](in []T, fn func(T) bool) bool
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
arr.some(fn)
```

```go [Standard Go]
// Manual loop or Go 1.21+ slices package
for i, v := range items {
    // transform v
}
```

```go [gox]
import "github.com/sahilkhaire/gox/slice"

slice.Some(arr, fn)
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
