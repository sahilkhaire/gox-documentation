---
title: "client"
package: "client"
gox-doc-version: "7"
---

<div class="package-hero"><div class="api-meta"><span class="api-badge node">Node: axios, fetch</span></div><h1>client</h1>
<p class="subtitle">Package client provides an axios/fetch-style HTTP client over net/http. Node equivalent: axios, node-fetch.</p><div class="import-line">import "github.com/sahilkhaire/gox/client"</div></div>
## Quick start

A minimal example to get going:

```go
import "github.com/sahilkhaire/gox/client"

res, err := client.Get(ctx, "https://api.example.com/users")
```

::: tip Full cookbook
See the [**client cookbook**](/packages/client/cookbook) for multi-step recipes and real-world patterns.
:::

## What's in this package

If you have used **axios, fetch** in Node.js, this package is your starting point. Browse **10 documented symbols** below — each page includes Node.js, standard Go, and gox side-by-side examples.

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/client/set-default-client">SetDefaultClient</a></td><td><span class="kind-pill">func</span></td><td>SetDefaultClient replaces the client used by Fetch.</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/client/client">Client</a></td><td><span class="kind-pill">type</span></td><td>Client is a reusable HTTP client with defaults (axios instance).</td></tr>
<tr><td><a href="/packages/client/request-opts">RequestOpts</a></td><td><span class="kind-pill">type</span></td><td>RequestOpts configures a single request.</td></tr>
<tr><td><a href="/packages/client/response">Response</a></td><td><span class="kind-pill">type</span></td><td>Response wraps an HTTP response.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/client/response-fetch">Response.Fetch</a></td><td><span class="kind-pill">method</span></td><td>Fetch is a fetch()-style helper using the package default client.</td></tr>
<tr><td><a href="/packages/client/request-opts-json-body">RequestOpts.JSONBody</a></td><td><span class="kind-pill">method</span></td><td>JSONBody returns RequestOpts with a JSON body.</td></tr>
<tr><td><a href="/packages/client/client-new">Client.New</a></td><td><span class="kind-pill">method</span></td><td>New creates a Client with optional timeout.</td></tr>
<tr><td><a href="/packages/client/request-opts-with-headers">RequestOpts.WithHeaders</a></td><td><span class="kind-pill">method</span></td><td>WithHeaders returns RequestOpts with headers.</td></tr>
<tr><td><a href="/packages/client/request-opts-with-method">RequestOpts.WithMethod</a></td><td><span class="kind-pill">method</span></td><td>WithMethod sets method on opts (helper for Fetch).</td></tr>
<tr><td><a href="/packages/client/request-opts-with-query">RequestOpts.WithQuery</a></td><td><span class="kind-pill">method</span></td><td>WithQuery returns RequestOpts with query values.</td></tr>
</tbody></table>

