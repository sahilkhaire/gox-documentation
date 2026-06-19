---
title: "err Cookbook"
package: "err"
gox-doc-version: "10"
---

<div class="api-meta"><span class="api-badge node">Node: http-errors</span><span class="api-badge import">github.com/sahilkhaire/gox/err</span></div>
# err Cookbook

http-errors style responses in gox/http handlers.

```go
import goxerr "github.com/sahilkhaire/gox/err"

app.Get("/users/{id}", func(c *goxhttp.Ctx) error {
    user, ok := store.Find(c.Param("id"))
    if !ok {
        return goxerr.NotFound("user not found")
    }
    return c.JSON(200, user)
})
```

← [Back to err overview](/packages/err/)
