---
title: "Array"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
gox-doc-version: "10"
---

<SymbolHeader pkg="validate" title="Array" node="zod, joi" import-path="github.com/sahilkhaire/gox/validate" />
## Overview

## Signature

<div class="signature-block">

```go
func Array(elem Field) *arrayField
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
if err := validator.Struct(v); err != nil { /* handle */ }
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

// validate
_ = validate.Array(/* args */)
```

:::

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

```go
import "github.com/sahilkhaire/gox/validate"

sch := validate.Object(map[string]validate.Field{
	"tags": validate.Array(validate.String().MinLen(1)).MinLen(1),
})
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/validate/bool">Bool</a><a class="related-chip" href="/packages/validate/float">Float</a><a class="related-chip" href="/packages/validate/int">Int</a>
</div>

← [Back to validate package overview](/packages/validate/)
