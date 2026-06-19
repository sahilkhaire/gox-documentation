---
title: "crypto"
package: "crypto"
gox-doc-version: "10"
---

<PackageOverview
  name="crypto"
  node="crypto, bcrypt"
  import-path="github.com/sahilkhaire/gox/crypto"
  subtitle="Package crypto provides Node crypto-style hashing, HMAC, random bytes, and bcrypt passwords. Node equivalent: crypto.createHash, crypto.randomBytes, bcrypt."
  :symbol-count=6
  :has-cookbook=true
  migration-link=""
  narrative="Node crypto and bcrypt patterns — hashing, HMAC, random bytes, and password helpers."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/crypto/sha256"><div class="featured-name">SHA256</div><div class="featured-summary">createHash</div></a>
<a class="featured-card" href="/packages/crypto/hash-password"><div class="featured-name">HashPassword</div><div class="featured-summary">bcrypt.hash</div></a>
<a class="featured-card" href="/packages/crypto/random-string"><div class="featured-name">RandomString</div><div class="featured-summary">crypto.randomBytes</div></a>
</div>

## Common use cases

- Hash passwords for storage
- Sign webhook payloads
- Generate secure tokens

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>crypto.createHash('sha256').update(d).digest('hex')</code></td><td><a href="/packages/crypto/sha256"><code>crypto.SHA256(d)</code></a></td></tr>
<tr><td><code>crypto.createHmac('sha256', key).update(d).digest('hex')</code></td><td><a href="/packages/crypto/hmacsha256"><code>crypto.HMACSHA256(d, key)</code></a></td></tr>
<tr><td><code>crypto.randomBytes(n)</code></td><td><a href="/packages/crypto/random-bytes"><code>crypto.RandomBytes(n)</code></a></td></tr>
<tr><td><code>bcrypt.hash(password, 10)</code></td><td><a href="/packages/crypto/hash-password"><code>crypto.HashPassword(password)</code></a></td></tr>
<tr><td><code>bcrypt.compare(password, hash)</code></td><td><a href="/packages/crypto/check-password"><code>crypto.CheckPassword(password, hash)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/crypto"

hash, err := crypto.HashPassword(password)
```

::: tip Full cookbook
See the [**crypto cookbook**](/packages/crypto/cookbook) for multi-step recipes and real-world patterns.
:::

## API reference

Browse **6 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/crypto/check-password">CheckPassword</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>CheckPassword reports whether password matches the bcrypt hash.</td></tr>
<tr><td><a href="/packages/crypto/hmacsha256">HMACSHA256</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>HMACSHA256 returns HMAC-SHA256 of data with key as lowercase hex.</td></tr>
<tr><td><a href="/packages/crypto/hash-password">HashPassword</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>HashPassword hashes password with bcrypt (bcrypt.hash).</td></tr>
<tr><td><a href="/packages/crypto/random-bytes">RandomBytes</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>RandomBytes returns n cryptographically secure random bytes (crypto.randomBytes).</td></tr>
<tr><td><a href="/packages/crypto/random-string">RandomString</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>RandomString returns a URL-safe base64 string with roughly n bytes of entropy (uses n random bytes, encoded without p...</td></tr>
<tr><td><a href="/packages/crypto/sha256">SHA256</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>SHA256 returns the SHA-256 digest of data as lowercase hex (crypto.createHash('sha256')).</td></tr>
</tbody></table>

