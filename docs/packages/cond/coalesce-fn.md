---
title: "CoalesceFn"
package: "cond"
import: "github.com/sahilkhaire/gox/cond"
node: "obj?.field ?? \"guest\""
gox-doc-version: "11"
---

<SymbolHeader pkg="cond" title="CoalesceFn" node="obj?.field ?? &quot;guest&quot;" import-path="github.com/sahilkhaire/gox/cond" />
## Overview

CoalesceFn returns the first non-zero result from the given functions, in order.

If you are coming from Node.js, the closest pattern is **`obj?.field ?? "guest"`**.

## Signature

<div class="signature-block">

```go
func CoalesceFn[T comparable](fns ...func() T) T
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
obj?.field ?? "guest"
```

```go [Standard Go]
result := fallback
if v != zero {
    result = v
}
```

```go [gox]
import "github.com/sahilkhaire/gox/cond"

cond.CoalesceFn(ptr, func(v T) R { ... }, fallback)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/cond"

cond.CoalesceFn(ptr, func(v T) R { ... }, fallback)
```

## Tips

Prefer explicit zero-value checks in performance-critical hot paths if the compiler cannot prove types.

## Standard library alternative

Use `if/else` for branching and explicit nil checks instead of `??`.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/cond/coalesce">Coalesce</a><a class="related-chip" href="/packages/cond/coalesce-ptr">CoalescePtr</a><a class="related-chip" href="/packages/cond/if">If</a>
</div>

← [Back to cond package overview](/packages/cond/)
