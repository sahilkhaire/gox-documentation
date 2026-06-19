---
title: "IsBlank"
package: "str"
import: "github.com/sahilkhaire/gox/str"
node: "!s.trim()"
gox-doc-version: "14"
---

<SymbolHeader pkg="str" title="IsBlank" node="!s.trim()" import-path="github.com/sahilkhaire/gox/str" />
## Overview

IsBlank reports whether s is empty or only whitespace.

If you are coming from Node.js, the closest pattern is **`!s.trim()`**.

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

## Example

```go
import "github.com/sahilkhaire/gox/str"

if str.IsBlank(input) { /* empty */ }
```

## Tips

Import `github.com/sahilkhaire/gox/str` and call `IsBlank` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
blank := strings.TrimSpace(s) == ""
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/str/camel">Camel</a><a class="related-chip" href="/packages/str/capitalize">Capitalize</a><a class="related-chip" href="/packages/str/pad-end">PadEnd</a>
</div>

← [Back to str package overview](/packages/str/)
