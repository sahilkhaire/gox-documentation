---
title: "Homedir"
package: "osutil"
import: "github.com/sahilkhaire/gox/osutil"
node: "os.homedir()"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: os.homedir()</span><span class="api-badge import">github.com/sahilkhaire/gox/osutil</span></div>
# Homedir

## Overview

Maps the Node.js pattern `os.homedir()` to gox `osutil.Homedir()`.

**Node.js equivalent:** `os.homedir()`

## Signature

```go
func Homedir() (string, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
os.homedir()
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/osutil"

osutil.Homedir()
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
- [Hostname](/packages/osutil/hostname)

← [Back to osutil package overview](/packages/osutil/)
