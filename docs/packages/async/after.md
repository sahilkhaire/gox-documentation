---
title: "After"
package: "async"
import: "github.com/sahilkhaire/gox/async"
node: "setTimeout(fn, d)"
gox-doc-version: "14"
---

<SymbolHeader pkg="async" title="After" node="setTimeout(fn, d)" import-path="github.com/sahilkhaire/gox/async" />
## Overview

After sends the time on the returned channel when duration elapses.

If you are coming from Node.js, the closest pattern is **`setTimeout(fn, d)`**.

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

## Example

```go
import "github.com/sahilkhaire/gox/async"

async.After(ctx, d, fn)
```

## Tips

All async helpers respect context cancellation — prefer them over raw goroutines when you need timeouts.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/async/all">All</a><a class="related-chip" href="/packages/async/race">Race</a><a class="related-chip" href="/packages/async/retry">Retry</a>
</div>

← [Back to async package overview](/packages/async/)
