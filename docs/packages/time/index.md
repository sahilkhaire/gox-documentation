---
title: "time"
package: "time"
gox-doc-version: "10"
---

<PackageOverview
  name="time"
  node="moment, dayjs (package timex)"
  import-path="github.com/sahilkhaire/gox/time"
  subtitle="Package timex provides date/time helpers similar to moment, dayjs, and date-fns. Import path: github.com/sahilkhaire/gox/time (package timex). Node equivalent: moment, dayjs, date-fns, ms"
  :symbol-count=14
  :has-cookbook=false
  migration-link=""
  narrative="moment/dayjs-style date parsing, formatting, and calendar boundaries. Import as timex to avoid clashing with stdlib time."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/time/now"><div class="featured-name">Now</div><div class="featured-summary">moment()</div></a>
<a class="featured-card" href="/packages/time/format"><div class="featured-name">Format</div><div class="featured-summary">moment.format</div></a>
<a class="featured-card" href="/packages/time/parse"><div class="featured-name">Parse</div><div class="featured-summary">moment.parse</div></a>
</div>

::: warning Package identifier
Import `github.com/sahilkhaire/gox/time` but use the identifier **`timex`** in code — avoids clashing with Go's standard `time` package.
:::

## Common use cases

- Format ISO timestamps
- Start/end of day calculations
- Parse duration strings

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>moment().format()</code></td><td><a href="/packages/time/format"><code>timex.Format(t, layout)</code></a></td></tr>
<tr><td><code>ms('2d')</code></td><td><a href="/packages/time/parse-duration"><code>timex.ParseDuration("2d")</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/time"

now := timex.Now()
formatted := timex.Format(now, timex.LayoutISO)
```

## API reference

Browse **14 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/time/add">Add</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/time/diff">Diff</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/time/end-of-day">EndOfDay</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/time/end-of-month">EndOfMonth</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/time/format">Format</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/time/now">Now</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/time/parse">Parse</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/time/parse-duration">ParseDuration</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/time/start-of-day">StartOfDay</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/time/start-of-month">StartOfMonth</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td></td></tr>
</tbody></table>

### Constants

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/time/layout-date">LayoutDate</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">const</span></td><td></td></tr>
<tr><td><a href="/packages/time/layout-date-time">LayoutDateTime</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">const</span></td><td></td></tr>
<tr><td><a href="/packages/time/layout-iso">LayoutISO</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">const</span></td><td></td></tr>
<tr><td><a href="/packages/time/layout-time">LayoutTime</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">const</span></td><td></td></tr>
</tbody></table>

