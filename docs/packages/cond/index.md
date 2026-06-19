---
title: "cond"
package: "cond"
gox-doc-version: "7"
---

<div class="package-hero"><div class="api-meta"><span class="api-badge node">Node: ternary ? :, nullish ??</span></div><h1>cond</h1>
<p class="subtitle">Package cond provides ternary and nullish-coalescing-style helpers, similar to JavaScript's ? : and ?? operators.</p><div class="import-line">import "github.com/sahilkhaire/gox/cond"</div></div>
## Quick start

A minimal example to get going:

```go
import "github.com/sahilkhaire/gox/cond"

label := cond.If(score >= 60, "pass", "fail")
```

::: tip Full cookbook
See the [**cond cookbook**](/packages/cond/cookbook) for multi-step recipes and real-world patterns.
:::

## What's in this package

If you have used **ternary ? :, nullish ??** in Node.js, this package is your starting point. Browse **5 documented symbols** below — each page includes Node.js, standard Go, and gox side-by-side examples.

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/cond/coalesce">Coalesce</a></td><td><span class="kind-pill">func</span></td><td>Coalesce returns the first value that is not the zero value (Node: a ?? b ?? c).</td></tr>
<tr><td><a href="/packages/cond/coalesce-fn">CoalesceFn</a></td><td><span class="kind-pill">func</span></td><td>CoalesceFn returns the first non-zero result from the given functions, in order.</td></tr>
<tr><td><a href="/packages/cond/coalesce-ptr">CoalescePtr</a></td><td><span class="kind-pill">func</span></td><td>CoalescePtr returns the first non-nil pointer's value, or zero if all nil.</td></tr>
<tr><td><a href="/packages/cond/if">If</a></td><td><span class="kind-pill">func</span></td><td>If returns a when cond is true, otherwise b (Node: cond ? a : b).</td></tr>
<tr><td><a href="/packages/cond/if-lazy">IfLazy</a></td><td><span class="kind-pill">func</span></td><td>IfLazy evaluates a or b lazily via thunks (Node: cond ? a() : b()).</td></tr>
</tbody></table>

