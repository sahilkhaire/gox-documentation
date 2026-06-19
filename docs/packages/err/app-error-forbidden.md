---
title: "AppError.Forbidden"
package: "err"
import: "github.com/sahilkhaire/gox/err"
node: "createError(403, msg)"
gox-doc-version: "11"
---

<SymbolHeader pkg="err" title="AppError.Forbidden" node="createError(403, msg)" import-path="github.com/sahilkhaire/gox/err" />
## Overview

Forbidden returns a 403 AppError.

If you are coming from Node.js, the closest pattern is **`createError(403, msg)`**.

Method on **`AppError`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func Forbidden(message string) *AppError
```

</div>

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

## Example

```go
import "github.com/sahilkhaire/gox/err"

return err.Forbidden("forbidden")
```

## Tips

Obtain a `AppError` value first (see constructors on the package overview), then call `Forbidden`.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/err/app-error-bad-request">AppError.BadRequest</a><a class="related-chip" href="/packages/err/app-error-internal">AppError.Internal</a><a class="related-chip" href="/packages/err/app-error-new">AppError.New</a>
</div>

← [Back to err package overview](/packages/err/)
