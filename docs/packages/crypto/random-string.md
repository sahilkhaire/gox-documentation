---
title: "RandomString"
package: "crypto"
import: "github.com/sahilkhaire/gox/crypto"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: crypto, bcrypt</span><span class="api-badge import">github.com/sahilkhaire/gox/crypto</span></div>
# RandomString

## Overview

RandomString returns a URL-safe base64 string with roughly n bytes of entropy
(uses n random bytes, encoded without p...

## Signature

```go
func RandomString(n int) (string, error)
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
import "github.com/sahilkhaire/gox/crypto"

// crypto
_ = crypto.RandomString(/* args */)
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
