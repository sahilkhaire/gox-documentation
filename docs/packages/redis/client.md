---
title: "Client"
package: "redis"
import: "github.com/sahilkhaire/gox/redis"
gox-doc-version: "10"
---

<SymbolHeader pkg="redis" title="Client" node="ioredis" import-path="github.com/sahilkhaire/gox/redis" />
## Overview

Client wraps go-redis Client.

## Signature

<div class="signature-block">

```go
type Client struct {
	RDB *goredis.Client
}
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

_ = redis.Client
```

:::

← [Back to redis package overview](/packages/redis/)
