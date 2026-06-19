---
title: "redis Cookbook"
package: "redis"
gox-doc-version: "10"
---

<div class="api-meta"><span class="api-badge node">Node: ioredis</span><span class="api-badge import">github.com/sahilkhaire/gox/redis</span></div>
# redis Cookbook

ioredis-style commands with context.

```go
import "github.com/sahilkhaire/gox/redis"

rdb, err := redis.New("localhost:6379")
defer rdb.Close()

err = rdb.Set(ctx, "key", "value", time.Hour)
val, err := rdb.Get(ctx, "key")
```

← [Back to redis overview](/packages/redis/)
