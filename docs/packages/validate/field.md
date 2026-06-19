---
title: "Field"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
gox-doc-version: "10"
---

<SymbolHeader pkg="validate" title="Field" node="zod, joi" import-path="github.com/sahilkhaire/gox/validate" />
## Overview

Field validates a single value in a programmatic schema.

## Signature

<div class="signature-block">

```go
type Field interface {
	// contains filtered or unexported methods
}
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

_ = validate.Field
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/validate/array">Array</a><a class="related-chip" href="/packages/validate/bool">Bool</a><a class="related-chip" href="/packages/validate/float">Float</a>
</div>

← [Back to validate package overview](/packages/validate/)
