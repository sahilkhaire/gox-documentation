---
title: "async"
package: "async"
gox-doc-version: "10"
---

<PackageOverview
  name="async"
  node="Promise.all, timers"
  import-path="github.com/sahilkhaire/gox/async"
  subtitle="Package async provides Promise-like concurrency helpers with context cancellation."
  :symbol-count=7
  :has-cookbook=true
  migration-link=""
  narrative="Promise.all, race, sleep, retry, and timeout patterns with context cancellation built in."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/async/all"><div class="featured-name">All</div><div class="featured-summary">Promise.all</div></a>
<a class="featured-card" href="/packages/async/sleep"><div class="featured-name">Sleep</div><div class="featured-summary">setTimeout</div></a>
<a class="featured-card" href="/packages/async/retry"><div class="featured-name">Retry</div><div class="featured-summary">Retry with backoff</div></a>
</div>

## Common use cases

- Fetch multiple resources in parallel
- Implement deadlines on outbound calls
- Retry flaky operations with backoff

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>Promise.all([a(), b()])</code></td><td><a href="/packages/async/all"><code>async.All(ctx, a, b)</code></a></td></tr>
<tr><td><code>Promise.race([a(), b()])</code></td><td><a href="/packages/async/race"><code>async.Race(ctx, a, b)</code></a></td></tr>
<tr><td><code>await sleep(ms)</code></td><td><a href="/packages/async/sleep"><code>async.Sleep(ctx, duration)</code></a></td></tr>
<tr><td><code>setTimeout(fn, d)</code></td><td><a href="/packages/async/after"><code>async.After(ctx, d, fn)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/async"

a, b, err := async.All(ctx, fetchA, fetchB)
```

::: tip Full cookbook
See the [**async cookbook**](/packages/async/cookbook) for multi-step recipes and real-world patterns.
:::

## API reference

Browse **7 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/async/after">After</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>After sends the time on the returned channel when duration elapses.</td></tr>
<tr><td><a href="/packages/async/all">All</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>All runs tasks concurrently and returns when all complete or ctx is cancelled.</td></tr>
<tr><td><a href="/packages/async/race">Race</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Race runs tasks concurrently and returns the first successful result or error.</td></tr>
<tr><td><a href="/packages/async/retry">Retry</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Retry calls fn until it succeeds or attempts are exhausted.</td></tr>
<tr><td><a href="/packages/async/sleep">Sleep</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Sleep waits until duration elapses or ctx is cancelled.</td></tr>
<tr><td><a href="/packages/async/timeout">Timeout</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Timeout runs fn with a deadline derived from ctx and timeout.</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/async/retry-config">RetryConfig</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>RetryConfig configures Retry backoff.</td></tr>
</tbody></table>

