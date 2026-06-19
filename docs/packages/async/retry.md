---
title: "Retry"
package: "async"
import: "github.com/sahilkhaire/gox/async"
gox-doc-version: "10"
---

<SymbolHeader pkg="async" title="Retry" node="Promise.all, timers" import-path="github.com/sahilkhaire/gox/async" />
## Overview

Retry calls fn until it succeeds or attempts are exhausted.

## Signature

<div class="signature-block">

```go
func Retry(ctx context.Context, cfg RetryConfig, fn func(context.Context) error) error
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
// Loop with backoff or context.WithTimeout
```

```go [gox]
import "github.com/sahilkhaire/gox/async"

// async
_ = async.Retry(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/async/after">After</a><a class="related-chip" href="/packages/async/all">All</a><a class="related-chip" href="/packages/async/race">Race</a>
</div>

← [Back to async package overview](/packages/async/)
