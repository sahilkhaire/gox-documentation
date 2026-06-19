---
title: "After"
package: "async"
import: "github.com/sahilkhaire/gox/async"
node: "setTimeout(fn, d)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: setTimeout(fn, d)</span><span class="api-badge import">github.com/sahilkhaire/gox/async</span></div>
# After

## Overview

Maps the Node.js pattern `setTimeout(fn, d)` to gox `async.After(ctx, d, fn)`.

**Node.js equivalent:** `setTimeout(fn, d)`

## Signature

```go
func After(d time.Duration) <-chan time.Time
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
setTimeout(fn, d)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/async"

async.After(ctx, d, fn)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [All](/packages/async/all)
- [Race](/packages/async/race)
- [Retry](/packages/async/retry)

← [Back to async package overview](/packages/async/)
