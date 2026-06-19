---
title: "http"
package: "http"
gox-doc-version: "11"
---

<PackageOverview
  name="http"
  node="express, cors, helmet, morgan"
  import-path="github.com/sahilkhaire/gox/http"
  subtitle="Package http provides an Express-like HTTP server on chi. Node equivalent: express, cors, helmet, morgan (Logger middleware), multer (ParseMultipart), express-session (SessionMiddleware), Server-Sent Events (SSE), express-rate-limit (RateLimit)."
  :symbol-count=24
/>
## API reference

Select a symbol below — each page explains what it does, shows Node.js vs Go comparisons, and includes a runnable example.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/http/save-uploaded-file">SaveUploadedFile</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>SaveUploadedFile saves an uploaded file to destPath.</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/http/app">App</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>App is the root HTTP application (Express app).</td></tr>
<tr><td><a href="/packages/http/cors-options">CORSOptions</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>CORSOptions configures the CORS middleware.</td></tr>
<tr><td><a href="/packages/http/ctx">Ctx</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Ctx wraps a single HTTP request and response.</td></tr>
<tr><td><a href="/packages/http/event-stream">EventStream</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>EventStream sends Server-Sent Events.</td></tr>
<tr><td><a href="/packages/http/handler">Handler</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Handler is an Express-style handler that returns an error for centralized handling.</td></tr>
<tr><td><a href="/packages/http/memory-store">MemoryStore</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>MemoryStore is an in-memory session store (express-session memory).</td></tr>
<tr><td><a href="/packages/http/middleware">Middleware</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Middleware wraps the next handler in the chain.</td></tr>
<tr><td><a href="/packages/http/multipart-form">MultipartForm</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>MultipartForm wraps a parsed multipart form.</td></tr>
<tr><td><a href="/packages/http/rate-limit-options">RateLimitOptions</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>RateLimitOptions configures per-key rate limiting (express-rate-limit).</td></tr>
<tr><td><a href="/packages/http/session-options">SessionOptions</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>SessionOptions configures session middleware.</td></tr>
<tr><td><a href="/packages/http/session-store">SessionStore</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>SessionStore persists session data.</td></tr>
<tr><td><a href="/packages/http/ws-handler">WSHandler</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>WSHandler handles a WebSocket connection in an Express-style handler.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/http/middleware-cors">Middleware.CORS</a></td><td><code class="node-cell">cors</code></td><td><span class="kind-pill">method</span></td><td>CORS adds Cross-Origin Resource Sharing headers.</td></tr>
<tr><td><a href="/packages/http/middleware-logger">Middleware.Logger</a></td><td><code class="node-cell">morgan</code></td><td><span class="kind-pill">method</span></td><td>Logger logs each request with slog (morgan-style).</td></tr>
<tr><td><a href="/packages/http/app-new">App.New</a></td><td><code class="node-cell">express()</code></td><td><span class="kind-pill">method</span></td><td>New creates a new App with a fresh chi router.</td></tr>
<tr><td><a href="/packages/http/memory-store-new-memory-store">MemoryStore.NewMemoryStore</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>NewMemoryStore creates a memory session store.</td></tr>
<tr><td><a href="/packages/http/multipart-form-parse-multipart">MultipartForm.ParseMultipart</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>ParseMultipart parses a multipart request (multer).</td></tr>
<tr><td><a href="/packages/http/middleware-rate-limit">Middleware.RateLimit</a></td><td><code class="node-cell">express-rate-limit</code></td><td><span class="kind-pill">method</span></td><td>RateLimit returns middleware that limits requests per key.</td></tr>
<tr><td><a href="/packages/http/middleware-recover">Middleware.Recover</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>Recover catches panics and returns 500.</td></tr>
<tr><td><a href="/packages/http/event-stream-sse">EventStream.SSE</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>SSE prepares the response for server-sent events.</td></tr>
<tr><td><a href="/packages/http/handler-sse-handler">Handler.SSEHandler</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>SSEHandler runs fn with an EventStream (Express-style handler).</td></tr>
<tr><td><a href="/packages/http/middleware-security">Middleware.Security</a></td><td><code class="node-cell">helmet</code></td><td><span class="kind-pill">method</span></td><td>Security sets common security headers (helmet-style).</td></tr>
<tr><td><a href="/packages/http/middleware-session-middleware">Middleware.SessionMiddleware</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>SessionMiddleware loads and saves session data on each request.</td></tr>
</tbody></table>

