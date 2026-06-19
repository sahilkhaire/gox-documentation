---
title: "Client.New"
package: "redis"
import: "github.com/sahilkhaire/gox/redis"
node: "new Redis()"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: new Redis()</span><span class="api-badge import">github.com/sahilkhaire/gox/redis</span></div>
# Client.New

## Overview

New connects to addr (host:port).

**Node.js equivalent:** `new Redis()`

## Signature

```go
func New(addr string) *Client
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const redis = new Redis();
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/redis"

rdb, err := redis.New("localhost:6379")
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Client.NewFromClient](/packages/redis/client-new-from-client)

← [Back to redis package overview](/packages/redis/)
