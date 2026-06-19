---
title: "crypto"
package: "crypto"
gox-doc-version: "14"
---

<PackageOverview
  name="crypto"
  node="crypto, bcrypt"
  import-path="github.com/sahilkhaire/gox/crypto"
  subtitle="Package crypto provides Node crypto-style hashing, HMAC, random bytes, and bcrypt passwords. Node equivalent: crypto.createHash, crypto.randomBytes, bcrypt."
  :symbol-count=6
/>
## API reference

Select a symbol below — each page explains what it does, shows Node.js vs Go comparisons, and includes a runnable example.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/crypto/check-password">CheckPassword</a></td><td><code class="node-cell">bcrypt.compare(password, hash)</code></td><td><span class="kind-pill">func</span></td><td>CheckPassword reports whether password matches the bcrypt hash.</td></tr>
<tr><td><a href="/packages/crypto/hmacsha256">HMACSHA256</a></td><td><code class="node-cell">crypto.createHmac('sha256', key).update(d).digest('hex')</code></td><td><span class="kind-pill">func</span></td><td>HMACSHA256 returns HMAC-SHA256 of data with key as lowercase hex.</td></tr>
<tr><td><a href="/packages/crypto/hash-password">HashPassword</a></td><td><code class="node-cell">bcrypt.hash(password, 10)</code></td><td><span class="kind-pill">func</span></td><td>HashPassword hashes password with bcrypt (bcrypt.hash).</td></tr>
<tr><td><a href="/packages/crypto/random-bytes">RandomBytes</a></td><td><code class="node-cell">crypto.randomBytes(n)</code></td><td><span class="kind-pill">func</span></td><td>RandomBytes returns n cryptographically secure random bytes (crypto.randomBytes).</td></tr>
<tr><td><a href="/packages/crypto/random-string">RandomString</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>RandomString returns a URL-safe base64 string with roughly n bytes of entropy (uses n random bytes, encoded without p...</td></tr>
<tr><td><a href="/packages/crypto/sha256">SHA256</a></td><td><code class="node-cell">crypto.createHash('sha256').update(d).digest('hex')</code></td><td><span class="kind-pill">func</span></td><td>SHA256 returns the SHA-256 digest of data as lowercase hex (crypto.createHash('sha256')).</td></tr>
</tbody></table>

