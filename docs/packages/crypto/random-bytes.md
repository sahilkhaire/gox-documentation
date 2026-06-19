---
title: "RandomBytes"
package: "crypto"
import: "github.com/sahilkhaire/gox/crypto"
node: "crypto.randomBytes(n)"
gox-doc-version: "10"
---

<SymbolHeader pkg="crypto" title="RandomBytes" node="crypto.randomBytes(n)" import-path="github.com/sahilkhaire/gox/crypto" />
## Overview

RandomBytes returns n cryptographically secure random bytes (crypto.randomBytes).

**Node.js equivalent:** `crypto.randomBytes(n)`

## Signature

<div class="signature-block">

```go
func RandomBytes(n int) ([]byte, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
crypto.randomBytes(n)
```

```go [Standard Go]
b := make([]byte, n)
_, err := rand.Read(b)
```

```go [gox]
import "github.com/sahilkhaire/gox/crypto"

crypto.RandomBytes(n)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/crypto/check-password">CheckPassword</a><a class="related-chip" href="/packages/crypto/hmacsha256">HMACSHA256</a><a class="related-chip" href="/packages/crypto/hash-password">HashPassword</a>
</div>

← [Back to crypto package overview](/packages/crypto/)
