---
title: "Middleware.Security"
package: "http"
import: "github.com/sahilkhaire/gox/http"
node: "helmet"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: helmet</span><span class="api-badge import">github.com/sahilkhaire/gox/http</span></div>
# Middleware.Security

## Overview

Maps the Node.js pattern `helmet` to gox `http.Security()`. Part of the http package — Node.js analog: express.

**Node.js equivalent:** `helmet`

## Signature

```go
func Security() Middleware
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
helmet
```

```go [Standard Go]
func handler(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(data)
}
```

```go [gox]
import "github.com/sahilkhaire/gox/http"

http.Security()
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Middleware.CORS](/packages/http/middleware-cors)
- [Middleware.Logger](/packages/http/middleware-logger)
- [Middleware.RateLimit](/packages/http/middleware-rate-limit)

← [Back to http package overview](/packages/http/)
