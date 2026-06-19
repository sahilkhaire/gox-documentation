---
title: "CheckPassword"
package: "crypto"
import: "github.com/sahilkhaire/gox/crypto"
node: "bcrypt.compare(password, hash)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: bcrypt.compare(password, hash)</span><span class="api-badge import">github.com/sahilkhaire/gox/crypto</span></div>
# CheckPassword

## Overview

Maps the Node.js pattern `bcrypt.compare(password, hash)` to gox `crypto.CheckPassword(password, hash)`.

**Node.js equivalent:** `bcrypt.compare(password, hash)`

## Signature

```go
func CheckPassword(password, hash string) bool
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
bcrypt.compare(password, hash)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/crypto"

crypto.CheckPassword(password, hash)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [HMACSHA256](/packages/crypto/hmacsha256)
- [HashPassword](/packages/crypto/hash-password)
- [RandomBytes](/packages/crypto/random-bytes)

← [Back to crypto package overview](/packages/crypto/)
