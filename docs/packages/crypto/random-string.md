---
title: "RandomString"
package: "crypto"
import: "github.com/sahilkhaire/gox/crypto"
gox-doc-version: "10"
---

<SymbolHeader pkg="crypto" title="RandomString" node="crypto, bcrypt" import-path="github.com/sahilkhaire/gox/crypto" />
## Overview

RandomString returns a URL-safe base64 string with roughly n bytes of entropy
(uses n random bytes, encoded without p...

## Signature

<div class="signature-block">

```go
func RandomString(n int) (string, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
b := make([]byte, n)
_, err := rand.Read(b)
```

```go [gox]
import "github.com/sahilkhaire/gox/crypto"

// crypto
_ = crypto.RandomString(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/crypto/check-password">CheckPassword</a><a class="related-chip" href="/packages/crypto/hmacsha256">HMACSHA256</a><a class="related-chip" href="/packages/crypto/hash-password">HashPassword</a>
</div>

← [Back to crypto package overview](/packages/crypto/)
