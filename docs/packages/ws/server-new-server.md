---
title: "Server.NewServer"
package: "ws"
import: "github.com/sahilkhaire/gox/ws"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: ws</span><span class="api-badge import">github.com/sahilkhaire/gox/ws</span></div>
# Server.NewServer

## Overview

NewServer creates a WebSocket server with the given upgrader (nil uses DefaultUpgrader).

## Signature

```go
func NewServer(up *Upgrader) *Server
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const wss = new WebSocket.Server({ server });
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/ws"

srv := ws.NewServer()
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

← [Back to ws package overview](/packages/ws/)
