---
title: "Client.New"
package: "redis"
import: "github.com/sahilkhaire/gox/redis"
node: "new Redis()"
gox-doc-version: "14"
---

<SymbolHeader pkg="redis" title="Client.New" node="new Redis()" import-path="github.com/sahilkhaire/gox/redis" />
## Overview

New connects to addr (host:port).

If you are coming from Node.js, the closest pattern is **`new Redis()`**.

Method on **`Client`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func New(addr string) *Client
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const redis = new Redis();
```

```go [Standard Go]
rdb := redis.NewClient(&redis.Options{Addr: addr})
val, err := rdb.Get(ctx, key).Result()
```

```go [gox]
import "github.com/sahilkhaire/gox/redis"

rdb, err := redis.New("localhost:6379")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/redis"

rdb, err := redis.New("localhost:6379")
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
<a class="related-chip" href="/packages/redis/client-new-from-client">Client.NewFromClient</a>
</div>

← [Back to redis package overview](/packages/redis/)
