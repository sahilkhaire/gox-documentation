---
title: "Get"
package: "maputil"
import: "github.com/sahilkhaire/gox/maputil"
node: "_.get(obj, \"a.b\")"
gox-doc-version: "10"
---

<SymbolHeader pkg="maputil" title="Get" node="_.get(obj, &quot;a.b&quot;)" import-path="github.com/sahilkhaire/gox/maputil" />
## Overview

Get reads a nested value from m using dot-separated path (lodash get).

**Node.js equivalent:** `_.get(obj, "a.b")`

## Signature

<div class="signature-block">

```go
func Get(m map[string]any, path string) (any, bool)
```

</div>

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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/maputil/clone">Clone</a><a class="related-chip" href="/packages/maputil/invert">Invert</a><a class="related-chip" href="/packages/maputil/keys">Keys</a>
</div>

← [Back to maputil package overview](/packages/maputil/)
