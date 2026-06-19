---
title: "compress"
package: "compress"
gox-doc-version: "10"
---

<PackageOverview
  name="compress"
  node="zlib"
  import-path="github.com/sahilkhaire/gox/compress"
  subtitle="Package compress provides gzip helpers similar to Node.js zlib. Node equivalent: zlib.gzip, zlib.gunzip"
  :symbol-count=4
  :has-cookbook=false
  migration-link=""
  narrative="zlib-style gzip and gunzip helpers for compressing HTTP payloads and archives."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/compress/gzip"><div class="featured-name">Gzip</div><div class="featured-summary">zlib.gzip</div></a>
<a class="featured-card" href="/packages/compress/gunzip"><div class="featured-name">Gunzip</div><div class="featured-summary">zlib.gunzip</div></a>
</div>

## Common use cases

- Compress API responses
- Decompress uploaded gzip files
- Shrink log batches

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>gzipSync</code></td><td><a href="/packages/compress/gzip"><code>compress.Gzip(data)</code></a></td></tr>
<tr><td><code>gunzipSync</code></td><td><a href="/packages/compress/gunzip"><code>compress.Gunzip(data)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/compress"

gz, err := compress.Gzip(data)
```

## API reference

Browse **4 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/compress/gunzip">Gunzip</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Gunzip decompresses gzip data.</td></tr>
<tr><td><a href="/packages/compress/gzip">Gzip</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Gzip compresses data with gzip.</td></tr>
<tr><td><a href="/packages/compress/gzip-reader">GzipReader</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>GzipReader wraps a gzip reader over src.</td></tr>
<tr><td><a href="/packages/compress/gzip-writer">GzipWriter</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>GzipWriter wraps dst with a gzip writer; caller must Close the writer.</td></tr>
</tbody></table>

