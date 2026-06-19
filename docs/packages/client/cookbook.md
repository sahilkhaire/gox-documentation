---
title: "client Cookbook"
package: "client"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: axios, fetch</span><span class="api-badge import">github.com/sahilkhaire/gox/client</span></div>
# client Cookbook

axios / fetch-style HTTP client.

```go
import "github.com/sahilkhaire/gox/client"

c := client.New()
res, err := c.Get(ctx, "https://api.example.com/users", client.WithQuery(url.Values{
    "page": {"1"},
}))
var users []User
if err := res.JSON(&users); err != nil { /* handle */ }
```

← [Back to client overview](/packages/client/)
