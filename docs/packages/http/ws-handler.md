---
title: "WSHandler"
package: "http"
import: "github.com/sahilkhaire/gox/http"
gox-doc-version: "14"
---

<SymbolHeader pkg="http" title="WSHandler" node="express, cors, helmet, morgan" import-path="github.com/sahilkhaire/gox/http" />
## Overview

WSHandler handles a WebSocket connection in an Express-style handler.

Part of the **`http`** package — Node.js analog: *express, cors, helmet, morgan*.

`WSHandler` is a type exported by gox. Methods on this type are documented separately.

## Signature

<div class="signature-block">

```go
type WSHandler func(*Ctx, *goxws.Conn) error
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

app.HandleWS("/ws", func(conn *ws.Conn) error { return nil })
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/http"

app.HandleWS("/ws", func(conn *ws.Conn) error { return nil })
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
<a class="related-chip" href="/packages/http/save-uploaded-file">SaveUploadedFile</a>
</div>

← [Back to http package overview](/packages/http/)
