---
title: "Sign"
package: "jwt"
import: "github.com/sahilkhaire/gox/jwt"
node: "jwt.sign(payload, secret, { expiresIn })"
gox-doc-version: "10"
---

<SymbolHeader pkg="jwt" title="Sign" node="jwt.sign(payload, secret, { expiresIn })" import-path="github.com/sahilkhaire/gox/jwt" />
## Overview

Sign builds a signed JWT from claims (jwt.sign).

**Node.js equivalent:** `jwt.sign(payload, secret, { expiresIn })`

## Signature

<div class="signature-block">

```go
func Sign(claims jwtlib.MapClaims, secret []byte, opts *SignOptions) (string, error)
```

</div>

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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/jwt/decode">Decode</a><a class="related-chip" href="/packages/jwt/verify">Verify</a>
</div>

← [Back to jwt package overview](/packages/jwt/)
