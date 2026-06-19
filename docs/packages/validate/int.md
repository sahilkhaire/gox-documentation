---
title: "Int"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
node: "z.number().min(18)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: z.number().min(18)</span><span class="api-badge import">github.com/sahilkhaire/gox/validate</span></div>
# Int

## Overview

Maps the Node.js pattern `z.number().min(18)` to gox `validate.Int().Min(18)`. Part of the validate package — Node.js analog: zod/joi.

**Node.js equivalent:** `z.number().min(18)`

## Signature

```go
func Int() *intField
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
z.number().min(18)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

validate.Int().Min(18)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Array](/packages/validate/array)
- [Bool](/packages/validate/bool)
- [Float](/packages/validate/float)

← [Back to validate package overview](/packages/validate/)
