---
title: "Some"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "arr.some(fn)"
gox-doc-version: "10"
---

<SymbolHeader pkg="slice" title="Some" node="arr.some(fn)" import-path="github.com/sahilkhaire/gox/slice" />
## Overview

Some reports whether fn is true for any element (Array.some).

**Node.js equivalent:** `arr.some(fn)`

## Signature

<div class="signature-block">

```go
func Some[T any](in []T, fn func(T) bool) bool
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
arr.some(fn)
```

```go [Standard Go]
// Manual loop or Go 1.21+ slices package
for _, v := range items {
    // transform or filter v
}
```

```go [gox]
import "github.com/sahilkhaire/gox/slice"

slice.Some(arr, fn)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/slice/chunk">Chunk</a><a class="related-chip" href="/packages/slice/contains">Contains</a><a class="related-chip" href="/packages/slice/every">Every</a>
</div>

← [Back to slice package overview](/packages/slice/)
