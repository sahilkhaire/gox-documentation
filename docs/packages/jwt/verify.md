---
title: "Verify"
package: "jwt"
import: "github.com/sahilkhaire/gox/jwt"
node: "jwt.verify(token, secret)"
gox-doc-version: "11"
---

<SymbolHeader pkg="jwt" title="Verify" node="jwt.verify(token, secret)" import-path="github.com/sahilkhaire/gox/jwt" />
## Overview

Verify parses and validates a JWT with secret (jwt.verify).

If you are coming from Node.js, the closest pattern is **`jwt.verify(token, secret)`**.

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

## Example

```go
import "github.com/sahilkhaire/gox/jwt"

claims, err := jwt.Verify(token, secret)
```

## Tips

Import `github.com/sahilkhaire/gox/jwt` and call `Verify` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/jwt/decode">Decode</a><a class="related-chip" href="/packages/jwt/sign">Sign</a>
</div>

← [Back to jwt package overview](/packages/jwt/)
