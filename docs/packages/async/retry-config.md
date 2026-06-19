---
title: "RetryConfig"
package: "async"
import: "github.com/sahilkhaire/gox/async"
gox-doc-version: "14"
---

<SymbolHeader pkg="async" title="RetryConfig" node="Promise.all, timers" import-path="github.com/sahilkhaire/gox/async" />
## Overview

RetryConfig configures Retry backoff.

Part of the **`async`** package — Node.js analog: *Promise.all, timers*.

`RetryConfig` is a type exported by gox. Methods on this type are documented separately.

## Signature

<div class="signature-block">

```go
type RetryConfig struct {
	Attempts int
	Initial  time.Duration
	Max      time.Duration
	Factor   float64
}
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical Promise.all, timers pattern in Node.js
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches.
```

```go [gox]
import "github.com/sahilkhaire/gox/async"

cfg := async.RetryConfig{MaxAttempts: 3, Delay: time.Second}
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/async"

cfg := async.RetryConfig{MaxAttempts: 3, Delay: time.Second}
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
