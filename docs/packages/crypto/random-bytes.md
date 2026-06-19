---
title: "RandomBytes"
package: "crypto"
import: "github.com/sahilkhaire/gox/crypto"
node: "crypto.randomBytes(n)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: crypto.randomBytes(n)</span><span class="api-badge import">github.com/sahilkhaire/gox/crypto</span></div>
# RandomBytes

## Overview

Maps the Node.js pattern `crypto.randomBytes(n)` to gox `crypto.RandomBytes(n)`.

**Node.js equivalent:** `crypto.randomBytes(n)`

## Signature

```go
func RandomBytes(n int) ([]byte, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
crypto.randomBytes(n)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/crypto"

crypto.RandomBytes(n)
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
- [HashPassword](/packages/crypto/hash-password)

← [Back to crypto package overview](/packages/crypto/)
