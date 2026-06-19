---
title: "GetClaims"
package: "auth"
import: "github.com/sahilkhaire/gox/auth"
gox-doc-version: "14"
---

<SymbolHeader pkg="auth" title="GetClaims" node="passport" import-path="github.com/sahilkhaire/gox/auth" />
## Overview

GetClaims returns JWT claims set by Bearer middleware.

Part of the **`auth`** package — Node.js analog: *passport*.

## Signature

<div class="signature-block">

```go
func GetClaims(c *goxhttp.Ctx) jwtlib.MapClaims
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical passport pattern in Node.js
```

```go [Standard Go]
// net/http middleware checking Authorization header
```

```go [gox]
import "github.com/sahilkhaire/gox/auth"

claims, ok := auth.GetClaims(c)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/auth"

claims, ok := auth.GetClaims(c)
```

## Tips

Import `github.com/sahilkhaire/gox/auth` and call `GetClaims` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/auth/api-key">APIKey</a><a class="related-chip" href="/packages/auth/basic">Basic</a><a class="related-chip" href="/packages/auth/bearer">Bearer</a>
</div>

← [Back to auth package overview](/packages/auth/)
