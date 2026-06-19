---
title: "Pascal"
package: "str"
import: "github.com/sahilkhaire/gox/str"
node: "PascalCase(s)"
gox-doc-version: "11"
---

<SymbolHeader pkg="str" title="Pascal" node="PascalCase(s)" import-path="github.com/sahilkhaire/gox/str" />
## Overview

Pascal converts to PascalCase.

If you are coming from Node.js, the closest pattern is **`PascalCase(s)`**.

## Signature

<div class="signature-block">

```go
func Pascal(s string) string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
_.startCase('foo bar').replace(/ /g, '');
```

```go [Standard Go]
// Manual string transformation with strings/unicode
```

```go [gox]
import "github.com/sahilkhaire/gox/str"

s := str.Pascal("foo bar")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/str"

s := str.Pascal("foo bar")
```

## Tips

Import `github.com/sahilkhaire/gox/str` and call `Pascal` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/str/camel">Camel</a><a class="related-chip" href="/packages/str/capitalize">Capitalize</a><a class="related-chip" href="/packages/str/is-blank">IsBlank</a>
</div>

← [Back to str package overview](/packages/str/)
