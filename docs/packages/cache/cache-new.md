---
title: "Cache.New"
package: "cache"
import: "github.com/sahilkhaire/gox/cache"
node: "new LRUCache({ max })"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: new LRUCache({ max })</span><span class="api-badge import">github.com/sahilkhaire/gox/cache</span></div>
# Cache.New

## Overview

New creates a cache that evicts least-recently-used entries when size exceeds maxSize.

**Node.js equivalent:** `new LRUCache({ max })`

## Signature

```go
func New(maxSize int) *Cache
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const cache = new LRUCache({ max: 100 });
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/cache"

c := cache.New(100)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

← [Back to cache package overview](/packages/cache/)
