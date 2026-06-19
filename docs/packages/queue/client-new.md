---
title: "Client.New"
package: "queue"
import: "github.com/sahilkhaire/gox/queue"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: bull</span><span class="api-badge import">github.com/sahilkhaire/gox/queue</span></div>
# Client.New

## Overview

New creates a queue client connected to redisAddr (host:port).

## Signature

```go
func New(redisAddr string) *Client
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const queue = new Queue('tasks');
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/queue"

client, err := queue.New("localhost:6379")
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

← [Back to queue package overview](/packages/queue/)
