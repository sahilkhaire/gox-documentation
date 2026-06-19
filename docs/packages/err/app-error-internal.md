---
title: "AppError.Internal"
package: "err"
import: "github.com/sahilkhaire/gox/err"
node: "createError(500, msg)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: createError(500, msg)</span><span class="api-badge import">github.com/sahilkhaire/gox/err</span></div>
# AppError.Internal

## Overview

Internal returns a 500 AppError.

**Node.js equivalent:** `createError(500, msg)`

## Signature

```go
func Internal(message string) *AppError
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
next(createError(500, 'internal'));
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/err"

return err.Internal("internal error")
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [AppError.BadRequest](/packages/err/app-error-bad-request)
- [AppError.Forbidden](/packages/err/app-error-forbidden)
- [AppError.New](/packages/err/app-error-new)

← [Back to err package overview](/packages/err/)
