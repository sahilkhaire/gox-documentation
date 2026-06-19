---
title: "Coalesce"
package: "cond"
import: "github.com/sahilkhaire/gox/cond"
node: "a ?? b ?? c"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: a ?? b ?? c</span><span class="api-badge import">github.com/sahilkhaire/gox/cond</span></div>
# Coalesce

## Overview

Coalesce returns the first value that is not the zero value (Node: a ?? b ?? c).

**Node.js equivalent:** `a ?? b ?? c`

## Signature

```go
func Coalesce[T comparable](vals ...T) T
```

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

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

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

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [CoalesceFn](/packages/cond/coalesce-fn)
- [CoalescePtr](/packages/cond/coalesce-ptr)
- [If](/packages/cond/if)

← [Back to cond package overview](/packages/cond/)
