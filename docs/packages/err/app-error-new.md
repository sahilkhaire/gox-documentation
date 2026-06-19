---
title: "AppError.New"
package: "err"
import: "github.com/sahilkhaire/gox/err"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: http-errors</span><span class="api-badge import">github.com/sahilkhaire/gox/err</span></div>
# AppError.New

## Overview

New creates an AppError with code and message.

## Signature

```go
func New(code int, message string) *AppError
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
import "github.com/sahilkhaire/gox/err"

var v AppError
v.New(/* args */)
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
