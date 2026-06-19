---
title: "Invert"
package: "maputil"
import: "github.com/sahilkhaire/gox/maputil"
node: "_.invert(obj)"
gox-doc-version: "11"
---

<SymbolHeader pkg="maputil" title="Invert" node="_.invert(obj)" import-path="github.com/sahilkhaire/gox/maputil" />
## Overview

Invert swaps keys and values; duplicate values overwrite (lodash invert).

If you are coming from Node.js, the closest pattern is **`_.invert(obj)`**.

## Signature

<div class="signature-block">

```go
func Invert[K, V comparable](m map[K]V) map[V]K
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
_.invert(obj)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches.
```

```go [gox]
import "github.com/sahilkhaire/gox/maputil"

maputil.Invert(obj)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/maputil"

maputil.Invert(obj)
```

## Tips

Import `github.com/sahilkhaire/gox/maputil` and call `Invert` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/maputil/clone">Clone</a><a class="related-chip" href="/packages/maputil/get">Get</a><a class="related-chip" href="/packages/maputil/keys">Keys</a>
</div>

← [Back to maputil package overview](/packages/maputil/)
