---
title: "Snake"
package: "str"
import: "github.com/sahilkhaire/gox/str"
node: "snake_case(s)"
gox-doc-version: "10"
---

<SymbolHeader pkg="str" title="Snake" node="snake_case(s)" import-path="github.com/sahilkhaire/gox/str" />
## Overview

Snake converts to snake_case.

**Node.js equivalent:** `snake_case(s)`

## Signature

<div class="signature-block">

```go
func Snake(s string) string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
_.snakeCase('FooBar');
```

```go [Standard Go]
s = strings.ToLower(strings.ReplaceAll(name, " ", "_"))
```

```go [gox]
import "github.com/sahilkhaire/gox/str"

s := str.Snake("FooBar")
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/str/camel">Camel</a><a class="related-chip" href="/packages/str/capitalize">Capitalize</a><a class="related-chip" href="/packages/str/is-blank">IsBlank</a>
</div>

← [Back to str package overview](/packages/str/)
