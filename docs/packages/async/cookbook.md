---
title: "async Cookbook"
package: "async"
gox-doc-version: "10"
---

<div class="api-meta"><span class="api-badge node">Node: Promise.all, timers</span><span class="api-badge import">github.com/sahilkhaire/gox/async</span></div>
# async Cookbook

Promise.all / race / sleep with context cancellation.

```go
import "github.com/sahilkhaire/gox/async"

a, b, err := async.All(ctx, fetchUsers, fetchPosts)
winner, err := async.Race(ctx, fastPath, slowPath)
if err := async.Sleep(ctx, 2*time.Second); err != nil { return err }
```

← [Back to async overview](/packages/async/)
