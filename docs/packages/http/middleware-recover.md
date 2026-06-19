---
title: "Middleware.Recover"
package: "http"
import: "github.com/sahilkhaire/gox/http"
gox-doc-version: "14"
---

<SymbolHeader pkg="http" title="Middleware.Recover" node="express, cors, helmet, morgan" import-path="github.com/sahilkhaire/gox/http" />
## Overview

Recover catches panics and returns 500.

Part of the **`http`** package — Node.js analog: *express, cors, helmet, morgan*.

Method on **`Middleware`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func Recover() Middleware
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

app.Use(http.Recover())
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/http"

app.Use(http.Recover())
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
<a class="related-chip" href="/packages/http/middleware-cors">Middleware.CORS</a><a class="related-chip" href="/packages/http/middleware-logger">Middleware.Logger</a><a class="related-chip" href="/packages/http/middleware-rate-limit">Middleware.RateLimit</a>
</div>

← [Back to http package overview](/packages/http/)
