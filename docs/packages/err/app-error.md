---
title: "AppError"
package: "err"
import: "github.com/sahilkhaire/gox/err"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: http-errors</span><span class="api-badge import">github.com/sahilkhaire/gox/err</span></div>
# AppError

## Overview

AppError is an error with an HTTP status code.

## Signature

```go
type AppError struct {
	Code    int
	Message string
	Cause   error
}
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

_ = err.AppError
```

:::

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

```go
import "github.com/sahilkhaire/gox/err"

base := errors.New("root")
e := Wrap(http.StatusTeapot, "wrapped", base)
var ae *AppError
```

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [As](/packages/err/as)
- [Is](/packages/err/is)

← [Back to err package overview](/packages/err/)
