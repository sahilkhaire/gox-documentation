---
title: "Bool"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: zod, joi</span><span class="api-badge import">github.com/sahilkhaire/gox/validate</span></div>
# Bool

## Overview

## Signature

```go
func Bool() *boolField
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
_ = validate.Bool(/* args */)
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
- [Float](/packages/validate/float)
- [Int](/packages/validate/int)

← [Back to validate package overview](/packages/validate/)
