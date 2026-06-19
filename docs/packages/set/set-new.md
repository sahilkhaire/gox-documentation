---
title: "Set.New"
package: "set"
import: "github.com/sahilkhaire/gox/set"
node: "new Set(arr)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: new Set(arr)</span><span class="api-badge import">github.com/sahilkhaire/gox/set</span></div>
# Set.New

## Overview

New creates a set from items.

**Node.js equivalent:** `new Set(arr)`

## Signature

```go
func New[T comparable](items ...T) Set[T]
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const s = new Set([1, 2, 3]);
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/set"

s := set.New(1, 2, 3)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Set.Difference](/packages/set/set-difference)
- [Set.Intersection](/packages/set/set-intersection)
- [Set.Union](/packages/set/set-union)

← [Back to set package overview](/packages/set/)
