---
title: "log"
package: "log"
gox-doc-version: "10"
---

<PackageOverview
  name="log"
  node="winston, pino"
  import-path="github.com/sahilkhaire/gox/log"
  subtitle="Package log provides winston/pino-style logging on top of log/slog. Node equivalent: winston, pino."
  :symbol-count=5
  :has-cookbook=true
  migration-link=""
  narrative="winston/pino-style structured logging with slog underneath."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/log/logger-new"><div class="featured-name">New</div><div class="featured-summary">winston.createLogger</div></a>
<a class="featured-card" href="/packages/log/logger-default"><div class="featured-name">Default</div><div class="featured-summary">Default logger</div></a>
</div>

## Common use cases

- Structured request logging
- Set default logger for packages
- Level-based log output

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>winston.createLogger()</code></td><td><a href="/packages/log/new"><code>log.New()</code></a></td></tr>
<tr><td><code>logger.info(msg, meta)</code></td><td><a href="/packages/log/info"><code>logger.Info(msg, keyvals...)</code></a></td></tr>
<tr><td><code>logger.child({ ... })</code></td><td><a href="/packages/log/with"><code>logger.With(fields...)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/log"

logger := log.New()
logger.Info("server started", "port", 8080)
```

::: tip Full cookbook
See the [**log cookbook**](/packages/log/cookbook) for multi-step recipes and real-world patterns.
:::

## API reference

Browse **5 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/log/set-default">SetDefault</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>SetDefault replaces the package default logger.</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/log/logger">Logger</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Logger wraps slog.Logger with familiar leveled helpers.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/log/logger-default">Logger.Default</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>Default returns the package-level default logger.</td></tr>
<tr><td><a href="/packages/log/logger-new">Logger.New</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>New returns a Logger writing JSON to stderr at info level.</td></tr>
<tr><td><a href="/packages/log/logger-new-with-level">Logger.NewWithLevel</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>NewWithLevel returns a Logger at the given level.</td></tr>
</tbody></table>

