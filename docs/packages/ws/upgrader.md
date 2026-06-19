---
title: "Upgrader"
package: "ws"
import: "github.com/sahilkhaire/gox/ws"
gox-doc-version: "14"
---

<SymbolHeader pkg="ws" title="Upgrader" node="ws" import-path="github.com/sahilkhaire/gox/ws" />
## Overview

Upgrader upgrades HTTP connections to WebSocket.

Part of the **`ws`** package — Node.js analog: *ws*.

`Upgrader` is a type exported by gox. Methods on this type are documented separately.

## Signature

<div class="signature-block">

```go
type Upgrader = websocket.Upgrader
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical ws pattern in Node.js
```

```go [Standard Go]
// gorilla/websocket Upgrader or Dialer
```

```go [gox]
import "github.com/sahilkhaire/gox/ws"

upgrader := ws.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/ws"

upgrader := ws.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
```

## Tips

Browse methods on this type in the sidebar for handler-style APIs and options structs.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to ws package overview](/packages/ws/)
