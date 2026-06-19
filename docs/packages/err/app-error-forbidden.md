---
title: "AppError.Forbidden"
package: "err"
import: "github.com/sahilkhaire/gox/err"
node: "createError(403, msg)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: createError(403, msg)</span><span class="api-badge import">github.com/sahilkhaire/gox/err</span></div>
# AppError.Forbidden

## Overview

Forbidden returns a 403 AppError.

**Node.js equivalent:** `createError(403, msg)`

## Signature

```go
func Forbidden(message string) *AppError
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
next(createError(403, 'forbidden'));
```

```go [Standard Go]
return fmt.Errorf("forbidden")
```

```go [gox]
import "github.com/sahilkhaire/gox/err"

return err.Forbidden("forbidden")
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
- [AppError.Internal](/packages/err/app-error-internal)
- [AppError.New](/packages/err/app-error-new)

← [Back to err package overview](/packages/err/)
