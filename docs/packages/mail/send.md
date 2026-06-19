---
title: "Send"
package: "mail"
import: "github.com/sahilkhaire/gox/mail"
node: "sendMail"
gox-doc-version: "14"
---

<SymbolHeader pkg="mail" title="Send" node="sendMail" import-path="github.com/sahilkhaire/gox/mail" />
## Overview

Send delivers msg using SMTP. ctx cancellation is honored before connecting.
Node: transporter.sendMail(msg)

If you are coming from Node.js, the closest pattern is **`sendMail`**.

## Signature

<div class="signature-block">

```go
func Send(ctx context.Context, msg Message, cfg SMTPConfig) error
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
sendMail
```

```go [Standard Go]
smtp.SendMail(addr, auth, from, to, msg)
```

```go [gox]
import "github.com/sahilkhaire/gox/mail"

mail.Send(ctx, msg, SMTPConfig)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/mail"

mail.Send(ctx, msg, SMTPConfig)
```

## Tips

Import `github.com/sahilkhaire/gox/mail` and call `Send` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
smtp.SendMail(addr, auth, from, to, msg)
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/mail/build-mime">BuildMIME</a>
</div>

← [Back to mail package overview](/packages/mail/)
