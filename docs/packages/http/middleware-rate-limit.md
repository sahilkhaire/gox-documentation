---
title: "Middleware.RateLimit"
package: "http"
import: "github.com/sahilkhaire/gox/http"
node: "express-rate-limit"
gox-doc-version: "10"
---

<SymbolHeader pkg="http" title="Middleware.RateLimit" node="express-rate-limit" import-path="github.com/sahilkhaire/gox/http" />
## Overview

RateLimit returns middleware that limits requests per key.

**Node.js equivalent:** `express-rate-limit`

## Signature

<div class="signature-block">

```go
func RateLimit(opts RateLimitOptions) Middleware
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
express-rate-limit
```

```go [Standard Go]
func handler(w http.ResponseWriter, r *http.Request) {
    // chi or net/http
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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/http/middleware-cors">Middleware.CORS</a><a class="related-chip" href="/packages/http/middleware-logger">Middleware.Logger</a><a class="related-chip" href="/packages/http/middleware-recover">Middleware.Recover</a>
</div>

← [Back to http package overview](/packages/http/)
