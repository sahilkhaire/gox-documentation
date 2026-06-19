---
title: "Flatten"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "_.flatten(arr)"
gox-doc-version: "10"
---

<SymbolHeader pkg="slice" title="Flatten" node="_.flatten(arr)" import-path="github.com/sahilkhaire/gox/slice" />
## Overview

Flatten flattens one level of nested slices (lodash flatten).

**Node.js equivalent:** `_.flatten(arr)`

## Signature

<div class="signature-block">

```go
func Flatten[T any](in [][]T) []T
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
_.flatten(arr)
```

```go [Standard Go]
// Manual loop or Go 1.21+ slices package
for _, v := range items {
    // transform or filter v
}
```

```go [gox]
import "github.com/sahilkhaire/gox/slice"

slice.Flatten(arr)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/slice/chunk">Chunk</a><a class="related-chip" href="/packages/slice/contains">Contains</a><a class="related-chip" href="/packages/slice/every">Every</a>
</div>

← [Back to slice package overview](/packages/slice/)
