---
title: "ValidateSchema"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
node: "schema.parse(data)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: schema.parse(data)</span><span class="api-badge import">github.com/sahilkhaire/gox/validate</span></div>
# ValidateSchema

## Overview

Maps the Node.js pattern `schema.parse(data)` to gox `validate.ValidateSchema(sch, data)`. Part of the validate package — Node.js analog: zod/joi.

**Node.js equivalent:** `schema.parse(data)`

## Signature

```go
func ValidateSchema(sch Schema, v any) error
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
schema.parse(data)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

validate.ValidateSchema(sch, data)
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
