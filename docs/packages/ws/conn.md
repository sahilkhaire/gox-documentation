---
title: "Conn"
package: "ws"
import: "github.com/sahilkhaire/gox/ws"
gox-doc-version: "10"
---

<SymbolHeader pkg="ws" title="Conn" node="ws" import-path="github.com/sahilkhaire/gox/ws" />
## Overview

Conn wraps a gorilla WebSocket connection with JSON helpers.

## Signature

<div class="signature-block">

```go
type Conn struct {
	*websocket.Conn
}
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

_ = ws.Conn
```

:::

← [Back to ws package overview](/packages/ws/)
