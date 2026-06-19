---
title: "fs"
package: "fs"
gox-doc-version: "10"
---

<PackageOverview
  name="fs"
  node="fs.promises"
  import-path="github.com/sahilkhaire/gox/fs"
  subtitle="Package fs provides fs.promises-style file operations with context.Context."
  :symbol-count=8
  :has-cookbook=true
  migration-link=""
  narrative="fs.promises-style file helpers with context-first I/O — read, write, copy, mkdir, and stat."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/fs/read-file"><div class="featured-name">ReadFile</div><div class="featured-summary">fs.readFile</div></a>
<a class="featured-card" href="/packages/fs/write-file"><div class="featured-name">WriteFile</div><div class="featured-summary">fs.writeFile</div></a>
<a class="featured-card" href="/packages/fs/exists"><div class="featured-name">Exists</div><div class="featured-summary">fs.access</div></a>
</div>

## Common use cases

- Read config files at startup
- Write uploaded content to disk
- Check existence before processing

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>fs.readFile(path)</code></td><td><a href="/packages/fs/read-file"><code>fs.ReadFile(ctx, path)</code></a></td></tr>
<tr><td><code>fs.writeFile(path, data)</code></td><td><a href="/packages/fs/write-file"><code>fs.WriteFile(ctx, path, data)</code></a></td></tr>
<tr><td><code>fs.mkdir(path)</code></td><td><a href="/packages/fs/mkdir"><code>fs.Mkdir(ctx, path)</code></a></td></tr>
<tr><td><code>fs.rm(path)</code></td><td><a href="/packages/fs/remove"><code>fs.Remove(ctx, path)</code></a></td></tr>
<tr><td><code>fs.exists(path)</code></td><td><a href="/packages/fs/exists"><code>fs.Exists(ctx, path)</code></a></td></tr>
<tr><td><code>fs.copyFile(src, dst)</code></td><td><a href="/packages/fs/copy"><code>fs.Copy(ctx, src, dst)</code></a></td></tr>
<tr><td><code>fs.stat(path)</code></td><td><a href="/packages/fs/stat"><code>fs.Stat(ctx, path)</code></a></td></tr>
<tr><td><code>fs.readdir(path)</code></td><td><a href="/packages/fs/read-dir"><code>fs.ReadDir(ctx, path)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/fs"

data, err := fs.ReadFile(ctx, "config.json")
```

::: tip Full cookbook
See the [**fs cookbook**](/packages/fs/cookbook) for multi-step recipes and real-world patterns.
:::

## API reference

Browse **8 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/fs/copy">Copy</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Copy copies src to dst (fs.promises.copyFile).</td></tr>
<tr><td><a href="/packages/fs/exists">Exists</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Exists reports whether path exists.</td></tr>
<tr><td><a href="/packages/fs/mkdir">Mkdir</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Mkdir creates a directory (fs.promises.mkdir).</td></tr>
<tr><td><a href="/packages/fs/read-dir">ReadDir</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>ReadDir reads directory entries (fs.promises.readdir).</td></tr>
<tr><td><a href="/packages/fs/read-file">ReadFile</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>ReadFile reads path (fs.promises.readFile).</td></tr>
<tr><td><a href="/packages/fs/remove">Remove</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Remove deletes path (fs.promises.rm).</td></tr>
<tr><td><a href="/packages/fs/stat">Stat</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Stat returns file info (fs.promises.stat).</td></tr>
<tr><td><a href="/packages/fs/write-file">WriteFile</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>WriteFile writes data to path (fs.promises.writeFile).</td></tr>
</tbody></table>

