---
title: "ws Cookbook"
package: "ws"
gox-doc-version: "10"
---

<div class="api-meta"><span class="api-badge node">Node: ws</span><span class="api-badge import">github.com/sahilkhaire/gox/ws</span></div>
# ws Cookbook

WebSocket server and client dial.

```go
import "github.com/sahilkhaire/gox/ws"

conn, err := ws.Dial(ctx, "wss://example.com/ws", nil)
srv := ws.NewServer()
srv.Handle("/chat", func(c *ws.Conn) error { /* ... */ })
```

← [Back to ws overview](/packages/ws/)
