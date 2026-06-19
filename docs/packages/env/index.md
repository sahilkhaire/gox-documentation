---
title: "env"
package: "env"
gox-doc-version: "10"
---

<PackageOverview
  name="env"
  node="dotenv / process.env"
  import-path="github.com/sahilkhaire/gox/env"
  subtitle="Package env provides process.env and dotenv-style configuration loading."
  :symbol-count=7
  :has-cookbook=true
  migration-link=""
  narrative="dotenv-style loading plus typed getters for strings, ints, booleans, and durations."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/env/load"><div class="featured-name">Load</div><div class="featured-summary">dotenv.config</div></a>
<a class="featured-card" href="/packages/env/get"><div class="featured-name">Get</div><div class="featured-summary">process.env</div></a>
<a class="featured-card" href="/packages/env/get-int"><div class="featured-name">GetInt</div><div class="featured-summary">Typed int env</div></a>
</div>

## Common use cases

- Load .env in local dev
- Read PORT with defaults
- Parse feature flags as booleans

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>require('dotenv').config()</code></td><td><a href="/packages/env/load"><code>env.Load()</code></a></td></tr>
<tr><td><code>process.env.KEY</code></td><td><a href="/packages/env/get"><code>env.Get("KEY", default)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/env"

port := env.Get("PORT", "8080")
```

::: tip Full cookbook
See the [**env cookbook**](/packages/env/cookbook) for multi-step recipes and real-world patterns.
:::

## API reference

Browse **7 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/env/get">Get</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Get returns the value for key (override, then os.Getenv).</td></tr>
<tr><td><a href="/packages/env/get-bool">GetBool</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>GetBool parses a bool environment variable.</td></tr>
<tr><td><a href="/packages/env/get-duration">GetDuration</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>GetDuration parses a duration environment variable.</td></tr>
<tr><td><a href="/packages/env/get-int">GetInt</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>GetInt parses an int environment variable.</td></tr>
<tr><td><a href="/packages/env/load">Load</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Load parses a dotenv file and sets variables in the environment and override layer.</td></tr>
<tr><td><a href="/packages/env/must-get">MustGet</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>MustGet returns the value or panics if missing.</td></tr>
<tr><td><a href="/packages/env/set">Set</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Set sets an override value (for tests).</td></tr>
</tbody></table>

