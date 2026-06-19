---
title: "event"
package: "event"
gox-doc-version: "10"
---

<PackageOverview
  name="event"
  node="EventEmitter"
  import-path="github.com/sahilkhaire/gox/event"
  subtitle="Package event provides a thread-safe event emitter similar to Node.js EventEmitter. Node equivalent: events.EventEmitter"
  :symbol-count=2
  :has-cookbook=false
  migration-link=""
  narrative="EventEmitter-style pub/sub for in-process events."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/event/emitter-new"><div class="featured-name">New</div><div class="featured-summary">EventEmitter</div></a>
<a class="featured-card" href="/packages/event/emitter"><div class="featured-name">On</div><div class="featured-summary">emitter.on</div></a>
<a class="featured-card" href="/packages/event/emitter"><div class="featured-name">Emit</div><div class="featured-summary">emitter.emit</div></a>
</div>

## Common use cases

- Decouple modules with events
- Notify listeners on state changes
- Build plugin hooks

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>emitter.on(event, fn)</code></td><td><a href="/packages/event/on"><code>On(event, handler)</code></a></td></tr>
<tr><td><code>emitter.once(event, fn)</code></td><td><a href="/packages/event/once"><code>Once(event, handler)</code></a></td></tr>
<tr><td><code>emitter.off(event, fn)</code></td><td><a href="/packages/event/off"><code>Off(event, handler)</code></a></td></tr>
<tr><td><code>emitter.emit(event, ...args)</code></td><td><a href="/packages/event/emit"><code>Emit(event, args...)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/event"

em := event.New()
em.On("ready", func() { /* ... */ })
```

## API reference

Browse **2 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/event/emitter">Emitter</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Emitter dispatches named events to registered handlers.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/event/emitter-new">Emitter.New</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>New returns an empty Emitter.</td></tr>
</tbody></table>

