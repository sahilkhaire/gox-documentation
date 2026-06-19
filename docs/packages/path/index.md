---
title: "path"
package: "path"
gox-doc-version: "10"
---

<PackageOverview
  name="path"
  node="path"
  import-path="github.com/sahilkhaire/gox/path"
  subtitle="Package path provides Node path module helpers using path/filepath."
  :symbol-count=6
  :has-cookbook=false
  migration-link=""
  narrative="Node path helpers on top of filepath — join, basename, extname, and absolute checks."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/path/join"><div class="featured-name">Join</div><div class="featured-summary">path.join</div></a>
<a class="featured-card" href="/packages/path/basename"><div class="featured-name">Basename</div><div class="featured-summary">path.basename</div></a>
<a class="featured-card" href="/packages/path/extname"><div class="featured-name">Extname</div><div class="featured-summary">path.extname</div></a>
</div>

## Common use cases

- Build cross-platform file paths
- Extract extensions for uploads
- Resolve relative config paths

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>path.join(...)</code></td><td><a href="/packages/path/join"><code>path.Join(...)</code></a></td></tr>
<tr><td><code>path.resolve(...)</code></td><td><a href="/packages/path/resolve"><code>path.Resolve(...)</code></a></td></tr>
<tr><td><code>path.basename(p)</code></td><td><a href="/packages/path/basename"><code>path.Basename(p)</code></a></td></tr>
<tr><td><code>path.extname(p)</code></td><td><a href="/packages/path/extname"><code>path.Extname(p)</code></a></td></tr>
<tr><td><code>path.dirname(p)</code></td><td><a href="/packages/path/dirname"><code>path.Dirname(p)</code></a></td></tr>
<tr><td><code>path.isAbsolute(p)</code></td><td><a href="/packages/path/is-abs"><code>path.IsAbs(p)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/path"

p := path.Join("var", "log", "app.log")
```

## API reference

Browse **6 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/path/basename">Basename</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Basename returns the last element (path.basename).</td></tr>
<tr><td><a href="/packages/path/dirname">Dirname</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Dirname returns the directory (path.dirname).</td></tr>
<tr><td><a href="/packages/path/extname">Extname</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Extname returns the extension including dot (path.extname).</td></tr>
<tr><td><a href="/packages/path/is-abs">IsAbs</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>IsAbs reports whether p is absolute (path.isAbsolute).</td></tr>
<tr><td><a href="/packages/path/join">Join</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Join joins path segments (path.join).</td></tr>
<tr><td><a href="/packages/path/resolve">Resolve</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Resolve resolves to an absolute path (path.resolve).</td></tr>
</tbody></table>

