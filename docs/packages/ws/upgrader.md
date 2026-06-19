---
title: "Upgrader"
package: "ws"
import: "github.com/sahilkhaire/gox/ws"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: ws</span><span class="api-badge import">github.com/sahilkhaire/gox/ws</span></div>
# Upgrader

## Overview

Upgrader upgrades HTTP connections to WebSocket.

## Signature

```go
type Upgrader = websocket.Upgrader
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/ws"

_ = ws.Upgrader
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

← [Back to ws package overview](/packages/ws/)
