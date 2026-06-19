---
title: "Server.NewServer"
package: "ws"
import: "github.com/sahilkhaire/gox/ws"
gox-doc-version: "11"
---

<SymbolHeader pkg="ws" title="Server.NewServer" node="ws" import-path="github.com/sahilkhaire/gox/ws" />
## Overview

NewServer creates a WebSocket server with the given upgrader (nil uses DefaultUpgrader).

Part of the **`ws`** package — Node.js analog: *ws*.

Method on **`Server`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func NewServer(up *Upgrader) *Server
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const wss = new WebSocket.Server({ server });
```

```go [Standard Go]
// gorilla/websocket Upgrader or Dialer
```

```go [gox]
import "github.com/sahilkhaire/gox/ws"

srv := ws.NewServer()
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/ws"

srv := ws.NewServer()
```

## Tips

Obtain a `Server` value first (see constructors on the package overview), then call `NewServer`.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to ws package overview](/packages/ws/)
