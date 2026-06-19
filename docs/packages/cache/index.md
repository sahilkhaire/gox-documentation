---
title: "cache"
package: "cache"
gox-doc-version: "10"
---

<PackageOverview
  name="cache"
  node="lru-cache"
  import-path="github.com/sahilkhaire/gox/cache"
  subtitle="Package cache provides an in-memory LRU cache with optional TTL per entry, similar to node-cache and lru-cache in Node.js. Node equivalent: node-cache, lru-cache"
  :symbol-count=2
  :has-cookbook=true
  migration-link=""
  narrative="lru-cache-style in-memory LRU with TTL-style eviction by capacity."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/cache/cache-new"><div class="featured-name">New</div><div class="featured-summary">new LRUCache</div></a>
<a class="featured-card" href="/packages/cache/cache"><div class="featured-name">Get</div><div class="featured-summary">cache.get</div></a>
<a class="featured-card" href="/packages/cache/cache"><div class="featured-name">Set</div><div class="featured-summary">cache.set</div></a>
</div>

## Common use cases

- Memoize expensive lookups
- Cache parsed config
- Bound memory for hot keys

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>new LRUCache({ max })</code></td><td><a href="/packages/cache/new"><code>cache.New(maxSize)</code></a></td></tr>
<tr><td><code>cache.set(k, v, ttl)</code></td><td><a href="/packages/cache/set"><code>Set(key, value, ttl)</code></a></td></tr>
<tr><td><code>cache.get(k)</code></td><td><a href="/packages/cache/get"><code>Get(key)</code></a></td></tr>
<tr><td><code>cache.del(k)</code></td><td><a href="/packages/cache/delete"><code>Delete(key)</code></a></td></tr>
<tr><td><code>cache.flushAll()</code></td><td><a href="/packages/cache/clear"><code>Clear()</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/cache"

c := cache.New(1000)
```

::: tip Full cookbook
See the [**cache cookbook**](/packages/cache/cookbook) for multi-step recipes and real-world patterns.
:::

## API reference

Browse **2 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/cache/cache">Cache</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Cache is a thread-safe LRU cache with per-key TTL.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/cache/cache-new">Cache.New</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>New creates a cache that evicts least-recently-used entries when size exceeds maxSize.</td></tr>
</tbody></table>

