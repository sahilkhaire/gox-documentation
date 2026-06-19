---
title: "Set"
package: "maputil"
import: "github.com/sahilkhaire/gox/maputil"
node: "_.set(obj, \"a.b\", v)"
gox-doc-version: "10"
---

<SymbolHeader pkg="maputil" title="Set" node="_.set(obj, &quot;a.b&quot;, v)" import-path="github.com/sahilkhaire/gox/maputil" />
## Overview

Set writes value at dot-separated path, creating intermediate maps (lodash set).

**Node.js equivalent:** `_.set(obj, "a.b", v)`

## Signature

<div class="signature-block">

```go
func Set(m map[string]any, path string, value any)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
_.set(obj, "a.b", v)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches.
```

```go [gox]
import "github.com/sahilkhaire/gox/maputil"

maputil.Set(obj, "a.b", v)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/maputil/clone">Clone</a><a class="related-chip" href="/packages/maputil/get">Get</a><a class="related-chip" href="/packages/maputil/invert">Invert</a>
</div>

← [Back to maputil package overview](/packages/maputil/)
