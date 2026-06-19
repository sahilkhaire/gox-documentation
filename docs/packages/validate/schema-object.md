---
title: "Schema.Object"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
node: "z.object({...})"
gox-doc-version: "11"
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
if err := validator.Struct(v); err != nil { /* handle */ }
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

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to validate package overview](/packages/validate/)
