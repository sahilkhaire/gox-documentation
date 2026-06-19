---
title: "Decode"
package: "jwt"
import: "github.com/sahilkhaire/gox/jwt"
node: "jwt.decode(token)"
gox-doc-version: "10"
---

<SymbolHeader pkg="jwt" title="Decode" node="jwt.decode(token)" import-path="github.com/sahilkhaire/gox/jwt" />
## Overview

Decode parses claims without verifying the signature (jwt.decode).

**Node.js equivalent:** `jwt.decode(token)`

## Signature

<div class="signature-block">

```go
func Decode(token string) (jwtlib.MapClaims, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
jwt.decode(token)
```

```go [Standard Go]
// github.com/golang-jwt/jwt/v5 token.Sign / Parse
```

```go [gox]
import "github.com/sahilkhaire/gox/jwt"

jwt.Decode(token)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/jwt/sign">Sign</a><a class="related-chip" href="/packages/jwt/verify">Verify</a>
</div>

← [Back to jwt package overview](/packages/jwt/)
