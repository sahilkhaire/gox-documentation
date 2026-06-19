---
title: "EncodeQuery"
package: "url"
import: "github.com/sahilkhaire/gox/url"
node: "querystring.stringify(obj)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: querystring.stringify(obj)</span><span class="api-badge import">github.com/sahilkhaire/gox/url</span></div>
# EncodeQuery

## Overview

Maps the Node.js pattern `querystring.stringify(obj)` to gox `url.EncodeQuery(obj)`.

**Node.js equivalent:** `querystring.stringify(obj)`

## Signature

```go
func EncodeQuery(v url.Values) string
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
querystring.stringify(obj)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/url"

url.EncodeQuery(obj)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Format](/packages/url/format)
- [ParseQuery](/packages/url/parse-query)
- [Resolve](/packages/url/resolve)

← [Back to url package overview](/packages/url/)
