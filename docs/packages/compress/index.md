---
title: "compress"
package: "compress"
gox-doc-version: "11"
---

<PackageOverview
  name="compress"
  node="zlib"
  import-path="github.com/sahilkhaire/gox/compress"
  subtitle="Package compress provides gzip helpers similar to Node.js zlib. Node equivalent: zlib.gzip, zlib.gunzip"
  :symbol-count=4
/>
## API reference

Select a symbol below — each page explains what it does, shows Node.js vs Go comparisons, and includes a runnable example.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/compress/gunzip">Gunzip</a></td><td><code class="node-cell">gunzipSync</code></td><td><span class="kind-pill">func</span></td><td>Gunzip decompresses gzip data.</td></tr>
<tr><td><a href="/packages/compress/gzip">Gzip</a></td><td><code class="node-cell">gzipSync</code></td><td><span class="kind-pill">func</span></td><td>Gzip compresses data with gzip.</td></tr>
<tr><td><a href="/packages/compress/gzip-reader">GzipReader</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>GzipReader wraps a gzip reader over src.</td></tr>
<tr><td><a href="/packages/compress/gzip-writer">GzipWriter</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>GzipWriter wraps dst with a gzip writer; caller must Close the writer.</td></tr>
</tbody></table>

