---
title: "CoalesceFn"
package: "cond"
import: "github.com/sahilkhaire/gox/cond"
node: "obj?.field ?? \"guest\""
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: obj?.field ?? "guest"</span><span class="api-badge import">github.com/sahilkhaire/gox/cond</span></div>
# CoalesceFn

## Overview

Maps the Node.js pattern `obj?.field ?? "guest"` to gox `cond.CoalesceFn(ptr, func(v T) R { ... }, fallback)`. Part of the cond package — Node.js analog: ternary ? :, nullish ??.

**Node.js equivalent:** `obj?.field ?? "guest"`

## Signature

```go
func CoalesceFn[T comparable](fns ...func() T) T
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
obj?.field ?? "guest"
```

```go [Standard Go]
if condition {
    result = a
} else {
    result = b
}
```

```go [gox]
import "github.com/sahilkhaire/gox/cond"

cond.CoalesceFn(ptr, func(v T) R { ... }, fallback)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Coalesce](/packages/cond/coalesce)
- [CoalescePtr](/packages/cond/coalesce-ptr)
- [If](/packages/cond/if)

← [Back to cond package overview](/packages/cond/)
