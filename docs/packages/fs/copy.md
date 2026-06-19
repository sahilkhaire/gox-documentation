---
title: "Copy"
package: "fs"
import: "github.com/sahilkhaire/gox/fs"
node: "fs.copyFile(src, dst)"
gox-doc-version: "10"
---

<SymbolHeader pkg="fs" title="Copy" node="fs.copyFile(src, dst)" import-path="github.com/sahilkhaire/gox/fs" />
## Overview

Copy copies src to dst (fs.promises.copyFile).

**Node.js equivalent:** `fs.copyFile(src, dst)`

## Signature

<div class="signature-block">

```go
func Copy(ctx context.Context, src, dst string) error
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
fs.copyFile(src, dst)
```

```go [Standard Go]
// os / io helpers — e.g. os.copy
```

```go [gox]
import "github.com/sahilkhaire/gox/fs"

fs.Copy(ctx, src, dst)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/fs/exists">Exists</a><a class="related-chip" href="/packages/fs/mkdir">Mkdir</a><a class="related-chip" href="/packages/fs/read-dir">ReadDir</a>
</div>

← [Back to fs package overview](/packages/fs/)
