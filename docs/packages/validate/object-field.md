---
title: "ObjectField"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
gox-doc-version: "14"
---

<SymbolHeader pkg="validate" title="ObjectField" node="zod, joi" import-path="github.com/sahilkhaire/gox/validate" />
## Overview

ObjectField nests an object schema as a field.

Part of the **`validate`** package — Node.js analog: *zod, joi*.

## Signature

<div class="signature-block">

```go
func ObjectField(fields map[string]Field) *objectField
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical zod, joi pattern in Node.js
```

```go [Standard Go]
field := validate.String().Email() // fluent schema builder
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

field := validate.ObjectField(map[string]validate.Field{"city": validate.String()})
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/validate"

field := validate.ObjectField(map[string]validate.Field{"city": validate.String()})
```

## Tips

Import `github.com/sahilkhaire/gox/validate` and call `ObjectField` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
field := validate.String().Email() // fluent schema builder
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/validate/array">Array</a><a class="related-chip" href="/packages/validate/bool">Bool</a><a class="related-chip" href="/packages/validate/float">Float</a>
</div>

← [Back to validate package overview](/packages/validate/)
