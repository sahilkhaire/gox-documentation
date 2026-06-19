---
title: "Remove"
package: "fs"
import: "github.com/sahilkhaire/gox/fs"
node: "fs.rm(path)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: fs.rm(path)</span><span class="api-badge import">github.com/sahilkhaire/gox/fs</span></div>
# Remove

## Overview

Maps the Node.js pattern `fs.rm(path)` to gox `fs.Remove(ctx, path)`.

**Node.js equivalent:** `fs.rm(path)`

## Signature

```go
func Remove(ctx context.Context, path string) error
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
fs.rm(path)
```

```go [Standard Go]
data, err := os.ReadFile(path)
```

```go [gox]
import "github.com/sahilkhaire/gox/fs"

fs.Remove(ctx, path)
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
