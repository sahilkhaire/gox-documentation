---
title: "env Cookbook"
package: "env"
gox-doc-version: "10"
---

<div class="api-meta"><span class="api-badge node">Node: dotenv / process.env</span><span class="api-badge import">github.com/sahilkhaire/gox/env</span></div>
# env Cookbook

dotenv-style configuration with typed getters.

```go
import "github.com/sahilkhaire/gox/env"

_ = env.Load() // reads .env from cwd

port := env.Get("PORT", "8080")
debug := env.GetBool("DEBUG", false)
timeout := env.GetDuration("TIMEOUT", 30*time.Second)
```

← [Back to env overview](/packages/env/)
