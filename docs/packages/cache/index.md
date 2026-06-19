---
title: "cache"
package: "cache"
gox-doc-version: "11"
---

<PackageOverview
  name="cache"
  node="lru-cache"
  import-path="github.com/sahilkhaire/gox/cache"
  subtitle="Package cache provides an in-memory LRU cache with optional TTL per entry, similar to node-cache and lru-cache in Node.js. Node equivalent: node-cache, lru-cache"
  :symbol-count=2
/>
## API reference

Select a symbol below — each page explains what it does, shows Node.js vs Go comparisons, and includes a runnable example.

<SymbolFilter placeholder="Filter symbols…" />

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/cache/cache">Cache</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Cache is a thread-safe LRU cache with per-key TTL.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/cache/cache-new">Cache.New</a></td><td><code class="node-cell">new LRUCache({ max })</code></td><td><span class="kind-pill">method</span></td><td>New creates a cache that evicts least-recently-used entries when size exceeds maxSize.</td></tr>
</tbody></table>

