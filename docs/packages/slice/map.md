---
title: "Map"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "arr.map(fn)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: arr.map(fn)</span><span class="api-badge import">github.com/sahilkhaire/gox/slice</span></div>
# Map

## Overview

Transforms each element of a slice, returning a new slice — identical mental model to `Array.prototype.map`. Uses Go generics so input and output types can differ safely.

**Node.js equivalent:** `arr.map(fn)`

## Signature

```go
func Map[T, U any](in []T, fn func(T) U) []U
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const names = users.map(u => u.name);
```

```go [Standard Go]
names := make([]string, len(users))
for i, u := range users {
    names[i] = u.Name
}
```

```go [gox]
import "github.com/sahilkhaire/gox/slice"

names := slice.Map(users, func(u User) string { return u.Name })
```

:::

## Tips

Chain with `slice.Filter` and `slice.Reduce` for lodash-style pipelines. Pre-allocates output slice for performance.

## Related APIs

- [Chunk](/packages/slice/chunk)
- [Contains](/packages/slice/contains)
- [Every](/packages/slice/every)

← [Back to slice package overview](/packages/slice/)
