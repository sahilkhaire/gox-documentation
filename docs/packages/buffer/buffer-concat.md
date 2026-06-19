---
title: "Buffer.Concat"
package: "buffer"
import: "github.com/sahilkhaire/gox/buffer"
node: "Buffer.concat(list)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: Buffer.concat(list)</span><span class="api-badge import">github.com/sahilkhaire/gox/buffer</span></div>
# Buffer.Concat

## Overview

Concat joins buffers (Buffer.concat).

**Node.js equivalent:** `Buffer.concat(list)`

## Signature

```go
func Concat(parts ...Buffer) Buffer
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
Buffer.concat([buf1, buf2]);
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/buffer"

combined := buffer.Concat(buf1, buf2)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Buffer.From](/packages/buffer/buffer-from)
- [Buffer.FromString](/packages/buffer/buffer-from-string)

← [Back to buffer package overview](/packages/buffer/)
