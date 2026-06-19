---
title: "If"
package: "cond"
import: "github.com/sahilkhaire/gox/cond"
node: "cond ? a : b"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: cond ? a : b</span><span class="api-badge import">github.com/sahilkhaire/gox/cond</span></div>
# If

## Overview

Returns `a` when `cond` is true, otherwise `b`. The closest equivalent to JavaScript's ternary operator — both branches must be the same type in Go.

**Node.js equivalent:** `cond ? a : b`

## Signature

```go
func If[T any](cond bool, a, b T) T
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const label = age >= 18 ? 'adult' : 'minor';
```

```go [Standard Go]
label := "minor"
if age >= 18 {
    label = "adult"
}
```

```go [gox]
import "github.com/sahilkhaire/gox/cond"

label := cond.If(age >= 18, "adult", "minor")
```

:::

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

```go
import "github.com/sahilkhaire/gox/cond"

adult := cond.If(age >= 18, "adult", "minor")
```

## Tips

Both branches are evaluated eagerly. Use `cond.IfLazy` when one branch is expensive and should be skipped.

## Related APIs

- [Coalesce](/packages/cond/coalesce)
- [CoalesceFn](/packages/cond/coalesce-fn)
- [CoalescePtr](/packages/cond/coalesce-ptr)

← [Back to cond package overview](/packages/cond/)
