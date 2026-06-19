---
title: "Middleware.Security"
package: "http"
import: "github.com/sahilkhaire/gox/http"
node: "helmet"
gox-doc-version: "10"
---

<SymbolHeader pkg="http" title="Middleware.Security" node="helmet" import-path="github.com/sahilkhaire/gox/http" />
## Overview

Security sets common security headers (helmet-style).

**Node.js equivalent:** `helmet`

## Signature

<div class="signature-block">

```go
func Security() Middleware
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
helmet
```

```go [Standard Go]
func handler(w http.ResponseWriter, r *http.Request) {
    // chi or net/http
}
```

```go [gox]
import "github.com/sahilkhaire/gox/http"

http.Security()
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/http/middleware-cors">Middleware.CORS</a><a class="related-chip" href="/packages/http/middleware-logger">Middleware.Logger</a><a class="related-chip" href="/packages/http/middleware-rate-limit">Middleware.RateLimit</a>
</div>

← [Back to http package overview](/packages/http/)
