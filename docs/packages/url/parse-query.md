---
title: "ParseQuery"
package: "url"
import: "github.com/sahilkhaire/gox/url"
node: "querystring.parse(qs)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: querystring.parse(qs)</span><span class="api-badge import">github.com/sahilkhaire/gox/url</span></div>
# ParseQuery

## Overview

ParseQuery parses a query string (querystring.parse).

**Node.js equivalent:** `querystring.parse(qs)`

## Signature

```go
func ParseQuery(query string) (url.Values, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
querystring.parse('a=1&b=2');
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/url"

vals, err := url.ParseQuery("a=1&b=2")
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
- [Resolve](/packages/url/resolve)

← [Back to url package overview](/packages/url/)
