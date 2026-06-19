---
title: "Exists"
package: "fs"
import: "github.com/sahilkhaire/gox/fs"
node: "fs.exists(path)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: fs.exists(path)</span><span class="api-badge import">github.com/sahilkhaire/gox/fs</span></div>
# Exists

## Overview

Maps the Node.js pattern `fs.exists(path)` to gox `fs.Exists(ctx, path)`.

**Node.js equivalent:** `fs.exists(path)`

## Signature

```go
func Exists(ctx context.Context, path string) (bool, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
fs.exists(path)
```

```go [Standard Go]
data, err := os.ReadFile(path)
```

```go [gox]
import "github.com/sahilkhaire/gox/fs"

fs.Exists(ctx, path)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Copy](/packages/fs/copy)
- [Mkdir](/packages/fs/mkdir)
- [ReadDir](/packages/fs/read-dir)

← [Back to fs package overview](/packages/fs/)
