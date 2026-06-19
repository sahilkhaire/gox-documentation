---
title: "Sleep"
package: "async"
import: "github.com/sahilkhaire/gox/async"
node: "await sleep(ms)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: await sleep(ms)</span><span class="api-badge import">github.com/sahilkhaire/gox/async</span></div>
# Sleep

## Overview

Sleep waits until duration elapses or ctx is cancelled.

**Node.js equivalent:** `await sleep(ms)`

## Signature

```go
func Sleep(ctx context.Context, d time.Duration) error
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
await new Promise(r => setTimeout(r, 1000));
```

```go [Standard Go]
select {
case <-time.After(time.Second):
case <-ctx.Done():
    return ctx.Err()
}
```

```go [gox]
import "github.com/sahilkhaire/gox/async"

if err := async.Sleep(ctx, time.Second); err != nil { return err }
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
