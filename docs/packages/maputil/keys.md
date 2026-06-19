---
title: "Keys"
package: "maputil"
import: "github.com/sahilkhaire/gox/maputil"
node: "Object.keys(obj)"
gox-doc-version: "14"
---

<SymbolHeader pkg="maputil" title="Keys" node="Object.keys(obj)" import-path="github.com/sahilkhaire/gox/maputil" />
## Overview

Keys returns map keys (Object.keys).

If you are coming from Node.js, the closest pattern is **`Object.keys(obj)`**.

## Signature

<div class="signature-block">

```go
func Keys[K comparable, V any](m map[K]V) []K
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
Object.keys(obj)
```

```go [Standard Go]
keys := make([]string, 0, len(m))
for k := range m { keys = append(keys, k) }
```

```go [gox]
import "github.com/sahilkhaire/gox/maputil"

maputil.Keys(obj)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/maputil"

maputil.Keys(obj)
```

## Tips

Import `github.com/sahilkhaire/gox/maputil` and call `Keys` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
keys := make([]string, 0, len(m))
for k := range m { keys = append(keys, k) }
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/maputil/clone">Clone</a><a class="related-chip" href="/packages/maputil/get">Get</a><a class="related-chip" href="/packages/maputil/invert">Invert</a>
</div>

← [Back to maputil package overview](/packages/maputil/)
