---
title: "AppError.Unauthorized"
package: "err"
import: "github.com/sahilkhaire/gox/err"
node: "createError(401, msg)"
gox-doc-version: "11"
---

<SymbolHeader pkg="err" title="AppError.Unauthorized" node="createError(401, msg)" import-path="github.com/sahilkhaire/gox/err" />
## Overview

Unauthorized returns a 401 AppError.

If you are coming from Node.js, the closest pattern is **`createError(401, msg)`**.

Method on **`AppError`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func Unauthorized(message string) *AppError
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
next(createError(401, 'unauthorized'));
```

```go [Standard Go]
return fmt.Errorf("%w: message", http.ErrUnauthorized)
```

```go [gox]
import "github.com/sahilkhaire/gox/err"

return err.Unauthorized("unauthorized")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/err"

return err.Unauthorized("unauthorized")
```

## Tips

Obtain a `AppError` value first (see constructors on the package overview), then call `Unauthorized`.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/err/app-error-bad-request">AppError.BadRequest</a><a class="related-chip" href="/packages/err/app-error-forbidden">AppError.Forbidden</a><a class="related-chip" href="/packages/err/app-error-internal">AppError.Internal</a>
</div>

← [Back to err package overview](/packages/err/)
