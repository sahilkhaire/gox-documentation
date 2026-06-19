---
title: "Conn.Dial"
package: "ws"
import: "github.com/sahilkhaire/gox/ws"
gox-doc-version: "14"
---

<SymbolHeader pkg="ws" title="Conn.Dial" node="ws" import-path="github.com/sahilkhaire/gox/ws" />
## Overview

Dial connects to the WebSocket url.

Part of the **`ws`** package — Node.js analog: *ws*.

Method on **`Conn`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func Dial(ctx context.Context, url string, header http.Header) (*Conn, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const ws = new WebSocket('wss://example.com');
```

```go [Standard Go]
// gorilla/websocket Upgrader or Dialer
```

```go [gox]
import "github.com/sahilkhaire/gox/ws"

conn, err := ws.Dial(ctx, "wss://example.com", nil)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/ws"

conn, err := ws.Dial(ctx, "wss://example.com", nil)
```

## Tips

Obtain a `Conn` value first (see constructors on the package overview), then call `Dial`.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/ws/conn-upgrade">Conn.Upgrade</a>
</div>

← [Back to ws package overview](/packages/ws/)
