---
title: "Merge"
package: "maputil"
import: "github.com/sahilkhaire/gox/maputil"
node: "_.merge(a, b)"
gox-doc-version: "14"
---

<SymbolHeader pkg="maputil" title="Merge" node="_.merge(a, b)" import-path="github.com/sahilkhaire/gox/maputil" />
## Overview

Merge copies keys from sources into dst, later maps override (lodash merge, shallow).

If you are coming from Node.js, the closest pattern is **`_.merge(a, b)`**.

## Signature

<div class="signature-block">

```go
func Merge[K comparable, V any](dst map[K]V, sources ...map[K]V) map[K]V
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
_.merge(a, b)
```

```go [Standard Go]
out := maps.Clone(base)
maps.Copy(out, overlay)
```

```go [gox]
import "github.com/sahilkhaire/gox/maputil"

maputil.Merge(a, b)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/maputil"

maputil.Merge(a, b)
```

## Tips

Import `github.com/sahilkhaire/gox/maputil` and call `Merge` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
out := maps.Clone(base)
maps.Copy(out, overlay)
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/maputil/clone">Clone</a><a class="related-chip" href="/packages/maputil/get">Get</a><a class="related-chip" href="/packages/maputil/invert">Invert</a>
</div>

← [Back to maputil package overview](/packages/maputil/)
