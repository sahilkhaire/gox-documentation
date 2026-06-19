---
title: "Schema"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
gox-doc-version: "14"
---

<SymbolHeader pkg="validate" title="Schema" node="zod, joi" import-path="github.com/sahilkhaire/gox/validate" />
## Overview

Schema is a map-based object schema (Joi.object).

Part of the **`validate`** package — Node.js analog: *zod, joi*.

`Schema` is a type exported by gox. Methods on this type are documented separately.

## Signature

<div class="signature-block">

```go
type Schema struct {
	Fields map[string]Field
}
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical zod, joi pattern in Node.js
```

```go [Standard Go]
// use github.com/go-playground/validator struct tags or manual checks
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

sch := validate.Object(map[string]validate.Field{
	"name":	validate.String().MinLen(2),
	"role":	validate.String().OneOf("admin", "user"),
})
err := validate.ValidateSchema(sch, map[string]any{"name": "a", "role": "guest"})
err = validate.ValidateSchema(sch, map[string]any{"name": "alice", "role": "admin"})
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/validate"

sch := validate.Object(map[string]validate.Field{
	"name":	validate.String().MinLen(2),
	"role":	validate.String().OneOf("admin", "user"),
})
err := validate.ValidateSchema(sch, map[string]any{"name": "a", "role": "guest"})
err = validate.ValidateSchema(sch, map[string]any{"name": "alice", "role": "admin"})
```

## Tips

Browse methods on this type in the sidebar for handler-style APIs and options structs.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/validate/array">Array</a><a class="related-chip" href="/packages/validate/bool">Bool</a><a class="related-chip" href="/packages/validate/float">Float</a>
</div>

← [Back to validate package overview](/packages/validate/)
