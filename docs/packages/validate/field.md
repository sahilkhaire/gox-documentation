---
title: "Field"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
gox-doc-version: "11"
---

<SymbolHeader pkg="validate" title="Field" node="zod, joi" import-path="github.com/sahilkhaire/gox/validate" />
## Overview

Field validates a single value in a programmatic schema.

Part of the **`validate`** package — Node.js analog: *zod, joi*.

`Field` is a type exported by gox. Methods on this type are documented separately.

## Signature

<div class="signature-block">

```go
type Field interface {
	// contains filtered or unexported methods
}
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical zod, joi pattern in Node.js
```

```go [Standard Go]
if err := validator.Struct(v); err != nil { /* handle */ }
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

_ = validate.Field
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/validate"

_ = validate.Field
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
