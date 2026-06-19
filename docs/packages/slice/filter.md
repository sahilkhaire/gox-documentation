---
title: "Filter"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "arr.filter(fn)"
gox-doc-version: "11"
---

<SymbolHeader pkg="slice" title="Filter" node="arr.filter(fn)" import-path="github.com/sahilkhaire/gox/slice" />
## Overview

Keeps elements matching a predicate — same as `Array.filter`. Returns a new slice containing only matching items.

If you are coming from Node.js, the closest pattern is **`arr.filter(fn)`**.

## Signature

<div class="signature-block">

```go
func Filter[T any](in []T, fn func(T) bool) []T
```

</div>

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

## Example

```go
import "github.com/sahilkhaire/gox/slice"

adults := slice.Filter(users, func(u User) bool { return u.Age >= 18 })
```

## Tips

Combine with `slice.Map` for filter-then-map pipelines without nested loops.

## Standard library alternative

Use a `for` loop or Go 1.21+ `slices` package helpers from the standard library.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/slice/chunk">Chunk</a><a class="related-chip" href="/packages/slice/contains">Contains</a><a class="related-chip" href="/packages/slice/every">Every</a>
</div>

← [Back to slice package overview](/packages/slice/)
