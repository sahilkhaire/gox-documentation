---
title: "exec"
package: "exec"
gox-doc-version: "14"
---

<PackageOverview
  name="exec"
  node="child_process"
  import-path="github.com/sahilkhaire/gox/exec"
  subtitle="Package exec runs external commands with context. Node equivalent: child_process."
  :symbol-count=3
/>
## API reference

Select a symbol below — each page explains what it does, shows Node.js vs Go comparisons, and includes a runnable example.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/exec/exec">Exec</a></td><td><code class="node-cell">exec('ls -la')</code></td><td><span class="kind-pill">func</span></td><td>Exec runs a shell-less command string split on whitespace (child_process.exec).</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/exec/cmd">Cmd</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Cmd wraps os/exec.Cmd with context cancellation.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/exec/cmd-command">Cmd.Command</a></td><td><code class="node-cell">spawn(cmd, args)</code></td><td><span class="kind-pill">method</span></td><td>Command builds a command (child_process.spawn / execFile).</td></tr>
</tbody></table>

