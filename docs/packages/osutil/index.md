---
title: "osutil"
package: "osutil"
gox-doc-version: "10"
---

<PackageOverview
  name="osutil"
  node="os"
  import-path="github.com/sahilkhaire/gox/osutil"
  subtitle="Package osutil exposes host and OS information. Node equivalent: os module helpers."
  :symbol-count=9
  :has-cookbook=false
  migration-link=""
  narrative="os module helpers — hostname, homedir, temp dirs, platform, and architecture."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/osutil/hostname"><div class="featured-name">Hostname</div><div class="featured-summary">os.hostname</div></a>
<a class="featured-card" href="/packages/osutil/homedir"><div class="featured-name">Homedir</div><div class="featured-summary">os.homedir</div></a>
<a class="featured-card" href="/packages/osutil/temp-dir"><div class="featured-name">TempDir</div><div class="featured-summary">os.tmpdir</div></a>
</div>

## Common use cases

- Detect runtime environment
- Resolve home directory paths
- Create temp workspaces

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>os.hostname()</code></td><td><a href="/packages/osutil/hostname"><code>osutil.Hostname()</code></a></td></tr>
<tr><td><code>os.homedir()</code></td><td><a href="/packages/osutil/homedir"><code>osutil.Homedir()</code></a></td></tr>
<tr><td><code>os.tmpdir()</code></td><td><a href="/packages/osutil/temp-dir"><code>osutil.TempDir()</code></a></td></tr>
<tr><td><code>os.cpus().length</code></td><td><a href="/packages/osutil/cp-us"><code>osutil.CPUs()</code></a></td></tr>
<tr><td><code>process.platform</code></td><td><a href="/packages/osutil/platform"><code>osutil.Platform()</code></a></td></tr>
<tr><td><code>process.arch</code></td><td><a href="/packages/osutil/arch"><code>osutil.Arch()</code></a></td></tr>
<tr><td><code>os.userInfo()</code></td><td><a href="/packages/osutil/user-info"><code>osutil.UserInfo()</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/osutil"

home, err := osutil.Homedir()
```

## API reference

Browse **9 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/osutil/arch">Arch</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Arch returns the architecture (runtime.GOARCH).</td></tr>
<tr><td><a href="/packages/osutil/cp-us">CPUs</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>CPUs returns the number of logical CPUs.</td></tr>
<tr><td><a href="/packages/osutil/homedir">Homedir</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Homedir returns the current user's home directory.</td></tr>
<tr><td><a href="/packages/osutil/hostname">Hostname</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Hostname returns the host name.</td></tr>
<tr><td><a href="/packages/osutil/platform">Platform</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Platform returns the OS (runtime.GOOS).</td></tr>
<tr><td><a href="/packages/osutil/temp-dir">TempDir</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>TempDir returns the default temp directory.</td></tr>
<tr><td><a href="/packages/osutil/user-info-numeric">UserInfoNumeric</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>UserInfoNumeric is like UserInfo but parses uid/gid as ints when possible.</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/osutil/user">User</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>User holds user identity fields when available (os.userInfo).</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/osutil/user-user-info">User.UserInfo</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>UserInfo returns the current user.</td></tr>
</tbody></table>

