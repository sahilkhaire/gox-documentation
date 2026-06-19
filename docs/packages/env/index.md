---
title: "env"
package: "env"
gox-doc-version: "7"
---

<div class="package-hero"><div class="api-meta"><span class="api-badge node">Node: dotenv / process.env</span></div><h1>env</h1>
<p class="subtitle">Package env provides process.env and dotenv-style configuration loading.</p><div class="import-line">import "github.com/sahilkhaire/gox/env"</div></div>
## Quick start

A minimal example to get going:

```go
import "github.com/sahilkhaire/gox/env"

port := env.Get("PORT", "8080")
```

::: tip Full cookbook
See the [**env cookbook**](/packages/env/cookbook) for multi-step recipes and real-world patterns.
:::

## What's in this package

If you have used **dotenv / process.env** in Node.js, this package is your starting point. Browse **7 documented symbols** below — each page includes Node.js, standard Go, and gox side-by-side examples.

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/env/get">Get</a></td><td><span class="kind-pill">func</span></td><td>Get returns the value for key (override, then os.Getenv).</td></tr>
<tr><td><a href="/packages/env/get-bool">GetBool</a></td><td><span class="kind-pill">func</span></td><td>GetBool parses a bool environment variable.</td></tr>
<tr><td><a href="/packages/env/get-duration">GetDuration</a></td><td><span class="kind-pill">func</span></td><td>GetDuration parses a duration environment variable.</td></tr>
<tr><td><a href="/packages/env/get-int">GetInt</a></td><td><span class="kind-pill">func</span></td><td>GetInt parses an int environment variable.</td></tr>
<tr><td><a href="/packages/env/load">Load</a></td><td><span class="kind-pill">func</span></td><td>Load parses a dotenv file and sets variables in the environment and override layer.</td></tr>
<tr><td><a href="/packages/env/must-get">MustGet</a></td><td><span class="kind-pill">func</span></td><td>MustGet returns the value or panics if missing.</td></tr>
<tr><td><a href="/packages/env/set">Set</a></td><td><span class="kind-pill">func</span></td><td>Set sets an override value (for tests).</td></tr>
</tbody></table>

