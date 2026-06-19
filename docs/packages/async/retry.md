---
title: "Retry"
package: "async"
import: "github.com/sahilkhaire/gox/async"
gox-doc-version: "14"
---

<SymbolHeader pkg="async" title="Retry" node="Promise.all, timers" import-path="github.com/sahilkhaire/gox/async" />
## Overview

Retry calls fn until it succeeds or attempts are exhausted.

Part of the **`async`** package — Node.js analog: *Promise.all, timers*.

## Signature

<div class="signature-block">

```go
func Retry(ctx context.Context, cfg RetryConfig, fn func(context.Context) error) error
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical Promise.all, timers pattern in Node.js
```

```go [Standard Go]
// Loop with backoff or context.WithTimeout
```

```go [gox]
import "github.com/sahilkhaire/gox/async"

err := async.Retry(ctx, async.RetryConfig{MaxAttempts: 3, Delay: time.Second}, func(ctx context.Context) error {
    return fetch(ctx)
})
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/async"

err := async.Retry(ctx, async.RetryConfig{MaxAttempts: 3, Delay: time.Second}, func(ctx context.Context) error {
    return fetch(ctx)
})
```

## Tips

All async helpers respect context cancellation — prefer them over raw goroutines when you need timeouts.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/async/after">After</a><a class="related-chip" href="/packages/async/all">All</a><a class="related-chip" href="/packages/async/race">Race</a>
</div>

← [Back to async package overview](/packages/async/)
