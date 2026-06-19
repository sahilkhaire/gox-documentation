---
title: "Pipe"
package: "stream"
import: "github.com/sahilkhaire/gox/stream"
node: "src.pipe(dst)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: src.pipe(dst)</span><span class="api-badge import">github.com/sahilkhaire/gox/stream</span></div>
# Pipe

## Overview

Maps the Node.js pattern `src.pipe(dst)` to gox `stream.Pipe(src, dst)`.

**Node.js equivalent:** `src.pipe(dst)`

## Signature

```go
func Pipe(src io.Reader, dst io.Writer) (int64, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
src.pipe(dst)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/stream"

stream.Pipe(src, dst)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [ReadAll](/packages/stream/read-all)
- [TeeReader](/packages/stream/tee-reader)
- [Transform](/packages/stream/transform)

← [Back to stream package overview](/packages/stream/)
