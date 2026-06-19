---
title: "PadEnd"
package: "str"
import: "github.com/sahilkhaire/gox/str"
node: "s.padEnd(n, ch)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: s.padEnd(n, ch)</span><span class="api-badge import">github.com/sahilkhaire/gox/str</span></div>
# PadEnd

## Overview

Maps the Node.js pattern `s.padEnd(n, ch)` to gox `str.PadEnd(s, n, ch)`.

**Node.js equivalent:** `s.padEnd(n, ch)`

## Signature

```go
func PadEnd(s string, length int, pad string) string
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
s.padEnd(n, ch)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/str"

str.PadEnd(s, n, ch)
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
