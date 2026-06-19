---
title: "osutil"
package: "osutil"
gox-doc-version: "11"
---

<PackageOverview
  name="osutil"
  node="os"
  import-path="github.com/sahilkhaire/gox/osutil"
  subtitle="Package osutil exposes host and OS information. Node equivalent: os module helpers."
  :symbol-count=9
/>
## API reference

Select a symbol below — each page explains what it does, shows Node.js vs Go comparisons, and includes a runnable example.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/osutil/arch">Arch</a></td><td><code class="node-cell">process.arch</code></td><td><span class="kind-pill">func</span></td><td>Arch returns the architecture (runtime.GOARCH).</td></tr>
<tr><td><a href="/packages/osutil/cp-us">CPUs</a></td><td><code class="node-cell">os.cpus().length</code></td><td><span class="kind-pill">func</span></td><td>CPUs returns the number of logical CPUs.</td></tr>
<tr><td><a href="/packages/osutil/homedir">Homedir</a></td><td><code class="node-cell">os.homedir()</code></td><td><span class="kind-pill">func</span></td><td>Homedir returns the current user's home directory.</td></tr>
<tr><td><a href="/packages/osutil/hostname">Hostname</a></td><td><code class="node-cell">os.hostname()</code></td><td><span class="kind-pill">func</span></td><td>Hostname returns the host name.</td></tr>
<tr><td><a href="/packages/osutil/platform">Platform</a></td><td><code class="node-cell">process.platform</code></td><td><span class="kind-pill">func</span></td><td>Platform returns the OS (runtime.GOOS).</td></tr>
<tr><td><a href="/packages/osutil/temp-dir">TempDir</a></td><td><code class="node-cell">os.tmpdir()</code></td><td><span class="kind-pill">func</span></td><td>TempDir returns the default temp directory.</td></tr>
<tr><td><a href="/packages/osutil/user-info-numeric">UserInfoNumeric</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>UserInfoNumeric is like UserInfo but parses uid/gid as ints when possible.</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/osutil/user">User</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>User holds user identity fields when available (os.userInfo).</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/osutil/user-user-info">User.UserInfo</a></td><td><code class="node-cell">os.userInfo()</code></td><td><span class="kind-pill">method</span></td><td>UserInfo returns the current user.</td></tr>
</tbody></table>

