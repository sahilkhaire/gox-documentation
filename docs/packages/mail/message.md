---
title: "Message"
package: "mail"
import: "github.com/sahilkhaire/gox/mail"
gox-doc-version: "10"
---

<SymbolHeader pkg="mail" title="Message" node="nodemailer" import-path="github.com/sahilkhaire/gox/mail" />
## Overview

Message is an outbound email.

## Signature

<div class="signature-block">

```go
type Message struct {
	From    string
	To      []string
	Subject string
	Body    string
	HTML    string
}
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
smtp.SendMail(addr, auth, from, to, msg)
```

```go [gox]
import "github.com/sahilkhaire/gox/mail"

_ = mail.Message
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/mail/build-mime">BuildMIME</a><a class="related-chip" href="/packages/mail/send">Send</a>
</div>

← [Back to mail package overview](/packages/mail/)
