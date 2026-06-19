---
title: "Decode"
package: "jwt"
import: "github.com/sahilkhaire/gox/jwt"
node: "jwt.decode(token)"
gox-doc-version: "14"
---

<SymbolHeader pkg="jwt" title="Decode" node="jwt.decode(token)" import-path="github.com/sahilkhaire/gox/jwt" />
## Overview

Decode parses claims without verifying the signature (jwt.decode).

If you are coming from Node.js, the closest pattern is **`jwt.decode(token)`**.

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

## Example

```go
import "github.com/sahilkhaire/gox/jwt"

jwt.Decode(token)
```

## Tips

Import `github.com/sahilkhaire/gox/jwt` and call `Decode` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/jwt/sign">Sign</a><a class="related-chip" href="/packages/jwt/verify">Verify</a>
</div>

← [Back to jwt package overview](/packages/jwt/)
