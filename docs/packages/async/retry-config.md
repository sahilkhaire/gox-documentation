---
title: "RetryConfig"
package: "async"
import: "github.com/sahilkhaire/gox/async"
gox-doc-version: "10"
---

<SymbolHeader pkg="async" title="RetryConfig" node="Promise.all, timers" import-path="github.com/sahilkhaire/gox/async" />
## Overview

RetryConfig configures Retry backoff.

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
// See package overview
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches.
```

```go [gox]
import "github.com/sahilkhaire/gox/async"

_ = async.RetryConfig
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/async/after">After</a><a class="related-chip" href="/packages/async/all">All</a><a class="related-chip" href="/packages/async/race">Race</a>
</div>

← [Back to async package overview](/packages/async/)
