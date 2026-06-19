---
title: "Array"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
gox-doc-version: "14"
---

<SymbolHeader pkg="validate" title="Array" node="zod, joi" import-path="github.com/sahilkhaire/gox/validate" />
## Overview

Part of the **`validate`** package — Node.js analog: *zod, joi*.

## Signature

<div class="signature-block">

```go
func Array(elem Field) *arrayField
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
z.array(z.string())
```

```go [Standard Go]
field := validate.String().Email() // fluent schema builder
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

sch := validate.Object(map[string]validate.Field{
	"tags": validate.Array(validate.String().MinLen(1)).MinLen(1),
})
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/validate"

sch := validate.Object(map[string]validate.Field{
	"tags": validate.Array(validate.String().MinLen(1)).MinLen(1),
})
```

## Tips

Import `github.com/sahilkhaire/gox/validate` and call `Array` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
field := validate.String().Email() // fluent schema builder
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/validate/bool">Bool</a><a class="related-chip" href="/packages/validate/float">Float</a><a class="related-chip" href="/packages/validate/int">Int</a>
</div>

← [Back to validate package overview](/packages/validate/)
