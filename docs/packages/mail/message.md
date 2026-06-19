---
title: "Message"
package: "mail"
import: "github.com/sahilkhaire/gox/mail"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: nodemailer</span><span class="api-badge import">github.com/sahilkhaire/gox/mail</span></div>
# Message

## Overview

Message is an outbound email.

## Signature

```go
type Message struct {
	From    string
	To      []string
	Subject string
	Body    string
	HTML    string
}
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/mail"

_ = mail.Message
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [BuildMIME](/packages/mail/build-mime)
- [Send](/packages/mail/send)

← [Back to mail package overview](/packages/mail/)
