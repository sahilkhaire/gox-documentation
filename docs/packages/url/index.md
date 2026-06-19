---
title: "url"
package: "url"
gox-doc-version: "10"
---

<PackageOverview
  name="url"
  node="url, querystring"
  import-path="github.com/sahilkhaire/gox/url"
  subtitle="Package url provides Node url and querystring helpers over net/url."
  :symbol-count=6
  :has-cookbook=false
  migration-link=""
  narrative="Parse, format, and query-string utilities matching Node url and querystring modules."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/url/url-parse"><div class="featured-name">Parse</div><div class="featured-summary">new URL()</div></a>
<a class="featured-card" href="/packages/url/parse-query"><div class="featured-name">ParseQuery</div><div class="featured-summary">querystring.parse</div></a>
<a class="featured-card" href="/packages/url/encode-query"><div class="featured-name">EncodeQuery</div><div class="featured-summary">querystring.stringify</div></a>
</div>

## Common use cases

- Parse redirect URLs
- Build query strings for APIs
- Resolve relative links

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>new URL(str)</code></td><td><a href="/packages/url/parse"><code>url.Parse(str)</code></a></td></tr>
<tr><td><code>url.format(obj)</code></td><td><a href="/packages/url/format"><code>url.Format(u)</code></a></td></tr>
<tr><td><code>url.resolve(base, rel)</code></td><td><a href="/packages/url/resolve"><code>url.Resolve(base, rel)</code></a></td></tr>
<tr><td><code>querystring.parse(qs)</code></td><td><a href="/packages/url/parse-query"><code>url.ParseQuery(qs)</code></a></td></tr>
<tr><td><code>querystring.stringify(obj)</code></td><td><a href="/packages/url/encode-query"><code>url.EncodeQuery(obj)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/url"

u, err := url.Parse("https://example.com/api?page=1")
```

## API reference

Browse **6 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/url/encode-query">EncodeQuery</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>EncodeQuery encodes values (querystring.stringify).</td></tr>
<tr><td><a href="/packages/url/format">Format</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Format returns the serialized URL (url.format).</td></tr>
<tr><td><a href="/packages/url/parse-query">ParseQuery</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>ParseQuery parses a query string (querystring.parse).</td></tr>
<tr><td><a href="/packages/url/resolve">Resolve</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Resolve resolves relative against base (url.resolve).</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/url/url">URL</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>URL wraps net/url.URL with Node-friendly helpers.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/url/url-parse">URL.Parse</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>Parse parses a URL string (url.parse).</td></tr>
</tbody></table>

