---
title: "WriteFile"
package: "fs"
import: "github.com/sahilkhaire/gox/fs"
node: "fs.writeFile(path, data)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: fs.writeFile(path, data)</span><span class="api-badge import">github.com/sahilkhaire/gox/fs</span></div>
# WriteFile

## Overview

Maps the Node.js pattern `fs.writeFile(path, data)` to gox `fs.WriteFile(ctx, path, data)`.

**Node.js equivalent:** `fs.writeFile(path, data)`

## Signature

```go
func WriteFile(ctx context.Context, path string, data []byte, perm os.FileMode) error
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
fs.writeFile(path, data)
```

```go [Standard Go]
data, err := os.ReadFile(path)
```

```go [gox]
import "github.com/sahilkhaire/gox/fs"

fs.WriteFile(ctx, path, data)
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
