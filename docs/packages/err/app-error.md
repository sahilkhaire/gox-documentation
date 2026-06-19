---
title: "AppError"
package: "err"
import: "github.com/sahilkhaire/gox/err"
gox-doc-version: "10"
---

<SymbolHeader pkg="err" title="AppError" node="http-errors" import-path="github.com/sahilkhaire/gox/err" />
## Overview

AppError is an error with an HTTP status code.

## Signature

<div class="signature-block">

```go
type AppError struct {
	Code    int
	Message string
	Cause   error
}
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches.
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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/err/as">As</a><a class="related-chip" href="/packages/err/is">Is</a>
</div>

← [Back to err package overview](/packages/err/)
