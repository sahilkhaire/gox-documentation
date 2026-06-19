---
title: "semver"
package: "semver"
gox-doc-version: "10"
---

<PackageOverview
  name="semver"
  node="semver"
  import-path="github.com/sahilkhaire/gox/semver"
  subtitle="Package semver parses and compares semantic versions, similar to npm semver. Node equivalent: semver"
  :symbol-count=5
  :has-cookbook=false
  migration-link=""
  narrative="semver module helpers — parse, compare, increment, and satisfy ranges."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/semver/version-parse"><div class="featured-name">Parse</div><div class="featured-summary">semver.parse</div></a>
<a class="featured-card" href="/packages/semver/compare"><div class="featured-name">Compare</div><div class="featured-summary">semver.compare</div></a>
<a class="featured-card" href="/packages/semver/satisfies"><div class="featured-name">Satisfies</div><div class="featured-summary">semver.satisfies</div></a>
</div>

## Common use cases

- Gate features by app version
- Compare dependency versions
- Bump version segments

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>semver.parse(v)</code></td><td><a href="/packages/semver/parse"><code>semver.Parse(v)</code></a></td></tr>
<tr><td><code>semver.compare(a, b)</code></td><td><a href="/packages/semver/compare"><code>semver.Compare(a, b)</code></a></td></tr>
<tr><td><code>semver.satisfies(v, range)</code></td><td><a href="/packages/semver/satisfies"><code>semver.Satisfies(v, constraint)</code></a></td></tr>
<tr><td><code>semver.inc(v, 'minor')</code></td><td><a href="/packages/semver/inc"><code>semver.Inc(v, part)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/semver"

v, err := semver.Parse("1.2.3")
```

## API reference

Browse **5 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/semver/compare">Compare</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Compare returns -1 if a &lt; b, 0 if equal, 1 if a &gt; b.</td></tr>
<tr><td><a href="/packages/semver/inc">Inc</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Inc increments part (major, minor, patch) and returns the new version string.</td></tr>
<tr><td><a href="/packages/semver/satisfies">Satisfies</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Satisfies reports whether version matches constraint (npm range syntax).</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/semver/version">Version</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Version is a semantic version.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/semver/version-parse">Version.Parse</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>Parse parses a semver string.</td></tr>
</tbody></table>

