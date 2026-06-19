---
title: "http"
package: "http"
gox-doc-version: "7"
---

<div class="package-hero"><div class="api-meta"><span class="api-badge node">Node: express, cors, helmet, morgan</span></div><h1>http</h1>
<p class="subtitle">Package http provides an Express-like HTTP server on chi. Node equivalent: express, cors, helmet, morgan (Logger middleware), multer (ParseMultipart), express-session (SessionMiddleware), Server-Sent Events (SSE), express-rate-limit (RateLimit).</p><div class="import-line">import "github.com/sahilkhaire/gox/http"</div></div>
## Quick start

A minimal example to get going:

```go
import "github.com/sahilkhaire/gox/http"

app := goxhttp.New()
app.Get("/health", func(c *goxhttp.Ctx) error {
    return c.JSON(200, map[string]string{"status": "ok"})
})
```

::: tip Full cookbook
See the [**http cookbook**](/packages/http/cookbook) for multi-step recipes and real-world patterns.
:::

## What's in this package

If you have used **express, cors, helmet, morgan** in Node.js, this package is your starting point. Browse **24 documented symbols** below — each page includes Node.js, standard Go, and gox side-by-side examples.

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/http/save-uploaded-file">SaveUploadedFile</a></td><td><span class="kind-pill">func</span></td><td>SaveUploadedFile saves an uploaded file to destPath.</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/http/app">App</a></td><td><span class="kind-pill">type</span></td><td>App is the root HTTP application (Express app).</td></tr>
<tr><td><a href="/packages/http/cors-options">CORSOptions</a></td><td><span class="kind-pill">type</span></td><td>CORSOptions configures the CORS middleware.</td></tr>
<tr><td><a href="/packages/http/ctx">Ctx</a></td><td><span class="kind-pill">type</span></td><td>Ctx wraps a single HTTP request and response.</td></tr>
<tr><td><a href="/packages/http/event-stream">EventStream</a></td><td><span class="kind-pill">type</span></td><td>EventStream sends Server-Sent Events.</td></tr>
<tr><td><a href="/packages/http/handler">Handler</a></td><td><span class="kind-pill">type</span></td><td>Handler is an Express-style handler that returns an error for centralized handling.</td></tr>
<tr><td><a href="/packages/http/memory-store">MemoryStore</a></td><td><span class="kind-pill">type</span></td><td>MemoryStore is an in-memory session store (express-session memory).</td></tr>
<tr><td><a href="/packages/http/middleware">Middleware</a></td><td><span class="kind-pill">type</span></td><td>Middleware wraps the next handler in the chain.</td></tr>
<tr><td><a href="/packages/http/multipart-form">MultipartForm</a></td><td><span class="kind-pill">type</span></td><td>MultipartForm wraps a parsed multipart form.</td></tr>
<tr><td><a href="/packages/http/rate-limit-options">RateLimitOptions</a></td><td><span class="kind-pill">type</span></td><td>RateLimitOptions configures per-key rate limiting (express-rate-limit).</td></tr>
<tr><td><a href="/packages/http/session-options">SessionOptions</a></td><td><span class="kind-pill">type</span></td><td>SessionOptions configures session middleware.</td></tr>
<tr><td><a href="/packages/http/session-store">SessionStore</a></td><td><span class="kind-pill">type</span></td><td>SessionStore persists session data.</td></tr>
<tr><td><a href="/packages/http/ws-handler">WSHandler</a></td><td><span class="kind-pill">type</span></td><td>WSHandler handles a WebSocket connection in an Express-style handler.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/http/middleware-cors">Middleware.CORS</a></td><td><span class="kind-pill">method</span></td><td>CORS adds Cross-Origin Resource Sharing headers.</td></tr>
<tr><td><a href="/packages/http/middleware-logger">Middleware.Logger</a></td><td><span class="kind-pill">method</span></td><td>Logger logs each request with slog (morgan-style).</td></tr>
<tr><td><a href="/packages/http/app-new">App.New</a></td><td><span class="kind-pill">method</span></td><td>New creates a new App with a fresh chi router.</td></tr>
<tr><td><a href="/packages/http/memory-store-new-memory-store">MemoryStore.NewMemoryStore</a></td><td><span class="kind-pill">method</span></td><td>NewMemoryStore creates a memory session store.</td></tr>
<tr><td><a href="/packages/http/multipart-form-parse-multipart">MultipartForm.ParseMultipart</a></td><td><span class="kind-pill">method</span></td><td>ParseMultipart parses a multipart request (multer).</td></tr>
<tr><td><a href="/packages/http/middleware-rate-limit">Middleware.RateLimit</a></td><td><span class="kind-pill">method</span></td><td>RateLimit returns middleware that limits requests per key.</td></tr>
<tr><td><a href="/packages/http/middleware-recover">Middleware.Recover</a></td><td><span class="kind-pill">method</span></td><td>Recover catches panics and returns 500.</td></tr>
<tr><td><a href="/packages/http/event-stream-sse">EventStream.SSE</a></td><td><span class="kind-pill">method</span></td><td>SSE prepares the response for server-sent events.</td></tr>
<tr><td><a href="/packages/http/handler-sse-handler">Handler.SSEHandler</a></td><td><span class="kind-pill">method</span></td><td>SSEHandler runs fn with an EventStream (Express-style handler).</td></tr>
<tr><td><a href="/packages/http/middleware-security">Middleware.Security</a></td><td><span class="kind-pill">method</span></td><td>Security sets common security headers (helmet-style).</td></tr>
<tr><td><a href="/packages/http/middleware-session-middleware">Middleware.SessionMiddleware</a></td><td><span class="kind-pill">method</span></td><td>SessionMiddleware loads and saves session data on each request.</td></tr>
</tbody></table>

