---
title: "Buffer"
package: "buffer"
import: "github.com/sahilkhaire/gox/buffer"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: Buffer</span><span class="api-badge import">github.com/sahilkhaire/gox/buffer</span></div>
# Buffer

## Overview

Buffer is a byte slice with helper methods.

## Signature

```go
type Buffer []byte
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
import "github.com/sahilkhaire/gox/buffer"

_ = buffer.Buffer
```

:::

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

```go
import "github.com/sahilkhaire/gox/buffer"

a := FromString("ab")
b := From([]byte("ab"))
c := Concat(FromString("a"), FromString("b"))
```

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Compare](/packages/buffer/compare)
- [Equals](/packages/buffer/equals)

← [Back to buffer package overview](/packages/buffer/)
