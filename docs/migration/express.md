# Express → gox/http

Side-by-side patterns for common Express apps.

## App setup

| Express | gox/http |
|---------|----------|
| `const express = require('express')` | `import goxhttp "github.com/sahilkhaire/gox/http"` |
| `const app = express()` | `app := goxhttp.New()` |
| `app.listen(3000)` | `app.Listen(":3000")` |

## Routes

| Express | gox/http |
|---------|----------|
| `app.get('/users', (req, res) => { ... })` | `app.Get("/users", func(c *goxhttp.Ctx) error { ... })` |
| `res.status(201).json({ id: 1 })` | `return c.JSON(201, map[string]int{"id": 1})` |
| `res.sendStatus(204)` | `return c.Status(204).JSON(204, nil)` |

## Middleware

| Express | gox/http |
|---------|----------|
| `app.use(morgan('combined'))` | `app.Use(goxhttp.Logger())` |
| `app.use(cors())` | `app.Use(goxhttp.CORS(goxhttp.CORSOptions{AllowOrigins: []string{"*"}}))` |
| `app.use(helmet())` | `app.Use(goxhttp.Security())` |
| `app.use((err, req, res, next) => ...)` | Return `error` from handlers; `*err.AppError` sets status |

## Body & query

| Express | gox/http |
|---------|----------|
| `req.body` (json) | `var body T; c.BindJSON(&body)` |
| `req.query` | `var q T; c.BindQuery(&q)` (uses `json` struct tags) |
| `req.params.id` | `c.Param("id")` |

## Routers

| Express | gox/http |
|---------|----------|
| `const router = express.Router()` | `api := app.Router()` |
| `app.use('/api', router)` | `app.Mount("/api", api)` |

## Errors (http-errors)

| Express | gox/http + gox/err |
|---------|-------------------|
| `next(createError(404, 'not found'))` | `return err.NotFound("not found")` |
| `next(err)` default 500 | Any other `error` → 500 JSON `{"error":"..."}` |

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

See `examples/http-basic/main.go` for a runnable server.
