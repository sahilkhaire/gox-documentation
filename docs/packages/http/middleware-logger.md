---
title: "Middleware.Logger"
package: "http"
import: "github.com/sahilkhaire/gox/http"
node: "morgan"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: morgan</span><span class="api-badge import">github.com/sahilkhaire/gox/http</span></div>
# Middleware.Logger

## Overview

Maps the Node.js pattern `morgan` to gox `http.Logger()`. Part of the http package — Node.js analog: express.

**Node.js equivalent:** `morgan`

## Signature

```go
func Logger() Middleware
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
morgan
```

```go [Standard Go]
func handler(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(data)
}
```

```go [gox]
import "github.com/sahilkhaire/gox/http"

http.Logger()
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
- [Middleware.RateLimit](/packages/http/middleware-rate-limit)
- [Middleware.Recover](/packages/http/middleware-recover)

← [Back to http package overview](/packages/http/)
