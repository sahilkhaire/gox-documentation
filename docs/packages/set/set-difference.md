---
title: "Set.Difference"
package: "set"
import: "github.com/sahilkhaire/gox/set"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: Set</span><span class="api-badge import">github.com/sahilkhaire/gox/set</span></div>
# Set.Difference

## Overview

Difference returns elements in a but not in b.

## Signature

```go
func Difference[T comparable](a, b Set[T]) Set[T]
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
import "github.com/sahilkhaire/gox/set"

var v Set
v.Difference(/* args */)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Set.Intersection](/packages/set/set-intersection)
- [Set.New](/packages/set/set-new)
- [Set.Union](/packages/set/set-union)

← [Back to set package overview](/packages/set/)
