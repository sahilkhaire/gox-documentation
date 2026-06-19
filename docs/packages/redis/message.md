---
title: "Message"
package: "redis"
import: "github.com/sahilkhaire/gox/redis"
gox-doc-version: "10"
---

<SymbolHeader pkg="redis" title="Message" node="ioredis" import-path="github.com/sahilkhaire/gox/redis" />
## Overview

Message is a pub/sub payload.

## Signature

<div class="signature-block">

```go
type Message struct {
	Channel string
	Payload string
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

_ = redis.Message
```

:::

← [Back to redis package overview](/packages/redis/)
