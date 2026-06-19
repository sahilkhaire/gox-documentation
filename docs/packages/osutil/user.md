---
title: "User"
package: "osutil"
import: "github.com/sahilkhaire/gox/osutil"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: os</span><span class="api-badge import">github.com/sahilkhaire/gox/osutil</span></div>
# User

## Overview

User holds user identity fields when available (os.userInfo).

## Signature

```go
type User struct {
	Username string
	Uid      string
	Gid      string
}
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
import "github.com/sahilkhaire/gox/osutil"

_ = osutil.User
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Arch](/packages/osutil/arch)
- [CPUs](/packages/osutil/cp-us)
- [Homedir](/packages/osutil/homedir)

← [Back to osutil package overview](/packages/osutil/)
