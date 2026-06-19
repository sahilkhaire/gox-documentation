---
title: "SaveUploadedFile"
package: "http"
import: "github.com/sahilkhaire/gox/http"
gox-doc-version: "10"
---

<SymbolHeader pkg="http" title="SaveUploadedFile" node="express, cors, helmet, morgan" import-path="github.com/sahilkhaire/gox/http" />
## Overview

SaveUploadedFile saves an uploaded file to destPath.

## Signature

<div class="signature-block">

```go
func SaveUploadedFile(file *multipart.FileHeader, destPath string) error
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

// http
_ = http.SaveUploadedFile(/* args */)
```

:::

← [Back to http package overview](/packages/http/)
