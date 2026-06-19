---
title: "String"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
node: "z.string().email()"
gox-doc-version: "14"
---

<SymbolHeader pkg="validate" title="String" node="z.string().email()" import-path="github.com/sahilkhaire/gox/validate" />
## Overview

If you are coming from Node.js, the closest pattern is **`z.string().email()`**.

## Signature

<div class="signature-block">

```go
func String() *stringField
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
z.string().email()
```

```go [Standard Go]
field := validate.String().Email() // fluent schema builder
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

field := validate.String().Required().Email().MinLen(3)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/validate"

field := validate.String().Required().Email().MinLen(3)
```

## Tips

Import `github.com/sahilkhaire/gox/validate` and call `String` directly. See the comparison below for the standard library equivalent.

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
