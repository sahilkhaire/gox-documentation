---
title: "IsAbs"
package: "path"
import: "github.com/sahilkhaire/gox/path"
node: "path.isAbsolute(p)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: path.isAbsolute(p)</span><span class="api-badge import">github.com/sahilkhaire/gox/path</span></div>
# IsAbs

## Overview

Maps the Node.js pattern `path.isAbsolute(p)` to gox `path.IsAbs(p)`.

**Node.js equivalent:** `path.isAbsolute(p)`

## Signature

```go
func IsAbs(p string) bool
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
path.isAbsolute(p)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/path"

path.IsAbs(p)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Basename](/packages/path/basename)
- [Dirname](/packages/path/dirname)
- [Extname](/packages/path/extname)

← [Back to path package overview](/packages/path/)
