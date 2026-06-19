---
title: "Message"
package: "redis"
import: "github.com/sahilkhaire/gox/redis"
gox-doc-version: "11"
---

<SymbolHeader pkg="redis" title="Message" node="ioredis" import-path="github.com/sahilkhaire/gox/redis" />
## Overview

Message is a pub/sub payload.

Part of the **`redis`** package — Node.js analog: *ioredis*.

`Message` is a type exported by gox. Methods on this type are documented separately.

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
// Typical ioredis pattern in Node.js
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

## Example

```go
import "github.com/sahilkhaire/gox/redis"

_ = redis.Message
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to redis package overview](/packages/redis/)
