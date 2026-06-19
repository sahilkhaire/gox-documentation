---
title: "Contains"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "arr.includes(x)"
gox-doc-version: "11"
---

<SymbolHeader pkg="slice" title="Contains" node="arr.includes(x)" import-path="github.com/sahilkhaire/gox/slice" />
## Overview

Contains reports whether v is in the slice (Array.includes).

If you are coming from Node.js, the closest pattern is **`arr.includes(x)`**.

## Signature

<div class="signature-block">

```go
func Contains[T comparable](in []T, v T) bool
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
arr.includes(x)
```

```go [Standard Go]
// Manual loop or Go 1.21+ slices package
for _, v := range items {
    // transform or filter v
}
```

```go [gox]
import "github.com/sahilkhaire/gox/slice"

slice.Contains(arr, x)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/slice"

slice.Contains(arr, x)
```

## Tips

Chain `Filter`, `Map`, and `Reduce` for lodash-style pipelines. Results are new slices — inputs are never mutated.

## Standard library alternative

Use a `for` loop or Go 1.21+ `slices` package helpers from the standard library.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/slice/chunk">Chunk</a><a class="related-chip" href="/packages/slice/every">Every</a><a class="related-chip" href="/packages/slice/filter">Filter</a>
</div>

← [Back to slice package overview](/packages/slice/)
