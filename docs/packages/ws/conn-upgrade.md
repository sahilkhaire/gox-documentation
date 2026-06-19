---
title: "Conn.Upgrade"
package: "ws"
import: "github.com/sahilkhaire/gox/ws"
gox-doc-version: "11"
---

<SymbolHeader pkg="ws" title="Conn.Upgrade" node="ws" import-path="github.com/sahilkhaire/gox/ws" />
## Overview

Upgrade upgrades the HTTP request on w/r using up (nil uses DefaultUpgrader).

Part of the **`ws`** package — Node.js analog: *ws*.

Method on **`Conn`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func Upgrade(w http.ResponseWriter, r *http.Request, up *Upgrader) (*Conn, error)
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

var v Conn
v.Upgrade(/* args */)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/ws"

var v Conn
v.Upgrade(/* args */)
```

## Tips

Obtain a `Conn` value first (see constructors on the package overview), then call `Upgrade`.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/ws/conn-dial">Conn.Dial</a>
</div>

← [Back to ws package overview](/packages/ws/)
