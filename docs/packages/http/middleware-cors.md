---
title: "Middleware.CORS"
package: "http"
import: "github.com/sahilkhaire/gox/http"
node: "cors"
gox-doc-version: "14"
---

<SymbolHeader pkg="http" title="Middleware.CORS" node="cors" import-path="github.com/sahilkhaire/gox/http" />
## Overview

CORS adds Cross-Origin Resource Sharing headers.

If you are coming from Node.js, the closest pattern is **`cors`**.

Method on **`Middleware`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func CORS(opts CORSOptions) Middleware
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
cors
```

```go [Standard Go]
func handler(w http.ResponseWriter, r *http.Request) {
    // chi or net/http
}
```

```go [gox]
import "github.com/sahilkhaire/gox/http"

http.CORS(opts)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/http"

http.CORS(opts)
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
<a class="related-chip" href="/packages/http/middleware-logger">Middleware.Logger</a><a class="related-chip" href="/packages/http/middleware-rate-limit">Middleware.RateLimit</a><a class="related-chip" href="/packages/http/middleware-recover">Middleware.Recover</a>
</div>

← [Back to http package overview](/packages/http/)
