---
title: "Pascal"
package: "str"
import: "github.com/sahilkhaire/gox/str"
node: "PascalCase(s)"
gox-doc-version: "10"
---

<SymbolHeader pkg="str" title="Pascal" node="PascalCase(s)" import-path="github.com/sahilkhaire/gox/str" />
## Overview

Pascal converts to PascalCase.

**Node.js equivalent:** `PascalCase(s)`

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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/str/camel">Camel</a><a class="related-chip" href="/packages/str/capitalize">Capitalize</a><a class="related-chip" href="/packages/str/is-blank">IsBlank</a>
</div>

← [Back to str package overview](/packages/str/)
