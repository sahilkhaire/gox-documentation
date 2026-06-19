---
title: "HashPassword"
package: "crypto"
import: "github.com/sahilkhaire/gox/crypto"
node: "bcrypt.hash(password, 10)"
gox-doc-version: "10"
---

<SymbolHeader pkg="crypto" title="HashPassword" node="bcrypt.hash(password, 10)" import-path="github.com/sahilkhaire/gox/crypto" />
## Overview

HashPassword hashes password with bcrypt (bcrypt.hash).

**Node.js equivalent:** `bcrypt.hash(password, 10)`

## Signature

<div class="signature-block">

```go
func HashPassword(password string) (string, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
await bcrypt.hash(password, 10);
```

```go [Standard Go]
hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
```

```go [gox]
import "github.com/sahilkhaire/gox/crypto"

hash, err := crypto.HashPassword(password)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/crypto/check-password">CheckPassword</a><a class="related-chip" href="/packages/crypto/hmacsha256">HMACSHA256</a><a class="related-chip" href="/packages/crypto/random-bytes">RandomBytes</a>
</div>

← [Back to crypto package overview](/packages/crypto/)
