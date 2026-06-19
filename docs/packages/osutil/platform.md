---
title: "Platform"
package: "osutil"
import: "github.com/sahilkhaire/gox/osutil"
node: "process.platform"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: process.platform</span><span class="api-badge import">github.com/sahilkhaire/gox/osutil</span></div>
# Platform

## Overview

Maps the Node.js pattern `process.platform` to gox `osutil.Platform()`.

**Node.js equivalent:** `process.platform`

## Signature

```go
func Platform() string
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
process.platform
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/osutil"

osutil.Platform()
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
