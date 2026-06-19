---
title: "semver"
package: "semver"
gox-doc-version: "11"
---

<PackageOverview
  name="semver"
  node="semver"
  import-path="github.com/sahilkhaire/gox/semver"
  subtitle="Package semver parses and compares semantic versions, similar to npm semver. Node equivalent: semver"
  :symbol-count=5
/>
## API reference

Select a symbol below — each page explains what it does, shows Node.js vs Go comparisons, and includes a runnable example.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/semver/compare">Compare</a></td><td><code class="node-cell">semver.compare(a, b)</code></td><td><span class="kind-pill">func</span></td><td>Compare returns -1 if a &lt; b, 0 if equal, 1 if a &gt; b.</td></tr>
<tr><td><a href="/packages/semver/inc">Inc</a></td><td><code class="node-cell">semver.inc(v, 'minor')</code></td><td><span class="kind-pill">func</span></td><td>Inc increments part (major, minor, patch) and returns the new version string.</td></tr>
<tr><td><a href="/packages/semver/satisfies">Satisfies</a></td><td><code class="node-cell">semver.satisfies(v, range)</code></td><td><span class="kind-pill">func</span></td><td>Satisfies reports whether version matches constraint (npm range syntax).</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/semver/version">Version</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Version is a semantic version.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/semver/version-parse">Version.Parse</a></td><td><code class="node-cell">semver.parse(v)</code></td><td><span class="kind-pill">method</span></td><td>Parse parses a semver string.</td></tr>
</tbody></table>

