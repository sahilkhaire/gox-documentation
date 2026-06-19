---
title: "ReadDir"
package: "fs"
import: "github.com/sahilkhaire/gox/fs"
node: "fs.readdir(path)"
gox-doc-version: "11"
---

<SymbolHeader pkg="fs" title="ReadDir" node="fs.readdir(path)" import-path="github.com/sahilkhaire/gox/fs" />
## Overview

ReadDir reads directory entries (fs.promises.readdir).

If you are coming from Node.js, the closest pattern is **`fs.readdir(path)`**.

## Signature

<div class="signature-block">

```go
func ReadDir(ctx context.Context, path string) ([]os.DirEntry, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
fs.readdir(path)
```

```go [Standard Go]
// os / io helpers — e.g. os.readdir
```

```go [gox]
import "github.com/sahilkhaire/gox/fs"

fs.ReadDir(ctx, path)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/fs"

fs.ReadDir(ctx, path)
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

Use `os.ReadFile`, `os.WriteFile`, etc. Pass `context.Context` manually for cancellation.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/fs/copy">Copy</a><a class="related-chip" href="/packages/fs/exists">Exists</a><a class="related-chip" href="/packages/fs/mkdir">Mkdir</a>
</div>

← [Back to fs package overview](/packages/fs/)
