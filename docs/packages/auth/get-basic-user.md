---
title: "GetBasicUser"
package: "auth"
import: "github.com/sahilkhaire/gox/auth"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: passport</span><span class="api-badge import">github.com/sahilkhaire/gox/auth</span></div>
# GetBasicUser

## Overview

GetBasicUser returns the username set by Basic middleware.

## Signature

```go
func GetBasicUser(c *goxhttp.Ctx) string
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
_ = auth.GetBasicUser(/* args */)
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
- [Basic](/packages/auth/basic)
- [Bearer](/packages/auth/bearer)

← [Back to auth package overview](/packages/auth/)
