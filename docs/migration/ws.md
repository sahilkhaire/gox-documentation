# WebSocket migration (`ws` → `gox/ws`)

Use `gox/ws` for low-level WebSocket servers and clients (similar to the [`ws`](https://www.npmjs.com/package/ws) package). For Socket.IO-style rooms and fallbacks, use raw WebSockets or a dedicated Go library and document app-specific patterns.

## Server

```go
srv := ws.NewServer(nil)
http.Handle("/socket", srv.HandleWS(func(c *ws.Conn) {
    for {
        var msg map[string]any
        if err := c.ReadJSON(&msg); err != nil {
            return
        }
        _ = c.WriteJSON(msg)
    }
}))
```

With `gox/http`:

```go
app := http.New()
app.HandleWS("/ws", nil, func(c *http.Ctx, conn *ws.Conn) error {
    // ...
    return nil
})
```

## Client

```go
conn, err := ws.Dial(ctx, "wss://example.com/ws", nil)
defer conn.Close()
_ = conn.Send(map[string]string{"type": "ping"})
```
