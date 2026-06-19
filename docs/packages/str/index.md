---
title: "str"
package: "str"
gox-doc-version: "10"
---

<PackageOverview
  name="str"
  node="string helpers"
  import-path="github.com/sahilkhaire/gox/str"
  subtitle="Package str provides lodash/String-style string helpers."
  :symbol-count=9
  :has-cookbook=true
  migration-link=""
  narrative="String casing, slugging, padding, and truncation helpers matching lodash and String prototype patterns."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/str/slug"><div class="featured-name">Slug</div><div class="featured-summary">URL-friendly slug</div></a>
<a class="featured-card" href="/packages/str/camel"><div class="featured-name">Camel</div><div class="featured-summary">camelCase</div></a>
<a class="featured-card" href="/packages/str/truncate"><div class="featured-name">Truncate</div><div class="featured-summary">_.truncate</div></a>
</div>

## Common use cases

- Generate URL slugs from titles
- Normalize casing for APIs
- Truncate UI labels safely by rune count

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>slugify(s)</code></td><td><a href="/packages/str/slug"><code>str.Slug(s)</code></a></td></tr>
<tr><td><code>camelCase(s)</code></td><td><a href="/packages/str/camel"><code>str.Camel(s)</code></a></td></tr>
<tr><td><code>snake_case(s)</code></td><td><a href="/packages/str/snake"><code>str.Snake(s)</code></a></td></tr>
<tr><td><code>PascalCase(s)</code></td><td><a href="/packages/str/pascal"><code>str.Pascal(s)</code></a></td></tr>
<tr><td><code>_.capitalize(s)</code></td><td><a href="/packages/str/capitalize"><code>str.Capitalize(s)</code></a></td></tr>
<tr><td><code>_.truncate(s, n)</code></td><td><a href="/packages/str/truncate"><code>str.Truncate(s, n)</code></a></td></tr>
<tr><td><code>s.padStart(n, ch)</code></td><td><a href="/packages/str/pad-start"><code>str.PadStart(s, n, ch)</code></a></td></tr>
<tr><td><code>s.padEnd(n, ch)</code></td><td><a href="/packages/str/pad-end"><code>str.PadEnd(s, n, ch)</code></a></td></tr>
<tr><td><code>!s.trim()</code></td><td><a href="/packages/str/is-blank"><code>str.IsBlank(s)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/str"

slug := str.Slug("Hello World!")
```

::: tip Full cookbook
See the [**str cookbook**](/packages/str/cookbook) for multi-step recipes and real-world patterns.
:::

## API reference

Browse **9 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/str/camel">Camel</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Camel converts kebab/snake to camelCase.</td></tr>
<tr><td><a href="/packages/str/capitalize">Capitalize</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Capitalize uppercases the first rune.</td></tr>
<tr><td><a href="/packages/str/is-blank">IsBlank</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>IsBlank reports whether s is empty or only whitespace.</td></tr>
<tr><td><a href="/packages/str/pad-end">PadEnd</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>PadEnd pads s on the right to length with pad (String.padEnd).</td></tr>
<tr><td><a href="/packages/str/pad-start">PadStart</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>PadStart pads s on the left to length with pad (String.padStart).</td></tr>
<tr><td><a href="/packages/str/pascal">Pascal</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Pascal converts to PascalCase.</td></tr>
<tr><td><a href="/packages/str/slug">Slug</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Slug converts s to a URL-friendly slug.</td></tr>
<tr><td><a href="/packages/str/snake">Snake</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Snake converts to snake_case.</td></tr>
<tr><td><a href="/packages/str/truncate">Truncate</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Truncate shortens s to max runes, adding suffix if needed.</td></tr>
</tbody></table>

