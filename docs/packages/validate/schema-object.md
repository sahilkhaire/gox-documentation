---
title: "Schema.Object"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
node: "z.object({...})"
gox-doc-version: "14"
---

<SymbolHeader pkg="validate" title="Schema.Object" node="z.object({...})" import-path="github.com/sahilkhaire/gox/validate" />
## Overview

Object builds a schema from field definitions.

If you are coming from Node.js, the closest pattern is **`z.object({...})`**.

Method on **`Schema`** — call it on a value of that type after constructing or receiving one from a constructor.

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
field := validate.String().Email() // fluent schema builder
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

sch := validate.Object(map[string]validate.Field{"name": validate.String().Required()})
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/validate"

sch := validate.Object(map[string]validate.Field{"name": validate.String().Required()})
```

## Tips

Obtain a `Schema` value first (see constructors on the package overview), then call `Object`.

## Standard library alternative

Use the standard library directly:

```go
field := validate.String().Email() // fluent schema builder
```

← [Back to validate package overview](/packages/validate/)
