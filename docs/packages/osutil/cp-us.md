---
title: "CPUs"
package: "osutil"
import: "github.com/sahilkhaire/gox/osutil"
node: "os.cpus().length"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: os.cpus().length</span><span class="api-badge import">github.com/sahilkhaire/gox/osutil</span></div>
# CPUs

## Overview

Maps the Node.js pattern `os.cpus().length` to gox `osutil.CPUs()`.

**Node.js equivalent:** `os.cpus().length`

## Signature

```go
func CPUs() int
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
os.cpus().length
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/osutil"

osutil.CPUs()
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
- [Homedir](/packages/osutil/homedir)
- [Hostname](/packages/osutil/hostname)

← [Back to osutil package overview](/packages/osutil/)
