---
title: "Capitalize"
package: "str"
import: "github.com/sahilkhaire/gox/str"
node: "_.capitalize(s)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: _.capitalize(s)</span><span class="api-badge import">github.com/sahilkhaire/gox/str</span></div>
# Capitalize

## Overview

Maps the Node.js pattern `_.capitalize(s)` to gox `str.Capitalize(s)`.

**Node.js equivalent:** `_.capitalize(s)`

## Signature

```go
func Capitalize(s string) string
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
_.capitalize(s)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/str"

str.Capitalize(s)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Camel](/packages/str/camel)
- [IsBlank](/packages/str/is-blank)
- [PadEnd](/packages/str/pad-end)

← [Back to str package overview](/packages/str/)
