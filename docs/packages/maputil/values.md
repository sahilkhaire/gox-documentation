---
title: "Values"
package: "maputil"
import: "github.com/sahilkhaire/gox/maputil"
node: "Object.values(obj)"
gox-doc-version: "10"
---

<SymbolHeader pkg="maputil" title="Values" node="Object.values(obj)" import-path="github.com/sahilkhaire/gox/maputil" />
## Overview

Values returns map values (Object.values).

**Node.js equivalent:** `Object.values(obj)`

## Signature

<div class="signature-block">

```go
func Values[K comparable, V any](m map[K]V) []V
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
Object.values(obj)
```

```go [Standard Go]
keys := make([]string, 0, len(m))
for k := range m { keys = append(keys, k) }
```

```go [gox]
import "github.com/sahilkhaire/gox/maputil"

maputil.Values(obj)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/maputil/clone">Clone</a><a class="related-chip" href="/packages/maputil/get">Get</a><a class="related-chip" href="/packages/maputil/invert">Invert</a>
</div>

← [Back to maputil package overview](/packages/maputil/)
