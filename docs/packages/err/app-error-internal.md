---
title: "AppError.Internal"
package: "err"
import: "github.com/sahilkhaire/gox/err"
node: "createError(500, msg)"
gox-doc-version: "10"
---

<SymbolHeader pkg="err" title="AppError.Internal" node="createError(500, msg)" import-path="github.com/sahilkhaire/gox/err" />
## Overview

Internal returns a 500 AppError.

**Node.js equivalent:** `createError(500, msg)`

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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/err/app-error-bad-request">AppError.BadRequest</a><a class="related-chip" href="/packages/err/app-error-forbidden">AppError.Forbidden</a><a class="related-chip" href="/packages/err/app-error-new">AppError.New</a>
</div>

← [Back to err package overview](/packages/err/)
