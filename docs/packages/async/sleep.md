---
title: "Sleep"
package: "async"
import: "github.com/sahilkhaire/gox/async"
node: "await sleep(ms)"
gox-doc-version: "10"
---

<SymbolHeader pkg="async" title="Sleep" node="await sleep(ms)" import-path="github.com/sahilkhaire/gox/async" />
## Overview

Sleep waits until duration elapses or ctx is cancelled.

**Node.js equivalent:** `await sleep(ms)`

## Signature

<div class="signature-block">

```go
func Sleep(ctx context.Context, d time.Duration) error
```

</div>

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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/async/after">After</a><a class="related-chip" href="/packages/async/all">All</a><a class="related-chip" href="/packages/async/race">Race</a>
</div>

← [Back to async package overview](/packages/async/)
