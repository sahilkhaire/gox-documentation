---
title: "slice"
package: "slice"
gox-doc-version: "7"
---

<div class="package-hero"><div class="api-meta"><span class="api-badge node">Node: lodash / Array.*</span></div><h1>slice</h1>
<p class="subtitle">Package slice provides lodash/Array-style collection helpers for slices.</p><div class="import-line">import "github.com/sahilkhaire/gox/slice"</div></div>
## Quick start

A minimal example to get going:

```go
import "github.com/sahilkhaire/gox/slice"

evens := slice.Filter(nums, func(n int) bool { return n%2 == 0 })
```

::: tip Full cookbook
See the [**slice cookbook**](/packages/slice/cookbook) for multi-step recipes and real-world patterns.
:::

## What's in this package

If you have used **lodash / Array.*** in Node.js, this package is your starting point. Browse **13 documented symbols** below — each page includes Node.js, standard Go, and gox side-by-side examples.

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/slice/chunk">Chunk</a></td><td><span class="kind-pill">func</span></td><td>Chunk splits into sub-slices of size (lodash chunk).</td></tr>
<tr><td><a href="/packages/slice/contains">Contains</a></td><td><span class="kind-pill">func</span></td><td>Contains reports whether v is in the slice (Array.includes).</td></tr>
<tr><td><a href="/packages/slice/every">Every</a></td><td><span class="kind-pill">func</span></td><td>Every reports whether fn is true for all elements (Array.every).</td></tr>
<tr><td><a href="/packages/slice/filter">Filter</a></td><td><span class="kind-pill">func</span></td><td>Filter keeps elements where fn returns true (Array.filter).</td></tr>
<tr><td><a href="/packages/slice/find">Find</a></td><td><span class="kind-pill">func</span></td><td>Find returns the first element matching fn, or zero and false (Array.find).</td></tr>
<tr><td><a href="/packages/slice/find-index">FindIndex</a></td><td><span class="kind-pill">func</span></td><td>FindIndex returns the index of the first match, or -1 (Array.findIndex).</td></tr>
<tr><td><a href="/packages/slice/flatten">Flatten</a></td><td><span class="kind-pill">func</span></td><td>Flatten flattens one level of nested slices (lodash flatten).</td></tr>
<tr><td><a href="/packages/slice/group-by">GroupBy</a></td><td><span class="kind-pill">func</span></td><td>GroupBy groups elements by key from fn (lodash groupBy).</td></tr>
<tr><td><a href="/packages/slice/map">Map</a></td><td><span class="kind-pill">func</span></td><td>Map applies fn to each element (Array.map).</td></tr>
<tr><td><a href="/packages/slice/reduce">Reduce</a></td><td><span class="kind-pill">func</span></td><td>Reduce folds the slice (Array.reduce).</td></tr>
<tr><td><a href="/packages/slice/some">Some</a></td><td><span class="kind-pill">func</span></td><td>Some reports whether fn is true for any element (Array.some).</td></tr>
<tr><td><a href="/packages/slice/sort-by">SortBy</a></td><td><span class="kind-pill">func</span></td><td>SortBy returns a sorted copy ordered by fn's key (lodash sortBy).</td></tr>
<tr><td><a href="/packages/slice/uniq">Uniq</a></td><td><span class="kind-pill">func</span></td><td>Uniq returns unique elements in first-seen order (lodash uniq).</td></tr>
</tbody></table>

