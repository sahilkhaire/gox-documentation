---
title: "If"
package: "cond"
import: "github.com/sahilkhaire/gox/cond"
node: "cond ? a : b"
gox-doc-version: "11"
---

<SymbolHeader pkg="cond" title="If" node="cond ? a : b" import-path="github.com/sahilkhaire/gox/cond" />
## Overview

Returns `a` when `cond` is true, otherwise `b`. The closest equivalent to JavaScript's ternary operator — both branches must be the same type in Go.

If you are coming from Node.js, the closest pattern is **`cond ? a : b`**.

## Signature

<div class="signature-block">

```go
func If[T any](cond bool, a, b T) T
```

</div>

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

## Example

```go
import "github.com/sahilkhaire/gox/cond"

adult := cond.If(age >= 18, "adult", "minor")
```

## Tips

Both branches are evaluated eagerly. Use `cond.IfLazy` when one branch is expensive and should be skipped.

## Standard library alternative

Use `if/else` for branching and explicit nil checks instead of `??`.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/cond/coalesce">Coalesce</a><a class="related-chip" href="/packages/cond/coalesce-fn">CoalesceFn</a><a class="related-chip" href="/packages/cond/coalesce-ptr">CoalescePtr</a>
</div>

← [Back to cond package overview](/packages/cond/)
