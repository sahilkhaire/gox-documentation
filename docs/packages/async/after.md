---
title: "After"
package: "async"
import: "github.com/sahilkhaire/gox/async"
node: "setTimeout(fn, d)"
gox-doc-version: "10"
---

<SymbolHeader pkg="async" title="After" node="setTimeout(fn, d)" import-path="github.com/sahilkhaire/gox/async" />
## Overview

After sends the time on the returned channel when duration elapses.

**Node.js equivalent:** `setTimeout(fn, d)`

## Signature

<div class="signature-block">

```go
func After(d time.Duration) <-chan time.Time
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
setTimeout(fn, d)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches.
```

```go [gox]
import "github.com/sahilkhaire/gox/async"

async.After(ctx, d, fn)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/async/all">All</a><a class="related-chip" href="/packages/async/race">Race</a><a class="related-chip" href="/packages/async/retry">Retry</a>
</div>

← [Back to async package overview](/packages/async/)
