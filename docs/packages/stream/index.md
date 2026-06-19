---
title: "stream"
package: "stream"
gox-doc-version: "11"
---

<PackageOverview
  name="stream"
  node="Node stream"
  import-path="github.com/sahilkhaire/gox/stream"
  subtitle="Package stream provides io helpers similar to Node.js streams. Node equivalent: stream, fs.createReadStream pipe"
  :symbol-count=4
/>
## API reference

Select a symbol below — each page explains what it does, shows Node.js vs Go comparisons, and includes a runnable example.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/stream/pipe">Pipe</a></td><td><code class="node-cell">src.pipe(dst)</code></td><td><span class="kind-pill">func</span></td><td>Pipe copies from src to dst until EOF or error.</td></tr>
<tr><td><a href="/packages/stream/read-all">ReadAll</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>ReadAll reads r until EOF or ctx is cancelled.</td></tr>
<tr><td><a href="/packages/stream/tee-reader">TeeReader</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>TeeReader returns a Reader that writes to w what it reads from r.</td></tr>
<tr><td><a href="/packages/stream/transform">Transform</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Transform returns a Reader that applies fn to each chunk from reader.</td></tr>
</tbody></table>

