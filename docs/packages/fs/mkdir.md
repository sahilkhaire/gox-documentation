---
title: "Mkdir"
package: "fs"
import: "github.com/sahilkhaire/gox/fs"
node: "fs.mkdir(path)"
gox-doc-version: "10"
---

<SymbolHeader pkg="fs" title="Mkdir" node="fs.mkdir(path)" import-path="github.com/sahilkhaire/gox/fs" />
## Overview

Mkdir creates a directory (fs.promises.mkdir).

**Node.js equivalent:** `fs.mkdir(path)`

## Signature

<div class="signature-block">

```go
func Mkdir(ctx context.Context, path string, perm os.FileMode) error
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
await fs.promises.mkdir('dir', { recursive: true });
```

```go [Standard Go]
err := os.MkdirAll("dir", 0755)
```

```go [gox]
import "github.com/sahilkhaire/gox/fs"

err := fs.Mkdir(ctx, "dir", 0755)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/fs/copy">Copy</a><a class="related-chip" href="/packages/fs/exists">Exists</a><a class="related-chip" href="/packages/fs/read-dir">ReadDir</a>
</div>

← [Back to fs package overview](/packages/fs/)
