---
title: "cache Cookbook"
package: "cache"
gox-doc-version: "10"
---

<div class="api-meta"><span class="api-badge node">Node: lru-cache</span><span class="api-badge import">github.com/sahilkhaire/gox/cache</span></div>
# cache Cookbook

In-memory LRU cache with capacity bound.

```go
import "github.com/sahilkhaire/gox/cache"

c := cache.New(1000)
c.Set("user:1", user, time.Hour)
if v, ok := c.Get("user:1"); ok { _ = v }
```

← [Back to cache overview](/packages/cache/)
