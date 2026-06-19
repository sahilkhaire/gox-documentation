---
title: "ReadFile"
package: "fs"
import: "github.com/sahilkhaire/gox/fs"
node: "fs.readFile(path)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: fs.readFile(path)</span><span class="api-badge import">github.com/sahilkhaire/gox/fs</span></div>
# ReadFile

## Overview

ReadFile reads path (fs.promises.readFile).

**Node.js equivalent:** `fs.readFile(path)`

## Signature

```go
func ReadFile(ctx context.Context, path string) ([]byte, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const data = await fs.promises.readFile('file.txt');
```

```go [Standard Go]
data, err := os.ReadFile("file.txt")
```

```go [gox]
import "github.com/sahilkhaire/gox/fs"

data, err := fs.ReadFile(ctx, "file.txt")
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
