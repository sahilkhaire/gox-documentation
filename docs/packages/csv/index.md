---
title: "csv"
package: "csv"
gox-doc-version: "10"
---

<PackageOverview
  name="csv"
  node="papaparse"
  import-path="github.com/sahilkhaire/gox/csv"
  subtitle="Package csv reads and writes CSV with header-row maps. Node equivalent: papaparse, csv-parse."
  :symbol-count=4
  :has-cookbook=false
  migration-link=""
  narrative="papaparse-style CSV read, write, and parse for tabular data import/export."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/csv/read"><div class="featured-name">Read</div><div class="featured-summary">Parse CSV file</div></a>
<a class="featured-card" href="/packages/csv/write"><div class="featured-name">Write</div><div class="featured-summary">Write CSV</div></a>
<a class="featured-card" href="/packages/csv/parse-string"><div class="featured-name">ParseString</div><div class="featured-summary">Parse inline CSV</div></a>
</div>

## Common use cases

- Import user CSV uploads
- Export reports
- Parse inline CSV strings

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/csv"

rows, err := csv.ReadAll(ctx, "data.csv")
```

## API reference

Browse **4 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/csv/parse-string">ParseString</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>ParseString parses CSV text with a header row.</td></tr>
<tr><td><a href="/packages/csv/read">Read</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Read parses CSV from r; the first row is used as column keys.</td></tr>
<tr><td><a href="/packages/csv/read-all">ReadAll</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>ReadAll reads a CSV file at path.</td></tr>
<tr><td><a href="/packages/csv/write">Write</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Write writes rows to w using columns for field order (header row).</td></tr>
</tbody></table>

