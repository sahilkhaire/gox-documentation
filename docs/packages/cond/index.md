---
title: "cond"
package: "cond"
gox-doc-version: "11"
---

<PackageOverview
  name="cond"
  node="ternary ? :, nullish ??"
  import-path="github.com/sahilkhaire/gox/cond"
  subtitle="Package cond provides ternary and nullish-coalescing-style helpers, similar to JavaScript's ? : and ?? operators."
  :symbol-count=5
/>
## API reference

Select a symbol below — each page explains what it does, shows Node.js vs Go comparisons, and includes a runnable example.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/cond/coalesce">Coalesce</a></td><td><code class="node-cell">a ?? b ?? c</code></td><td><span class="kind-pill">func</span></td><td>Coalesce returns the first value that is not the zero value (Node: a ?? b ?? c).</td></tr>
<tr><td><a href="/packages/cond/coalesce-fn">CoalesceFn</a></td><td><code class="node-cell">obj?.field ?? "guest"</code></td><td><span class="kind-pill">func</span></td><td>CoalesceFn returns the first non-zero result from the given functions, in order.</td></tr>
<tr><td><a href="/packages/cond/coalesce-ptr">CoalescePtr</a></td><td><code class="node-cell">ptr ?? fallback</code></td><td><span class="kind-pill">func</span></td><td>CoalescePtr returns the first non-nil pointer's value, or zero if all nil.</td></tr>
<tr><td><a href="/packages/cond/if">If</a></td><td><code class="node-cell">cond ? a : b</code></td><td><span class="kind-pill">func</span></td><td>If returns a when cond is true, otherwise b (Node: cond ? a : b).</td></tr>
<tr><td><a href="/packages/cond/if-lazy">IfLazy</a></td><td><code class="node-cell">cond ? f() : g()</code></td><td><span class="kind-pill">func</span></td><td>IfLazy evaluates a or b lazily via thunks (Node: cond ? a() : b()).</td></tr>
</tbody></table>

