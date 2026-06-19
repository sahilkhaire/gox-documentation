---
title: "Resolve"
package: "url"
import: "github.com/sahilkhaire/gox/url"
node: "url.resolve(base, rel)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: url.resolve(base, rel)</span><span class="api-badge import">github.com/sahilkhaire/gox/url</span></div>
# Resolve

## Overview

Maps the Node.js pattern `url.resolve(base, rel)` to gox `url.Resolve(base, rel)`.

**Node.js equivalent:** `url.resolve(base, rel)`

## Signature

```go
func Resolve(base, relative string) (string, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
url.resolve(base, rel)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/url"

url.Resolve(base, rel)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [EncodeQuery](/packages/url/encode-query)
- [Format](/packages/url/format)
- [ParseQuery](/packages/url/parse-query)

← [Back to url package overview](/packages/url/)
