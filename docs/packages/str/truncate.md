---
title: "Truncate"
package: "str"
import: "github.com/sahilkhaire/gox/str"
node: "_.truncate(s, n)"
gox-doc-version: "10"
---

<SymbolHeader pkg="str" title="Truncate" node="_.truncate(s, n)" import-path="github.com/sahilkhaire/gox/str" />
## Overview

Truncate shortens s to max runes, adding suffix if needed.

**Node.js equivalent:** `_.truncate(s, n)`

## Signature

<div class="signature-block">

```go
func Truncate(s string, max int, suffix string) string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
_.truncate(str, { length: 10 });
```

```go [Standard Go]
if len([]rune(s)) > 10 { s = string([]rune(s)[:10]) + "…" }
```

```go [gox]
import "github.com/sahilkhaire/gox/str"

short := str.Truncate(long, 10)
```

:::

## Tips

Truncates by rune count, not byte length — safe for Unicode.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/str/camel">Camel</a><a class="related-chip" href="/packages/str/capitalize">Capitalize</a><a class="related-chip" href="/packages/str/is-blank">IsBlank</a>
</div>

← [Back to str package overview](/packages/str/)
