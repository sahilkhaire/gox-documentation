---
title: "db"
package: "db"
gox-doc-version: "14"
---

<PackageOverview
  name="db"
  node="knex"
  import-path="github.com/sahilkhaire/gox/db"
  subtitle="Package db provides a Knex-like query builder on sqlx. Node equivalent: knex, pg, mysql2"
  :symbol-count=5
/>
## API reference

Select a symbol below — each page explains what it does, shows Node.js vs Go comparisons, and includes a runnable example.

<SymbolFilter placeholder="Filter symbols…" />

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/db/db">DB</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>DB wraps *sqlx.DB with a chainable query API.</td></tr>
<tr><td><a href="/packages/db/query">Query</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Query is a chainable SELECT builder.</td></tr>
<tr><td><a href="/packages/db/tx">Tx</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Tx wraps a sqlx transaction.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/db/db-must-open">DB.MustOpen</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>MustOpen connects or panics.</td></tr>
<tr><td><a href="/packages/db/db-open">DB.Open</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>Open connects using driver name and DSN.</td></tr>
</tbody></table>

