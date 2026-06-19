---
title: "Transform"
package: "stream"
import: "github.com/sahilkhaire/gox/stream"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: Node stream</span><span class="api-badge import">github.com/sahilkhaire/gox/stream</span></div>
# Transform

## Overview

Transform returns a Reader that applies fn to each chunk from reader.

## Signature

```go
func Transform(reader io.Reader, fn func([]byte) ([]byte, error)) io.Reader
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/stream"

// stream
_ = stream.Transform(/* args */)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Pipe](/packages/stream/pipe)
- [ReadAll](/packages/stream/read-all)
- [TeeReader](/packages/stream/tee-reader)

← [Back to stream package overview](/packages/stream/)
