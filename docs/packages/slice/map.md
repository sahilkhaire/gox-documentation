---
title: "Map"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "arr.map(fn)"
gox-doc-version: "11"
---

<SymbolHeader pkg="slice" title="Map" node="arr.map(fn)" import-path="github.com/sahilkhaire/gox/slice" />
## Overview

Transforms each element of a slice, returning a new slice — identical mental model to `Array.prototype.map`. Uses Go generics so input and output types can differ safely.

If you are coming from Node.js, the closest pattern is **`arr.map(fn)`**.

## Signature

<div class="signature-block">

```go
func Map[T, U any](in []T, fn func(T) U) []U
```

</div>

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

## Example

```go
import "github.com/sahilkhaire/gox/slice"

names := slice.Map(users, func(u User) string { return u.Name })
```

## Tips

Chain with `slice.Filter` and `slice.Reduce` for lodash-style pipelines. Pre-allocates output slice for performance.

## Standard library alternative

Use a `for` loop or Go 1.21+ `slices` package helpers from the standard library.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/slice/chunk">Chunk</a><a class="related-chip" href="/packages/slice/contains">Contains</a><a class="related-chip" href="/packages/slice/every">Every</a>
</div>

← [Back to slice package overview](/packages/slice/)
