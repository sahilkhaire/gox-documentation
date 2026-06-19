---
title: "BuildMIME"
package: "mail"
import: "github.com/sahilkhaire/gox/mail"
gox-doc-version: "14"
---

<SymbolHeader pkg="mail" title="BuildMIME" node="nodemailer" import-path="github.com/sahilkhaire/gox/mail" />
## Overview

BuildMIME exposes MIME encoding for tests and custom transports.

Part of the **`mail`** package — Node.js analog: *nodemailer*.

## Signature

<div class="signature-block">

```go
func BuildMIME(msg Message) []byte
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

raw := mail.BuildMIME(msg)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/mail"

raw := mail.BuildMIME(msg)
```

## Tips

Import `github.com/sahilkhaire/gox/mail` and call `BuildMIME` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
smtp.SendMail(addr, auth, from, to, msg)
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/mail/send">Send</a>
</div>

← [Back to mail package overview](/packages/mail/)
