---
title: "User.UserInfo"
package: "osutil"
import: "github.com/sahilkhaire/gox/osutil"
node: "os.userInfo()"
gox-doc-version: "14"
---

<SymbolHeader pkg="osutil" title="User.UserInfo" node="os.userInfo()" import-path="github.com/sahilkhaire/gox/osutil" />
## Overview

UserInfo returns the current user.

If you are coming from Node.js, the closest pattern is **`os.userInfo()`**.

Method on **`User`** — call it on a value of that type after constructing or receiving one from a constructor.

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

## Example

```go
import "github.com/sahilkhaire/gox/osutil"

osutil.UserInfo()
```

## Tips

Obtain a `User` value first (see constructors on the package overview), then call `UserInfo`.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to osutil package overview](/packages/osutil/)
