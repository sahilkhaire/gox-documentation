---
title: "queue"
package: "queue"
gox-doc-version: "11"
---

<PackageOverview
  name="queue"
  node="bull"
  import-path="github.com/sahilkhaire/gox/queue"
  subtitle="Package queue provides a Redis-backed task queue similar to Bull/BullMQ, backed by github.com/hibiken/asynq. Node equivalent: bull, bullmq"
  :symbol-count=8
/>
## API reference

Select a symbol below — each page explains what it does, shows Node.js vs Go comparisons, and includes a runnable example.

<SymbolFilter placeholder="Filter symbols…" />

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/queue/client">Client</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Client enqueues background tasks.</td></tr>
<tr><td><a href="/packages/queue/enqueue-opts">EnqueueOpts</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>EnqueueOpts configures task enqueue behavior.</td></tr>
<tr><td><a href="/packages/queue/handler">Handler</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Handler processes a task payload.</td></tr>
<tr><td><a href="/packages/queue/serve-mux">ServeMux</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>ServeMux routes tasks by type name.</td></tr>
<tr><td><a href="/packages/queue/worker">Worker</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Worker processes tasks from Redis.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/queue/client-new">Client.New</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>New creates a queue client connected to redisAddr (host:port).</td></tr>
<tr><td><a href="/packages/queue/serve-mux-new-serve-mux">ServeMux.NewServeMux</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>NewServeMux returns an empty task router.</td></tr>
<tr><td><a href="/packages/queue/worker-new-worker">Worker.NewWorker</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>NewWorker creates a worker that dispatches to mux.</td></tr>
</tbody></table>

