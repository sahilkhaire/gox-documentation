---
title: "CORSOptions"
package: "http"
import: "github.com/sahilkhaire/gox/http"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: express, cors, helmet, morgan</span><span class="api-badge import">github.com/sahilkhaire/gox/http</span></div>
# CORSOptions

## Overview

CORSOptions configures the CORS middleware.

## Signature

```go
type CORSOptions struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
	MaxAge           int
}
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

_ = http.CORSOptions
```

:::

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

```go
import "github.com/sahilkhaire/gox/http"

app := New()
app.Use(CORS(CORSOptions{AllowOrigins: []string{"https://app.example"}}))
app.Get("/", func(c *Ctx) error { return c.Status(204).JSON(204, nil) })
req := httptest.NewRequest(http.MethodOptions, "/", nil)
req.Header.Set("Origin", "https://app.example")
rec := httptest.NewRecorder()
```

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [SaveUploadedFile](/packages/http/save-uploaded-file)

← [Back to http package overview](/packages/http/)
