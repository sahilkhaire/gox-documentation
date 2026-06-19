---
title: "ws"
package: "ws"
gox-doc-version: "7"
---

<div class="package-hero"><div class="api-meta"><span class="api-badge node">Node: ws</span></div><h1>ws</h1>
<p class="subtitle">Package ws provides WebSocket client and server helpers. Node equivalent: ws, lightweight socket.io patterns.</p><div class="import-line">import "github.com/sahilkhaire/gox/ws"</div></div>
## What's in this package

If you have used **ws** in Node.js, this package is your starting point. Browse **8 documented symbols** below — each page includes Node.js, standard Go, and gox side-by-side examples.

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/ws/conn">Conn</a></td><td><span class="kind-pill">type</span></td><td>Conn wraps a gorilla WebSocket connection with JSON helpers.</td></tr>
<tr><td><a href="/packages/ws/server">Server</a></td><td><span class="kind-pill">type</span></td><td>Server registers WebSocket routes.</td></tr>
<tr><td><a href="/packages/ws/upgrader">Upgrader</a></td><td><span class="kind-pill">type</span></td><td>Upgrader upgrades HTTP connections to WebSocket.</td></tr>
<tr><td><a href="/packages/ws/ws-handler">WSHandler</a></td><td><span class="kind-pill">type</span></td><td>WSHandler is a WebSocket connection handler.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/ws/conn-dial">Conn.Dial</a></td><td><span class="kind-pill">method</span></td><td>Dial connects to the WebSocket url.</td></tr>
<tr><td><a href="/packages/ws/server-new-server">Server.NewServer</a></td><td><span class="kind-pill">method</span></td><td>NewServer creates a WebSocket server with the given upgrader (nil uses DefaultUpgrader).</td></tr>
<tr><td><a href="/packages/ws/conn-upgrade">Conn.Upgrade</a></td><td><span class="kind-pill">method</span></td><td>Upgrade upgrades the HTTP request on w/r using up (nil uses DefaultUpgrader).</td></tr>
</tbody></table>

### Variables

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/ws/default-upgrader">DefaultUpgrader</a></td><td><span class="kind-pill">var</span></td><td>DefaultUpgrader allows all origins (suitable for tests; tighten in production).</td></tr>
</tbody></table>

