---
title: "buffer"
package: "buffer"
gox-doc-version: "10"
---

<PackageOverview
  name="buffer"
  node="Buffer"
  import-path="github.com/sahilkhaire/gox/buffer"
  subtitle="Package buffer provides Node Buffer-style helpers for byte slices."
  :symbol-count=6
  :has-cookbook=false
  migration-link=""
  narrative="Buffer-style byte helpers for concatenation, comparison, and construction from strings or slices."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/buffer/buffer-from"><div class="featured-name">From</div><div class="featured-summary">Buffer.from</div></a>
<a class="featured-card" href="/packages/buffer/buffer-concat"><div class="featured-name">Concat</div><div class="featured-summary">Buffer.concat</div></a>
<a class="featured-card" href="/packages/buffer/equals"><div class="featured-name">Equals</div><div class="featured-summary">Compare bytes</div></a>
</div>

## Common use cases

- Concatenate binary payloads
- Compare constant-time byte slices
- Build buffers from UTF-8 strings

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>Buffer.from(data)</code></td><td><a href="/packages/buffer/from"><code>buffer.From(data)</code></a></td></tr>
<tr><td><code>Buffer.from(str)</code></td><td><a href="/packages/buffer/from-string"><code>buffer.FromString(str)</code></a></td></tr>
<tr><td><code>Buffer.concat(list)</code></td><td><a href="/packages/buffer/concat"><code>buffer.Concat(buffers...)</code></a></td></tr>
<tr><td><code>buf.compare(other)</code></td><td><a href="/packages/buffer/compare"><code>buffer.Compare(a, b)</code></a></td></tr>
<tr><td><code>buf.equals(other)</code></td><td><a href="/packages/buffer/equals"><code>buffer.Equals(a, b)</code></a></td></tr>
<tr><td><code>buf.toString()</code></td><td><a href="/packages/buffer/to-string"><code>buffer.ToString(b)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/buffer"

b := buffer.FromString("hello")
```

## API reference

Browse **6 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/buffer/compare">Compare</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Compare compares two buffers lexicographically (Buffer.compare).</td></tr>
<tr><td><a href="/packages/buffer/equals">Equals</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Equals reports whether a and b are equal (timing-safe enough for content equality).</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/buffer/buffer">Buffer</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Buffer is a byte slice with helper methods.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/buffer/buffer-concat">Buffer.Concat</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>Concat joins buffers (Buffer.concat).</td></tr>
<tr><td><a href="/packages/buffer/buffer-from">Buffer.From</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>From allocates a copy of data (Buffer.from).</td></tr>
<tr><td><a href="/packages/buffer/buffer-from-string">Buffer.FromString</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>FromString creates a buffer from a string.</td></tr>
</tbody></table>

