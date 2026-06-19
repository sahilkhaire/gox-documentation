---
title: "jwt Cookbook"
package: "jwt"
gox-doc-version: "10"
---

<div class="api-meta"><span class="api-badge node">Node: jsonwebtoken</span><span class="api-badge import">github.com/sahilkhaire/gox/jwt</span></div>
# jwt Cookbook

jsonwebtoken-style sign and verify.

```go
import "github.com/sahilkhaire/gox/jwt"

token, err := jwt.Sign(
    map[string]any{"sub": "user-123"},
    secret,
    jwt.SignOptions{ExpiresIn: time.Hour},
)
claims, err := jwt.Verify(token, secret)
```

← [Back to jwt overview](/packages/jwt/)
