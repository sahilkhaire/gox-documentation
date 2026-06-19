---
title: "SortBy"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "_.sortBy(arr, fn)"
gox-doc-version: "14"
---

<SymbolHeader pkg="slice" title="SortBy" node="_.sortBy(arr, fn)" import-path="github.com/sahilkhaire/gox/slice" />
## Overview

SortBy returns a sorted copy ordered by fn's key (lodash sortBy).

If you are coming from Node.js, the closest pattern is **`_.sortBy(arr, fn)`**.

## Signature

<div class="signature-block">

```go
func SortBy[T any, K cmp.Ordered](in []T, fn func(T) K) []T
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
_.sortBy(arr, fn)
```

```go [Standard Go]
// Manual loop or Go 1.21+ slices package
for _, v := range items {
    // transform or filter v
}
```

```go [gox]
import "github.com/sahilkhaire/gox/slice"

slice.SortBy(arr, fn)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/slice"

slice.SortBy(arr, fn)
```

## Tips

Chain `Filter`, `Map`, and `Reduce` for lodash-style pipelines. Results are new slices — inputs are never mutated.

## Standard library alternative

Use a `for` loop or Go 1.21+ `slices` package helpers from the standard library.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/slice/chunk">Chunk</a><a class="related-chip" href="/packages/slice/contains">Contains</a><a class="related-chip" href="/packages/slice/every">Every</a>
</div>

← [Back to slice package overview](/packages/slice/)
