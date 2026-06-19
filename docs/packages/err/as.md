---
title: "As"
package: "err"
import: "github.com/sahilkhaire/gox/err"
gox-doc-version: "10"
---

<SymbolHeader pkg="err" title="As" node="http-errors" import-path="github.com/sahilkhaire/gox/err" />
## Overview

As finds the first error in the chain assignable to target.

## Signature

<div class="signature-block">

```go
func As(err error, target any) bool
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
_ = err.As(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/err/is">Is</a>
</div>

← [Back to err package overview](/packages/err/)
