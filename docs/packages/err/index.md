---
title: "err"
package: "err"
gox-doc-version: "14"
---

<PackageOverview
  name="err"
  node="http-errors"
  import-path="github.com/sahilkhaire/gox/err"
  subtitle="Package err provides application errors with HTTP status codes (Express-style)."
  :symbol-count=10
/>
## API reference

Select a symbol below — each page explains what it does, shows Node.js vs Go comparisons, and includes a runnable example.

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
<tr><td><a href="/packages/err/app-error-bad-request">AppError.BadRequest</a></td><td><code class="node-cell">createError(400, msg)</code></td><td><span class="kind-pill">method</span></td><td>BadRequest returns a 400 AppError.</td></tr>
<tr><td><a href="/packages/err/app-error-forbidden">AppError.Forbidden</a></td><td><code class="node-cell">createError(403, msg)</code></td><td><span class="kind-pill">method</span></td><td>Forbidden returns a 403 AppError.</td></tr>
<tr><td><a href="/packages/err/app-error-internal">AppError.Internal</a></td><td><code class="node-cell">createError(500, msg)</code></td><td><span class="kind-pill">method</span></td><td>Internal returns a 500 AppError.</td></tr>
<tr><td><a href="/packages/err/app-error-new">AppError.New</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>New creates an AppError with code and message.</td></tr>
<tr><td><a href="/packages/err/app-error-not-found">AppError.NotFound</a></td><td><code class="node-cell">createError(404, msg)</code></td><td><span class="kind-pill">method</span></td><td>NotFound returns a 404 AppError.</td></tr>
<tr><td><a href="/packages/err/app-error-unauthorized">AppError.Unauthorized</a></td><td><code class="node-cell">createError(401, msg)</code></td><td><span class="kind-pill">method</span></td><td>Unauthorized returns a 401 AppError.</td></tr>
<tr><td><a href="/packages/err/app-error-wrap">AppError.Wrap</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>Wrap wraps cause with code and message.</td></tr>
</tbody></table>

