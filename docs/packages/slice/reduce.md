---
title: "Reduce"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "arr.reduce(fn, init)"
gox-doc-version: "10"
---

<SymbolHeader pkg="slice" title="Reduce" node="arr.reduce(fn, init)" import-path="github.com/sahilkhaire/gox/slice" />
## Overview

Reduce folds the slice (Array.reduce).

**Node.js equivalent:** `arr.reduce(fn, init)`

## Signature

<div class="signature-block">

```go
func Reduce[T, U any](in []T, init U, fn func(U, T) U) U
```

</div>

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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/slice/chunk">Chunk</a><a class="related-chip" href="/packages/slice/contains">Contains</a><a class="related-chip" href="/packages/slice/every">Every</a>
</div>

← [Back to slice package overview](/packages/slice/)
