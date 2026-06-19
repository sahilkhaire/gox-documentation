---
title: "exec"
package: "exec"
gox-doc-version: "10"
---

<PackageOverview
  name="exec"
  node="child_process"
  import-path="github.com/sahilkhaire/gox/exec"
  subtitle="Package exec runs external commands with context. Node equivalent: child_process."
  :symbol-count=3
  :has-cookbook=false
  migration-link=""
  narrative="child_process-style command execution with context timeouts and captured stdout/stderr."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/exec/cmd-command"><div class="featured-name">Command</div><div class="featured-summary">spawn</div></a>
<a class="featured-card" href="/packages/exec/exec"><div class="featured-name">Exec</div><div class="featured-summary">execFile</div></a>
</div>

## Common use cases

- Run CLI tools from Go services
- Execute migration scripts
- Capture subprocess output

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>spawn(cmd, args)</code></td><td><a href="/packages/exec/command"><code>exec.Command(ctx, name, args...)</code></a></td></tr>
<tr><td><code>exec('ls -la')</code></td><td><a href="/packages/exec/exec"><code>exec.Exec(ctx, command)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/exec"

out, err := exec.Command(ctx, "git", "status").Output()
```

## API reference

Browse **3 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/exec/exec">Exec</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Exec runs a shell-less command string split on whitespace (child_process.exec).</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/exec/cmd">Cmd</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Cmd wraps os/exec.Cmd with context cancellation.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/exec/cmd-command">Cmd.Command</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>Command builds a command (child_process.spawn / execFile).</td></tr>
</tbody></table>

