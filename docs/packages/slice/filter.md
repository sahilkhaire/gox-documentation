---
title: "Filter"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "arr.filter(fn)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: arr.filter(fn)</span><span class="api-badge import">github.com/sahilkhaire/gox/slice</span></div>
# Filter

## Overview

Keeps elements matching a predicate — same as `Array.filter`. Returns a new slice containing only matching items.

**Node.js equivalent:** `arr.filter(fn)`

## Signature

```go
func Filter[T any](in []T, fn func(T) bool) []T
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const adults = users.filter(u => u.age >= 18);
```

```go [Standard Go]
var adults []User
for _, u := range users {
    if u.Age >= 18 {
        adults = append(adults, u)
    }
}
```

```go [gox]
import "github.com/sahilkhaire/gox/slice"

adults := slice.Filter(users, func(u User) bool { return u.Age >= 18 })
```

:::

## Tips

Combine with `slice.Map` for filter-then-map pipelines without nested loops.

## Related APIs

- [Chunk](/packages/slice/chunk)
- [Contains](/packages/slice/contains)
- [Every](/packages/slice/every)

← [Back to slice package overview](/packages/slice/)
