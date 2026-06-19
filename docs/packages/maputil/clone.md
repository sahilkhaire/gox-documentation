---
title: "Clone"
package: "maputil"
import: "github.com/sahilkhaire/gox/maputil"
node: "_.cloneDeep(obj)"
gox-doc-version: "14"
---

<SymbolHeader pkg="maputil" title="Clone" node="_.cloneDeep(obj)" import-path="github.com/sahilkhaire/gox/maputil" />
## Overview

Clone returns a shallow copy of m (lodash clone).

If you are coming from Node.js, the closest pattern is **`_.cloneDeep(obj)`**.

## Signature

<div class="signature-block">

```go
func Clone[K comparable, V any](m map[K]V) map[K]V
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
_.cloneDeep(obj)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches.
```

```go [gox]
import "github.com/sahilkhaire/gox/maputil"

maputil.Clone(obj)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/maputil"

maputil.Clone(obj)
```

## Tips

Import `github.com/sahilkhaire/gox/maputil` and call `Clone` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/maputil/get">Get</a><a class="related-chip" href="/packages/maputil/invert">Invert</a><a class="related-chip" href="/packages/maputil/keys">Keys</a>
</div>

← [Back to maputil package overview](/packages/maputil/)
