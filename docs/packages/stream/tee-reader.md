---
title: "TeeReader"
package: "stream"
import: "github.com/sahilkhaire/gox/stream"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: Node stream</span><span class="api-badge import">github.com/sahilkhaire/gox/stream</span></div>
# TeeReader

## Overview

TeeReader returns a Reader that writes to w what it reads from r.

## Signature

```go
func TeeReader(r io.Reader, w io.Writer) io.Reader
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
_ = stream.TeeReader(/* args */)
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
- [Transform](/packages/stream/transform)

← [Back to stream package overview](/packages/stream/)
