---
title: "Schema.Object"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
node: "z.object({...})"
gox-doc-version: "10"
---

<SymbolHeader pkg="validate" title="Schema.Object" node="z.object({...})" import-path="github.com/sahilkhaire/gox/validate" />
## Overview

Object builds a schema from field definitions.

**Node.js equivalent:** `z.object({...})`

## Signature

<div class="signature-block">

```go
func Object(fields map[string]Field) Schema
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
z.object({ name: z.string() });
```

```go [Standard Go]
if err := validator.Struct(v); err != nil { /* handle */ }
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

sch := validate.Object(map[string]validate.Field{"name": validate.String().Required()})
```

:::

← [Back to validate package overview](/packages/validate/)
