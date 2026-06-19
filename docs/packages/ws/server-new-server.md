---
title: "Server.NewServer"
package: "ws"
import: "github.com/sahilkhaire/gox/ws"
gox-doc-version: "10"
---

<SymbolHeader pkg="ws" title="Server.NewServer" node="ws" import-path="github.com/sahilkhaire/gox/ws" />
## Overview

NewServer creates a WebSocket server with the given upgrader (nil uses DefaultUpgrader).

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

← [Back to ws package overview](/packages/ws/)
