---
title: "log Cookbook"
package: "log"
gox-doc-version: "10"
---

<div class="api-meta"><span class="api-badge node">Node: winston, pino</span><span class="api-badge import">github.com/sahilkhaire/gox/log</span></div>
# log Cookbook

Structured logging with slog underneath.

```go
import "github.com/sahilkhaire/gox/log"

logger := log.New()
log.SetDefault(logger)
logger.Info("request", "method", "GET", "path", "/health")
```

← [Back to log overview](/packages/log/)
