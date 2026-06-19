---
title: "SignOptions"
package: "jwt"
import: "github.com/sahilkhaire/gox/jwt"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: jsonwebtoken</span><span class="api-badge import">github.com/sahilkhaire/gox/jwt</span></div>
# SignOptions

## Overview

SignOptions configures token signing.

## Signature

```go
type SignOptions struct {
	ExpiresIn time.Duration
	Method    jwtlib.SigningMethod
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
import "github.com/sahilkhaire/gox/jwt"

_ = jwt.SignOptions
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Decode](/packages/jwt/decode)
- [Sign](/packages/jwt/sign)
- [Verify](/packages/jwt/verify)

← [Back to jwt package overview](/packages/jwt/)
