---
title: "Conn.Dial"
package: "ws"
import: "github.com/sahilkhaire/gox/ws"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: ws</span><span class="api-badge import">github.com/sahilkhaire/gox/ws</span></div>
# Conn.Dial

## Overview

Dial connects to the WebSocket url.

## Signature

```go
func Dial(ctx context.Context, url string, header http.Header) (*Conn, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const ws = new WebSocket('wss://example.com');
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/ws"

conn, err := ws.Dial(ctx, "wss://example.com", nil)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Conn.Upgrade](/packages/ws/conn-upgrade)

← [Back to ws package overview](/packages/ws/)
