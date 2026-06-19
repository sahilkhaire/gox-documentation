---
title: "SaveUploadedFile"
package: "http"
import: "github.com/sahilkhaire/gox/http"
gox-doc-version: "14"
---

<SymbolHeader pkg="http" title="SaveUploadedFile" node="express, cors, helmet, morgan" import-path="github.com/sahilkhaire/gox/http" />
## Overview

SaveUploadedFile saves an uploaded file to destPath.

Part of the **`http`** package — Node.js analog: *express, cors, helmet, morgan*.

## Signature

<div class="signature-block">

```go
func SaveUploadedFile(file *multipart.FileHeader, destPath string) error
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

path, err := http.SaveUploadedFile(ctx, file, "./uploads")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/http"

path, err := http.SaveUploadedFile(ctx, file, "./uploads")
```

## Tips

Stack `Logger`, `Recover`, and `Security` middleware the way you would morgan + helmet in Express.

## Standard library alternative

Use the standard library directly:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // chi or net/http
}
```

← [Back to http package overview](/packages/http/)
