---
title: "CORSOptions"
package: "http"
import: "github.com/sahilkhaire/gox/http"
gox-doc-version: "11"
---

<SymbolHeader pkg="http" title="CORSOptions" node="express, cors, helmet, morgan" import-path="github.com/sahilkhaire/gox/http" />
## Overview

CORSOptions configures the CORS middleware.

Part of the **`http`** package — Node.js analog: *express, cors, helmet, morgan*.

`CORSOptions` is a type exported by gox. Methods on this type are documented separately.

## Signature

<div class="signature-block">

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

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical express, cors, helmet, morgan pattern in Node.js
```

```go [Standard Go]
func handler(w http.ResponseWriter, r *http.Request) {
    // chi or net/http
}
```

```go [gox]
import "github.com/sahilkhaire/gox/http"

_ = http.CORSOptions
```

:::

## Example

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

Stack `Logger`, `Recover`, and `Security` middleware the way you would morgan + helmet in Express.

## Standard library alternative

Use `net/http` with handler functions `func(w http.ResponseWriter, r *http.Request)` or a router like chi/echo directly.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/http/save-uploaded-file">SaveUploadedFile</a>
</div>

← [Back to http package overview](/packages/http/)
