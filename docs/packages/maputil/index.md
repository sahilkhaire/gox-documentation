---
title: "maputil"
package: "maputil"
gox-doc-version: "7"
---

<div class="package-hero"><div class="api-meta"><span class="api-badge node">Node: lodash object utils</span></div><h1>maputil</h1>
<p class="subtitle">Package maputil provides lodash-style object helpers for maps.</p><div class="import-line">import "github.com/sahilkhaire/gox/maputil"</div></div>
## Quick start

A minimal example to get going:

```go
import "github.com/sahilkhaire/gox/maputil"

picked := maputil.Pick(cfg, "host", "port")
```

::: tip Full cookbook
See the [**maputil cookbook**](/packages/maputil/cookbook) for multi-step recipes and real-world patterns.
:::

## What's in this package

If you have used **lodash object utils** in Node.js, this package is your starting point. Browse **9 documented symbols** below — each page includes Node.js, standard Go, and gox side-by-side examples.

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/maputil/clone">Clone</a></td><td><span class="kind-pill">func</span></td><td>Clone returns a shallow copy of m (lodash clone).</td></tr>
<tr><td><a href="/packages/maputil/get">Get</a></td><td><span class="kind-pill">func</span></td><td>Get reads a nested value from m using dot-separated path (lodash get).</td></tr>
<tr><td><a href="/packages/maputil/invert">Invert</a></td><td><span class="kind-pill">func</span></td><td>Invert swaps keys and values; duplicate values overwrite (lodash invert).</td></tr>
<tr><td><a href="/packages/maputil/keys">Keys</a></td><td><span class="kind-pill">func</span></td><td>Keys returns map keys (Object.keys).</td></tr>
<tr><td><a href="/packages/maputil/merge">Merge</a></td><td><span class="kind-pill">func</span></td><td>Merge copies keys from sources into dst, later maps override (lodash merge, shallow).</td></tr>
<tr><td><a href="/packages/maputil/omit">Omit</a></td><td><span class="kind-pill">func</span></td><td>Omit returns a new map without the given keys (lodash omit).</td></tr>
<tr><td><a href="/packages/maputil/pick">Pick</a></td><td><span class="kind-pill">func</span></td><td>Pick returns a new map with only the given keys (lodash pick).</td></tr>
<tr><td><a href="/packages/maputil/set">Set</a></td><td><span class="kind-pill">func</span></td><td>Set writes value at dot-separated path, creating intermediate maps (lodash set).</td></tr>
<tr><td><a href="/packages/maputil/values">Values</a></td><td><span class="kind-pill">func</span></td><td>Values returns map values (Object.values).</td></tr>
</tbody></table>

