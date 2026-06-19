---
title: "GroupBy"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "_.groupBy(arr, fn)"
gox-doc-version: "11"
---

<SymbolHeader pkg="slice" title="GroupBy" node="_.groupBy(arr, fn)" import-path="github.com/sahilkhaire/gox/slice" />
## Overview

GroupBy groups elements by key from fn (lodash groupBy).

If you are coming from Node.js, the closest pattern is **`_.groupBy(arr, fn)`**.

## Signature

<div class="signature-block">

```go
func GroupBy[T any, K comparable](in []T, fn func(T) K) map[K][]T
```

</div>

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

## Example

```go
import "github.com/sahilkhaire/gox/slice"

byRole := slice.GroupBy(users, func(u User) string { return u.Role })
```

## Tips

Chain `Filter`, `Map`, and `Reduce` for lodash-style pipelines. Results are new slices — inputs are never mutated.

## Standard library alternative

Use a `for` loop or Go 1.21+ `slices` package helpers from the standard library.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/slice/chunk">Chunk</a><a class="related-chip" href="/packages/slice/contains">Contains</a><a class="related-chip" href="/packages/slice/every">Every</a>
</div>

← [Back to slice package overview](/packages/slice/)
