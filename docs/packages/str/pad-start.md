---
title: "PadStart"
package: "str"
import: "github.com/sahilkhaire/gox/str"
node: "s.padStart(n, ch)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: s.padStart(n, ch)</span><span class="api-badge import">github.com/sahilkhaire/gox/str</span></div>
# PadStart

## Overview

Maps the Node.js pattern `s.padStart(n, ch)` to gox `str.PadStart(s, n, ch)`.

**Node.js equivalent:** `s.padStart(n, ch)`

## Signature

```go
func PadStart(s string, length int, pad string) string
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
s.padStart(n, ch)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/str"

str.PadStart(s, n, ch)
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
- [Capitalize](/packages/str/capitalize)
- [IsBlank](/packages/str/is-blank)

← [Back to str package overview](/packages/str/)
