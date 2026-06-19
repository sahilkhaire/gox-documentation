---
title: "Pick"
package: "maputil"
import: "github.com/sahilkhaire/gox/maputil"
node: "_.pick(obj, keys)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: _.pick(obj, keys)</span><span class="api-badge import">github.com/sahilkhaire/gox/maputil</span></div>
# Pick

## Overview

Pick returns a new map with only the given keys (lodash pick).

**Node.js equivalent:** `_.pick(obj, keys)`

## Signature

```go
func Pick[K comparable, V any](m map[K]V, keys ...K) map[K]V
```

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

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Clone](/packages/maputil/clone)
- [Get](/packages/maputil/get)
- [Invert](/packages/maputil/invert)

← [Back to maputil package overview](/packages/maputil/)
