---
title: "All"
package: "async"
import: "github.com/sahilkhaire/gox/async"
node: "Promise.all([a(), b()])"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: Promise.all([a(), b()])</span><span class="api-badge import">github.com/sahilkhaire/gox/async</span></div>
# All

## Overview

All runs tasks concurrently and returns when all complete or ctx is cancelled.

**Node.js equivalent:** `Promise.all([a(), b()])`

## Signature

```go
func All(ctx context.Context, tasks ...func(context.Context) error) error
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const [a, b] = await Promise.all([fetchA(), fetchB()]);
```

```go [Standard Go]
errg, ctx := errgroup.WithContext(ctx)
// run goroutines...
```

```go [gox]
import "github.com/sahilkhaire/gox/async"

a, b, err := async.All(ctx, fetchA, fetchB)
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
- [Race](/packages/async/race)
- [Retry](/packages/async/retry)

← [Back to async package overview](/packages/async/)
