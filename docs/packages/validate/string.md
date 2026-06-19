---
title: "String"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
node: "z.string().email()"
gox-doc-version: "11"
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
if err := validator.Struct(v); err != nil { /* handle */ }
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

validate.String().Email()
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/validate"

validate.String().Email()
```

## Tips

Import `github.com/sahilkhaire/gox/validate` and call `String` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/validate/array">Array</a><a class="related-chip" href="/packages/validate/bool">Bool</a><a class="related-chip" href="/packages/validate/float">Float</a>
</div>

← [Back to validate package overview](/packages/validate/)
