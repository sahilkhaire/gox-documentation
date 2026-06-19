---
title: "Uniq"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "_.uniq(arr)"
gox-doc-version: "11"
---

<SymbolHeader pkg="slice" title="Uniq" node="_.uniq(arr)" import-path="github.com/sahilkhaire/gox/slice" />
## Overview

Uniq returns unique elements in first-seen order (lodash uniq).

If you are coming from Node.js, the closest pattern is **`_.uniq(arr)`**.

## Signature

<div class="signature-block">

```go
func Uniq[T comparable](in []T) []T
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const unique = _.uniq(items);
```

```go [Standard Go]
seen := make(map[T]struct{})
var unique []T
for _, v := range items { /* dedupe */ }
```

```go [gox]
import "github.com/sahilkhaire/gox/slice"

unique := slice.Uniq(items)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/slice"

unique := slice.Uniq(items)
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
