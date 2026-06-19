---
title: "cron"
package: "cron"
gox-doc-version: "10"
---

<PackageOverview
  name="cron"
  node="node-cron"
  import-path="github.com/sahilkhaire/gox/cron"
  subtitle="Package cron schedules recurring jobs using cron expressions, similar to node-cron. Node equivalent: node-cron"
  :symbol-count=4
  :has-cookbook=true
  migration-link=""
  narrative="node-cron-style schedulers for recurring jobs."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/cron/scheduler-new"><div class="featured-name">New</div><div class="featured-summary">node-cron</div></a>
<a class="featured-card" href="/packages/cron/scheduler"><div class="featured-name">Schedule</div><div class="featured-summary">cron.schedule</div></a>
<a class="featured-card" href="/packages/cron/parse"><div class="featured-name">Parse</div><div class="featured-summary">Parse expression</div></a>
</div>

## Common use cases

- Run cleanup every hour
- Schedule report generation
- Parse cron expressions

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>cron.schedule(expr, fn)</code></td><td><a href="/packages/cron/schedule"><code>Schedule(expr, fn)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/cron"

sched := cron.New()
sched.Schedule("0 * * * *", cleanup)
```

::: tip Full cookbook
See the [**cron cookbook**](/packages/cron/cookbook) for multi-step recipes and real-world patterns.
:::

## API reference

Browse **4 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/cron/parse">Parse</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Parse validates a cron expression (with optional seconds field).</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/cron/job-id">JobID</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>JobID identifies a scheduled job.</td></tr>
<tr><td><a href="/packages/cron/scheduler">Scheduler</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Scheduler runs cron jobs in the background.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/cron/scheduler-new">Scheduler.New</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>New creates a scheduler and starts its internal runner.</td></tr>
</tbody></table>

