---
title: "Camel"
package: "str"
import: "github.com/sahilkhaire/gox/str"
node: "camelCase(s)"
gox-doc-version: "10"
---

<SymbolHeader pkg="str" title="Camel" node="camelCase(s)" import-path="github.com/sahilkhaire/gox/str" />
## Overview

Converts kebab-case or snake_case strings to camelCase — common when mapping JSON field names to Go struct tags.

## Signature

<div class="signature-block">

```go
func Camel(s string) string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
_.camelCase('foo_bar');
```

```go [Standard Go]
// Manual rune walk or regexp replace
s = "fooBar"
```

```go [gox]
import "github.com/sahilkhaire/gox/str"

s := str.Camel("foo_bar")
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/str/capitalize">Capitalize</a><a class="related-chip" href="/packages/str/is-blank">IsBlank</a><a class="related-chip" href="/packages/str/pad-end">PadEnd</a>
</div>

← [Back to str package overview](/packages/str/)
