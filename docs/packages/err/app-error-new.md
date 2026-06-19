---
title: "AppError.New"
package: "err"
import: "github.com/sahilkhaire/gox/err"
gox-doc-version: "14"
---

<SymbolHeader pkg="err" title="AppError.New" node="http-errors" import-path="github.com/sahilkhaire/gox/err" />
## Overview

New creates an AppError with code and message.

Part of the **`err`** package — Node.js analog: *http-errors*.

Method on **`AppError`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func New(code int, message string) *AppError
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

e := err.AppError.New(404, "not found")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/err"

e := err.AppError.New(404, "not found")
```

## Tips

Obtain a `AppError` value first (see constructors on the package overview), then call `New`.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/err/app-error-bad-request">AppError.BadRequest</a><a class="related-chip" href="/packages/err/app-error-forbidden">AppError.Forbidden</a><a class="related-chip" href="/packages/err/app-error-internal">AppError.Internal</a>
</div>

← [Back to err package overview](/packages/err/)
