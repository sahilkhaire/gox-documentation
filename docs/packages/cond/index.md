---
title: "cond"
package: "cond"
gox-doc-version: "10"
---

<PackageOverview
  name="cond"
  node="ternary ? :, nullish ??"
  import-path="github.com/sahilkhaire/gox/cond"
  subtitle="Package cond provides ternary and nullish-coalescing-style helpers, similar to JavaScript's ? : and ?? operators."
  :symbol-count=5
  :has-cookbook=true
  migration-link=""
  narrative="Replace JavaScript ternary and nullish coalescing with typed, generic helpers. Ideal for config defaults, labels, and safe fallbacks without nested if/else noise."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/cond/if"><div class="featured-name">If</div><div class="featured-summary">Ternary ? : operator</div></a>
<a class="featured-card" href="/packages/cond/coalesce"><div class="featured-name">Coalesce</div><div class="featured-summary">Nullish ?? chain</div></a>
<a class="featured-card" href="/packages/cond/if-lazy"><div class="featured-name">IfLazy</div><div class="featured-summary">Lazy ternary branches</div></a>
</div>

## Common use cases

- Pass/fail labels from boolean conditions
- Default values when strings or pointers are empty
- Lazy evaluation when one branch is expensive

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>cond ? a : b</code></td><td><a href="/packages/cond/if"><code>cond.If(cond, a, b)</code></a></td></tr>
<tr><td><code>cond ? f() : g()</code></td><td><a href="/packages/cond/if-lazy"><code>cond.IfLazy(cond, f, g)</code></a></td></tr>
<tr><td><code>a ?? b ?? c</code></td><td><a href="/packages/cond/coalesce"><code>cond.Coalesce(a, b, c)</code></a></td></tr>
<tr><td><code>ptr ?? fallback</code></td><td><a href="/packages/cond/coalesce-ptr"><code>cond.CoalescePtr(ptr, fallback)</code></a></td></tr>
<tr><td><code>obj?.field ?? "guest"</code></td><td><a href="/packages/cond/coalesce-fn"><code>cond.CoalesceFn(ptr, func(v T) R { ... }, fallback)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/cond"

label := cond.If(score >= 60, "pass", "fail")
```

::: tip Full cookbook
See the [**cond cookbook**](/packages/cond/cookbook) for multi-step recipes and real-world patterns.
:::

## API reference

Browse **5 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/cond/coalesce">Coalesce</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Coalesce returns the first value that is not the zero value (Node: a ?? b ?? c).</td></tr>
<tr><td><a href="/packages/cond/coalesce-fn">CoalesceFn</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>CoalesceFn returns the first non-zero result from the given functions, in order.</td></tr>
<tr><td><a href="/packages/cond/coalesce-ptr">CoalescePtr</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>CoalescePtr returns the first non-nil pointer's value, or zero if all nil.</td></tr>
<tr><td><a href="/packages/cond/if">If</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>If returns a when cond is true, otherwise b (Node: cond ? a : b).</td></tr>
<tr><td><a href="/packages/cond/if-lazy">IfLazy</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>IfLazy evaluates a or b lazily via thunks (Node: cond ? a() : b()).</td></tr>
</tbody></table>

