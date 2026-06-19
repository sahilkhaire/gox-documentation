---
title: "AppError"
package: "err"
import: "github.com/sahilkhaire/gox/err"
gox-doc-version: "14"
---

<SymbolHeader pkg="err" title="AppError" node="http-errors" import-path="github.com/sahilkhaire/gox/err" />
## Overview

AppError is an error with an HTTP status code.

Part of the **`err`** package — Node.js analog: *http-errors*.

`AppError` is a type exported by gox. Methods on this type are documented separately.

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
// Typical http-errors pattern in Node.js
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches.
```

```go [gox]
import "github.com/sahilkhaire/gox/err"

base := errors.New("root")
e := Wrap(http.StatusTeapot, "wrapped", base)
var ae *AppError
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/err"

base := errors.New("root")
e := Wrap(http.StatusTeapot, "wrapped", base)
var ae *AppError
```

## Tips

Browse methods on this type in the sidebar for handler-style APIs and options structs.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/err/as">As</a><a class="related-chip" href="/packages/err/is">Is</a>
</div>

← [Back to err package overview](/packages/err/)
