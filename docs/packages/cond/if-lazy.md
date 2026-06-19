---
title: "IfLazy"
package: "cond"
import: "github.com/sahilkhaire/gox/cond"
node: "cond ? f() : g()"
gox-doc-version: "11"
---

<SymbolHeader pkg="cond" title="IfLazy" node="cond ? f() : g()" import-path="github.com/sahilkhaire/gox/cond" />
## Overview

IfLazy evaluates a or b lazily via thunks (Node: cond ? a() : b()).

If you are coming from Node.js, the closest pattern is **`cond ? f() : g()`**.

## Signature

<div class="signature-block">

```go
func IfLazy[T any](cond bool, a, b func() T) T
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const v = cond ? expensive() : fallback();
```

```go [Standard Go]
var v T
if cond {
    v = expensive()
} else {
    v = fallback()
}
```

```go [gox]
import "github.com/sahilkhaire/gox/cond"

v := cond.IfLazy(cond, expensive, fallback)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/cond"

called := 0
got := IfLazy(false, func() int { called++; return 1 }, func() int { return 2 })
got = IfLazy(true, func() int { return 3 }, func() int { called++; return 4 })
```

## Tips

Import `github.com/sahilkhaire/gox/cond` and call `IfLazy` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use `if/else` for branching and explicit nil checks instead of `??`.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/cond/coalesce">Coalesce</a><a class="related-chip" href="/packages/cond/coalesce-fn">CoalesceFn</a><a class="related-chip" href="/packages/cond/coalesce-ptr">CoalescePtr</a>
</div>

← [Back to cond package overview](/packages/cond/)
