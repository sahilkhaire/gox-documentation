---
title: "Send"
package: "mail"
import: "github.com/sahilkhaire/gox/mail"
node: "sendMail"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: sendMail</span><span class="api-badge import">github.com/sahilkhaire/gox/mail</span></div>
# Send

## Overview

Maps the Node.js pattern `sendMail` to gox `mail.Send(ctx, msg, SMTPConfig)`.

**Node.js equivalent:** `sendMail`

## Signature

```go
func Send(ctx context.Context, msg Message, cfg SMTPConfig) error
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
sendMail
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/mail"

mail.Send(ctx, msg, SMTPConfig)
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

← [Back to mail package overview](/packages/mail/)
