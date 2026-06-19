---
title: "Values"
package: "maputil"
import: "github.com/sahilkhaire/gox/maputil"
node: "Object.values(obj)"
gox-doc-version: "14"
---

<SymbolHeader pkg="maputil" title="Values" node="Object.values(obj)" import-path="github.com/sahilkhaire/gox/maputil" />
## Overview

Values returns map values (Object.values).

If you are coming from Node.js, the closest pattern is **`Object.values(obj)`**.

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

## Example

```go
import "github.com/sahilkhaire/gox/maputil"

maputil.Values(obj)
```

## Tips

Import `github.com/sahilkhaire/gox/maputil` and call `Values` directly. See the comparison below for the standard library equivalent.

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
