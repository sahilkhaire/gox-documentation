---
title: "User"
package: "osutil"
import: "github.com/sahilkhaire/gox/osutil"
gox-doc-version: "11"
---

<SymbolHeader pkg="osutil" title="User" node="os" import-path="github.com/sahilkhaire/gox/osutil" />
## Overview

User holds user identity fields when available (os.userInfo).

Part of the **`osutil`** package — Node.js analog: *os*.

`User` is a type exported by gox. Methods on this type are documented separately.

## Signature

<div class="signature-block">

```go
type User struct {
	Username string
	Uid      string
	Gid      string
}
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical os pattern in Node.js
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches.
```

```go [gox]
import "github.com/sahilkhaire/gox/osutil"

_ = osutil.User
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/osutil"

_ = osutil.User
```

## Tips

Browse methods on this type in the sidebar for handler-style APIs and options structs.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/osutil/arch">Arch</a><a class="related-chip" href="/packages/osutil/cp-us">CPUs</a><a class="related-chip" href="/packages/osutil/homedir">Homedir</a>
</div>

← [Back to osutil package overview](/packages/osutil/)
