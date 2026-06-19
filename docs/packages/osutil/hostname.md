---
title: "Hostname"
package: "osutil"
import: "github.com/sahilkhaire/gox/osutil"
node: "os.hostname()"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: os.hostname()</span><span class="api-badge import">github.com/sahilkhaire/gox/osutil</span></div>
# Hostname

## Overview

Hostname returns the host name.

**Node.js equivalent:** `os.hostname()`

## Signature

```go
func Hostname() (string, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
os.hostname();
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/osutil"

name, err := osutil.Hostname()
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
