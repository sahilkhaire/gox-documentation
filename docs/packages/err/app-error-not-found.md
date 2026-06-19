---
title: "AppError.NotFound"
package: "err"
import: "github.com/sahilkhaire/gox/err"
node: "createError(404, msg)"
gox-doc-version: "10"
---

<SymbolHeader pkg="err" title="AppError.NotFound" node="createError(404, msg)" import-path="github.com/sahilkhaire/gox/err" />
## Overview

Returns a typed 404 error for HTTP handlers — maps to createError(404) from http-errors.

## Signature

<div class="signature-block">

```go
func NotFound(message string) *AppError
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
next(createError(404, 'not found'));
```

```go [Standard Go]
return fmt.Errorf("not found: %w", ErrNotFound)
```

```go [gox]
import "github.com/sahilkhaire/gox/err"

return err.NotFound("not found")
```

:::

## Tips

Return from gox/http handlers; status code is inferred automatically.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/err/app-error-bad-request">AppError.BadRequest</a><a class="related-chip" href="/packages/err/app-error-forbidden">AppError.Forbidden</a><a class="related-chip" href="/packages/err/app-error-internal">AppError.Internal</a>
</div>

← [Back to err package overview](/packages/err/)
