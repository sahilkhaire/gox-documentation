---
title: "As"
package: "err"
import: "github.com/sahilkhaire/gox/err"
gox-doc-version: "11"
---

<SymbolHeader pkg="err" title="As" node="http-errors" import-path="github.com/sahilkhaire/gox/err" />
## Overview

As finds the first error in the chain assignable to target.

Part of the **`err`** package — Node.js analog: *http-errors*.

## Signature

<div class="signature-block">

```go
func As(err error, target any) bool
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical http-errors pattern in Node.js
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

## Example

```go
import "github.com/sahilkhaire/gox/err"

// err
_ = err.As(/* args */)
```

## Tips

Import `github.com/sahilkhaire/gox/err` and call `As` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/err/is">Is</a>
</div>

← [Back to err package overview](/packages/err/)
