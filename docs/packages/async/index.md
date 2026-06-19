---
title: "async"
package: "async"
gox-doc-version: "7"
---

<div class="package-hero"><div class="api-meta"><span class="api-badge node">Node: Promise.all, timers</span></div><h1>async</h1>
<p class="subtitle">Package async provides Promise-like concurrency helpers with context cancellation.</p><div class="import-line">import "github.com/sahilkhaire/gox/async"</div></div>
::: tip Full cookbook
See the [**async cookbook**](/packages/async/cookbook) for multi-step recipes and real-world patterns.
:::

## What's in this package

If you have used **Promise.all, timers** in Node.js, this package is your starting point. Browse **7 documented symbols** below — each page includes Node.js, standard Go, and gox side-by-side examples.

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/async/after">After</a></td><td><span class="kind-pill">func</span></td><td>After sends the time on the returned channel when duration elapses.</td></tr>
<tr><td><a href="/packages/async/all">All</a></td><td><span class="kind-pill">func</span></td><td>All runs tasks concurrently and returns when all complete or ctx is cancelled.</td></tr>
<tr><td><a href="/packages/async/race">Race</a></td><td><span class="kind-pill">func</span></td><td>Race runs tasks concurrently and returns the first successful result or error.</td></tr>
<tr><td><a href="/packages/async/retry">Retry</a></td><td><span class="kind-pill">func</span></td><td>Retry calls fn until it succeeds or attempts are exhausted.</td></tr>
<tr><td><a href="/packages/async/sleep">Sleep</a></td><td><span class="kind-pill">func</span></td><td>Sleep waits until duration elapses or ctx is cancelled.</td></tr>
<tr><td><a href="/packages/async/timeout">Timeout</a></td><td><span class="kind-pill">func</span></td><td>Timeout runs fn with a deadline derived from ctx and timeout.</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/async/retry-config">RetryConfig</a></td><td><span class="kind-pill">type</span></td><td>RetryConfig configures Retry backoff.</td></tr>
</tbody></table>

