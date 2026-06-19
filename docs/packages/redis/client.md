---
title: "Client"
package: "redis"
import: "github.com/sahilkhaire/gox/redis"
gox-doc-version: "14"
---

<SymbolHeader pkg="redis" title="Client" node="ioredis" import-path="github.com/sahilkhaire/gox/redis" />
## Overview

Client wraps go-redis Client.

Part of the **`redis`** package — Node.js analog: *ioredis*.

`Client` is a type exported by gox. Methods on this type are documented separately.

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
// Typical ioredis pattern in Node.js
```

```go [Standard Go]
rdb := redis.NewClient(&redis.Options{Addr: addr})
val, err := rdb.Get(ctx, key).Result()
```

```go [gox]
import "github.com/sahilkhaire/gox/redis"

rdb, err := redis.New("localhost:6379")
if err != nil {
    return err
}
val, err := rdb.Get(ctx, "session:1")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/redis"

rdb, err := redis.New("localhost:6379")
if err != nil {
    return err
}
val, err := rdb.Get(ctx, "session:1")
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

Use the standard library directly:

```go
rdb := redis.NewClient(&redis.Options{Addr: addr})
val, err := rdb.Get(ctx, key).Result()
```

← [Back to redis package overview](/packages/redis/)
