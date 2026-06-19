---
title: "json"
package: "json"
gox-doc-version: "10"
---

<PackageOverview
  name="json"
  node="JSON.parse/stringify"
  import-path="github.com/sahilkhaire/gox/json"
  subtitle="Package json provides JSON parse/stringify helpers. Node equivalent: JSON.parse, JSON.stringify"
  :symbol-count=7
  :has-cookbook=true
  migration-link=""
  narrative="JSON.parse/stringify with generics, pretty printing, and file helpers."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/json/parse"><div class="featured-name">Parse</div><div class="featured-summary">JSON.parse</div></a>
<a class="featured-card" href="/packages/json/stringify"><div class="featured-name">Stringify</div><div class="featured-summary">JSON.stringify</div></a>
<a class="featured-card" href="/packages/json/pretty"><div class="featured-name">Pretty</div><div class="featured-summary">Pretty print</div></a>
</div>

## Common use cases

- Parse API payloads into structs
- Pretty-print debug output
- Read/write JSON config files

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>JSON.parse(str)</code></td><td><a href="/packages/json/parse"><code>json.Parse([]byte, &v)</code></a></td></tr>
<tr><td><code>JSON.stringify(obj)</code></td><td><a href="/packages/json/stringify"><code>json.Stringify(v)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/json"

obj, err := json.Parse[Config](jsonStr)
```

::: tip Full cookbook
See the [**json cookbook**](/packages/json/cookbook) for multi-step recipes and real-world patterns.
:::

## API reference

Browse **7 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/json/must-parse">MustParse</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>MustParse unmarshals or panics.</td></tr>
<tr><td><a href="/packages/json/must-stringify">MustStringify</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>MustStringify marshals or panics.</td></tr>
<tr><td><a href="/packages/json/parse">Parse</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Parse unmarshals data into v (JSON.parse).</td></tr>
<tr><td><a href="/packages/json/parse-file">ParseFile</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>ParseFile reads path and unmarshals into v.</td></tr>
<tr><td><a href="/packages/json/pretty">Pretty</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Pretty returns indented JSON.</td></tr>
<tr><td><a href="/packages/json/stringify">Stringify</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Stringify marshals v to a string (JSON.stringify).</td></tr>
<tr><td><a href="/packages/json/write-file">WriteFile</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>WriteFile marshals v and writes to path (helper for examples).</td></tr>
</tbody></table>

