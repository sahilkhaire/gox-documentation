---
title: "CORSOptions"
package: "http"
import: "github.com/sahilkhaire/gox/http"
gox-doc-version: "10"
---

<SymbolHeader pkg="http" title="CORSOptions" node="express, cors, helmet, morgan" import-path="github.com/sahilkhaire/gox/http" />
## Overview

CORSOptions configures the CORS middleware.

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
// See package overview
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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/http/save-uploaded-file">SaveUploadedFile</a>
</div>

← [Back to http package overview](/packages/http/)
