---
title: "Retry"
package: "async"
import: "github.com/sahilkhaire/gox/async"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: Promise.all, timers</span><span class="api-badge import">github.com/sahilkhaire/gox/async</span></div>
# Retry

## Overview

Retry calls fn until it succeeds or attempts are exhausted.

## Signature

```go
func Retry(ctx context.Context, cfg RetryConfig, fn func(context.Context) error) error
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
import "github.com/sahilkhaire/gox/async"

// async
_ = async.Retry(/* args */)
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
- [Race](/packages/async/race)

← [Back to async package overview](/packages/async/)
