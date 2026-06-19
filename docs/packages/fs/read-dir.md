---
title: "ReadDir"
package: "fs"
import: "github.com/sahilkhaire/gox/fs"
node: "fs.readdir(path)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: fs.readdir(path)</span><span class="api-badge import">github.com/sahilkhaire/gox/fs</span></div>
# ReadDir

## Overview

Maps the Node.js pattern `fs.readdir(path)` to gox `fs.ReadDir(ctx, path)`.

**Node.js equivalent:** `fs.readdir(path)`

## Signature

```go
func ReadDir(ctx context.Context, path string) ([]os.DirEntry, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
fs.readdir(path)
```

```go [Standard Go]
data, err := os.ReadFile(path)
```

```go [gox]
import "github.com/sahilkhaire/gox/fs"

fs.ReadDir(ctx, path)
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
- [Mkdir](/packages/fs/mkdir)

← [Back to fs package overview](/packages/fs/)
