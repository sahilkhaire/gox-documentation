---
title: "CoalescePtr"
package: "cond"
import: "github.com/sahilkhaire/gox/cond"
node: "ptr ?? fallback"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: ptr ?? fallback</span><span class="api-badge import">github.com/sahilkhaire/gox/cond</span></div>
# CoalescePtr

## Overview

CoalescePtr returns the first non-nil pointer's value, or zero if all nil.

**Node.js equivalent:** `ptr ?? fallback`

## Signature

```go
func CoalescePtr[T any](vals ...*T) T
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const name = user?.name ?? 'guest';
```

```go [Standard Go]
name := "guest"
if user != nil {
    name = user.Name
}
```

```go [gox]
import "github.com/sahilkhaire/gox/cond"

name := cond.CoalescePtr(&user.Name, "guest")
```

:::

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

```go
import "github.com/sahilkhaire/gox/cond"

a, b := 1, 2
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
- [If](/packages/cond/if)

← [Back to cond package overview](/packages/cond/)
