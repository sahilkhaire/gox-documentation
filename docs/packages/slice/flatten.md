---
title: "Flatten"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "_.flatten(arr)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: _.flatten(arr)</span><span class="api-badge import">github.com/sahilkhaire/gox/slice</span></div>
# Flatten

## Overview

Maps the Node.js pattern `_.flatten(arr)` to gox `slice.Flatten(arr)`. Part of the slice package — Node.js analog: lodash / Array.*.

**Node.js equivalent:** `_.flatten(arr)`

## Signature

```go
func Flatten[T any](in [][]T) []T
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
_.flatten(arr)
```

```go [Standard Go]
// Manual loop or Go 1.21+ slices package
for i, v := range items {
    // transform v
}
```

```go [gox]
import "github.com/sahilkhaire/gox/slice"

slice.Flatten(arr)
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
