---
title: "Camel"
package: "str"
import: "github.com/sahilkhaire/gox/str"
node: "camelCase(s)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: camelCase(s)</span><span class="api-badge import">github.com/sahilkhaire/gox/str</span></div>
# Camel

## Overview

Camel converts kebab/snake to camelCase.

**Node.js equivalent:** `camelCase(s)`

## Signature

```go
func Camel(s string) string
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
_.camelCase('foo_bar');
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/str"

s := str.Camel("foo_bar")
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Capitalize](/packages/str/capitalize)
- [IsBlank](/packages/str/is-blank)
- [PadEnd](/packages/str/pad-end)

← [Back to str package overview](/packages/str/)
