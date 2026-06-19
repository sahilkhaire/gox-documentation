---
title: "Client.NewFromClient"
package: "redis"
import: "github.com/sahilkhaire/gox/redis"
gox-doc-version: "11"
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

var v Client
v.NewFromClient(/* args */)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/redis"

var v Client
v.NewFromClient(/* args */)
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/redis/client-new">Client.New</a>
</div>

← [Back to redis package overview](/packages/redis/)
