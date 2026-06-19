---
title: "String"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
node: "z.string().email()"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: z.string().email()</span><span class="api-badge import">github.com/sahilkhaire/gox/validate</span></div>
# String

## Overview

Maps the Node.js pattern `z.string().email()` to gox `validate.String().Email()`. Part of the validate package — Node.js analog: zod/joi.

**Node.js equivalent:** `z.string().email()`

## Signature

```go
func String() *stringField
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
z.string().email()
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

validate.String().Email()
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
