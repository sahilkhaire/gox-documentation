---
title: "Client.NewFromClient"
package: "redis"
import: "github.com/sahilkhaire/gox/redis"
gox-doc-version: "14"
---

<SymbolHeader pkg="redis" title="Client.NewFromClient" node="ioredis" import-path="github.com/sahilkhaire/gox/redis" />
## Overview

NewFromClient wraps an existing client.

Part of the **`redis`** package — Node.js analog: *ioredis*.

Method on **`Client`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func NewFromClient(c *goredis.Client) *Client
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical ioredis pattern in Node.js
```

```go [Standard Go]
rdb := redis.NewClient(&redis.Options{Addr: addr})
val, err := rdb.Get(ctx, key).Result()
```

```go [gox]
import "github.com/sahilkhaire/gox/redis"

wrapper := redis.NewFromClient(rdb)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/redis"

wrapper := redis.NewFromClient(rdb)
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

Use the standard library directly:

```go
rdb := redis.NewClient(&redis.Options{Addr: addr})
val, err := rdb.Get(ctx, key).Result()
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/redis/client-new">Client.New</a>
</div>

← [Back to redis package overview](/packages/redis/)
