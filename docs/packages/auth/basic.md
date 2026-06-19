---
title: "Basic"
package: "auth"
import: "github.com/sahilkhaire/gox/auth"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: passport</span><span class="api-badge import">github.com/sahilkhaire/gox/auth</span></div>
# Basic

## Overview

Basic validates HTTP Basic credentials.

## Signature

```go
func Basic(validate func(user, pass string) bool) goxhttp.Middleware
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/auth"

// auth
_ = auth.Basic(/* args */)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [APIKey](/packages/auth/api-key)
- [Bearer](/packages/auth/bearer)
- [GetAPIKey](/packages/auth/get-api-key)

← [Back to auth package overview](/packages/auth/)
