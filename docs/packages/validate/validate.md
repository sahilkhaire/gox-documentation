---
title: "Validate"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
gox-doc-version: "10"
---

<SymbolHeader pkg="validate" title="Validate" node="zod, joi" import-path="github.com/sahilkhaire/gox/validate" />
## Overview

Validate runs struct-tag validation on structs.

## Signature

<div class="signature-block">

```go
func Validate(v any) error
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
schema.validate(data);
```

```go [Standard Go]
if err := validator.Struct(v); err != nil { /* handle */ }
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

if err := validate.Validate(&v); err != nil { /* handle */ }
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/validate/array">Array</a><a class="related-chip" href="/packages/validate/bool">Bool</a><a class="related-chip" href="/packages/validate/float">Float</a>
</div>

← [Back to validate package overview](/packages/validate/)
