---
title: "jwt"
package: "jwt"
gox-doc-version: "7"
---

<div class="package-hero"><div class="api-meta"><span class="api-badge node">Node: jsonwebtoken</span></div><h1>jwt</h1>
<p class="subtitle">Package jwt provides sign, verify, and decode helpers around JSON Web Tokens. Node equivalent: jsonwebtoken.</p><div class="import-line">import "github.com/sahilkhaire/gox/jwt"</div></div>
## Quick start

A minimal example to get going:

```go
import "github.com/sahilkhaire/gox/jwt"

token, err := jwt.Sign(claims, secret, jwt.SignOptions{ExpiresIn: time.Hour})
```

::: tip Full cookbook
See the [**jwt cookbook**](/packages/jwt/cookbook) for multi-step recipes and real-world patterns.
:::

## What's in this package

If you have used **jsonwebtoken** in Node.js, this package is your starting point. Browse **4 documented symbols** below — each page includes Node.js, standard Go, and gox side-by-side examples.

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/jwt/decode">Decode</a></td><td><span class="kind-pill">func</span></td><td>Decode parses claims without verifying the signature (jwt.decode).</td></tr>
<tr><td><a href="/packages/jwt/sign">Sign</a></td><td><span class="kind-pill">func</span></td><td>Sign builds a signed JWT from claims (jwt.sign).</td></tr>
<tr><td><a href="/packages/jwt/verify">Verify</a></td><td><span class="kind-pill">func</span></td><td>Verify parses and validates a JWT with secret (jwt.verify).</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/jwt/sign-options">SignOptions</a></td><td><span class="kind-pill">type</span></td><td>SignOptions configures token signing.</td></tr>
</tbody></table>

