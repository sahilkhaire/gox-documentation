---
title: "AppError.Internal"
package: "err"
import: "github.com/sahilkhaire/gox/err"
node: "createError(500, msg)"
gox-doc-version: "14"
---

<SymbolHeader pkg="err" title="AppError.Internal" node="createError(500, msg)" import-path="github.com/sahilkhaire/gox/err" />
## Overview

Internal returns a 500 AppError.

If you are coming from Node.js, the closest pattern is **`createError(500, msg)`**.

Method on **`AppError`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func Internal(message string) *AppError
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
next(createError(500, 'internal'));
```

```go [Standard Go]
return fmt.Errorf("%w: message", http.ErrInternal)
```

```go [gox]
import "github.com/sahilkhaire/gox/err"

return err.Internal("internal error")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/err"

return err.Internal("internal error")
```

## Tips

Obtain a `AppError` value first (see constructors on the package overview), then call `Internal`.

## Standard library alternative

Use the standard library directly:

```go
return fmt.Errorf("%w: message", http.ErrInternal)
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/err/app-error-bad-request">AppError.BadRequest</a><a class="related-chip" href="/packages/err/app-error-forbidden">AppError.Forbidden</a><a class="related-chip" href="/packages/err/app-error-new">AppError.New</a>
</div>

← [Back to err package overview](/packages/err/)
