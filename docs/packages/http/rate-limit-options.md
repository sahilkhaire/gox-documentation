---
title: "RateLimitOptions"
package: "http"
import: "github.com/sahilkhaire/gox/http"
gox-doc-version: "10"
---

<SymbolHeader pkg="http" title="RateLimitOptions" node="express, cors, helmet, morgan" import-path="github.com/sahilkhaire/gox/http" />
## Overview

RateLimitOptions configures per-key rate limiting (express-rate-limit).

## Signature

<div class="signature-block">

```go
type RateLimitOptions struct {
	Requests int
	Window   time.Duration
	Burst    int
	Key      func(*Ctx) string
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

_ = http.RateLimitOptions
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/http/save-uploaded-file">SaveUploadedFile</a>
</div>

← [Back to http package overview](/packages/http/)
