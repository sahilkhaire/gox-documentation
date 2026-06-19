---
title: "User"
package: "osutil"
import: "github.com/sahilkhaire/gox/osutil"
gox-doc-version: "10"
---

<SymbolHeader pkg="osutil" title="User" node="os" import-path="github.com/sahilkhaire/gox/osutil" />
## Overview

User holds user identity fields when available (os.userInfo).

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
// See package overview
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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/osutil/arch">Arch</a><a class="related-chip" href="/packages/osutil/cp-us">CPUs</a><a class="related-chip" href="/packages/osutil/homedir">Homedir</a>
</div>

← [Back to osutil package overview](/packages/osutil/)
