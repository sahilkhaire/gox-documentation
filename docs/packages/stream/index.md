---
title: "stream"
package: "stream"
gox-doc-version: "10"
---

<PackageOverview
  name="stream"
  node="Node stream"
  import-path="github.com/sahilkhaire/gox/stream"
  subtitle="Package stream provides io helpers similar to Node.js streams. Node equivalent: stream, fs.createReadStream pipe"
  :symbol-count=4
  :has-cookbook=false
  migration-link=""
  narrative="Node stream patterns — pipe, tee, transform, and read-all for io.Reader chains."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/stream/pipe"><div class="featured-name">Pipe</div><div class="featured-summary">stream.pipe</div></a>
<a class="featured-card" href="/packages/stream/read-all"><div class="featured-name">ReadAll</div><div class="featured-summary">stream consumers</div></a>
<a class="featured-card" href="/packages/stream/transform"><div class="featured-name">Transform</div><div class="featured-summary">Transform stream</div></a>
</div>

## Common use cases

- Pipe HTTP bodies to storage
- Duplicate readers for checksum + upload
- Drain streams to bytes

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>src.pipe(dst)</code></td><td><a href="/packages/stream/pipe"><code>stream.Pipe(src, dst)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/stream"

n, err := stream.ReadAll(reader)
```

## API reference

Browse **4 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/stream/pipe">Pipe</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Pipe copies from src to dst until EOF or error.</td></tr>
<tr><td><a href="/packages/stream/read-all">ReadAll</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>ReadAll reads r until EOF or ctx is cancelled.</td></tr>
<tr><td><a href="/packages/stream/tee-reader">TeeReader</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>TeeReader returns a Reader that writes to w what it reads from r.</td></tr>
<tr><td><a href="/packages/stream/transform">Transform</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Transform returns a Reader that applies fn to each chunk from reader.</td></tr>
</tbody></table>

