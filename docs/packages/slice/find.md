---
title: "Find"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "arr.find(fn)"
gox-doc-version: "11"
---

<SymbolHeader pkg="slice" title="Find" node="arr.find(fn)" import-path="github.com/sahilkhaire/gox/slice" />
## Overview

Find returns the first element matching fn, or zero and false (Array.find).

If you are coming from Node.js, the closest pattern is **`arr.find(fn)`**.

## Signature

<div class="signature-block">

```go
func Find[T any](in []T, fn func(T) bool) (T, bool)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const user = users.find(u => u.id === id);
```

```go [Standard Go]
var found User
var ok bool
for _, u := range users {
    if u.ID == id {
        found, ok = u, true
        break
    }
}
```

```go [gox]
import "github.com/sahilkhaire/gox/slice"

user, ok := slice.Find(users, func(u User) bool { return u.ID == id })
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/slice"

user, ok := slice.Find(users, func(u User) bool { return u.ID == id })
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
