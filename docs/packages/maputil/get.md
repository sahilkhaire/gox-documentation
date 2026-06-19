---
title: "Get"
package: "maputil"
import: "github.com/sahilkhaire/gox/maputil"
node: "_.get(obj, \"a.b\")"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: _.get(obj, "a.b")</span><span class="api-badge import">github.com/sahilkhaire/gox/maputil</span></div>
# Get

## Overview

Get reads a nested value from m using dot-separated path (lodash get).

**Node.js equivalent:** `_.get(obj, "a.b")`

## Signature

```go
func Get(m map[string]any, path string) (any, bool)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const city = _.get(obj, 'address.city', 'unknown');
```

```go [Standard Go]
city := "unknown"
if addr, ok := obj["address"].(map[string]any); ok {
    if c, ok := addr["city"].(string); ok {
        city = c
    }
}
```

```go [gox]
import "github.com/sahilkhaire/gox/maputil"

city, _ := maputil.Get[string](obj, "address.city")
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
- [Invert](/packages/maputil/invert)
- [Keys](/packages/maputil/keys)

← [Back to maputil package overview](/packages/maputil/)
