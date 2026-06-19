---
title: "jwt"
package: "jwt"
gox-doc-version: "10"
---

<PackageOverview
  name="jwt"
  node="jsonwebtoken"
  import-path="github.com/sahilkhaire/gox/jwt"
  subtitle="Package jwt provides sign, verify, and decode helpers around JSON Web Tokens. Node equivalent: jsonwebtoken."
  :symbol-count=4
  :has-cookbook=true
  migration-link=""
  narrative="jsonwebtoken-style sign, verify, and decode with golang-jwt underneath."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/jwt/sign"><div class="featured-name">Sign</div><div class="featured-summary">jwt.sign</div></a>
<a class="featured-card" href="/packages/jwt/verify"><div class="featured-name">Verify</div><div class="featured-summary">jwt.verify</div></a>
<a class="featured-card" href="/packages/jwt/decode"><div class="featured-name">Decode</div><div class="featured-summary">jwt.decode</div></a>
</div>

## Common use cases

- Issue session tokens
- Verify Bearer tokens in middleware
- Inspect claims without verifying (decode)

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>jwt.sign(payload, secret, { expiresIn })</code></td><td><a href="/packages/jwt/sign"><code>jwt.Sign(claims, secret, &jwt.SignOptions{ExpiresIn: d})</code></a></td></tr>
<tr><td><code>jwt.verify(token, secret)</code></td><td><a href="/packages/jwt/verify"><code>jwt.Verify(token, secret)</code></a></td></tr>
<tr><td><code>jwt.decode(token)</code></td><td><a href="/packages/jwt/decode"><code>jwt.Decode(token)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/jwt"

token, err := jwt.Sign(claims, secret, jwt.SignOptions{ExpiresIn: time.Hour})
```

::: tip Full cookbook
See the [**jwt cookbook**](/packages/jwt/cookbook) for multi-step recipes and real-world patterns.
:::

## API reference

Browse **4 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/jwt/decode">Decode</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Decode parses claims without verifying the signature (jwt.decode).</td></tr>
<tr><td><a href="/packages/jwt/sign">Sign</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Sign builds a signed JWT from claims (jwt.sign).</td></tr>
<tr><td><a href="/packages/jwt/verify">Verify</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Verify parses and validates a JWT with secret (jwt.verify).</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/jwt/sign-options">SignOptions</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>SignOptions configures token signing.</td></tr>
</tbody></table>

