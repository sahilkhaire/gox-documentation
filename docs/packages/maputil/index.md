---
title: "maputil"
package: "maputil"
gox-doc-version: "10"
---

<PackageOverview
  name="maputil"
  node="lodash object utils"
  import-path="github.com/sahilkhaire/gox/maputil"
  subtitle="Package maputil provides lodash-style object helpers for maps."
  :symbol-count=9
  :has-cookbook=true
  migration-link=""
  narrative="Object utilities from lodash — pick, omit, merge, deep get/set — for config maps and dynamic JSON."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/maputil/pick"><div class="featured-name">Pick</div><div class="featured-summary">_.pick</div></a>
<a class="featured-card" href="/packages/maputil/get"><div class="featured-name">Get</div><div class="featured-summary">_.get path</div></a>
<a class="featured-card" href="/packages/maputil/merge"><div class="featured-name">Merge</div><div class="featured-summary">_.merge</div></a>
</div>

## Common use cases

- Pick allowed keys from request bodies
- Merge defaults with overrides
- Read nested paths like address.city

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>_.pick(obj, keys)</code></td><td><a href="/packages/maputil/pick"><code>maputil.Pick(obj, keys...)</code></a></td></tr>
<tr><td><code>_.omit(obj, keys)</code></td><td><a href="/packages/maputil/omit"><code>maputil.Omit(obj, keys...)</code></a></td></tr>
<tr><td><code>_.merge(a, b)</code></td><td><a href="/packages/maputil/merge"><code>maputil.Merge(a, b)</code></a></td></tr>
<tr><td><code>_.cloneDeep(obj)</code></td><td><a href="/packages/maputil/clone"><code>maputil.Clone(obj)</code></a></td></tr>
<tr><td><code>Object.keys(obj)</code></td><td><a href="/packages/maputil/keys"><code>maputil.Keys(obj)</code></a></td></tr>
<tr><td><code>Object.values(obj)</code></td><td><a href="/packages/maputil/values"><code>maputil.Values(obj)</code></a></td></tr>
<tr><td><code>_.invert(obj)</code></td><td><a href="/packages/maputil/invert"><code>maputil.Invert(obj)</code></a></td></tr>
<tr><td><code>_.get(obj, "a.b")</code></td><td><a href="/packages/maputil/get"><code>maputil.Get(obj, "a.b")</code></a></td></tr>
<tr><td><code>_.set(obj, "a.b", v)</code></td><td><a href="/packages/maputil/set"><code>maputil.Set(obj, "a.b", v)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/maputil"

picked := maputil.Pick(cfg, "host", "port")
```

::: tip Full cookbook
See the [**maputil cookbook**](/packages/maputil/cookbook) for multi-step recipes and real-world patterns.
:::

## API reference

Browse **9 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/maputil/clone">Clone</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Clone returns a shallow copy of m (lodash clone).</td></tr>
<tr><td><a href="/packages/maputil/get">Get</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Get reads a nested value from m using dot-separated path (lodash get).</td></tr>
<tr><td><a href="/packages/maputil/invert">Invert</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Invert swaps keys and values; duplicate values overwrite (lodash invert).</td></tr>
<tr><td><a href="/packages/maputil/keys">Keys</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Keys returns map keys (Object.keys).</td></tr>
<tr><td><a href="/packages/maputil/merge">Merge</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Merge copies keys from sources into dst, later maps override (lodash merge, shallow).</td></tr>
<tr><td><a href="/packages/maputil/omit">Omit</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Omit returns a new map without the given keys (lodash omit).</td></tr>
<tr><td><a href="/packages/maputil/pick">Pick</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Pick returns a new map with only the given keys (lodash pick).</td></tr>
<tr><td><a href="/packages/maputil/set">Set</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Set writes value at dot-separated path, creating intermediate maps (lodash set).</td></tr>
<tr><td><a href="/packages/maputil/values">Values</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Values returns map values (Object.values).</td></tr>
</tbody></table>

