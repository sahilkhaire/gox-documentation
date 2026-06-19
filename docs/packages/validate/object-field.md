---
title: "ObjectField"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: zod, joi</span><span class="api-badge import">github.com/sahilkhaire/gox/validate</span></div>
# ObjectField

## Overview

ObjectField nests an object schema as a field.

## Signature

```go
func ObjectField(fields map[string]Field) *objectField
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

// validate
_ = validate.ObjectField(/* args */)
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
