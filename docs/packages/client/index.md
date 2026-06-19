---
title: "client"
package: "client"
gox-doc-version: "10"
---

<PackageOverview
  name="client"
  node="axios, fetch"
  import-path="github.com/sahilkhaire/gox/client"
  subtitle="Package client provides an axios/fetch-style HTTP client over net/http. Node equivalent: axios, node-fetch."
  :symbol-count=10
  :has-cookbook=true
  migration-link="/migration/axios"
  narrative="axios/fetch-style HTTP client with query params, headers, JSON bodies, and response helpers."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/client/client"><div class="featured-name">Get</div><div class="featured-summary">axios.get</div></a>
<a class="featured-card" href="/packages/client/response-fetch"><div class="featured-name">Fetch</div><div class="featured-summary">fetch()</div></a>
<a class="featured-card" href="/packages/client/client-new"><div class="featured-name">New</div><div class="featured-summary">axios.create</div></a>
</div>

## Common use cases

- Call third-party APIs
- POST JSON with typed responses
- Set default client for services

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>axios.create()</code></td><td><a href="/packages/client/new"><code>client.New()</code></a></td></tr>
<tr><td><code>axios.get/post/...</code></td><td><a href="/packages/client/get/post/put/patch/delete"><code>client.Get/Post/Put/Patch/Delete(ctx, url, opts)</code></a></td></tr>
<tr><td><code>axios.request(config)</code></td><td><a href="/packages/client/request"><code>client.Request(ctx, opts)</code></a></td></tr>
<tr><td><code>fetch(url)</code></td><td><a href="/packages/client/fetch"><code>client.Fetch(ctx, url, opts)</code></a></td></tr>
<tr><td><code>response.status</code></td><td><a href="/packages/client/status-code"><code>resp.StatusCode()</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/client"

res, err := client.Get(ctx, "https://api.example.com/users")
```

::: tip Full cookbook
See the [**client cookbook**](/packages/client/cookbook) for multi-step recipes and real-world patterns.
:::

::: info Migration guide
Coming from Node.js? Read the [**migration guide**](/migration/axios) for side-by-side patterns.
:::

## API reference

Browse **10 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/client/set-default-client">SetDefaultClient</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>SetDefaultClient replaces the client used by Fetch.</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/client/client">Client</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Client is a reusable HTTP client with defaults (axios instance).</td></tr>
<tr><td><a href="/packages/client/request-opts">RequestOpts</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>RequestOpts configures a single request.</td></tr>
<tr><td><a href="/packages/client/response">Response</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Response wraps an HTTP response.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/client/response-fetch">Response.Fetch</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>Fetch is a fetch()-style helper using the package default client.</td></tr>
<tr><td><a href="/packages/client/request-opts-json-body">RequestOpts.JSONBody</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>JSONBody returns RequestOpts with a JSON body.</td></tr>
<tr><td><a href="/packages/client/client-new">Client.New</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>New creates a Client with optional timeout.</td></tr>
<tr><td><a href="/packages/client/request-opts-with-headers">RequestOpts.WithHeaders</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>WithHeaders returns RequestOpts with headers.</td></tr>
<tr><td><a href="/packages/client/request-opts-with-method">RequestOpts.WithMethod</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>WithMethod sets method on opts (helper for Fetch).</td></tr>
<tr><td><a href="/packages/client/request-opts-with-query">RequestOpts.WithQuery</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>WithQuery returns RequestOpts with query values.</td></tr>
</tbody></table>

