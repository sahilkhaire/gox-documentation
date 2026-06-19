---
title: "Schema.Object"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
node: "z.object({...})"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: z.object({...})</span><span class="api-badge import">github.com/sahilkhaire/gox/validate</span></div>
# Schema.Object

## Overview

Object builds a schema from field definitions.

**Node.js equivalent:** `z.object({...})`

## Signature

```go
func Object(fields map[string]Field) Schema
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
z.object({ name: z.string() });
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

sch := validate.Object(map[string]validate.Field{"name": validate.String().Required()})
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

← [Back to validate package overview](/packages/validate/)
