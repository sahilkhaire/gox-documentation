---
title: "Compare"
package: "buffer"
import: "github.com/sahilkhaire/gox/buffer"
node: "buf.compare(other)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: buf.compare(other)</span><span class="api-badge import">github.com/sahilkhaire/gox/buffer</span></div>
# Compare

## Overview

Maps the Node.js pattern `buf.compare(other)` to gox `buffer.Compare(a, b)`.

**Node.js equivalent:** `buf.compare(other)`

## Signature

```go
func Compare(a, b Buffer) int
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
buf.compare(other)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/buffer"

buffer.Compare(a, b)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Equals](/packages/buffer/equals)

← [Back to buffer package overview](/packages/buffer/)
