---
title: "ValidateSchema"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
node: "schema.parse(data)"
gox-doc-version: "11"
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
schema.parse(data)
```

```go [Standard Go]
if err := validator.Struct(v); err != nil { /* handle */ }
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

validate.ValidateSchema(sch, data)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/validate"

validate.ValidateSchema(sch, data)
```

## Tips

Import `github.com/sahilkhaire/gox/validate` and call `ValidateSchema` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/validate/array">Array</a><a class="related-chip" href="/packages/validate/bool">Bool</a><a class="related-chip" href="/packages/validate/float">Float</a>
</div>

← [Back to validate package overview](/packages/validate/)
