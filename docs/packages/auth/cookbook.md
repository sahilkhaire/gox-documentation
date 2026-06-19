---
title: "auth Cookbook"
package: "auth"
gox-doc-version: "10"
---

<div class="api-meta"><span class="api-badge node">Node: passport</span><span class="api-badge import">github.com/sahilkhaire/gox/auth</span></div>
# auth Cookbook

Passport-style Bearer, API key, and Basic middleware.

```go
import "github.com/sahilkhaire/gox/auth"

app.Use(auth.Bearer(jwtSecret))
app.Use(auth.APIKey(auth.APIKeyOpts{Header: "X-API-Key", Keys: keys}))
claims, _ := auth.GetClaims(ctx)
```

← [Back to auth overview](/packages/auth/)
