---
title: "SessionOptions"
package: "http"
import: "github.com/sahilkhaire/gox/http"
gox-doc-version: "10"
---

<SymbolHeader pkg="http" title="SessionOptions" node="express, cors, helmet, morgan" import-path="github.com/sahilkhaire/gox/http" />
## Overview

SessionOptions configures session middleware.

## Signature

<div class="signature-block">

```go
type SessionOptions struct {
	CookieName string
	MaxAge     time.Duration
	Path       string
	HttpOnly   bool
	Secure     bool
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

_ = http.SessionOptions
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/http/save-uploaded-file">SaveUploadedFile</a>
</div>

← [Back to http package overview](/packages/http/)
