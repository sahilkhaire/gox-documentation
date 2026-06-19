---
title: "mail"
package: "mail"
gox-doc-version: "10"
---

<PackageOverview
  name="mail"
  node="nodemailer"
  import-path="github.com/sahilkhaire/gox/mail"
  subtitle="Package mail sends email via SMTP, similar to nodemailer. Node equivalent: nodemailer"
  :symbol-count=4
  :has-cookbook=false
  migration-link=""
  narrative="nodemailer-style SMTP sending and MIME building."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/mail/send"><div class="featured-name">Send</div><div class="featured-summary">nodemailer.send</div></a>
<a class="featured-card" href="/packages/mail/build-mime"><div class="featured-name">BuildMIME</div><div class="featured-summary">Build MIME</div></a>
<a class="featured-card" href="/packages/mail/message"><div class="featured-name">Message</div><div class="featured-summary">Email message</div></a>
</div>

## Common use cases

- Send transactional email
- Build multipart messages
- Configure SMTP transport

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>sendMail</code></td><td><a href="/packages/mail/send"><code>mail.Send(ctx, msg, SMTPConfig)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/mail"

err := mail.Send(ctx, cfg, msg)
```

## API reference

Browse **4 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/mail/build-mime">BuildMIME</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>BuildMIME exposes MIME encoding for tests and custom transports.</td></tr>
<tr><td><a href="/packages/mail/send">Send</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Send delivers msg using SMTP.</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/mail/message">Message</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Message is an outbound email.</td></tr>
<tr><td><a href="/packages/mail/smtp-config">SMTPConfig</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>SMTPConfig holds SMTP connection settings.</td></tr>
</tbody></table>

