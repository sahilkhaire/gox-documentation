---
title: "Middleware.Security"
package: "http"
import: "github.com/sahilkhaire/gox/http"
node: "helmet"
gox-doc-version: "11"
---

<SymbolHeader pkg="http" title="Middleware.Security" node="helmet" import-path="github.com/sahilkhaire/gox/http" />
## Overview

Security sets common security headers (helmet-style).

If you are coming from Node.js, the closest pattern is **`helmet`**.

Method on **`Middleware`** — call it on a value of that type after constructing or receiving one from a constructor.

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

## Example

```go
import "github.com/sahilkhaire/gox/http"

http.Security()
```

## Tips

Stack `Logger`, `Recover`, and `Security` middleware the way you would morgan + helmet in Express.

## Standard library alternative

Use `net/http` with handler functions `func(w http.ResponseWriter, r *http.Request)` or a router like chi/echo directly.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/http/middleware-cors">Middleware.CORS</a><a class="related-chip" href="/packages/http/middleware-logger">Middleware.Logger</a><a class="related-chip" href="/packages/http/middleware-rate-limit">Middleware.RateLimit</a>
</div>

← [Back to http package overview](/packages/http/)
