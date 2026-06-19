---
title: "log"
package: "log"
gox-doc-version: "14"
---

<PackageOverview
  name="log"
  node="winston, pino"
  import-path="github.com/sahilkhaire/gox/log"
  subtitle="Package log provides winston/pino-style logging on top of log/slog. Node equivalent: winston, pino."
  :symbol-count=5
/>
## API reference

Select a symbol below — each page explains what it does, shows Node.js vs Go comparisons, and includes a runnable example.

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
<tr><td><a href="/packages/log/logger-new">Logger.New</a></td><td><code class="node-cell">winston.createLogger()</code></td><td><span class="kind-pill">method</span></td><td>New returns a Logger writing JSON to stderr at info level.</td></tr>
<tr><td><a href="/packages/log/logger-new-with-level">Logger.NewWithLevel</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>NewWithLevel returns a Logger at the given level.</td></tr>
</tbody></table>

