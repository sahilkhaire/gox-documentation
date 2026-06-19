---
title: "Sign"
package: "jwt"
import: "github.com/sahilkhaire/gox/jwt"
node: "jwt.sign(payload, secret, { expiresIn })"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: jwt.sign(payload, secret, { expiresIn })</span><span class="api-badge import">github.com/sahilkhaire/gox/jwt</span></div>
# Sign

## Overview

Sign builds a signed JWT from claims (jwt.sign).

**Node.js equivalent:** `jwt.sign(payload, secret, { expiresIn })`

## Signature

```go
func Sign(claims jwtlib.MapClaims, secret []byte, opts *SignOptions) (string, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const token = jwt.sign(payload, secret);
```

```go [Standard Go]
// use golang-jwt directly
```

```go [gox]
import "github.com/sahilkhaire/gox/jwt"

token, err := jwt.Sign(claims, secret, jwt.SignOptions{})
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
- [Verify](/packages/jwt/verify)

← [Back to jwt package overview](/packages/jwt/)
