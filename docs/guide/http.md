# HTTP package guide

The [`gox/http`](/packages/http/) package mirrors Express.js patterns: apps, routers, middleware, JSON handlers, and built-in security helpers.

## Core routing

| Express | gox |
|---------|-----|
| `express()` | `http.New()` |
| `app.get('/path', handler)` | `app.Get("/path", handler)` |
| `res.json(data)` | `return c.JSON(200, data)` |
| `req.body` | `c.BindJSON(&v)` |
| `req.params.id` | `c.Param("id")` |

See: [http.New](/packages/http/new), [http.Ctx.JSON](/packages/http/ctx-json)

## Middleware

| npm | gox |
|-----|-----|
| `cors` | [http.CORS](/packages/http/cors) |
| `helmet` | [http.Security](/packages/http/security) |
| `morgan` | [http.Logger](/packages/http/logger) |
| error handler | Return `error` from handlers; use [gox/err](/packages/err/) for status codes |

## Advanced features

| npm | gox |
|-----|-----|
| `multer` | [http.ParseMultipart](/packages/http/parse-multipart), [http.SaveUploadedFile](/packages/http/save-uploaded-file) |
| `express-session` | [http.SessionMiddleware](/packages/http/session-middleware), [http.MemoryStore](/packages/http/memory-store) |
| Server-Sent Events | [http.SSE](/packages/http/sse), [http.SSEHandler](/packages/http/sse-handler) |
| `express-rate-limit` | [http.RateLimit](/packages/http/rate-limit) |
| WebSocket upgrade | [http.HandleWS](/packages/http/handle-ws), [gox/ws](/packages/ws/) |

## Example

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
```

See the [Express migration guide](/migration/express) for a full side-by-side walkthrough.
