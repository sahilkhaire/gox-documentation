---
title: "mongo"
package: "mongo"
gox-doc-version: "10"
---

<PackageOverview
  name="mongo"
  node="mongoose"
  import-path="github.com/sahilkhaire/gox/mongo"
  subtitle="Package mongo provides a thin mongoose-like layer on the official driver. Node equivalent: mongoose, mongodb"
  :symbol-count=8
  :has-cookbook=true
  migration-link=""
  narrative="mongoose-style MongoDB client, collections, and filter builders."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/mongo/client-connect"><div class="featured-name">Connect</div><div class="featured-summary">mongoose.connect</div></a>
<a class="featured-card" href="/packages/mongo/collection"><div class="featured-name">Collection</div><div class="featured-summary">Model collection</div></a>
<a class="featured-card" href="/packages/mongo/eq"><div class="featured-name">Eq</div><div class="featured-summary">Filter helper</div></a>
</div>

## Common use cases

- Connect with URI
- Find documents by field
- Use Eq/In/Gt filter helpers

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>mongoose.connect(uri)</code></td><td><a href="/packages/mongo/connect"><code>mongo.Connect(ctx, uri)</code></a></td></tr>
<tr><td><code>db.collection('users')</code></td><td><a href="/packages/mongo/db"><code>client.DB(name).Collection(name)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/mongo"

client, err := mongo.Connect(ctx, uri)
```

::: tip Full cookbook
See the [**mongo cookbook**](/packages/mongo/cookbook) for multi-step recipes and real-world patterns.
:::

## API reference

Browse **8 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/mongo/eq">Eq</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Eq matches equality on field.</td></tr>
<tr><td><a href="/packages/mongo/gt">Gt</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Gt matches field &gt; value.</td></tr>
<tr><td><a href="/packages/mongo/in">In</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>In matches field in values.</td></tr>
<tr><td><a href="/packages/mongo/set">Set</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Set returns bson.M for $set updates.</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/mongo/client">Client</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Client wraps mongo.Client.</td></tr>
<tr><td><a href="/packages/mongo/collection">Collection</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Collection wraps mongo.Collection with helper methods.</td></tr>
<tr><td><a href="/packages/mongo/database">Database</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Database wraps mongo.Database.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/mongo/client-connect">Client.Connect</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>Connect dials uri.</td></tr>
</tbody></table>

