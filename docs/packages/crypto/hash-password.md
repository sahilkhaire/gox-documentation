---
title: "HashPassword"
package: "crypto"
import: "github.com/sahilkhaire/gox/crypto"
node: "bcrypt.hash(password, 10)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: bcrypt.hash(password, 10)</span><span class="api-badge import">github.com/sahilkhaire/gox/crypto</span></div>
# HashPassword

## Overview

HashPassword hashes password with bcrypt (bcrypt.hash).

**Node.js equivalent:** `bcrypt.hash(password, 10)`

## Signature

```go
func HashPassword(password string) (string, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
await bcrypt.hash(password, 10);
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/crypto"

hash, err := crypto.HashPassword(password)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [CheckPassword](/packages/crypto/check-password)
- [HMACSHA256](/packages/crypto/hmacsha256)
- [RandomBytes](/packages/crypto/random-bytes)

← [Back to crypto package overview](/packages/crypto/)
