---
title: "Race"
package: "async"
import: "github.com/sahilkhaire/gox/async"
node: "Promise.race([a(), b()])"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: Promise.race([a(), b()])</span><span class="api-badge import">github.com/sahilkhaire/gox/async</span></div>
# Race

## Overview

Maps the Node.js pattern `Promise.race([a(), b()])` to gox `async.Race(ctx, a, b)`.

**Node.js equivalent:** `Promise.race([a(), b()])`

## Signature

```go
func Race[T any](ctx context.Context, tasks ...func(context.Context) (T, error)) (T, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
Promise.race([a(), b()])
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/async"

async.Race(ctx, a, b)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [After](/packages/async/after)
- [All](/packages/async/all)
- [Retry](/packages/async/retry)

← [Back to async package overview](/packages/async/)
