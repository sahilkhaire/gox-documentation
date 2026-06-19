---
title: "Snake"
package: "str"
import: "github.com/sahilkhaire/gox/str"
node: "snake_case(s)"
gox-doc-version: "11"
---

<SymbolHeader pkg="str" title="Snake" node="snake_case(s)" import-path="github.com/sahilkhaire/gox/str" />
## Overview

Snake converts to snake_case.

If you are coming from Node.js, the closest pattern is **`snake_case(s)`**.

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

## Example

```go
import "github.com/sahilkhaire/gox/str"

s := str.Snake("FooBar")
```

## Tips

Import `github.com/sahilkhaire/gox/str` and call `Snake` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/str/camel">Camel</a><a class="related-chip" href="/packages/str/capitalize">Capitalize</a><a class="related-chip" href="/packages/str/is-blank">IsBlank</a>
</div>

← [Back to str package overview](/packages/str/)
