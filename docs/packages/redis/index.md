---
title: "redis"
package: "redis"
gox-doc-version: "10"
---

<PackageOverview
  name="redis"
  node="ioredis"
  import-path="github.com/sahilkhaire/gox/redis"
  subtitle="Package redis provides an ioredis-like wrapper around go-redis. Node equivalent: ioredis, node-redis"
  :symbol-count=4
  :has-cookbook=true
  migration-link=""
  narrative="ioredis-style Redis client with get/set, pub/sub, and context-aware commands."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/redis/client-new"><div class="featured-name">New</div><div class="featured-summary">new Redis()</div></a>
<a class="featured-card" href="/packages/redis/client"><div class="featured-name">Get</div><div class="featured-summary">redis.get</div></a>
<a class="featured-card" href="/packages/redis/client"><div class="featured-name">Set</div><div class="featured-summary">redis.set</div></a>
</div>

## Common use cases

- Cache hot keys
- Pub/sub event fan-out
- Session or rate-limit storage

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>new Redis()</code></td><td><a href="/packages/redis/new"><code>redis.New(addr)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/redis"

rdb, err := redis.New("localhost:6379")
```

::: tip Full cookbook
See the [**redis cookbook**](/packages/redis/cookbook) for multi-step recipes and real-world patterns.
:::

## API reference

Browse **4 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/redis/client">Client</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Client wraps go-redis Client.</td></tr>
<tr><td><a href="/packages/redis/message">Message</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Message is a pub/sub payload.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/redis/client-new">Client.New</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>New connects to addr (host:port).</td></tr>
<tr><td><a href="/packages/redis/client-new-from-client">Client.NewFromClient</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>NewFromClient wraps an existing client.</td></tr>
</tbody></table>

