---
title: "Client.NewFromClient"
package: "redis"
import: "github.com/sahilkhaire/gox/redis"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: ioredis</span><span class="api-badge import">github.com/sahilkhaire/gox/redis</span></div>
# Client.NewFromClient

## Overview

NewFromClient wraps an existing client.

## Signature

```go
func NewFromClient(c *goredis.Client) *Client
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/redis"

var v Client
v.NewFromClient(/* args */)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Client.New](/packages/redis/client-new)

← [Back to redis package overview](/packages/redis/)
