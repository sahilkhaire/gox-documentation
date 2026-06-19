---
title: "PadStart"
package: "str"
import: "github.com/sahilkhaire/gox/str"
node: "s.padStart(n, ch)"
gox-doc-version: "11"
---

<SymbolHeader pkg="str" title="PadStart" node="s.padStart(n, ch)" import-path="github.com/sahilkhaire/gox/str" />
## Overview

PadStart pads s on the left to length with pad (String.padStart).

If you are coming from Node.js, the closest pattern is **`s.padStart(n, ch)`**.

## Signature

<div class="signature-block">

```go
func PadStart(s string, length int, pad string) string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
s.padStart(8, '0');
```

```go [Standard Go]
s = fmt.Sprintf("%*s", width, s) // or custom padding
```

```go [gox]
import "github.com/sahilkhaire/gox/str"

s := str.PadStart("42", 8, "0")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/str"

s := str.PadStart("42", 8, "0")
```

## Tips

Import `github.com/sahilkhaire/gox/str` and call `PadStart` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/str/camel">Camel</a><a class="related-chip" href="/packages/str/capitalize">Capitalize</a><a class="related-chip" href="/packages/str/is-blank">IsBlank</a>
</div>

← [Back to str package overview](/packages/str/)
