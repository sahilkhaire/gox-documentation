---
title: "Server"
package: "ws"
import: "github.com/sahilkhaire/gox/ws"
gox-doc-version: "10"
---

<SymbolHeader pkg="ws" title="Server" node="ws" import-path="github.com/sahilkhaire/gox/ws" />
## Overview

Server registers WebSocket routes.

## Signature

<div class="signature-block">

```go
type Server struct {
	Upgrader Upgrader
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

_ = ws.Server
```

:::

← [Back to ws package overview](/packages/ws/)
