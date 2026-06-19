---
title: "Buffer.FromString"
package: "buffer"
import: "github.com/sahilkhaire/gox/buffer"
node: "Buffer.from(str)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: Buffer.from(str)</span><span class="api-badge import">github.com/sahilkhaire/gox/buffer</span></div>
# Buffer.FromString

## Overview

Maps the Node.js pattern `Buffer.from(str)` to gox `buffer.FromString(str)`.

**Node.js equivalent:** `Buffer.from(str)`

## Signature

```go
func FromString(s string) Buffer
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
Buffer.from(str)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/buffer"

buffer.FromString(str)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Buffer.Concat](/packages/buffer/buffer-concat)
- [Buffer.From](/packages/buffer/buffer-from)

← [Back to buffer package overview](/packages/buffer/)
