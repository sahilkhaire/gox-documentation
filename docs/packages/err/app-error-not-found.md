---
title: "AppError.NotFound"
package: "err"
import: "github.com/sahilkhaire/gox/err"
node: "createError(404, msg)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: createError(404, msg)</span><span class="api-badge import">github.com/sahilkhaire/gox/err</span></div>
# AppError.NotFound

## Overview

NotFound returns a 404 AppError.

**Node.js equivalent:** `createError(404, msg)`

## Signature

```go
func NotFound(message string) *AppError
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
next(createError(404, 'not found'));
```

```go [Standard Go]
return fmt.Errorf("not found: %w", ErrNotFound)
```

```go [gox]
import "github.com/sahilkhaire/gox/err"

return err.NotFound("not found")
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
