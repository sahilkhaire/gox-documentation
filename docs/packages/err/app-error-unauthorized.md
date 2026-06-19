---
title: "AppError.Unauthorized"
package: "err"
import: "github.com/sahilkhaire/gox/err"
node: "createError(401, msg)"
gox-doc-version: "10"
---

<SymbolHeader pkg="err" title="AppError.Unauthorized" node="createError(401, msg)" import-path="github.com/sahilkhaire/gox/err" />
## Overview

Unauthorized returns a 401 AppError.

**Node.js equivalent:** `createError(401, msg)`

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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/err/app-error-bad-request">AppError.BadRequest</a><a class="related-chip" href="/packages/err/app-error-forbidden">AppError.Forbidden</a><a class="related-chip" href="/packages/err/app-error-internal">AppError.Internal</a>
</div>

← [Back to err package overview](/packages/err/)
