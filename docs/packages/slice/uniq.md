---
title: "Uniq"
package: "slice"
import: "github.com/sahilkhaire/gox/slice"
node: "_.uniq(arr)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: _.uniq(arr)</span><span class="api-badge import">github.com/sahilkhaire/gox/slice</span></div>
# Uniq

## Overview

Uniq returns unique elements in first-seen order (lodash uniq).

**Node.js equivalent:** `_.uniq(arr)`

## Signature

```go
func Uniq[T comparable](in []T) []T
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const unique = _.uniq(items);
```

```go [Standard Go]
seen := make(map[T]struct{})
var unique []T
for _, v := range items { /* dedupe */ }
```

```go [gox]
import "github.com/sahilkhaire/gox/slice"

unique := slice.Uniq(items)
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
