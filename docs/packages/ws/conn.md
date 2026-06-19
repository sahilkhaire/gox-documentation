---
title: "Conn"
package: "ws"
import: "github.com/sahilkhaire/gox/ws"
gox-doc-version: "14"
---

<SymbolHeader pkg="ws" title="Conn" node="ws" import-path="github.com/sahilkhaire/gox/ws" />
## Overview

Conn wraps a gorilla WebSocket connection with JSON helpers.

Part of the **`ws`** package — Node.js analog: *ws*.

`Conn` is a type exported by gox. Methods on this type are documented separately.

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
// Typical ws pattern in Node.js
```

```go [Standard Go]
// gorilla/websocket Upgrader or Dialer
```

```go [gox]
import "github.com/sahilkhaire/gox/ws"

if err := conn.WriteJSON(map[string]string{"event": "ping"}); err != nil {
    return err
}
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/ws"

if err := conn.WriteJSON(map[string]string{"event": "ping"}); err != nil {
    return err
}
```

## Tips

Browse methods on this type in the sidebar for handler-style APIs and options structs.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to ws package overview](/packages/ws/)
