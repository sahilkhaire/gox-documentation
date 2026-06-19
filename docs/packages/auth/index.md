---
title: "auth"
package: "auth"
gox-doc-version: "10"
---

<PackageOverview
  name="auth"
  node="passport"
  import-path="github.com/sahilkhaire/gox/auth"
  subtitle="Package auth provides Passport-style authentication middleware for gox/http. Node equivalent: passport, passport-jwt, passport-http-bearer."
  :symbol-count=8
  :has-cookbook=true
  migration-link="/migration/passport"
  narrative="Passport-style Bearer, API key, and Basic auth middleware with context claims."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/auth/bearer"><div class="featured-name">Bearer</div><div class="featured-summary">JWT Bearer auth</div></a>
<a class="featured-card" href="/packages/auth/api-key"><div class="featured-name">APIKey</div><div class="featured-summary">API key header</div></a>
<a class="featured-card" href="/packages/auth/basic"><div class="featured-name">Basic</div><div class="featured-summary">Basic auth</div></a>
</div>

## Common use cases

- Protect routes with JWT Bearer
- Accept API keys from headers
- Read authenticated user from context

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/auth"

app.Use(auth.Bearer(secret))
```

::: tip Full cookbook
See the [**auth cookbook**](/packages/auth/cookbook) for multi-step recipes and real-world patterns.
:::

::: info Migration guide
Coming from Node.js? Read the [**migration guide**](/migration/passport) for side-by-side patterns.
:::

## API reference

Browse **8 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

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

