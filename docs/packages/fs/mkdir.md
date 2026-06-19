---
title: "Mkdir"
package: "fs"
import: "github.com/sahilkhaire/gox/fs"
node: "fs.mkdir(path)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: fs.mkdir(path)</span><span class="api-badge import">github.com/sahilkhaire/gox/fs</span></div>
# Mkdir

## Overview

Maps the Node.js pattern `fs.mkdir(path)` to gox `fs.Mkdir(ctx, path)`.

**Node.js equivalent:** `fs.mkdir(path)`

## Signature

```go
func Mkdir(ctx context.Context, path string, perm os.FileMode) error
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
fs.mkdir(path)
```

```go [Standard Go]
data, err := os.ReadFile(path)
```

```go [gox]
import "github.com/sahilkhaire/gox/fs"

fs.Mkdir(ctx, path)
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
- [Exists](/packages/fs/exists)
- [ReadDir](/packages/fs/read-dir)

← [Back to fs package overview](/packages/fs/)
