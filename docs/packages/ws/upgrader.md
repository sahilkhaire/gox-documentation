---
title: "Upgrader"
package: "ws"
import: "github.com/sahilkhaire/gox/ws"
gox-doc-version: "10"
---

<SymbolHeader pkg="ws" title="Upgrader" node="ws" import-path="github.com/sahilkhaire/gox/ws" />
## Overview

Upgrader upgrades HTTP connections to WebSocket.

## Signature

<div class="signature-block">

```go
type Upgrader = websocket.Upgrader
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

_ = ws.Upgrader
```

:::

← [Back to ws package overview](/packages/ws/)
