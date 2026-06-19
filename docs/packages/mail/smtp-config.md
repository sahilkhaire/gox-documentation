---
title: "SMTPConfig"
package: "mail"
import: "github.com/sahilkhaire/gox/mail"
gox-doc-version: "11"
---

<SymbolHeader pkg="mail" title="SMTPConfig" node="nodemailer" import-path="github.com/sahilkhaire/gox/mail" />
## Overview

SMTPConfig holds SMTP connection settings.

Part of the **`mail`** package — Node.js analog: *nodemailer*.

`SMTPConfig` is a type exported by gox. Methods on this type are documented separately.

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
// Typical nodemailer pattern in Node.js
```

```go [Standard Go]
smtp.SendMail(addr, auth, from, to, msg)
```

```go [gox]
import "github.com/sahilkhaire/gox/mail"

_ = mail.SMTPConfig
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/mail"

_ = mail.SMTPConfig
```

## Tips

Browse methods on this type in the sidebar for handler-style APIs and options structs.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/mail/build-mime">BuildMIME</a><a class="related-chip" href="/packages/mail/send">Send</a>
</div>

← [Back to mail package overview](/packages/mail/)
