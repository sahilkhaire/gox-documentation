---
title: "SignOptions"
package: "jwt"
import: "github.com/sahilkhaire/gox/jwt"
gox-doc-version: "10"
---

<SymbolHeader pkg="jwt" title="SignOptions" node="jsonwebtoken" import-path="github.com/sahilkhaire/gox/jwt" />
## Overview

SignOptions configures token signing.

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
// See package overview
```

```go [Standard Go]
// github.com/golang-jwt/jwt/v5 token.Sign / Parse
```

```go [gox]
import "github.com/sahilkhaire/gox/jwt"

_ = jwt.SignOptions
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/jwt/decode">Decode</a><a class="related-chip" href="/packages/jwt/sign">Sign</a><a class="related-chip" href="/packages/jwt/verify">Verify</a>
</div>

← [Back to jwt package overview](/packages/jwt/)
