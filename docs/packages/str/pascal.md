---
title: "Pascal"
package: "str"
import: "github.com/sahilkhaire/gox/str"
node: "PascalCase(s)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: PascalCase(s)</span><span class="api-badge import">github.com/sahilkhaire/gox/str</span></div>
# Pascal

## Overview

Maps the Node.js pattern `PascalCase(s)` to gox `str.Pascal(s)`.

**Node.js equivalent:** `PascalCase(s)`

## Signature

```go
func Pascal(s string) string
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
PascalCase(s)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/str"

str.Pascal(s)
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
