---
title: "BuildMIME"
package: "mail"
import: "github.com/sahilkhaire/gox/mail"
gox-doc-version: "10"
---

<SymbolHeader pkg="mail" title="BuildMIME" node="nodemailer" import-path="github.com/sahilkhaire/gox/mail" />
## Overview

BuildMIME exposes MIME encoding for tests and custom transports.

## Signature

<div class="signature-block">

```go
func BuildMIME(msg Message) []byte
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

// mail
_ = mail.BuildMIME(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/mail/send">Send</a>
</div>

← [Back to mail package overview](/packages/mail/)
