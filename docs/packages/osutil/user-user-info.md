---
title: "User.UserInfo"
package: "osutil"
import: "github.com/sahilkhaire/gox/osutil"
node: "os.userInfo()"
gox-doc-version: "10"
---

<SymbolHeader pkg="osutil" title="User.UserInfo" node="os.userInfo()" import-path="github.com/sahilkhaire/gox/osutil" />
## Overview

UserInfo returns the current user.

**Node.js equivalent:** `os.userInfo()`

## Signature

<div class="signature-block">

```go
func UserInfo() (User, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
os.userInfo()
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches.
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

← [Back to osutil package overview](/packages/osutil/)
