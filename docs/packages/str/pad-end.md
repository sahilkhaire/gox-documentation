---
title: "PadEnd"
package: "str"
import: "github.com/sahilkhaire/gox/str"
node: "s.padEnd(n, ch)"
gox-doc-version: "10"
---

<SymbolHeader pkg="str" title="PadEnd" node="s.padEnd(n, ch)" import-path="github.com/sahilkhaire/gox/str" />
## Overview

PadEnd pads s on the right to length with pad (String.padEnd).

**Node.js equivalent:** `s.padEnd(n, ch)`

## Signature

<div class="signature-block">

```go
func PadEnd(s string, length int, pad string) string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
s.padEnd(8, '-');
```

```go [Standard Go]
s = fmt.Sprintf("%*s", width, s) // or custom padding
```

```go [gox]
import "github.com/sahilkhaire/gox/str"

s := str.PadEnd("go", 8, "-")
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/str/camel">Camel</a><a class="related-chip" href="/packages/str/capitalize">Capitalize</a><a class="related-chip" href="/packages/str/is-blank">IsBlank</a>
</div>

← [Back to str package overview](/packages/str/)
