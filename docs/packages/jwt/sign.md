---
title: "Sign"
package: "jwt"
import: "github.com/sahilkhaire/gox/jwt"
node: "jwt.sign(payload, secret, { expiresIn })"
gox-doc-version: "11"
---

<SymbolHeader pkg="jwt" title="Sign" node="jwt.sign(payload, secret, { expiresIn })" import-path="github.com/sahilkhaire/gox/jwt" />
## Overview

Sign builds a signed JWT from claims (jwt.sign).

If you are coming from Node.js, the closest pattern is **`jwt.sign(payload, secret, { expiresIn })`**.

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

## Example

```go
import "github.com/sahilkhaire/gox/jwt"

token, err := jwt.Sign(claims, secret, jwt.SignOptions{})
```

## Tips

Import `github.com/sahilkhaire/gox/jwt` and call `Sign` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/jwt/decode">Decode</a><a class="related-chip" href="/packages/jwt/verify">Verify</a>
</div>

← [Back to jwt package overview](/packages/jwt/)
