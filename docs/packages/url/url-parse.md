---
title: "URL.Parse"
package: "url"
import: "github.com/sahilkhaire/gox/url"
node: "new URL(str)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: new URL(str)</span><span class="api-badge import">github.com/sahilkhaire/gox/url</span></div>
# URL.Parse

## Overview

Parse parses a URL string (url.parse).

**Node.js equivalent:** `new URL(str)`

## Signature

```go
func Parse(raw string) (*URL, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
new URL('https://example.com/path');
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/url"

u, err := url.Parse("https://example.com/path")
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

← [Back to url package overview](/packages/url/)
