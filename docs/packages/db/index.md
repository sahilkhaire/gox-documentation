---
title: "db"
package: "db"
gox-doc-version: "7"
---

<div class="package-hero"><div class="api-meta"><span class="api-badge node">Node: knex</span></div><h1>db</h1>
<p class="subtitle">Package db provides a Knex-like query builder on sqlx. Node equivalent: knex, pg, mysql2</p><div class="import-line">import "github.com/sahilkhaire/gox/db"</div></div>
## Quick start

A minimal example to get going:

```go
import "github.com/sahilkhaire/gox/db"

row, err := db.From("users").WhereEq("id", id).First(ctx, &user)
```

::: tip Full cookbook
See the [**db cookbook**](/packages/db/cookbook) for multi-step recipes and real-world patterns.
:::

## What's in this package

If you have used **knex** in Node.js, this package is your starting point. Browse **5 documented symbols** below — each page includes Node.js, standard Go, and gox side-by-side examples.

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/db/db">DB</a></td><td><span class="kind-pill">type</span></td><td>DB wraps *sqlx.DB with a chainable query API.</td></tr>
<tr><td><a href="/packages/db/query">Query</a></td><td><span class="kind-pill">type</span></td><td>Query is a chainable SELECT builder.</td></tr>
<tr><td><a href="/packages/db/tx">Tx</a></td><td><span class="kind-pill">type</span></td><td>Tx wraps a sqlx transaction.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/db/db-must-open">DB.MustOpen</a></td><td><span class="kind-pill">method</span></td><td>MustOpen connects or panics.</td></tr>
<tr><td><a href="/packages/db/db-open">DB.Open</a></td><td><span class="kind-pill">method</span></td><td>Open connects using driver name and DSN.</td></tr>
</tbody></table>

