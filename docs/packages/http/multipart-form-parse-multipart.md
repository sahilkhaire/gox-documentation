---
title: "MultipartForm.ParseMultipart"
package: "http"
import: "github.com/sahilkhaire/gox/http"
gox-doc-version: "11"
---

<SymbolHeader pkg="http" title="MultipartForm.ParseMultipart" node="express, cors, helmet, morgan" import-path="github.com/sahilkhaire/gox/http" />
## Overview

ParseMultipart parses a multipart request (multer).

Part of the **`http`** package — Node.js analog: *express, cors, helmet, morgan*.

Method on **`MultipartForm`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func ParseMultipart(c *Ctx, maxMemory int64) (*MultipartForm, error)
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

var v MultipartForm
v.ParseMultipart(/* args */)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/http"

app := New()
dir := t.TempDir()
app.Post("/upload", func(c *Ctx) error {
	form, err := ParseMultipart(c, 1<<20)
	if err != nil {
		return err
	}
	files := form.Files("file")
	if len(files) != 1 {
		return goxerr.BadRequest("missing file")
	}
	dest := filepath.Join(dir, "out.txt")
	if err := SaveUploadedFile(files[0], dest); err != nil {
		return err
	}
	return c.JSON(200, map[string]string{"ok": "1"})
})
var body bytes.Buffer
w := multipart.NewWriter(&body)
part, _ := w.CreateFormFile("file", "x.txt")
```

## Tips

Stack `Logger`, `Recover`, and `Security` middleware the way you would morgan + helmet in Express.

## Standard library alternative

Use `net/http` with handler functions `func(w http.ResponseWriter, r *http.Request)` or a router like chi/echo directly.

← [Back to http package overview](/packages/http/)
