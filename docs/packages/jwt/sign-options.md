---
title: "SignOptions"
package: "jwt"
import: "github.com/sahilkhaire/gox/jwt"
gox-doc-version: "14"
---

<SymbolHeader pkg="jwt" title="SignOptions" node="jsonwebtoken" import-path="github.com/sahilkhaire/gox/jwt" />
## Overview

SignOptions configures token signing.

Part of the **`jwt`** package — Node.js analog: *jsonwebtoken*.

`SignOptions` is a type exported by gox. Methods on this type are documented separately.

## Signature

<div class="signature-block">

```go
type SignOptions struct {
	ExpiresIn time.Duration
	Method    jwtlib.SigningMethod
}
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical jsonwebtoken pattern in Node.js
```

```go [Standard Go]
// github.com/golang-jwt/jwt/v5 token.Sign / Parse
```

```go [gox]
import "github.com/sahilkhaire/gox/jwt"

opts := &jwt.SignOptions{ExpiresIn: time.Hour}
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/jwt"

opts := &jwt.SignOptions{ExpiresIn: time.Hour}
```

## Tips

Browse methods on this type in the sidebar for handler-style APIs and options structs.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/jwt/decode">Decode</a><a class="related-chip" href="/packages/jwt/sign">Sign</a><a class="related-chip" href="/packages/jwt/verify">Verify</a>
</div>

← [Back to jwt package overview](/packages/jwt/)
