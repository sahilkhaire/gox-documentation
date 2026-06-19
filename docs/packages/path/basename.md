---
title: "Basename"
package: "path"
import: "github.com/sahilkhaire/gox/path"
node: "path.basename(p)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: path.basename(p)</span><span class="api-badge import">github.com/sahilkhaire/gox/path</span></div>
# Basename

## Overview

Basename returns the last element (path.basename).

**Node.js equivalent:** `path.basename(p)`

## Signature

```go
func Basename(p string) string
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
path.basename('/foo/bar.txt');
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/path"

name := path.Basename("/foo/bar.txt")
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Dirname](/packages/path/dirname)
- [Extname](/packages/path/extname)
- [IsAbs](/packages/path/is-abs)

← [Back to path package overview](/packages/path/)
