---
title: "Copy"
package: "fs"
import: "github.com/sahilkhaire/gox/fs"
node: "fs.copyFile(src, dst)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: fs.copyFile(src, dst)</span><span class="api-badge import">github.com/sahilkhaire/gox/fs</span></div>
# Copy

## Overview

Maps the Node.js pattern `fs.copyFile(src, dst)` to gox `fs.Copy(ctx, src, dst)`.

**Node.js equivalent:** `fs.copyFile(src, dst)`

## Signature

```go
func Copy(ctx context.Context, src, dst string) error
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
fs.copyFile(src, dst)
```

```go [Standard Go]
data, err := os.ReadFile(path)
```

```go [gox]
import "github.com/sahilkhaire/gox/fs"

fs.Copy(ctx, src, dst)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Exists](/packages/fs/exists)
- [Mkdir](/packages/fs/mkdir)
- [ReadDir](/packages/fs/read-dir)

← [Back to fs package overview](/packages/fs/)
