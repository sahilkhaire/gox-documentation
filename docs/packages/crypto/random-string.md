---
title: "RandomString"
package: "crypto"
import: "github.com/sahilkhaire/gox/crypto"
gox-doc-version: "14"
---

<SymbolHeader pkg="crypto" title="RandomString" node="crypto, bcrypt" import-path="github.com/sahilkhaire/gox/crypto" />
## Overview

RandomString returns a URL-safe base64 string with roughly n bytes of entropy
(uses n random bytes, encoded without padding).

Part of the **`crypto`** package — Node.js analog: *crypto, bcrypt*.

## Signature

<div class="signature-block">

```go
func RandomString(n int) (string, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical crypto, bcrypt pattern in Node.js
```

```go [Standard Go]
b := make([]byte, n)
_, err := rand.Read(b)
```

```go [gox]
import "github.com/sahilkhaire/gox/crypto"

token, err := crypto.RandomString(32)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/crypto"

token, err := crypto.RandomString(32)
```

## Tips

Import `github.com/sahilkhaire/gox/crypto` and call `RandomString` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
b := make([]byte, n)
_, err := rand.Read(b)
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/crypto/check-password">CheckPassword</a><a class="related-chip" href="/packages/crypto/hmacsha256">HMACSHA256</a><a class="related-chip" href="/packages/crypto/hash-password">HashPassword</a>
</div>

← [Back to crypto package overview](/packages/crypto/)
