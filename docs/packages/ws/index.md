---
title: "ws"
package: "ws"
gox-doc-version: "10"
---

<PackageOverview
  name="ws"
  node="ws"
  import-path="github.com/sahilkhaire/gox/ws"
  subtitle="Package ws provides WebSocket client and server helpers. Node equivalent: ws, lightweight socket.io patterns."
  :symbol-count=8
  :has-cookbook=true
  migration-link="/migration/ws"
  narrative="WebSocket server, client dial, and upgrade helpers built on gorilla/websocket."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/ws/server-new-server"><div class="featured-name">NewServer</div><div class="featured-summary">WebSocket.Server</div></a>
<a class="featured-card" href="/packages/ws/conn-dial"><div class="featured-name">Dial</div><div class="featured-summary">new WebSocket()</div></a>
<a class="featured-card" href="/packages/ws/conn-upgrade"><div class="featured-name">Upgrade</div><div class="featured-summary">Upgrade handler</div></a>
</div>

## Common use cases

- Real-time chat backends
- Upgrade HTTP handlers to WS
- Connect to external WS APIs

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/ws"

conn, err := ws.Dial(ctx, "wss://example.com/ws", nil)
```

::: tip Full cookbook
See the [**ws cookbook**](/packages/ws/cookbook) for multi-step recipes and real-world patterns.
:::

::: info Migration guide
Coming from Node.js? Read the [**migration guide**](/migration/ws) for side-by-side patterns.
:::

## API reference

Browse **8 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/ws/conn">Conn</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Conn wraps a gorilla WebSocket connection with JSON helpers.</td></tr>
<tr><td><a href="/packages/ws/server">Server</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Server registers WebSocket routes.</td></tr>
<tr><td><a href="/packages/ws/upgrader">Upgrader</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Upgrader upgrades HTTP connections to WebSocket.</td></tr>
<tr><td><a href="/packages/ws/ws-handler">WSHandler</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>WSHandler is a WebSocket connection handler.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/ws/conn-dial">Conn.Dial</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>Dial connects to the WebSocket url.</td></tr>
<tr><td><a href="/packages/ws/server-new-server">Server.NewServer</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>NewServer creates a WebSocket server with the given upgrader (nil uses DefaultUpgrader).</td></tr>
<tr><td><a href="/packages/ws/conn-upgrade">Conn.Upgrade</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>Upgrade upgrades the HTTP request on w/r using up (nil uses DefaultUpgrader).</td></tr>
</tbody></table>

### Variables

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/ws/default-upgrader">DefaultUpgrader</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">var</span></td><td>DefaultUpgrader allows all origins (suitable for tests; tighten in production).</td></tr>
</tbody></table>

