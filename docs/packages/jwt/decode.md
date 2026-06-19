---
title: "Decode"
package: "jwt"
import: "github.com/sahilkhaire/gox/jwt"
node: "jwt.decode(token)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: jwt.decode(token)</span><span class="api-badge import">github.com/sahilkhaire/gox/jwt</span></div>
# Decode

## Overview

Maps the Node.js pattern `jwt.decode(token)` to gox `jwt.Decode(token)`.

**Node.js equivalent:** `jwt.decode(token)`

## Signature

```go
func Decode(token string) (jwtlib.MapClaims, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
jwt.decode(token)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/jwt"

jwt.Decode(token)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Sign](/packages/jwt/sign)
- [Verify](/packages/jwt/verify)

← [Back to jwt package overview](/packages/jwt/)
