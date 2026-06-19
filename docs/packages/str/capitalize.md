---
title: "Capitalize"
package: "str"
import: "github.com/sahilkhaire/gox/str"
node: "_.capitalize(s)"
gox-doc-version: "14"
---

<SymbolHeader pkg="str" title="Capitalize" node="_.capitalize(s)" import-path="github.com/sahilkhaire/gox/str" />
## Overview

Capitalize uppercases the first rune.

If you are coming from Node.js, the closest pattern is **`_.capitalize(s)`**.

## Signature

<div class="signature-block">

```go
func Capitalize(s string) string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
_.capitalize('hello');
```

```go [Standard Go]
// Manual string transformation with strings/unicode
```

```go [gox]
import "github.com/sahilkhaire/gox/str"

s := str.Capitalize("hello")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/str"

s := str.Capitalize("hello")
```

## Tips

Import `github.com/sahilkhaire/gox/str` and call `Capitalize` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/str/camel">Camel</a><a class="related-chip" href="/packages/str/is-blank">IsBlank</a><a class="related-chip" href="/packages/str/pad-end">PadEnd</a>
</div>

← [Back to str package overview](/packages/str/)
