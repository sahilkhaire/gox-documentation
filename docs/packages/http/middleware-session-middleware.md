---
title: "Middleware.SessionMiddleware"
package: "http"
import: "github.com/sahilkhaire/gox/http"
gox-doc-version: "10"
---

<SymbolHeader pkg="http" title="Middleware.SessionMiddleware" node="express, cors, helmet, morgan" import-path="github.com/sahilkhaire/gox/http" />
## Overview

SessionMiddleware loads and saves session data on each request.

## Signature

<div class="signature-block">

```go
func SessionMiddleware(store SessionStore, opts SessionOptions) Middleware
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
func handler(w http.ResponseWriter, r *http.Request) {
    // chi or net/http
}
```

```go [gox]
import "github.com/sahilkhaire/gox/http"

var v Middleware
v.SessionMiddleware(/* args */)
```

:::

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

```go
import "github.com/sahilkhaire/gox/http"

store := NewMemoryStore()
app := New()
app.Use(SessionMiddleware(store, SessionOptions{CookieName: "sid"}))
app.Get("/set", func(c *Ctx) error {
	c.Session()["user"] = "alice"
	return c.JSON(200, map[string]string{"ok": "1"})
})
app.Get("/get", func(c *Ctx) error {
	u, _ := c.Session()["user"].(string)
	return c.JSON(200, map[string]string{"user": u})
})
rec := httptest.NewRecorder()
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/http/middleware-cors">Middleware.CORS</a><a class="related-chip" href="/packages/http/middleware-logger">Middleware.Logger</a><a class="related-chip" href="/packages/http/middleware-rate-limit">Middleware.RateLimit</a>
</div>

← [Back to http package overview](/packages/http/)
