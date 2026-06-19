---
title: "mail"
package: "mail"
gox-doc-version: "11"
---

<PackageOverview
  name="mail"
  node="nodemailer"
  import-path="github.com/sahilkhaire/gox/mail"
  subtitle="Package mail sends email via SMTP, similar to nodemailer. Node equivalent: nodemailer"
  :symbol-count=4
/>
## API reference

Select a symbol below — each page explains what it does, shows Node.js vs Go comparisons, and includes a runnable example.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/mail/build-mime">BuildMIME</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>BuildMIME exposes MIME encoding for tests and custom transports.</td></tr>
<tr><td><a href="/packages/mail/send">Send</a></td><td><code class="node-cell">sendMail</code></td><td><span class="kind-pill">func</span></td><td>Send delivers msg using SMTP.</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/mail/message">Message</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Message is an outbound email.</td></tr>
<tr><td><a href="/packages/mail/smtp-config">SMTPConfig</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>SMTPConfig holds SMTP connection settings.</td></tr>
</tbody></table>

