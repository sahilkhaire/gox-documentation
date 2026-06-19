---
title: "EventStream.SSE"
package: "http"
import: "github.com/sahilkhaire/gox/http"
gox-doc-version: "10"
---

<SymbolHeader pkg="http" title="EventStream.SSE" node="express, cors, helmet, morgan" import-path="github.com/sahilkhaire/gox/http" />
## Overview

SSE prepares the response for server-sent events.

## Signature

<div class="signature-block">

```go
func SSE(c *Ctx) (*EventStream, error)
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

var v EventStream
v.SSE(/* args */)
```

:::

← [Back to http package overview](/packages/http/)
