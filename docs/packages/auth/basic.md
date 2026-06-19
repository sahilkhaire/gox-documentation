---
title: "Basic"
package: "auth"
import: "github.com/sahilkhaire/gox/auth"
gox-doc-version: "14"
---

<SymbolHeader pkg="auth" title="Basic" node="passport" import-path="github.com/sahilkhaire/gox/auth" />
## Overview

Basic validates HTTP Basic credentials.

Part of the **`auth`** package — Node.js analog: *passport*.

## Signature

<div class="signature-block">

```go
func Basic(validate func(user, pass string) bool) goxhttp.Middleware
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

app.Use(auth.Basic(func(user, pass string) bool { return user == "admin" }))
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/auth"

app.Use(auth.Basic(func(user, pass string) bool { return user == "admin" }))
```

## Tips

Import `github.com/sahilkhaire/gox/auth` and call `Basic` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/auth/api-key">APIKey</a><a class="related-chip" href="/packages/auth/bearer">Bearer</a><a class="related-chip" href="/packages/auth/get-api-key">GetAPIKey</a>
</div>

← [Back to auth package overview](/packages/auth/)
