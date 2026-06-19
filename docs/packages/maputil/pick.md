---
title: "Pick"
package: "maputil"
import: "github.com/sahilkhaire/gox/maputil"
node: "_.pick(obj, keys)"
gox-doc-version: "14"
---

<SymbolHeader pkg="maputil" title="Pick" node="_.pick(obj, keys)" import-path="github.com/sahilkhaire/gox/maputil" />
## Overview

Pick returns a new map with only the given keys (lodash pick).

If you are coming from Node.js, the closest pattern is **`_.pick(obj, keys)`**.

## Signature

<div class="signature-block">

```go
func Pick[K comparable, V any](m map[K]V, keys ...K) map[K]V
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const subset = _.pick(obj, ['a', 'b']);
```

```go [Standard Go]
subset := map[string]any{}
for _, k := range []string{"a", "b"} {
    if v, ok := obj[k]; ok {
        subset[k] = v
    }
}
```

```go [gox]
import "github.com/sahilkhaire/gox/maputil"

subset := maputil.Pick(obj, "a", "b")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/maputil"

subset := maputil.Pick(obj, "a", "b")
```

## Tips

Import `github.com/sahilkhaire/gox/maputil` and call `Pick` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
subset := map[string]any{}
for _, k := range []string{"a", "b"} {
    if v, ok := obj[k]; ok {
        subset[k] = v
    }
}
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/maputil/clone">Clone</a><a class="related-chip" href="/packages/maputil/get">Get</a><a class="related-chip" href="/packages/maputil/invert">Invert</a>
</div>

← [Back to maputil package overview](/packages/maputil/)
