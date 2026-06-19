---
title: "MultipartForm.ParseMultipart"
package: "http"
import: "github.com/sahilkhaire/gox/http"
gox-doc-version: "10"
---

<SymbolHeader pkg="http" title="MultipartForm.ParseMultipart" node="express, cors, helmet, morgan" import-path="github.com/sahilkhaire/gox/http" />
## Overview

ParseMultipart parses a multipart request (multer).

## Signature

<div class="signature-block">

```go
func ParseMultipart(c *Ctx, maxMemory int64) (*MultipartForm, error)
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

var v MultipartForm
v.ParseMultipart(/* args */)
```

:::

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

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

← [Back to http package overview](/packages/http/)
