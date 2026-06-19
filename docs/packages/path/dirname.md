---
title: "Dirname"
package: "path"
import: "github.com/sahilkhaire/gox/path"
node: "path.dirname(p)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: path.dirname(p)</span><span class="api-badge import">github.com/sahilkhaire/gox/path</span></div>
# Dirname

## Overview

Maps the Node.js pattern `path.dirname(p)` to gox `path.Dirname(p)`.

**Node.js equivalent:** `path.dirname(p)`

## Signature

```go
func Dirname(p string) string
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
path.dirname(p)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/path"

path.Dirname(p)
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
- [Extname](/packages/path/extname)
- [IsAbs](/packages/path/is-abs)

← [Back to path package overview](/packages/path/)
