---
title: "Int"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
node: "z.number().min(18)"
gox-doc-version: "10"
---

<SymbolHeader pkg="validate" title="Int" node="z.number().min(18)" import-path="github.com/sahilkhaire/gox/validate" />
## Overview

**Node.js equivalent:** `z.number().min(18)`

## Signature

<div class="signature-block">

```go
func Int() *intField
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
z.number().min(18)
```

```go [Standard Go]
if err := validator.Struct(v); err != nil { /* handle */ }
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

validate.Int().Min(18)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/validate/array">Array</a><a class="related-chip" href="/packages/validate/bool">Bool</a><a class="related-chip" href="/packages/validate/float">Float</a>
</div>

← [Back to validate package overview](/packages/validate/)
