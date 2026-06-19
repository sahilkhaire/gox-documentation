---
title: "Omit"
package: "maputil"
import: "github.com/sahilkhaire/gox/maputil"
node: "_.omit(obj, keys)"
gox-doc-version: "10"
---

<SymbolHeader pkg="maputil" title="Omit" node="_.omit(obj, keys)" import-path="github.com/sahilkhaire/gox/maputil" />
## Overview

Omit returns a new map without the given keys (lodash omit).

**Node.js equivalent:** `_.omit(obj, keys)`

## Signature

<div class="signature-block">

```go
func Omit[K comparable, V any](m map[K]V, keys ...K) map[K]V
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
_.omit(obj, keys)
```

```go [Standard Go]
out := make(map[string]any)
for k, v := range obj { /* copy keys */ }
```

```go [gox]
import "github.com/sahilkhaire/gox/maputil"

maputil.Omit(obj, keys...)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/maputil/clone">Clone</a><a class="related-chip" href="/packages/maputil/get">Get</a><a class="related-chip" href="/packages/maputil/invert">Invert</a>
</div>

← [Back to maputil package overview](/packages/maputil/)
