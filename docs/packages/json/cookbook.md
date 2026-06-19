---
title: "json Cookbook"
package: "json"
gox-doc-version: "10"
---

<div class="api-meta"><span class="api-badge node">Node: JSON.parse/stringify</span><span class="api-badge import">github.com/sahilkhaire/gox/json</span></div>
# json Cookbook

Typed JSON.parse/stringify with file helpers.

```go
import "github.com/sahilkhaire/gox/json"

cfg, err := json.ParseFile[Config](ctx, "config.json")
out, err := json.Stringify(response)
pretty, err := json.Pretty(debugPayload)
```

← [Back to json overview](/packages/json/)
