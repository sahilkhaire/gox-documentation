---
title: "crypto Cookbook"
package: "crypto"
gox-doc-version: "10"
---

<div class="api-meta"><span class="api-badge node">Node: crypto, bcrypt</span><span class="api-badge import">github.com/sahilkhaire/gox/crypto</span></div>
# crypto Cookbook

Hashing, HMAC, and bcrypt password helpers.

```go
import "github.com/sahilkhaire/gox/crypto"

hash, err := crypto.HashPassword(password)
ok := crypto.CheckPassword(password, hash)
digest := crypto.SHA256([]byte(payload))
```

← [Back to crypto overview](/packages/crypto/)
