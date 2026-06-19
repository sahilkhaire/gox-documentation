---
title: "Verify"
package: "jwt"
import: "github.com/sahilkhaire/gox/jwt"
node: "jwt.verify(token, secret)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: jwt.verify(token, secret)</span><span class="api-badge import">github.com/sahilkhaire/gox/jwt</span></div>
# Verify

## Overview

Verify parses and validates a JWT with secret (jwt.verify).

**Node.js equivalent:** `jwt.verify(token, secret)`

## Signature

```go
func Verify(token string, secret []byte) (jwtlib.MapClaims, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const payload = jwt.verify(token, secret);
```

```go [Standard Go]
// use golang-jwt directly
```

```go [gox]
import "github.com/sahilkhaire/gox/jwt"

claims, err := jwt.Verify(token, secret)
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

← [Back to jwt package overview](/packages/jwt/)
