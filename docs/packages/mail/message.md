---
title: "Message"
package: "mail"
import: "github.com/sahilkhaire/gox/mail"
gox-doc-version: "11"
---

<SymbolHeader pkg="mail" title="Message" node="nodemailer" import-path="github.com/sahilkhaire/gox/mail" />
## Overview

Message is an outbound email.

Part of the **`mail`** package — Node.js analog: *nodemailer*.

`Message` is a type exported by gox. Methods on this type are documented separately.

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
// Typical nodemailer pattern in Node.js
```

```go [Standard Go]
smtp.SendMail(addr, auth, from, to, msg)
```

```go [gox]
import "github.com/sahilkhaire/gox/mail"

_ = mail.Message
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/mail"

_ = mail.Message
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
