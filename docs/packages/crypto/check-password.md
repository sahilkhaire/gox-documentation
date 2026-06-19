---
title: "CheckPassword"
package: "crypto"
import: "github.com/sahilkhaire/gox/crypto"
node: "bcrypt.compare(password, hash)"
gox-doc-version: "11"
---

<SymbolHeader pkg="crypto" title="CheckPassword" node="bcrypt.compare(password, hash)" import-path="github.com/sahilkhaire/gox/crypto" />
## Overview

CheckPassword reports whether password matches the bcrypt hash.

If you are coming from Node.js, the closest pattern is **`bcrypt.compare(password, hash)`**.

## Signature

<div class="signature-block">

```go
func CheckPassword(password, hash string) bool
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
bcrypt.compare(password, hash)
```

```go [Standard Go]
hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
```

```go [gox]
import "github.com/sahilkhaire/gox/crypto"

crypto.CheckPassword(password, hash)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/crypto"

crypto.CheckPassword(password, hash)
```

## Tips

Import `github.com/sahilkhaire/gox/crypto` and call `CheckPassword` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/crypto/hmacsha256">HMACSHA256</a><a class="related-chip" href="/packages/crypto/hash-password">HashPassword</a><a class="related-chip" href="/packages/crypto/random-bytes">RandomBytes</a>
</div>

← [Back to crypto package overview](/packages/crypto/)
