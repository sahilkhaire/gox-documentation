---
title: "HMACSHA256"
package: "crypto"
import: "github.com/sahilkhaire/gox/crypto"
node: "crypto.createHmac('sha256', key).update(d).digest('hex')"
gox-doc-version: "10"
---

<SymbolHeader pkg="crypto" title="HMACSHA256" node="crypto.createHmac('sha256', key).update(d).digest('hex')" import-path="github.com/sahilkhaire/gox/crypto" />
## Overview

HMACSHA256 returns HMAC-SHA256 of data with key as lowercase hex.

**Node.js equivalent:** `crypto.createHmac('sha256', key).update(d).digest('hex')`

## Signature

<div class="signature-block">

```go
func HMACSHA256(data, key []byte) string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
crypto.createHmac('sha256', key).update(d).digest('hex')
```

```go [Standard Go]
h := sha256.Sum256(data)
// or hmac.New(sha256.New, key)
```

```go [gox]
import "github.com/sahilkhaire/gox/crypto"

crypto.HMACSHA256(d, key)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/crypto/check-password">CheckPassword</a><a class="related-chip" href="/packages/crypto/hash-password">HashPassword</a><a class="related-chip" href="/packages/crypto/random-bytes">RandomBytes</a>
</div>

← [Back to crypto package overview](/packages/crypto/)
