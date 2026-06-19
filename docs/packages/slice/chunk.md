---
title: "Chunk"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "_.chunk(arr, n)"
gox-doc-version: "14"
---

<SymbolHeader pkg="slice" title="Chunk" node="_.chunk(arr, n)" import-path="github.com/sahilkhaire/gox/slice" />
## Overview

Chunk splits into sub-slices of size (lodash chunk).

If you are coming from Node.js, the closest pattern is **`_.chunk(arr, n)`**.

## Signature

<div class="signature-block">

```go
func Chunk[T any](in []T, size int) [][]T
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
_.chunk(arr, n)
```

```go [Standard Go]
// Manual loop or Go 1.21+ slices package
for _, v := range items {
    // transform or filter v
}
```

```go [gox]
import "github.com/sahilkhaire/gox/slice"

slice.Chunk(arr, n)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/slice"

slice.Chunk(arr, n)
```

## Tips

Chain `Filter`, `Map`, and `Reduce` for lodash-style pipelines. Results are new slices — inputs are never mutated.

## Standard library alternative

Use a `for` loop or Go 1.21+ `slices` package helpers from the standard library.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/slice/contains">Contains</a><a class="related-chip" href="/packages/slice/every">Every</a><a class="related-chip" href="/packages/slice/filter">Filter</a>
</div>

← [Back to slice package overview](/packages/slice/)
