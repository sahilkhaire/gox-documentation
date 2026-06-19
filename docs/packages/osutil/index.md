---
title: "osutil"
package: "osutil"
gox-doc-version: "7"
---

<div class="package-hero"><div class="api-meta"><span class="api-badge node">Node: os</span></div><h1>osutil</h1>
<p class="subtitle">Package osutil exposes host and OS information. Node equivalent: os module helpers.</p><div class="import-line">import "github.com/sahilkhaire/gox/osutil"</div></div>
## What's in this package

If you have used **os** in Node.js, this package is your starting point. Browse **9 documented symbols** below — each page includes Node.js, standard Go, and gox side-by-side examples.

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/osutil/arch">Arch</a></td><td><span class="kind-pill">func</span></td><td>Arch returns the architecture (runtime.GOARCH).</td></tr>
<tr><td><a href="/packages/osutil/cp-us">CPUs</a></td><td><span class="kind-pill">func</span></td><td>CPUs returns the number of logical CPUs.</td></tr>
<tr><td><a href="/packages/osutil/homedir">Homedir</a></td><td><span class="kind-pill">func</span></td><td>Homedir returns the current user's home directory.</td></tr>
<tr><td><a href="/packages/osutil/hostname">Hostname</a></td><td><span class="kind-pill">func</span></td><td>Hostname returns the host name.</td></tr>
<tr><td><a href="/packages/osutil/platform">Platform</a></td><td><span class="kind-pill">func</span></td><td>Platform returns the OS (runtime.GOOS).</td></tr>
<tr><td><a href="/packages/osutil/temp-dir">TempDir</a></td><td><span class="kind-pill">func</span></td><td>TempDir returns the default temp directory.</td></tr>
<tr><td><a href="/packages/osutil/user-info-numeric">UserInfoNumeric</a></td><td><span class="kind-pill">func</span></td><td>UserInfoNumeric is like UserInfo but parses uid/gid as ints when possible.</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/osutil/user">User</a></td><td><span class="kind-pill">type</span></td><td>User holds user identity fields when available (os.userInfo).</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/osutil/user-user-info">User.UserInfo</a></td><td><span class="kind-pill">method</span></td><td>UserInfo returns the current user.</td></tr>
</tbody></table>

