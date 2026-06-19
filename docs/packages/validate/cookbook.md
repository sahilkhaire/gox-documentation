---
title: "validate Cookbook"
package: "validate"
gox-doc-version: "10"
---

<div class="api-meta"><span class="api-badge node">Node: zod, joi</span><span class="api-badge import">github.com/sahilkhaire/gox/validate</span></div>
# validate Cookbook

Zod/Joi-style schemas plus struct tag validation.

## Struct tags

```go
type User struct {
    Email string `json:"email" validate:"required,email"`
    Age   int    `json:"age" validate:"gte=18"`
}
if err := validate.Validate(&user); err != nil { /* 400 */ }
```

## Fluent object schema

```go
sch := validate.Object(map[string]validate.Field{
    "email": validate.String().Required().Email(),
    "age":   validate.Int().Min(18),
})
if err := validate.ValidateSchema(sch, data); err != nil { /* handle */ }
```

← [Back to validate overview](/packages/validate/)
