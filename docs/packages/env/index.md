---
title: "env"
package: "env"
gox-doc-version: "11"
---

<PackageOverview
  name="env"
  node="dotenv / process.env"
  import-path="github.com/sahilkhaire/gox/env"
  subtitle="Package env provides process.env and dotenv-style configuration loading."
  :symbol-count=7
/>
## API reference

Select a symbol below — each page explains what it does, shows Node.js vs Go comparisons, and includes a runnable example.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/env/get">Get</a></td><td><code class="node-cell">process.env.KEY</code></td><td><span class="kind-pill">func</span></td><td>Get returns the value for key (override, then os.Getenv).</td></tr>
<tr><td><a href="/packages/env/get-bool">GetBool</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>GetBool parses a bool environment variable.</td></tr>
<tr><td><a href="/packages/env/get-duration">GetDuration</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>GetDuration parses a duration environment variable.</td></tr>
<tr><td><a href="/packages/env/get-int">GetInt</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>GetInt parses an int environment variable.</td></tr>
<tr><td><a href="/packages/env/load">Load</a></td><td><code class="node-cell">require('dotenv').config()</code></td><td><span class="kind-pill">func</span></td><td>Load parses a dotenv file and sets variables in the environment and override layer.</td></tr>
<tr><td><a href="/packages/env/must-get">MustGet</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>MustGet returns the value or panics if missing.</td></tr>
<tr><td><a href="/packages/env/set">Set</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Set sets an override value (for tests).</td></tr>
</tbody></table>

