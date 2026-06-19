---
title: "Find"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "arr.find(fn)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: arr.find(fn)</span><span class="api-badge import">github.com/sahilkhaire/gox/slice</span></div>
# Find

## Overview

Find returns the first element matching fn, or zero and false (Array.find).

**Node.js equivalent:** `arr.find(fn)`

## Signature

```go
func Find[T any](in []T, fn func(T) bool) (T, bool)
```

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
