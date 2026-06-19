---
title: "MustValidate"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
gox-doc-version: "10"
---

<SymbolHeader pkg="validate" title="MustValidate" node="zod, joi" import-path="github.com/sahilkhaire/gox/validate" />
## Overview

MustValidate panics if validation fails.

## Signature

<div class="signature-block">

```go
func MustValidate(v any)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
if err := validator.Struct(v); err != nil { /* handle */ }
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

// validate
_ = validate.MustValidate(/* args */)
```

:::

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

```go
import "github.com/sahilkhaire/gox/validate"

validate.MustValidate(signup{})
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/validate/array">Array</a><a class="related-chip" href="/packages/validate/bool">Bool</a><a class="related-chip" href="/packages/validate/float">Float</a>
</div>

← [Back to validate package overview](/packages/validate/)
