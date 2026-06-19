---
title: "AppError.BadRequest"
package: "err"
import: "github.com/sahilkhaire/gox/err"
node: "createError(400, msg)"
gox-doc-version: "10"
---

<SymbolHeader pkg="err" title="AppError.BadRequest" node="createError(400, msg)" import-path="github.com/sahilkhaire/gox/err" />
## Overview

BadRequest returns a 400 AppError.

**Node.js equivalent:** `createError(400, msg)`

## Signature

<div class="signature-block">

```go
func BadRequest(message string) *AppError
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
next(createError(400, 'bad request'));
```

```go [Standard Go]
return fmt.Errorf("%w: message", http.ErrBadRequest)
```

```go [gox]
import "github.com/sahilkhaire/gox/err"

return err.BadRequest("bad request")
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/err/app-error-forbidden">AppError.Forbidden</a><a class="related-chip" href="/packages/err/app-error-internal">AppError.Internal</a><a class="related-chip" href="/packages/err/app-error-new">AppError.New</a>
</div>

← [Back to err package overview](/packages/err/)
