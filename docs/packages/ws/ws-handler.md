---
title: "WSHandler"
package: "ws"
import: "github.com/sahilkhaire/gox/ws"
gox-doc-version: "10"
---

<SymbolHeader pkg="ws" title="WSHandler" node="ws" import-path="github.com/sahilkhaire/gox/ws" />
## Overview

WSHandler is a WebSocket connection handler.

## Signature

<div class="signature-block">

```go
type WSHandler func(*Conn)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
// gorilla/websocket Upgrader or Dialer
```

```go [gox]
import "github.com/sahilkhaire/gox/ws"

_ = ws.WSHandler
```

:::

← [Back to ws package overview](/packages/ws/)
