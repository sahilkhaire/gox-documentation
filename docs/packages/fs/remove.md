---
title: "Remove"
package: "fs"
import: "github.com/sahilkhaire/gox/fs"
node: "fs.rm(path)"
gox-doc-version: "10"
---

<SymbolHeader pkg="fs" title="Remove" node="fs.rm(path)" import-path="github.com/sahilkhaire/gox/fs" />
## Overview

Remove deletes path (fs.promises.rm).

**Node.js equivalent:** `fs.rm(path)`

## Signature

<div class="signature-block">

```go
func Remove(ctx context.Context, path string) error
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
fs.rm(path)
```

```go [Standard Go]
// os / io helpers — e.g. os.remove
```

```go [gox]
import "github.com/sahilkhaire/gox/fs"

fs.Remove(ctx, path)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/fs/copy">Copy</a><a class="related-chip" href="/packages/fs/exists">Exists</a><a class="related-chip" href="/packages/fs/mkdir">Mkdir</a>
</div>

← [Back to fs package overview](/packages/fs/)
