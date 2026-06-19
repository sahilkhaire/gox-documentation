---
title: "HMACSHA256"
package: "crypto"
import: "github.com/sahilkhaire/gox/crypto"
node: "crypto.createHmac('sha256', key).update(d).digest('hex')"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: crypto.createHmac('sha256', key).update(d).digest('hex')</span><span class="api-badge import">github.com/sahilkhaire/gox/crypto</span></div>
# HMACSHA256

## Overview

Maps the Node.js pattern `crypto.createHmac('sha256', key).update(d).digest('hex')` to gox `crypto.HMACSHA256(d, key)`.

**Node.js equivalent:** `crypto.createHmac('sha256', key).update(d).digest('hex')`

## Signature

```go
func HMACSHA256(data, key []byte) string
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
crypto.createHmac('sha256', key).update(d).digest('hex')
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/crypto"

crypto.HMACSHA256(d, key)
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
- [HashPassword](/packages/crypto/hash-password)
- [RandomBytes](/packages/crypto/random-bytes)

← [Back to crypto package overview](/packages/crypto/)
