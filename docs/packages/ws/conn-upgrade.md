---
title: "Conn.Upgrade"
package: "ws"
import: "github.com/sahilkhaire/gox/ws"
gox-doc-version: "10"
---

<SymbolHeader pkg="ws" title="Conn.Upgrade" node="ws" import-path="github.com/sahilkhaire/gox/ws" />
## Overview

Upgrade upgrades the HTTP request on w/r using up (nil uses DefaultUpgrader).

## Signature

<div class="signature-block">

```go
func Upgrade(w http.ResponseWriter, r *http.Request, up *Upgrader) (*Conn, error)
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

var v Conn
v.Upgrade(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/ws/conn-dial">Conn.Dial</a>
</div>

← [Back to ws package overview](/packages/ws/)
