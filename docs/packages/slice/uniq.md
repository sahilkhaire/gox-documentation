---
title: "Uniq"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "_.uniq(arr)"
gox-doc-version: "10"
---

<SymbolHeader pkg="slice" title="Uniq" node="_.uniq(arr)" import-path="github.com/sahilkhaire/gox/slice" />
## Overview

Uniq returns unique elements in first-seen order (lodash uniq).

**Node.js equivalent:** `_.uniq(arr)`

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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/slice/chunk">Chunk</a><a class="related-chip" href="/packages/slice/contains">Contains</a><a class="related-chip" href="/packages/slice/every">Every</a>
</div>

← [Back to slice package overview](/packages/slice/)
