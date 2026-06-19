---
title: "IsBlank"
package: "str"
import: "github.com/sahilkhaire/gox/str"
node: "!s.trim()"
gox-doc-version: "10"
---

<SymbolHeader pkg="str" title="IsBlank" node="!s.trim()" import-path="github.com/sahilkhaire/gox/str" />
## Overview

IsBlank reports whether s is empty or only whitespace.

**Node.js equivalent:** `!s.trim()`

## Signature

<div class="signature-block">

```go
func IsBlank(s string) bool
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
!s.trim();
```

```go [Standard Go]
blank := strings.TrimSpace(s) == ""
```

```go [gox]
import "github.com/sahilkhaire/gox/str"

if str.IsBlank(input) { /* empty */ }
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/str/camel">Camel</a><a class="related-chip" href="/packages/str/capitalize">Capitalize</a><a class="related-chip" href="/packages/str/pad-end">PadEnd</a>
</div>

← [Back to str package overview](/packages/str/)
