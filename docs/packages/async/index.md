---
title: "async"
package: "async"
gox-doc-version: "11"
---

<PackageOverview
  name="async"
  node="Promise.all, timers"
  import-path="github.com/sahilkhaire/gox/async"
  subtitle="Package async provides Promise-like concurrency helpers with context cancellation."
  :symbol-count=7
/>
## API reference

Select a symbol below — each page explains what it does, shows Node.js vs Go comparisons, and includes a runnable example.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/async/after">After</a></td><td><code class="node-cell">setTimeout(fn, d)</code></td><td><span class="kind-pill">func</span></td><td>After sends the time on the returned channel when duration elapses.</td></tr>
<tr><td><a href="/packages/async/all">All</a></td><td><code class="node-cell">Promise.all([a(), b()])</code></td><td><span class="kind-pill">func</span></td><td>All runs tasks concurrently and returns when all complete or ctx is cancelled.</td></tr>
<tr><td><a href="/packages/async/race">Race</a></td><td><code class="node-cell">Promise.race([a(), b()])</code></td><td><span class="kind-pill">func</span></td><td>Race runs tasks concurrently and returns the first successful result or error.</td></tr>
<tr><td><a href="/packages/async/retry">Retry</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Retry calls fn until it succeeds or attempts are exhausted.</td></tr>
<tr><td><a href="/packages/async/sleep">Sleep</a></td><td><code class="node-cell">await sleep(ms)</code></td><td><span class="kind-pill">func</span></td><td>Sleep waits until duration elapses or ctx is cancelled.</td></tr>
<tr><td><a href="/packages/async/timeout">Timeout</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Timeout runs fn with a deadline derived from ctx and timeout.</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/async/retry-config">RetryConfig</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>RetryConfig configures Retry backoff.</td></tr>
</tbody></table>

