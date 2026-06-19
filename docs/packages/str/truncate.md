---
title: "Truncate"
package: "str"
import: "github.com/sahilkhaire/gox/str"
node: "_.truncate(s, n)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: _.truncate(s, n)</span><span class="api-badge import">github.com/sahilkhaire/gox/str</span></div>
# Truncate

## Overview

Truncate shortens s to max runes, adding suffix if needed.

**Node.js equivalent:** `_.truncate(s, n)`

## Signature

```go
func Truncate(s string, max int, suffix string) string
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
_.truncate(str, { length: 10 });
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/str"

short := str.Truncate(long, 10)
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
