---
title: "App.New"
package: "http"
import: "github.com/sahilkhaire/gox/http"
node: "express()"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: express()</span><span class="api-badge import">github.com/sahilkhaire/gox/http</span></div>
# App.New

## Overview

Creates a new Express-style application with chi router underneath. Register routes with `Get`/`Post`/… and global middleware with `Use`.

**Node.js equivalent:** `express()`

## Signature

```go
func New() *App
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const app = express();
```

```go [Standard Go]
mux := http.NewServeMux()
http.ListenAndServe(":8080", mux)
```

```go [gox]
import "github.com/sahilkhaire/gox/http"

app := goxhttp.New()
```

:::

## Tips

Stack `Logger`, `Recover`, and `Security` middleware like you would morgan + helmet in Express. See the [HTTP guide](/guide/http).

← [Back to http package overview](/packages/http/)
