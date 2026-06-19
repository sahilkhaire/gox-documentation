---
title: "archive"
package: "archive"
gox-doc-version: "10"
---

<PackageOverview
  name="archive"
  node="archiver"
  import-path="github.com/sahilkhaire/gox/archive"
  subtitle="Package archive provides zip and tar.gz helpers. Node equivalent: archiver, adm-zip, tar."
  :symbol-count=6
  :has-cookbook=false
  migration-link=""
  narrative="archiver-style zip and tar create/extract for bundling files."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/archive/zip-create"><div class="featured-name">ZipCreate</div><div class="featured-summary">Create zip</div></a>
<a class="featured-card" href="/packages/archive/zip-extract"><div class="featured-name">ZipExtract</div><div class="featured-summary">Extract zip</div></a>
<a class="featured-card" href="/packages/archive/tar-create"><div class="featured-name">TarCreate</div><div class="featured-summary">Create tar</div></a>
</div>

## Common use cases

- Bundle download artifacts
- Extract uploaded archives
- Stream tar entries

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/archive"

err := archive.ZipCreate("out.zip", entries)
```

## API reference

Browse **6 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/archive/tar-create">TarCreate</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>TarCreate builds a gzip-compressed tar archive.</td></tr>
<tr><td><a href="/packages/archive/tar-extract">TarExtract</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>TarExtract extracts a gzip-compressed tar archive into destDir.</td></tr>
<tr><td><a href="/packages/archive/zip-create">ZipCreate</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>ZipCreate builds a zip archive from a map of path -&gt; contents.</td></tr>
<tr><td><a href="/packages/archive/zip-create-entries">ZipCreateEntries</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>ZipCreateEntries builds a zip archive from file entries.</td></tr>
<tr><td><a href="/packages/archive/zip-extract">ZipExtract</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>ZipExtract writes zip data into destDir.</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/archive/file-entry">FileEntry</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>FileEntry is a named file for archive creation.</td></tr>
</tbody></table>

