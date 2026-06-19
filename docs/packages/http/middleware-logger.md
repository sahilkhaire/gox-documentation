---
title: "Middleware.Logger"
package: "http"
import: "github.com/sahilkhaire/gox/http"
node: "morgan"
gox-doc-version: "14"
---

<SymbolHeader pkg="http" title="Middleware.Logger" node="morgan" import-path="github.com/sahilkhaire/gox/http" />
## Overview

Logger logs each request with slog (morgan-style).

If you are coming from Node.js, the closest pattern is **`morgan`**.

Method on **`Middleware`** — call it on a value of that type after constructing or receiving one from a constructor.

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

## Example

```go
import "github.com/sahilkhaire/gox/http"

http.Logger()
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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/http/middleware-cors">Middleware.CORS</a><a class="related-chip" href="/packages/http/middleware-rate-limit">Middleware.RateLimit</a><a class="related-chip" href="/packages/http/middleware-recover">Middleware.Recover</a>
</div>

← [Back to http package overview](/packages/http/)
