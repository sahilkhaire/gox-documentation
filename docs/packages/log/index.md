---
title: "log"
package: "log"
gox-doc-version: "7"
---

<div class="package-hero"><div class="api-meta"><span class="api-badge node">Node: winston, pino</span></div><h1>log</h1>
<p class="subtitle">Package log provides winston/pino-style logging on top of log/slog. Node equivalent: winston, pino.</p><div class="import-line">import "github.com/sahilkhaire/gox/log"</div></div>
## What's in this package

If you have used **winston, pino** in Node.js, this package is your starting point. Browse **5 documented symbols** below — each page includes Node.js, standard Go, and gox side-by-side examples.

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/log/set-default">SetDefault</a></td><td><span class="kind-pill">func</span></td><td>SetDefault replaces the package default logger.</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/log/logger">Logger</a></td><td><span class="kind-pill">type</span></td><td>Logger wraps slog.Logger with familiar leveled helpers.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/log/logger-default">Logger.Default</a></td><td><span class="kind-pill">method</span></td><td>Default returns the package-level default logger.</td></tr>
<tr><td><a href="/packages/log/logger-new">Logger.New</a></td><td><span class="kind-pill">method</span></td><td>New returns a Logger writing JSON to stderr at info level.</td></tr>
<tr><td><a href="/packages/log/logger-new-with-level">Logger.NewWithLevel</a></td><td><span class="kind-pill">method</span></td><td>NewWithLevel returns a Logger at the given level.</td></tr>
</tbody></table>

