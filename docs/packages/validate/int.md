---
title: "Int"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
node: "z.number().min(18)"
gox-doc-version: "14"
---

<SymbolHeader pkg="validate" title="Int" node="z.number().min(18)" import-path="github.com/sahilkhaire/gox/validate" />
## Overview

If you are coming from Node.js, the closest pattern is **`z.number().min(18)`**.

## Signature

<div class="signature-block">

```go
func Int() *intField
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
z.number().min(18)
```

```go [Standard Go]
field := validate.String().Email() // fluent schema builder
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

field := validate.Int().Min(18)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/validate"

field := validate.Int().Min(18)
```

## Tips

Import `github.com/sahilkhaire/gox/validate` and call `Int` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
field := validate.String().Email() // fluent schema builder
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/validate/array">Array</a><a class="related-chip" href="/packages/validate/bool">Bool</a><a class="related-chip" href="/packages/validate/float">Float</a>
</div>

← [Back to validate package overview](/packages/validate/)
