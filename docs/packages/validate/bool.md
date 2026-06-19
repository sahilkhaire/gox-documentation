---
title: "Bool"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
gox-doc-version: "10"
---

<SymbolHeader pkg="validate" title="Bool" node="zod, joi" import-path="github.com/sahilkhaire/gox/validate" />
## Overview

## Signature

<div class="signature-block">

```go
func Bool() *boolField
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
_ = validate.Bool(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/validate/array">Array</a><a class="related-chip" href="/packages/validate/float">Float</a><a class="related-chip" href="/packages/validate/int">Int</a>
</div>

← [Back to validate package overview](/packages/validate/)
