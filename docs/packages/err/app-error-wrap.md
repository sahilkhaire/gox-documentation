---
title: "AppError.Wrap"
package: "err"
import: "github.com/sahilkhaire/gox/err"
gox-doc-version: "10"
---

<SymbolHeader pkg="err" title="AppError.Wrap" node="http-errors" import-path="github.com/sahilkhaire/gox/err" />
## Overview

Wrap wraps cause with code and message.

## Signature

<div class="signature-block">

```go
func Wrap(code int, message string, cause error) *AppError
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/err/app-error-bad-request">AppError.BadRequest</a><a class="related-chip" href="/packages/err/app-error-forbidden">AppError.Forbidden</a><a class="related-chip" href="/packages/err/app-error-internal">AppError.Internal</a>
</div>

← [Back to err package overview](/packages/err/)
