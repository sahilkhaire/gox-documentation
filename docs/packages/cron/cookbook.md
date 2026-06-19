---
title: "cron Cookbook"
package: "cron"
gox-doc-version: "10"
---

<div class="api-meta"><span class="api-badge node">Node: node-cron</span><span class="api-badge import">github.com/sahilkhaire/gox/cron</span></div>
# cron Cookbook

Schedule recurring jobs with cron expressions.

```go
import "github.com/sahilkhaire/gox/cron"

sched := cron.New()
sched.Schedule("0 * * * *", func() { cleanup() })
sched.Start()
```

← [Back to cron overview](/packages/cron/)
