---
title: "IfLazy"
package: "cond"
import: "github.com/sahilkhaire/gox/cond"
node: "cond ? f() : g()"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: cond ? f() : g()</span><span class="api-badge import">github.com/sahilkhaire/gox/cond</span></div>
# IfLazy

## Overview

IfLazy evaluates a or b lazily via thunks (Node: cond ? a() : b()).

**Node.js equivalent:** `cond ? f() : g()`

## Signature

```go
func IfLazy[T any](cond bool, a, b func() T) T
```

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

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

```go
import "github.com/sahilkhaire/gox/cond"

called := 0
got := IfLazy(false, func() int { called++; return 1 }, func() int { return 2 })
got = IfLazy(true, func() int { return 3 }, func() int { called++; return 4 })
```

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Coalesce](/packages/cond/coalesce)
- [CoalesceFn](/packages/cond/coalesce-fn)
- [CoalescePtr](/packages/cond/coalesce-ptr)

← [Back to cond package overview](/packages/cond/)
