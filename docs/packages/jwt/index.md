---
title: "jwt"
package: "jwt"
gox-doc-version: "14"
---

<PackageOverview
  name="jwt"
  node="jsonwebtoken"
  import-path="github.com/sahilkhaire/gox/jwt"
  subtitle="Package jwt provides sign, verify, and decode helpers around JSON Web Tokens. Node equivalent: jsonwebtoken."
  :symbol-count=4
/>
## API reference

Select a symbol below — each page explains what it does, shows Node.js vs Go comparisons, and includes a runnable example.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/jwt/decode">Decode</a></td><td><code class="node-cell">jwt.decode(token)</code></td><td><span class="kind-pill">func</span></td><td>Decode parses claims without verifying the signature (jwt.decode).</td></tr>
<tr><td><a href="/packages/jwt/sign">Sign</a></td><td><code class="node-cell">jwt.sign(payload, secret, { expiresIn })</code></td><td><span class="kind-pill">func</span></td><td>Sign builds a signed JWT from claims (jwt.sign).</td></tr>
<tr><td><a href="/packages/jwt/verify">Verify</a></td><td><code class="node-cell">jwt.verify(token, secret)</code></td><td><span class="kind-pill">func</span></td><td>Verify parses and validates a JWT with secret (jwt.verify).</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/jwt/sign-options">SignOptions</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>SignOptions configures token signing.</td></tr>
</tbody></table>

