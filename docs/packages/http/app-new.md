---
title: "App.New"
package: "http"
import: "github.com/sahilkhaire/gox/http"
node: "express()"
gox-doc-version: "14"
---

<SymbolHeader pkg="http" title="App.New" node="express()" import-path="github.com/sahilkhaire/gox/http" />
## Overview

Creates a new Express-style application with chi router underneath. Register routes with `Get`/`Post`/… and global middleware with `Use`.

If you are coming from Node.js, the closest pattern is **`express()`**.

Method on **`App`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func New() *App
```

</div>

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

## Example

```go
import "github.com/sahilkhaire/gox/http"

app := goxhttp.New()
```

## Tips

Stack `Logger`, `Recover`, and `Security` middleware like you would morgan + helmet in Express. See the [HTTP guide](/guide/http).

## Standard library alternative

Use the standard library directly:

```go
mux := http.NewServeMux()
http.ListenAndServe(":8080", mux)
```

← [Back to http package overview](/packages/http/)
