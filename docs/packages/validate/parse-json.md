---
title: "ParseJSON"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
gox-doc-version: "14"
---

<SymbolHeader pkg="validate" title="ParseJSON" node="zod, joi" import-path="github.com/sahilkhaire/gox/validate" />
## Overview

ParseJSON decodes JSON from r into v and validates struct tags on v.

Part of the **`validate`** package — Node.js analog: *zod, joi*.

## Signature

<div class="signature-block">

```go
func ParseJSON(r io.Reader, v any) error
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
JSON.parse(str); schema.parse(JSON.parse(str))
```

```go [Standard Go]
if err := validator.New().Struct(v); err != nil {
    return err
}
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

var payload signup
if err := validate.ParseJSON(r, &payload); err != nil {
	return err
}
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/validate"

var payload signup
if err := validate.ParseJSON(r, &payload); err != nil {
	return err
}
```

## Tips

Import `github.com/sahilkhaire/gox/validate` and call `ParseJSON` directly. See the comparison below for the standard library equivalent.

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
