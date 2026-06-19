---
title: "id"
package: "id"
gox-doc-version: "10"
---

<PackageOverview
  name="id"
  node="uuid, nanoid"
  import-path="github.com/sahilkhaire/gox/id"
  subtitle="Package id provides UUID and NanoID-style unique identifiers. Node equivalent: uuid, nanoid."
  :symbol-count=4
  :has-cookbook=false
  migration-link=""
  narrative="uuid and nanoid generation and parsing for identifiers."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/id/new-uuid"><div class="featured-name">NewUUID</div><div class="featured-summary">uuid.v4</div></a>
<a class="featured-card" href="/packages/id/new-nano-id"><div class="featured-name">NewNanoID</div><div class="featured-summary">nanoid</div></a>
<a class="featured-card" href="/packages/id/parse-uuid"><div class="featured-name">ParseUUID</div><div class="featured-summary">Parse UUID string</div></a>
</div>

## Common use cases

- Generate request IDs
- Parse UUID path params
- Short public IDs with nanoid

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>uuid.v4()</code></td><td><a href="/packages/id/new-uuid"><code>id.NewUUID()</code></a></td></tr>
<tr><td><code>uuid.parse(s)</code></td><td><a href="/packages/id/parse-uuid"><code>id.ParseUUID(s)</code></a></td></tr>
<tr><td><code>nanoid(size)</code></td><td><a href="/packages/id/new-nano-id"><code>id.NewNanoID(size)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/id"

uid := id.NewUUID()
```

## API reference

Browse **4 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/id/must-parse-uuid">MustParseUUID</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>MustParseUUID parses s or panics.</td></tr>
<tr><td><a href="/packages/id/new-nano-id">NewNanoID</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>NewNanoID returns a URL-safe ID using the default nanoid alphabet.</td></tr>
<tr><td><a href="/packages/id/new-uuid">NewUUID</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>NewUUID returns a random UUID string (uuid.v4).</td></tr>
<tr><td><a href="/packages/id/parse-uuid">ParseUUID</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>ParseUUID parses a UUID string.</td></tr>
</tbody></table>

