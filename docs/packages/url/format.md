---
title: "Format"
package: "url"
import: "github.com/sahilkhaire/gox/url"
node: "url.format(obj)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: url.format(obj)</span><span class="api-badge import">github.com/sahilkhaire/gox/url</span></div>
# Format

## Overview

Maps the Node.js pattern `url.format(obj)` to gox `url.Format(u)`.

**Node.js equivalent:** `url.format(obj)`

## Signature

```go
func Format(u *URL) string
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
url.format(obj)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/url"

url.Format(u)
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
- [ParseQuery](/packages/url/parse-query)
- [Resolve](/packages/url/resolve)

← [Back to url package overview](/packages/url/)
