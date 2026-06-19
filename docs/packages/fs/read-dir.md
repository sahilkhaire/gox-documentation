---
title: "ReadDir"
package: "fs"
import: "github.com/sahilkhaire/gox/fs"
node: "fs.readdir(path)"
gox-doc-version: "10"
---

<SymbolHeader pkg="fs" title="ReadDir" node="fs.readdir(path)" import-path="github.com/sahilkhaire/gox/fs" />
## Overview

ReadDir reads directory entries (fs.promises.readdir).

**Node.js equivalent:** `fs.readdir(path)`

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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/fs/copy">Copy</a><a class="related-chip" href="/packages/fs/exists">Exists</a><a class="related-chip" href="/packages/fs/mkdir">Mkdir</a>
</div>

← [Back to fs package overview](/packages/fs/)
