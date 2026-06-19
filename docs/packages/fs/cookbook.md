---
title: "fs Cookbook"
package: "fs"
gox-doc-version: "10"
---

<div class="api-meta"><span class="api-badge node">Node: fs.promises</span><span class="api-badge import">github.com/sahilkhaire/gox/fs</span></div>
# fs Cookbook

Context-aware file I/O — read, write, exists, mkdir.

```go
import "github.com/sahilkhaire/gox/fs"

data, err := fs.ReadFile(ctx, "config.json")
if ok, _ := fs.Exists(ctx, "uploads"); !ok {
    err = fs.Mkdir(ctx, "uploads", 0755)
}
err = fs.WriteFile(ctx, "out.json", data)
```

← [Back to fs overview](/packages/fs/)
