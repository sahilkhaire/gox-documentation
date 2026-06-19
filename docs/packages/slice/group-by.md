---
title: "GroupBy"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "_.groupBy(arr, fn)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: _.groupBy(arr, fn)</span><span class="api-badge import">github.com/sahilkhaire/gox/slice</span></div>
# GroupBy

## Overview

GroupBy groups elements by key from fn (lodash groupBy).

**Node.js equivalent:** `_.groupBy(arr, fn)`

## Signature

```go
func GroupBy[T any, K comparable](in []T, fn func(T) K) map[K][]T
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const byRole = _.groupBy(users, u => u.role);
```

```go [Standard Go]
byRole := make(map[string][]User)
for _, u := range users {
    byRole[u.Role] = append(byRole[u.Role], u)
}
```

```go [gox]
import "github.com/sahilkhaire/gox/slice"

byRole := slice.GroupBy(users, func(u User) string { return u.Role })
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Chunk](/packages/slice/chunk)
- [Contains](/packages/slice/contains)
- [Every](/packages/slice/every)

← [Back to slice package overview](/packages/slice/)
