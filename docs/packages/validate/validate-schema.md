---
title: "ValidateSchema"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
node: "schema.parse(data)"
gox-doc-version: "14"
---

<SymbolHeader pkg="validate" title="ValidateSchema" node="schema.parse(data)" import-path="github.com/sahilkhaire/gox/validate" />
## Overview

ValidateSchema validates v against sch (map or struct).

If you are coming from Node.js, the closest pattern is **`schema.parse(data)`**.

## Signature

<div class="signature-block">

```go
func ValidateSchema(sch Schema, v any) error
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
schema.parse(data);
```

```go [Standard Go]
// validate fields manually or map to struct tags
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

sch := validate.Object(map[string]validate.Field{
	"name": validate.String().MinLen(2),
	"role": validate.String().OneOf("admin", "user"),
})
if err := validate.ValidateSchema(sch, map[string]any{"name": "alice", "role": "admin"}); err != nil {
	return err
}
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/validate"

sch := validate.Object(map[string]validate.Field{
	"name": validate.String().MinLen(2),
	"role": validate.String().OneOf("admin", "user"),
})
if err := validate.ValidateSchema(sch, map[string]any{"name": "alice", "role": "admin"}); err != nil {
	return err
}
```

## Tips

Import `github.com/sahilkhaire/gox/validate` and call `ValidateSchema` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
// validate fields manually or map to struct tags
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/validate/array">Array</a><a class="related-chip" href="/packages/validate/bool">Bool</a><a class="related-chip" href="/packages/validate/float">Float</a>
</div>

← [Back to validate package overview](/packages/validate/)
