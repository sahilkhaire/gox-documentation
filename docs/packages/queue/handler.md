---
title: "Handler"
package: "queue"
import: "github.com/sahilkhaire/gox/queue"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: bull</span><span class="api-badge import">github.com/sahilkhaire/gox/queue</span></div>
# Handler

## Overview

Handler processes a task payload.

## Signature

```go
type Handler func(ctx context.Context, payload []byte) error
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
import "github.com/sahilkhaire/gox/queue"

_ = queue.Handler
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

← [Back to queue package overview](/packages/queue/)
