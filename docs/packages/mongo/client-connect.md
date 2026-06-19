---
title: "Client.Connect"
package: "mongo"
import: "github.com/sahilkhaire/gox/mongo"
node: "mongoose.connect(uri)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: mongoose.connect(uri)</span><span class="api-badge import">github.com/sahilkhaire/gox/mongo</span></div>
# Client.Connect

## Overview

Connect dials uri.

**Node.js equivalent:** `mongoose.connect(uri)`

## Signature

```go
func Connect(ctx context.Context, uri string) (*Client, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
await mongoose.connect(uri);
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/mongo"

client, err := mongo.Connect(ctx, uri)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

← [Back to mongo package overview](/packages/mongo/)
