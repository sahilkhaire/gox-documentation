---
title: "Reduce"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "arr.reduce(fn, init)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: arr.reduce(fn, init)</span><span class="api-badge import">github.com/sahilkhaire/gox/slice</span></div>
# Reduce

## Overview

Reduce folds the slice (Array.reduce).

**Node.js equivalent:** `arr.reduce(fn, init)`

## Signature

```go
func Reduce[T, U any](in []T, init U, fn func(U, T) U) U
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const total = nums.reduce((acc, n) => acc + n, 0);
```

```go [Standard Go]
total := 0
for _, n := range nums {
    total += n
}
```

```go [gox]
import "github.com/sahilkhaire/gox/slice"

total := slice.Reduce(nums, 0, func(acc, n int) int { return acc + n })
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
