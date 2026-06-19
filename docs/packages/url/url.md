---
title: "URL"
package: "url"
import: "github.com/sahilkhaire/gox/url"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: url, querystring</span><span class="api-badge import">github.com/sahilkhaire/gox/url</span></div>
# URL

## Overview

URL wraps net/url.URL with Node-friendly helpers.

## Signature

```go
type URL struct {
	*url.URL
}
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/url"

_ = url.URL
```

:::

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

```go
import "github.com/sahilkhaire/gox/url"

u, err := Parse("https://example.com/a?x=1")
res, err := Resolve("https://example.com/foo/", "bar")
vals, err := ParseQuery("a=1&b=2")
```

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
