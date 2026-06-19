---
title: "All"
package: "async"
import: "github.com/sahilkhaire/gox/async"
node: "Promise.all([a(), b()])"
gox-doc-version: "11"
---

<SymbolHeader pkg="async" title="All" node="Promise.all([a(), b()])" import-path="github.com/sahilkhaire/gox/async" />
## Overview

All runs tasks concurrently and returns when all complete or ctx is cancelled.

If you are coming from Node.js, the closest pattern is **`Promise.all([a(), b()])`**.

## Signature

<div class="signature-block">

```go
func All(ctx context.Context, tasks ...func(context.Context) error) error
```

</div>

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

## Example

```go
import "github.com/sahilkhaire/gox/async"

a, b, err := async.All(ctx, fetchA, fetchB)
```

## Tips

All async helpers respect context cancellation — prefer them over raw goroutines when you need timeouts.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/async/after">After</a><a class="related-chip" href="/packages/async/race">Race</a><a class="related-chip" href="/packages/async/retry">Retry</a>
</div>

← [Back to async package overview](/packages/async/)
