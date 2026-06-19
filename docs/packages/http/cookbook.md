---
title: "http Cookbook"
package: "http"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: express, cors, helmet, morgan</span><span class="api-badge import">github.com/sahilkhaire/gox/http</span></div>
# http Cookbook

Express-style apps with middleware, JSON, routers, and error handling.

## Minimal API server

```go
import (
    goxerr "github.com/sahilkhaire/gox/err"
    goxhttp "github.com/sahilkhaire/gox/http"
)

app := goxhttp.New()
app.Use(goxhttp.Logger(), goxhttp.Recover(), goxhttp.Security())

app.Get("/health", func(c *goxhttp.Ctx) error {
    return c.JSON(200, map[string]string{"status": "ok"})
})

app.Post("/users", func(c *goxhttp.Ctx) error {
    var u struct{ Name string `json:"name"` }
    if err := c.BindJSON(&u); err != nil {
        return goxerr.BadRequest("invalid json")
    }
    return c.JSON(201, u)
})

app.Listen(":3000")
```

## Mounted sub-router

```go
api := app.Router()
api.Get("/users/{id}", func(c *goxhttp.Ctx) error {
    return c.JSON(200, map[string]string{"id": c.Param("id")})
})
app.Mount("/api", api)
```

## CORS + typed errors

```go
app.Use(goxhttp.CORS(goxhttp.CORSOptions{AllowOrigins: []string{"https://app.example"}}))
app.Get("/missing", func(c *goxhttp.Ctx) error {
    return goxerr.NotFound("not found") // → 404 JSON automatically
})
```

See also: [HTTP guide](/guide/http) · [Express migration](/migration/express)



← [Back to http overview](/packages/http/)
