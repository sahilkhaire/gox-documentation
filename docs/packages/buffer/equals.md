---
title: "Equals"
package: "buffer"
import: "github.com/sahilkhaire/gox/buffer"
node: "buf.equals(other)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: buf.equals(other)</span><span class="api-badge import">github.com/sahilkhaire/gox/buffer</span></div>
# Equals

## Overview

Maps the Node.js pattern `buf.equals(other)` to gox `buffer.Equals(a, b)`.

**Node.js equivalent:** `buf.equals(other)`

## Signature

```go
func Equals(a, b Buffer) bool
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
buf.equals(other)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/buffer"

buffer.Equals(a, b)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Compare](/packages/buffer/compare)

← [Back to buffer package overview](/packages/buffer/)
