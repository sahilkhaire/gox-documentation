---
title: "EventStream.SSE"
package: "http"
import: "github.com/sahilkhaire/gox/http"
gox-doc-version: "14"
---

<SymbolHeader pkg="http" title="EventStream.SSE" node="express, cors, helmet, morgan" import-path="github.com/sahilkhaire/gox/http" />
## Overview

SSE prepares the response for server-sent events.

Part of the **`http`** package — Node.js analog: *express, cors, helmet, morgan*.

Method on **`EventStream`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func SSE(c *Ctx) (*EventStream, error)
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

stream, err := http.SSE(w, r)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/http"

stream, err := http.SSE(w, r)
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

← [Back to http package overview](/packages/http/)
