---
title: "Verify"
package: "jwt"
import: "github.com/sahilkhaire/gox/jwt"
node: "jwt.verify(token, secret)"
gox-doc-version: "10"
---

<SymbolHeader pkg="jwt" title="Verify" node="jwt.verify(token, secret)" import-path="github.com/sahilkhaire/gox/jwt" />
## Overview

Verify parses and validates a JWT with secret (jwt.verify).

**Node.js equivalent:** `jwt.verify(token, secret)`

## Signature

<div class="signature-block">

```go
func Verify(token string, secret []byte) (jwtlib.MapClaims, error)
```

</div>

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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/jwt/decode">Decode</a><a class="related-chip" href="/packages/jwt/sign">Sign</a>
</div>

← [Back to jwt package overview](/packages/jwt/)
