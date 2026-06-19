---
title: "auth"
package: "auth"
gox-doc-version: "11"
---

<PackageOverview
  name="auth"
  node="passport"
  import-path="github.com/sahilkhaire/gox/auth"
  subtitle="Package auth provides Passport-style authentication middleware for gox/http. Node equivalent: passport, passport-jwt, passport-http-bearer."
  :symbol-count=8
/>
## API reference

Select a symbol below — each page explains what it does, shows Node.js vs Go comparisons, and includes a runnable example.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/auth/api-key">APIKey</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>APIKey validates a header API key and stores it on the request context.</td></tr>
<tr><td><a href="/packages/auth/basic">Basic</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Basic validates HTTP Basic credentials.</td></tr>
<tr><td><a href="/packages/auth/bearer">Bearer</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Bearer validates Authorization: Bearer &lt;jwt&gt; and stores claims on the request context.</td></tr>
<tr><td><a href="/packages/auth/get-api-key">GetAPIKey</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>GetAPIKey returns the API key set by APIKey middleware.</td></tr>
<tr><td><a href="/packages/auth/get-basic-user">GetBasicUser</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>GetBasicUser returns the username set by Basic middleware.</td></tr>
<tr><td><a href="/packages/auth/get-claims">GetClaims</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>GetClaims returns JWT claims set by Bearer middleware.</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/auth/api-key-opts">APIKeyOpts</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>APIKeyOpts configures API key middleware.</td></tr>
<tr><td><a href="/packages/auth/bearer-options">BearerOptions</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>BearerOptions configures JWT bearer middleware.</td></tr>
</tbody></table>

