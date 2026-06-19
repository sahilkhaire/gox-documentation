---
title: "crypto"
package: "crypto"
gox-doc-version: "7"
---

<div class="package-hero"><div class="api-meta"><span class="api-badge node">Node: crypto, bcrypt</span></div><h1>crypto</h1>
<p class="subtitle">Package crypto provides Node crypto-style hashing, HMAC, random bytes, and bcrypt passwords. Node equivalent: crypto.createHash, crypto.randomBytes, bcrypt.</p><div class="import-line">import "github.com/sahilkhaire/gox/crypto"</div></div>
## What's in this package

If you have used **crypto, bcrypt** in Node.js, this package is your starting point. Browse **6 documented symbols** below — each page includes Node.js, standard Go, and gox side-by-side examples.

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/crypto/check-password">CheckPassword</a></td><td><span class="kind-pill">func</span></td><td>CheckPassword reports whether password matches the bcrypt hash.</td></tr>
<tr><td><a href="/packages/crypto/hmacsha256">HMACSHA256</a></td><td><span class="kind-pill">func</span></td><td>HMACSHA256 returns HMAC-SHA256 of data with key as lowercase hex.</td></tr>
<tr><td><a href="/packages/crypto/hash-password">HashPassword</a></td><td><span class="kind-pill">func</span></td><td>HashPassword hashes password with bcrypt (bcrypt.hash).</td></tr>
<tr><td><a href="/packages/crypto/random-bytes">RandomBytes</a></td><td><span class="kind-pill">func</span></td><td>RandomBytes returns n cryptographically secure random bytes (crypto.randomBytes).</td></tr>
<tr><td><a href="/packages/crypto/random-string">RandomString</a></td><td><span class="kind-pill">func</span></td><td>RandomString returns a URL-safe base64 string with roughly n bytes of entropy (uses n random bytes, encoded without p...</td></tr>
<tr><td><a href="/packages/crypto/sha256">SHA256</a></td><td><span class="kind-pill">func</span></td><td>SHA256 returns the SHA-256 digest of data as lowercase hex (crypto.createHash('sha256')).</td></tr>
</tbody></table>

