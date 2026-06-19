---
title: "db"
package: "db"
gox-doc-version: "10"
---

<PackageOverview
  name="db"
  node="knex"
  import-path="github.com/sahilkhaire/gox/db"
  subtitle="Package db provides a Knex-like query builder on sqlx. Node equivalent: knex, pg, mysql2"
  :symbol-count=5
  :has-cookbook=true
  migration-link="/migration/knex"
  narrative="Knex-style chainable SQL on sqlx — queries, inserts, transactions, and escape hatches to raw SQL."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/db/db-open"><div class="featured-name">Open</div><div class="featured-summary">knex()</div></a>
<a class="featured-card" href="/packages/db/query"><div class="featured-name">From</div><div class="featured-summary">knex('table')</div></a>
<a class="featured-card" href="/packages/db/tx"><div class="featured-name">Transaction</div><div class="featured-summary">knex.transaction</div></a>
</div>

## Common use cases

- CRUD with WhereEq chains
- Run migrations in transactions
- Access underlying *sqlx.DB* when needed

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>.first()</code></td><td><a href="/packages/db/first"><code>.First(ctx, &row)</code></a></td></tr>
<tr><td><code>.insert()</code></td><td><a href="/packages/db/insert"><code>db.Insert(ctx, table, row)</code></a></td></tr>
<tr><td><code>.update()</code></td><td><a href="/packages/db/update"><code>db.Update(ctx, table, set, where)</code></a></td></tr>
<tr><td><code>.del()</code></td><td><a href="/packages/db/delete"><code>db.Delete(ctx, table, where)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/db"

row, err := db.From("users").WhereEq("id", id).First(ctx, &user)
```

::: tip Full cookbook
See the [**db cookbook**](/packages/db/cookbook) for multi-step recipes and real-world patterns.
:::

::: info Migration guide
Coming from Node.js? Read the [**migration guide**](/migration/knex) for side-by-side patterns.
:::

## API reference

Browse **5 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

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

