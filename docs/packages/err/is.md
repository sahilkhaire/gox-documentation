---
title: "Is"
package: "err"
import: "github.com/sahilkhaire/gox/err"
gox-doc-version: "10"
---

<SymbolHeader pkg="err" title="Is" node="http-errors" import-path="github.com/sahilkhaire/gox/err" />
## Overview

Is reports whether err matches target in the error chain.

## Signature

<div class="signature-block">

```go
func Is(err, target error) bool
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

// err
_ = err.Is(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/err/as">As</a>
</div>

← [Back to err package overview](/packages/err/)
