---
title: "GetBasicUser"
package: "auth"
import: "github.com/sahilkhaire/gox/auth"
gox-doc-version: "11"
---

<SymbolHeader pkg="auth" title="GetBasicUser" node="passport" import-path="github.com/sahilkhaire/gox/auth" />
## Overview

GetBasicUser returns the username set by Basic middleware.

Part of the **`auth`** package — Node.js analog: *passport*.

## Signature

<div class="signature-block">

```go
func GetBasicUser(c *goxhttp.Ctx) string
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

// auth
_ = auth.GetBasicUser(/* args */)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/auth"

// auth
_ = auth.GetBasicUser(/* args */)
```

## Tips

Import `github.com/sahilkhaire/gox/auth` and call `GetBasicUser` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/auth/api-key">APIKey</a><a class="related-chip" href="/packages/auth/basic">Basic</a><a class="related-chip" href="/packages/auth/bearer">Bearer</a>
</div>

← [Back to auth package overview](/packages/auth/)
