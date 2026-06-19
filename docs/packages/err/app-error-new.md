---
title: "AppError.New"
package: "err"
import: "github.com/sahilkhaire/gox/err"
gox-doc-version: "10"
---

<SymbolHeader pkg="err" title="AppError.New" node="http-errors" import-path="github.com/sahilkhaire/gox/err" />
## Overview

New creates an AppError with code and message.

## Signature

<div class="signature-block">

```go
func New(code int, message string) *AppError
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches.
```

```go [gox]
import "github.com/sahilkhaire/gox/err"

var v AppError
v.New(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/err/app-error-bad-request">AppError.BadRequest</a><a class="related-chip" href="/packages/err/app-error-forbidden">AppError.Forbidden</a><a class="related-chip" href="/packages/err/app-error-internal">AppError.Internal</a>
</div>

← [Back to err package overview](/packages/err/)
