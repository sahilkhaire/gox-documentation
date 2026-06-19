---
title: "slice Cookbook"
package: "slice"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: lodash / Array.*</span><span class="api-badge import">github.com/sahilkhaire/gox/slice</span></div>
# slice Cookbook

lodash / Array-style pipelines with Go generics.

## Filter → map pipeline

```go
import "github.com/sahilkhaire/gox/slice"

adults := slice.Filter(users, func(u User) bool { return u.Age >= 18 })
names := slice.Map(adults, func(u User) string { return u.Name })
```

## Reduce to sum (from tests)

```go
in := []int{1, 2, 3, 4}
total := slice.Reduce(
    slice.Filter(slice.Map(in, func(x int) int { return x * 2 }), func(x int) bool { return x > 4 }),
    0,
    func(a, b int) int { return a + b },
) // 14
```

## Group by role

```go
byRole := slice.GroupBy(users, func(u User) string { return u.Role })
```

## Unique IDs preserving order

```go
ids := slice.Uniq([]string{"a", "b", "a", "c"})
```

← [Back to slice overview](/packages/slice/)
