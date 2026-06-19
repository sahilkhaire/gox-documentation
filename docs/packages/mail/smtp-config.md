---
title: "SMTPConfig"
package: "mail"
import: "github.com/sahilkhaire/gox/mail"
gox-doc-version: "10"
---

<SymbolHeader pkg="mail" title="SMTPConfig" node="nodemailer" import-path="github.com/sahilkhaire/gox/mail" />
## Overview

SMTPConfig holds SMTP connection settings.

## Signature

<div class="signature-block">

```go
type SMTPConfig struct {
	Host string
	Port int
	User string
	Pass string
	TLS  bool // use implicit TLS (SMTPS) when true
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

_ = mail.SMTPConfig
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/mail/build-mime">BuildMIME</a><a class="related-chip" href="/packages/mail/send">Send</a>
</div>

← [Back to mail package overview](/packages/mail/)
