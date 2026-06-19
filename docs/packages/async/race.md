---
title: "Race"
package: "async"
import: "github.com/sahilkhaire/gox/async"
node: "Promise.race([a(), b()])"
gox-doc-version: "11"
---

<SymbolHeader pkg="async" title="Race" node="Promise.race([a(), b()])" import-path="github.com/sahilkhaire/gox/async" />
## Overview

Race runs tasks concurrently and returns the first successful result or error.

If you are coming from Node.js, the closest pattern is **`Promise.race([a(), b()])`**.

## Signature

<div class="signature-block">

```go
func Race[T any](ctx context.Context, tasks ...func(context.Context) (T, error)) (T, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
Promise.race([a(), b()])
```

```go [Standard Go]
select {
case <-doneA:
case <-doneB:
case <-ctx.Done():
}
```

```go [gox]
import "github.com/sahilkhaire/gox/async"

async.Race(ctx, a, b)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/async"

async.Race(ctx, a, b)
```

## Tips

All async helpers respect context cancellation — prefer them over raw goroutines when you need timeouts.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/async/after">After</a><a class="related-chip" href="/packages/async/all">All</a><a class="related-chip" href="/packages/async/retry">Retry</a>
</div>

← [Back to async package overview](/packages/async/)
