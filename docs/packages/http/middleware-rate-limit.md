---
title: "Middleware.RateLimit"
package: "http"
import: "github.com/sahilkhaire/gox/http"
node: "express-rate-limit"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: express-rate-limit</span><span class="api-badge import">github.com/sahilkhaire/gox/http</span></div>
# Middleware.RateLimit

## Overview

Maps the Node.js pattern `express-rate-limit` to gox `http.RateLimit`. Part of the http package — Node.js analog: express.

**Node.js equivalent:** `express-rate-limit`

## Signature

```go
func RateLimit(opts RateLimitOptions) Middleware
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
express-rate-limit
```

```go [Standard Go]
func handler(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(data)
}
```

```go [gox]
import "github.com/sahilkhaire/gox/http"

http.RateLimit
```

:::

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

```go
import "github.com/sahilkhaire/gox/http"

app := New()
app.Use(RateLimit(RateLimitOptions{
	Requests:	2,
	Window:		time.Second,
	Burst:		2,
	Key:		func(c *Ctx) string { return "test" },
}))
app.Get("/x", func(c *Ctx) error {
	return c.JSON(200, map[string]string{"ok": "1"})
})
rec := httptest.NewRecorder()
app.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/x", nil))
```

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Middleware.CORS](/packages/http/middleware-cors)
- [Middleware.Logger](/packages/http/middleware-logger)
- [Middleware.Recover](/packages/http/middleware-recover)

← [Back to http package overview](/packages/http/)
