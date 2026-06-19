---
title: "Buffer.From"
package: "buffer"
import: "github.com/sahilkhaire/gox/buffer"
node: "Buffer.from(data)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: Buffer.from(data)</span><span class="api-badge import">github.com/sahilkhaire/gox/buffer</span></div>
# Buffer.From

## Overview

From allocates a copy of data (Buffer.from).

**Node.js equivalent:** `Buffer.from(data)`

## Signature

```go
func From(data []byte) Buffer
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
Buffer.from([1, 2, 3]);
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/buffer"

b := buffer.From([]byte{1, 2, 3})
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
- [Buffer.FromString](/packages/buffer/buffer-from-string)

← [Back to buffer package overview](/packages/buffer/)
