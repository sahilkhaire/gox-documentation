---
title: "set"
package: "set"
gox-doc-version: "10"
---

<PackageOverview
  name="set"
  node="Set"
  import-path="github.com/sahilkhaire/gox/set"
  subtitle="Package set provides JavaScript Set-like operations on generic comparable sets."
  :symbol-count=5
  :has-cookbook=false
  migration-link=""
  narrative="Set operations with union, intersection, and difference — a typed alternative to JavaScript Set utilities."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/set/set-new"><div class="featured-name">New</div><div class="featured-summary">new Set()</div></a>
<a class="featured-card" href="/packages/set/set-union"><div class="featured-name">Union</div><div class="featured-summary">Set union</div></a>
<a class="featured-card" href="/packages/set/set-intersection"><div class="featured-name">Intersection</div><div class="featured-summary">Set intersection</div></a>
</div>

## Common use cases

- Dedupe tag lists
- Compare permission sets
- Merge unique IDs from multiple sources

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>new Set(arr)</code></td><td><a href="/packages/set/new"><code>set.New(arr...)</code></a></td></tr>
<tr><td><code>set.add(v)</code></td><td><a href="/packages/set/add"><code>s.Add(v)</code></a></td></tr>
<tr><td><code>set.has(v)</code></td><td><a href="/packages/set/has"><code>s.Has(v)</code></a></td></tr>
<tr><td><code>set.delete(v)</code></td><td><a href="/packages/set/delete"><code>s.Delete(v)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/set"

u := set.Union(tagsA, tagsB)
```

## API reference

Browse **5 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/set/set">Set</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Set is a set of comparable values.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/set/set-difference">Set.Difference</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>Difference returns elements in a but not in b.</td></tr>
<tr><td><a href="/packages/set/set-intersection">Set.Intersection</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>Intersection returns elements in both a and b.</td></tr>
<tr><td><a href="/packages/set/set-new">Set.New</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>New creates a set from items.</td></tr>
<tr><td><a href="/packages/set/set-union">Set.Union</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>Union returns elements in a or b.</td></tr>
</tbody></table>

