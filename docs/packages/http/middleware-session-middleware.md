---
title: "Middleware.SessionMiddleware"
package: "http"
import: "github.com/sahilkhaire/gox/http"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: express, cors, helmet, morgan</span><span class="api-badge import">github.com/sahilkhaire/gox/http</span></div>
# Middleware.SessionMiddleware

## Overview

SessionMiddleware loads and saves session data on each request.

## Signature

```go
func SessionMiddleware(store SessionStore, opts SessionOptions) Middleware
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
func handler(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(data)
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
