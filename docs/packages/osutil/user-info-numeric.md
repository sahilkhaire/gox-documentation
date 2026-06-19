---
title: "UserInfoNumeric"
package: "osutil"
import: "github.com/sahilkhaire/gox/osutil"
gox-doc-version: "10"
---

<SymbolHeader pkg="osutil" title="UserInfoNumeric" node="os" import-path="github.com/sahilkhaire/gox/osutil" />
## Overview

UserInfoNumeric is like UserInfo but parses uid/gid as ints when possible.

## Signature

<div class="signature-block">

```go
func UserInfoNumeric() (username string, uid, gid int, err error)
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

// osutil
_ = osutil.UserInfoNumeric(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/osutil/arch">Arch</a><a class="related-chip" href="/packages/osutil/cp-us">CPUs</a><a class="related-chip" href="/packages/osutil/homedir">Homedir</a>
</div>

← [Back to osutil package overview](/packages/osutil/)
