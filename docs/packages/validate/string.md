---
title: "String"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
node: "z.string().email()"
gox-doc-version: "10"
---

<SymbolHeader pkg="validate" title="String" node="z.string().email()" import-path="github.com/sahilkhaire/gox/validate" />
## Overview

**Node.js equivalent:** `z.string().email()`

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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/validate/array">Array</a><a class="related-chip" href="/packages/validate/bool">Bool</a><a class="related-chip" href="/packages/validate/float">Float</a>
</div>

← [Back to validate package overview](/packages/validate/)
