---
title: "User.UserInfo"
package: "osutil"
import: "github.com/sahilkhaire/gox/osutil"
node: "os.userInfo()"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: os.userInfo()</span><span class="api-badge import">github.com/sahilkhaire/gox/osutil</span></div>
# User.UserInfo

## Overview

Maps the Node.js pattern `os.userInfo()` to gox `osutil.UserInfo()`.

**Node.js equivalent:** `os.userInfo()`

## Signature

```go
func UserInfo() (User, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
os.userInfo()
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/osutil"

osutil.UserInfo()
```

:::

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

```go
import "github.com/sahilkhaire/gox/osutil"

info, err := UserInfo()
```

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

← [Back to osutil package overview](/packages/osutil/)
