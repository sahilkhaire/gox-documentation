---
title: "FindIndex"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "arr.findIndex(fn)"
gox-doc-version: "11"
---

<SymbolHeader pkg="slice" title="FindIndex" node="arr.findIndex(fn)" import-path="github.com/sahilkhaire/gox/slice" />
## Overview

FindIndex returns the index of the first match, or -1 (Array.findIndex).

If you are coming from Node.js, the closest pattern is **`arr.findIndex(fn)`**.

## Signature

<div class="signature-block">

```go
func FindIndex[T any](in []T, fn func(T) bool) int
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
arr.findIndex(fn)
```

```go [Standard Go]
// Manual loop or Go 1.21+ slices package
for _, v := range items {
    // transform or filter v
}
```

```go [gox]
import "github.com/sahilkhaire/gox/slice"

slice.FindIndex(arr, fn)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/slice"

slice.FindIndex(arr, fn)
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
