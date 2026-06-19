---
title: "Arch"
package: "osutil"
import: "github.com/sahilkhaire/gox/osutil"
node: "process.arch"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: process.arch</span><span class="api-badge import">github.com/sahilkhaire/gox/osutil</span></div>
# Arch

## Overview

Maps the Node.js pattern `process.arch` to gox `osutil.Arch()`.

**Node.js equivalent:** `process.arch`

## Signature

```go
func Arch() string
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
process.arch
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/osutil"

osutil.Arch()
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [CPUs](/packages/osutil/cp-us)
- [Homedir](/packages/osutil/homedir)
- [Hostname](/packages/osutil/hostname)

← [Back to osutil package overview](/packages/osutil/)
