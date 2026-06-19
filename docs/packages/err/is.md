---
title: "Is"
package: "err"
import: "github.com/sahilkhaire/gox/err"
gox-doc-version: "14"
---

<SymbolHeader pkg="err" title="Is" node="http-errors" import-path="github.com/sahilkhaire/gox/err" />
## Overview

Is reports whether err matches target in the error chain.

Part of the **`err`** package — Node.js analog: *http-errors*.

## Signature

<div class="signature-block">

```go
func Is(err, target error) bool
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

if err.Is(err, target) { /* ... */ }
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/err"

if err.Is(err, target) { /* ... */ }
```

## Tips

Import `github.com/sahilkhaire/gox/err` and call `Is` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
errors.Is(err, target)
errors.As(err, &target)
fmt.Errorf("context: %w", err)
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/err/as">As</a>
</div>

← [Back to err package overview](/packages/err/)
