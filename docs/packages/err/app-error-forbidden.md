---
title: "AppError.Forbidden"
package: "err"
import: "github.com/sahilkhaire/gox/err"
node: "createError(403, msg)"
gox-doc-version: "10"
---

<SymbolHeader pkg="err" title="AppError.Forbidden" node="createError(403, msg)" import-path="github.com/sahilkhaire/gox/err" />
## Overview

Forbidden returns a 403 AppError.

**Node.js equivalent:** `createError(403, msg)`

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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/err/app-error-bad-request">AppError.BadRequest</a><a class="related-chip" href="/packages/err/app-error-internal">AppError.Internal</a><a class="related-chip" href="/packages/err/app-error-new">AppError.New</a>
</div>

← [Back to err package overview](/packages/err/)
