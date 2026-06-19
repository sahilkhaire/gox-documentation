---
title: "Coalesce"
package: "cond"
import: "github.com/sahilkhaire/gox/cond"
node: "a ?? b ?? c"
gox-doc-version: "11"
---

<SymbolHeader pkg="cond" title="Coalesce" node="a ?? b ?? c" import-path="github.com/sahilkhaire/gox/cond" />
## Overview

Coalesce returns the first value that is not the zero value (Node: a ?? b ?? c).

If you are coming from Node.js, the closest pattern is **`a ?? b ?? c`**.

## Signature

<div class="signature-block">

```go
func Coalesce[T comparable](vals ...T) T
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const name = maybeName ?? 'guest';
```

```go [Standard Go]
name := maybeName
if name == "" {
    name = "guest"
}
```

```go [gox]
import "github.com/sahilkhaire/gox/cond"

name := cond.Coalesce(maybeName, "guest")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/cond"

tests := []struct {
	in	[]int
	want	int
}{
	{[]int{0, 0, 3}, 3},
	{[]int{1, 2}, 1},
	{[]int{}, 0},
}
```

## Tips

Prefer explicit zero-value checks in performance-critical hot paths if the compiler cannot prove types.

## Standard library alternative

Use `if/else` for branching and explicit nil checks instead of `??`.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/cond/coalesce-fn">CoalesceFn</a><a class="related-chip" href="/packages/cond/coalesce-ptr">CoalescePtr</a><a class="related-chip" href="/packages/cond/if">If</a>
</div>

← [Back to cond package overview](/packages/cond/)
