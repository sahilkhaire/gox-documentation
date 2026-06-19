---
title: "Middleware.Logger"
package: "http"
import: "github.com/sahilkhaire/gox/http"
node: "morgan"
gox-doc-version: "10"
---

<SymbolHeader pkg="http" title="Middleware.Logger" node="morgan" import-path="github.com/sahilkhaire/gox/http" />
## Overview

Logger logs each request with slog (morgan-style).

**Node.js equivalent:** `morgan`

## Signature

<div class="signature-block">

```go
func Logger() Middleware
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
morgan
```

```go [Standard Go]
func handler(w http.ResponseWriter, r *http.Request) {
    // chi or net/http
}
```

```go [gox]
import "github.com/sahilkhaire/gox/http"

http.Logger()
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/http/middleware-cors">Middleware.CORS</a><a class="related-chip" href="/packages/http/middleware-rate-limit">Middleware.RateLimit</a><a class="related-chip" href="/packages/http/middleware-recover">Middleware.Recover</a>
</div>

← [Back to http package overview](/packages/http/)
