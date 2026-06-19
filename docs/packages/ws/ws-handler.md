---
title: "WSHandler"
package: "ws"
import: "github.com/sahilkhaire/gox/ws"
gox-doc-version: "14"
---

<SymbolHeader pkg="ws" title="WSHandler" node="ws" import-path="github.com/sahilkhaire/gox/ws" />
## Overview

WSHandler is a WebSocket connection handler.

Part of the **`ws`** package — Node.js analog: *ws*.

`WSHandler` is a type exported by gox. Methods on this type are documented separately.

## Signature

<div class="signature-block">

```go
type WSHandler func(*Conn)
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

var _ ws.WSHandler = func(conn *ws.Conn) error { return conn.WriteJSON(map[string]string{"ok": "1"}) }
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/ws"

var _ ws.WSHandler = func(conn *ws.Conn) error { return conn.WriteJSON(map[string]string{"ok": "1"}) }
```

## Tips

Browse methods on this type in the sidebar for handler-style APIs and options structs.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to ws package overview](/packages/ws/)
