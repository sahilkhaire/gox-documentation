---
title: "queue"
package: "queue"
gox-doc-version: "7"
---

<div class="package-hero"><div class="api-meta"><span class="api-badge node">Node: bull</span></div><h1>queue</h1>
<p class="subtitle">Package queue provides a Redis-backed task queue similar to Bull/BullMQ, backed by github.com/hibiken/asynq. Node equivalent: bull, bullmq</p><div class="import-line">import "github.com/sahilkhaire/gox/queue"</div></div>
## What's in this package

If you have used **bull** in Node.js, this package is your starting point. Browse **8 documented symbols** below — each page includes Node.js, standard Go, and gox side-by-side examples.

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/queue/client">Client</a></td><td><span class="kind-pill">type</span></td><td>Client enqueues background tasks.</td></tr>
<tr><td><a href="/packages/queue/enqueue-opts">EnqueueOpts</a></td><td><span class="kind-pill">type</span></td><td>EnqueueOpts configures task enqueue behavior.</td></tr>
<tr><td><a href="/packages/queue/handler">Handler</a></td><td><span class="kind-pill">type</span></td><td>Handler processes a task payload.</td></tr>
<tr><td><a href="/packages/queue/serve-mux">ServeMux</a></td><td><span class="kind-pill">type</span></td><td>ServeMux routes tasks by type name.</td></tr>
<tr><td><a href="/packages/queue/worker">Worker</a></td><td><span class="kind-pill">type</span></td><td>Worker processes tasks from Redis.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/queue/client-new">Client.New</a></td><td><span class="kind-pill">method</span></td><td>New creates a queue client connected to redisAddr (host:port).</td></tr>
<tr><td><a href="/packages/queue/serve-mux-new-serve-mux">ServeMux.NewServeMux</a></td><td><span class="kind-pill">method</span></td><td>NewServeMux returns an empty task router.</td></tr>
<tr><td><a href="/packages/queue/worker-new-worker">Worker.NewWorker</a></td><td><span class="kind-pill">method</span></td><td>NewWorker creates a worker that dispatches to mux.</td></tr>
</tbody></table>

