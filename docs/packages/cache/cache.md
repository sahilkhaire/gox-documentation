---
title: "Cache"
package: "cache"
import: "github.com/sahilkhaire/gox/cache"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: lru-cache</span><span class="api-badge import">github.com/sahilkhaire/gox/cache</span></div>
# Cache

## Overview

Cache is a thread-safe LRU cache with per-key TTL.

## Signature

```go
type Cache struct {
	// contains filtered or unexported fields
}
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/cache"

_ = cache.Cache
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

← [Back to cache package overview](/packages/cache/)
