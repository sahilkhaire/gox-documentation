---
title: "AppError.Unauthorized"
package: "err"
import: "github.com/sahilkhaire/gox/err"
node: "createError(401, msg)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: createError(401, msg)</span><span class="api-badge import">github.com/sahilkhaire/gox/err</span></div>
# AppError.Unauthorized

## Overview

Unauthorized returns a 401 AppError.

**Node.js equivalent:** `createError(401, msg)`

## Signature

```go
func Unauthorized(message string) *AppError
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
next(createError(401, 'unauthorized'));
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/err"

return err.Unauthorized("unauthorized")
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
- [AppError.Internal](/packages/err/app-error-internal)

← [Back to err package overview](/packages/err/)
