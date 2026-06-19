---
title: "Conn.Dial"
package: "ws"
import: "github.com/sahilkhaire/gox/ws"
gox-doc-version: "10"
---

<SymbolHeader pkg="ws" title="Conn.Dial" node="ws" import-path="github.com/sahilkhaire/gox/ws" />
## Overview

Dial connects to the WebSocket url.

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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/ws/conn-upgrade">Conn.Upgrade</a>
</div>

← [Back to ws package overview](/packages/ws/)
