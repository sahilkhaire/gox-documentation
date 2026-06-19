---
title: "ParseJSON"
package: "validate"
import: "github.com/sahilkhaire/gox/validate"
gox-doc-version: "11"
---

<SymbolHeader pkg="validate" title="ParseJSON" node="zod, joi" import-path="github.com/sahilkhaire/gox/validate" />
## Overview

ParseJSON decodes JSON from r into v and validates struct tags on v.

Part of the **`validate`** package — Node.js analog: *zod, joi*.

## Signature

<div class="signature-block">

```go
func ParseJSON(r io.Reader, v any) error
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical zod, joi pattern in Node.js
```

```go [Standard Go]
if err := validator.Struct(v); err != nil { /* handle */ }
```

```go [gox]
import "github.com/sahilkhaire/gox/validate"

// validate
_ = validate.ParseJSON(/* args */)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/validate"

// validate
_ = validate.ParseJSON(/* args */)
```

## Tips

Import `github.com/sahilkhaire/gox/validate` and call `ParseJSON` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/validate/array">Array</a><a class="related-chip" href="/packages/validate/bool">Bool</a><a class="related-chip" href="/packages/validate/float">Float</a>
</div>

← [Back to validate package overview](/packages/validate/)
