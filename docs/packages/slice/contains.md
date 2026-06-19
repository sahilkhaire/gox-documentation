---
title: "Contains"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "arr.includes(x)"
gox-doc-version: "10"
---

<SymbolHeader pkg="slice" title="Contains" node="arr.includes(x)" import-path="github.com/sahilkhaire/gox/slice" />
## Overview

Contains reports whether v is in the slice (Array.includes).

**Node.js equivalent:** `arr.includes(x)`

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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/slice/chunk">Chunk</a><a class="related-chip" href="/packages/slice/every">Every</a><a class="related-chip" href="/packages/slice/filter">Filter</a>
</div>

← [Back to slice package overview](/packages/slice/)
