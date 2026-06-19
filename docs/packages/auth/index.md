---
title: "auth"
package: "auth"
gox-doc-version: "7"
---

<div class="package-hero"><div class="api-meta"><span class="api-badge node">Node: passport</span></div><h1>auth</h1>
<p class="subtitle">Package auth provides Passport-style authentication middleware for gox/http. Node equivalent: passport, passport-jwt, passport-http-bearer.</p><div class="import-line">import "github.com/sahilkhaire/gox/auth"</div></div>
## What's in this package

If you have used **passport** in Node.js, this package is your starting point. Browse **8 documented symbols** below — each page includes Node.js, standard Go, and gox side-by-side examples.

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/auth/api-key">APIKey</a></td><td><span class="kind-pill">func</span></td><td>APIKey validates a header API key and stores it on the request context.</td></tr>
<tr><td><a href="/packages/auth/basic">Basic</a></td><td><span class="kind-pill">func</span></td><td>Basic validates HTTP Basic credentials.</td></tr>
<tr><td><a href="/packages/auth/bearer">Bearer</a></td><td><span class="kind-pill">func</span></td><td>Bearer validates Authorization: Bearer &lt;jwt&gt; and stores claims on the request context.</td></tr>
<tr><td><a href="/packages/auth/get-api-key">GetAPIKey</a></td><td><span class="kind-pill">func</span></td><td>GetAPIKey returns the API key set by APIKey middleware.</td></tr>
<tr><td><a href="/packages/auth/get-basic-user">GetBasicUser</a></td><td><span class="kind-pill">func</span></td><td>GetBasicUser returns the username set by Basic middleware.</td></tr>
<tr><td><a href="/packages/auth/get-claims">GetClaims</a></td><td><span class="kind-pill">func</span></td><td>GetClaims returns JWT claims set by Bearer middleware.</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/auth/api-key-opts">APIKeyOpts</a></td><td><span class="kind-pill">type</span></td><td>APIKeyOpts configures API key middleware.</td></tr>
<tr><td><a href="/packages/auth/bearer-options">BearerOptions</a></td><td><span class="kind-pill">type</span></td><td>BearerOptions configures JWT bearer middleware.</td></tr>
</tbody></table>

