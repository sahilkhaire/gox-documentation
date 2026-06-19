---
title: "Client.NewFromClient"
package: "redis"
import: "github.com/sahilkhaire/gox/redis"
gox-doc-version: "10"
---

<SymbolHeader pkg="redis" title="Client.NewFromClient" node="ioredis" import-path="github.com/sahilkhaire/gox/redis" />
## Overview

NewFromClient wraps an existing client.

## Signature

<div class="signature-block">

```go
func NewFromClient(c *goredis.Client) *Client
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/redis/client-new">Client.New</a>
</div>

← [Back to redis package overview](/packages/redis/)
