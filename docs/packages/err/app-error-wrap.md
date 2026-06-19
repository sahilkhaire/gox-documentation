---
title: "AppError.Wrap"
package: "err"
import: "github.com/sahilkhaire/gox/err"
gox-doc-version: "11"
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

var v AppError
v.Wrap(/* args */)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/err"

var v AppError
v.Wrap(/* args */)
```

## Tips

Obtain a `AppError` value first (see constructors on the package overview), then call `Wrap`.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/err/app-error-bad-request">AppError.BadRequest</a><a class="related-chip" href="/packages/err/app-error-forbidden">AppError.Forbidden</a><a class="related-chip" href="/packages/err/app-error-internal">AppError.Internal</a>
</div>

← [Back to err package overview](/packages/err/)
