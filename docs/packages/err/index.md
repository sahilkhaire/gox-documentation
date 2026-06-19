---
title: "err"
package: "err"
gox-doc-version: "10"
---

<PackageOverview
  name="err"
  node="http-errors"
  import-path="github.com/sahilkhaire/gox/err"
  subtitle="Package err provides application errors with HTTP status codes (Express-style)."
  :symbol-count=10
  :has-cookbook=true
  migration-link=""
  narrative="http-errors style typed HTTP errors with status codes — return from handlers for automatic JSON responses."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/err/app-error-not-found"><div class="featured-name">NotFound</div><div class="featured-summary">404 error</div></a>
<a class="featured-card" href="/packages/err/app-error-bad-request"><div class="featured-name">BadRequest</div><div class="featured-summary">400 error</div></a>
<a class="featured-card" href="/packages/err/app-error-unauthorized"><div class="featured-name">Unauthorized</div><div class="featured-summary">401 error</div></a>
</div>

## Common use cases

- Return 404 from REST handlers
- Wrap validation failures as 400
- Map domain errors to status codes

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>createError(400, msg)</code></td><td><a href="/packages/err/bad-request"><code>err.BadRequest(msg)</code></a></td></tr>
<tr><td><code>createError(404, msg)</code></td><td><a href="/packages/err/not-found"><code>err.NotFound(msg)</code></a></td></tr>
<tr><td><code>createError(401, msg)</code></td><td><a href="/packages/err/unauthorized"><code>err.Unauthorized(msg)</code></a></td></tr>
<tr><td><code>createError(403, msg)</code></td><td><a href="/packages/err/forbidden"><code>err.Forbidden(msg)</code></a></td></tr>
<tr><td><code>createError(500, msg)</code></td><td><a href="/packages/err/internal"><code>err.Internal(msg)</code></a></td></tr>
<tr><td><code>err.statusCode</code></td><td><a href="/packages/err/status-code"><code>err.StatusCode(err)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/err"

return err.NotFound("user not found")
```

::: tip Full cookbook
See the [**err cookbook**](/packages/err/cookbook) for multi-step recipes and real-world patterns.
:::

## API reference

Browse **10 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/err/as">As</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>As finds the first error in the chain assignable to target.</td></tr>
<tr><td><a href="/packages/err/is">Is</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Is reports whether err matches target in the error chain.</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/err/app-error">AppError</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>AppError is an error with an HTTP status code.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/err/app-error-bad-request">AppError.BadRequest</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>BadRequest returns a 400 AppError.</td></tr>
<tr><td><a href="/packages/err/app-error-forbidden">AppError.Forbidden</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>Forbidden returns a 403 AppError.</td></tr>
<tr><td><a href="/packages/err/app-error-internal">AppError.Internal</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>Internal returns a 500 AppError.</td></tr>
<tr><td><a href="/packages/err/app-error-new">AppError.New</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>New creates an AppError with code and message.</td></tr>
<tr><td><a href="/packages/err/app-error-not-found">AppError.NotFound</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>NotFound returns a 404 AppError.</td></tr>
<tr><td><a href="/packages/err/app-error-unauthorized">AppError.Unauthorized</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>Unauthorized returns a 401 AppError.</td></tr>
<tr><td><a href="/packages/err/app-error-wrap">AppError.Wrap</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>Wrap wraps cause with code and message.</td></tr>
</tbody></table>

