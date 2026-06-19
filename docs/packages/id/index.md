---
title: "id"
package: "id"
gox-doc-version: "14"
---

<PackageOverview
  name="id"
  node="uuid, nanoid"
  import-path="github.com/sahilkhaire/gox/id"
  subtitle="Package id provides UUID and NanoID-style unique identifiers. Node equivalent: uuid, nanoid."
  :symbol-count=4
/>
## API reference

Select a symbol below — each page explains what it does, shows Node.js vs Go comparisons, and includes a runnable example.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/id/must-parse-uuid">MustParseUUID</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>MustParseUUID parses s or panics.</td></tr>
<tr><td><a href="/packages/id/new-nano-id">NewNanoID</a></td><td><code class="node-cell">nanoid(size)</code></td><td><span class="kind-pill">func</span></td><td>NewNanoID returns a URL-safe ID using the default nanoid alphabet.</td></tr>
<tr><td><a href="/packages/id/new-uuid">NewUUID</a></td><td><code class="node-cell">uuid.v4()</code></td><td><span class="kind-pill">func</span></td><td>NewUUID returns a random UUID string (uuid.v4).</td></tr>
<tr><td><a href="/packages/id/parse-uuid">ParseUUID</a></td><td><code class="node-cell">uuid.parse(s)</code></td><td><span class="kind-pill">func</span></td><td>ParseUUID parses a UUID string.</td></tr>
</tbody></table>

