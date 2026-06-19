---
title: "Validate"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
gox-doc-version: "14"
---

<SymbolHeader pkg="validate" title="Validate" node="zod, joi" import-path="github.com/sahilkhaire/gox/validate" />
## Overview

Validate runs struct-tag validation on structs.

Part of the **`validate`** package — Node.js analog: *zod, joi*.

## Signature

<div class="signature-block">

```go
func Validate(v any) error
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
schema.validate(data);
```

```go [Standard Go]
if err := validator.New().Struct(v); err != nil {
    return err
}
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

type signup struct {
	Email string `validate:"required,email"`
	Age   int    `validate:"gte=18"`
}
if err := validate.Validate(signup{Email: "a@b.com", Age: 20}); err != nil {
	return err
}
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/validate"

type signup struct {
	Email string `validate:"required,email"`
	Age   int    `validate:"gte=18"`
}
if err := validate.Validate(signup{Email: "a@b.com", Age: 20}); err != nil {
	return err
}
```

## Tips

Import `github.com/sahilkhaire/gox/validate` and call `Validate` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
if err := validator.New().Struct(v); err != nil {
    return err
}
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/validate/array">Array</a><a class="related-chip" href="/packages/validate/bool">Bool</a><a class="related-chip" href="/packages/validate/float">Float</a>
</div>

← [Back to validate package overview](/packages/validate/)
