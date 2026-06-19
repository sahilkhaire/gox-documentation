---
title: "Every"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "arr.every(fn)"
gox-doc-version: "14"
---

<SymbolHeader pkg="slice" title="Every" node="arr.every(fn)" import-path="github.com/sahilkhaire/gox/slice" />
## Overview

Every reports whether fn is true for all elements (Array.every).

If you are coming from Node.js, the closest pattern is **`arr.every(fn)`**.

## Signature

<div class="signature-block">

```go
func Every[T any](in []T, fn func(T) bool) bool
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
arr.every(fn)
```

```go [Standard Go]
// Manual loop or Go 1.21+ slices package
for _, v := range items {
    // transform or filter v
}
```

```go [gox]
import "github.com/sahilkhaire/gox/slice"

slice.Every(arr, fn)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/slice"

slice.Every(arr, fn)
```

## Tips

Chain `Filter`, `Map`, and `Reduce` for lodash-style pipelines. Results are new slices — inputs are never mutated.

## Standard library alternative

Use a `for` loop or Go 1.21+ `slices` package helpers from the standard library.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/slice/chunk">Chunk</a><a class="related-chip" href="/packages/slice/contains">Contains</a><a class="related-chip" href="/packages/slice/filter">Filter</a>
</div>

← [Back to slice package overview](/packages/slice/)
