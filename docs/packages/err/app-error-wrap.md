---
title: "AppError.Wrap"
package: "err"
import: "github.com/sahilkhaire/gox/err"
gox-doc-version: "14"
---

<SymbolHeader pkg="err" title="AppError.Wrap" node="http-errors" import-path="github.com/sahilkhaire/gox/err" />
## Overview

Wrap wraps cause with code and message.

Part of the **`err`** package — Node.js analog: *http-errors*.

Method on **`AppError`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func Wrap(code int, message string, cause error) *AppError
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical http-errors pattern in Node.js
```

```go [Standard Go]
errors.Is(err, target)
errors.As(err, &target)
fmt.Errorf("context: %w", err)
```

```go [gox]
import "github.com/sahilkhaire/gox/err"

return err.AppError.Wrap(cause, 500, "internal error")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/err"

return err.AppError.Wrap(cause, 500, "internal error")
```

## Tips

Obtain a `AppError` value first (see constructors on the package overview), then call `Wrap`.

## Standard library alternative

Use the standard library directly:

```go
errors.Is(err, target)
errors.As(err, &target)
fmt.Errorf("context: %w", err)
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/err/app-error-bad-request">AppError.BadRequest</a><a class="related-chip" href="/packages/err/app-error-forbidden">AppError.Forbidden</a><a class="related-chip" href="/packages/err/app-error-internal">AppError.Internal</a>
</div>

← [Back to err package overview](/packages/err/)
