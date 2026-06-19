---
title: "slice"
package: "slice"
gox-doc-version: "10"
---

<PackageOverview
  name="slice"
  node="lodash / Array.*"
  import-path="github.com/sahilkhaire/gox/slice"
  subtitle="Package slice provides lodash/Array-style collection helpers for slices."
  :symbol-count=13
  :has-cookbook=true
  migration-link=""
  narrative="lodash and Array methods as generic Go functions — map, filter, reduce, groupBy, and more with compile-time type safety."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/slice/map"><div class="featured-name">Map</div><div class="featured-summary">Array.map</div></a>
<a class="featured-card" href="/packages/slice/filter"><div class="featured-name">Filter</div><div class="featured-summary">Array.filter</div></a>
<a class="featured-card" href="/packages/slice/group-by"><div class="featured-name">GroupBy</div><div class="featured-summary">lodash groupBy</div></a>
</div>

## Common use cases

- Transform API response arrays
- Filter collections before mapping
- Group records by key for dashboards

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>arr.map(fn)</code></td><td><a href="/packages/slice/map"><code>slice.Map(arr, fn)</code></a></td></tr>
<tr><td><code>arr.filter(fn)</code></td><td><a href="/packages/slice/filter"><code>slice.Filter(arr, fn)</code></a></td></tr>
<tr><td><code>arr.reduce(fn, init)</code></td><td><a href="/packages/slice/reduce"><code>slice.Reduce(arr, init, fn)</code></a></td></tr>
<tr><td><code>arr.find(fn)</code></td><td><a href="/packages/slice/find"><code>slice.Find(arr, fn)</code></a></td></tr>
<tr><td><code>arr.findIndex(fn)</code></td><td><a href="/packages/slice/find-index"><code>slice.FindIndex(arr, fn)</code></a></td></tr>
<tr><td><code>_.groupBy(arr, fn)</code></td><td><a href="/packages/slice/group-by"><code>slice.GroupBy(arr, fn)</code></a></td></tr>
<tr><td><code>_.uniq(arr)</code></td><td><a href="/packages/slice/uniq"><code>slice.Uniq(arr)</code></a></td></tr>
<tr><td><code>_.chunk(arr, n)</code></td><td><a href="/packages/slice/chunk"><code>slice.Chunk(arr, n)</code></a></td></tr>
<tr><td><code>_.flatten(arr)</code></td><td><a href="/packages/slice/flatten"><code>slice.Flatten(arr)</code></a></td></tr>
<tr><td><code>_.sortBy(arr, fn)</code></td><td><a href="/packages/slice/sort-by"><code>slice.SortBy(arr, fn)</code></a></td></tr>
<tr><td><code>arr.includes(x)</code></td><td><a href="/packages/slice/contains"><code>slice.Contains(arr, x)</code></a></td></tr>
<tr><td><code>arr.every(fn)</code></td><td><a href="/packages/slice/every"><code>slice.Every(arr, fn)</code></a></td></tr>
</tbody></table>

::: info
Showing 12 of 13 mappings. Browse the symbol table below for the full API.
:::

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/slice"

evens := slice.Filter(nums, func(n int) bool { return n%2 == 0 })
```

::: tip Full cookbook
See the [**slice cookbook**](/packages/slice/cookbook) for multi-step recipes and real-world patterns.
:::

## API reference

Browse **13 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/slice/chunk">Chunk</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Chunk splits into sub-slices of size (lodash chunk).</td></tr>
<tr><td><a href="/packages/slice/contains">Contains</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Contains reports whether v is in the slice (Array.includes).</td></tr>
<tr><td><a href="/packages/slice/every">Every</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Every reports whether fn is true for all elements (Array.every).</td></tr>
<tr><td><a href="/packages/slice/filter">Filter</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Filter keeps elements where fn returns true (Array.filter).</td></tr>
<tr><td><a href="/packages/slice/find">Find</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Find returns the first element matching fn, or zero and false (Array.find).</td></tr>
<tr><td><a href="/packages/slice/find-index">FindIndex</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>FindIndex returns the index of the first match, or -1 (Array.findIndex).</td></tr>
<tr><td><a href="/packages/slice/flatten">Flatten</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Flatten flattens one level of nested slices (lodash flatten).</td></tr>
<tr><td><a href="/packages/slice/group-by">GroupBy</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>GroupBy groups elements by key from fn (lodash groupBy).</td></tr>
<tr><td><a href="/packages/slice/map">Map</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Map applies fn to each element (Array.map).</td></tr>
<tr><td><a href="/packages/slice/reduce">Reduce</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Reduce folds the slice (Array.reduce).</td></tr>
<tr><td><a href="/packages/slice/some">Some</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Some reports whether fn is true for any element (Array.some).</td></tr>
<tr><td><a href="/packages/slice/sort-by">SortBy</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>SortBy returns a sorted copy ordered by fn's key (lodash sortBy).</td></tr>
<tr><td><a href="/packages/slice/uniq">Uniq</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Uniq returns unique elements in first-seen order (lodash uniq).</td></tr>
</tbody></table>

