---
title: "Schema"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
gox-doc-version: "10"
---

<SymbolHeader pkg="validate" title="Schema" node="zod, joi" import-path="github.com/sahilkhaire/gox/validate" />
## Overview

Schema is a map-based object schema (Joi.object).

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
// See package overview
```

```go [Standard Go]
if err := validator.Struct(v); err != nil { /* handle */ }
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

_ = validate.Schema
```

:::

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

```go
import "github.com/sahilkhaire/gox/validate"

sch := validate.Object(map[string]validate.Field{
	"name":	validate.String().MinLen(2),
	"role":	validate.String().OneOf("admin", "user"),
})
err := validate.ValidateSchema(sch, map[string]any{"name": "a", "role": "guest"})
err = validate.ValidateSchema(sch, map[string]any{"name": "alice", "role": "admin"})
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/validate/array">Array</a><a class="related-chip" href="/packages/validate/bool">Bool</a><a class="related-chip" href="/packages/validate/float">Float</a>
</div>

← [Back to validate package overview](/packages/validate/)
