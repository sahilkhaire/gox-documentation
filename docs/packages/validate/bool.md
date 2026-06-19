---
title: "Bool"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
gox-doc-version: "14"
---

<SymbolHeader pkg="validate" title="Bool" node="zod, joi" import-path="github.com/sahilkhaire/gox/validate" />
## Overview

Part of the **`validate`** package — Node.js analog: *zod, joi*.

## Signature

<div class="signature-block">

```go
func Bool() *boolField
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
z.boolean()
```

```go [Standard Go]
field := validate.String().Email() // fluent schema builder
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

field := validate.Bool().Required()
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/validate"

field := validate.Bool().Required()
```

## Tips

Import `github.com/sahilkhaire/gox/validate` and call `Bool` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
field := validate.String().Email() // fluent schema builder
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/validate/array">Array</a><a class="related-chip" href="/packages/validate/float">Float</a><a class="related-chip" href="/packages/validate/int">Int</a>
</div>

← [Back to validate package overview](/packages/validate/)
