---
title: "WriteFile"
package: "fs"
import: "github.com/sahilkhaire/gox/fs"
node: "fs.writeFile(path, data)"
gox-doc-version: "10"
---

<SymbolHeader pkg="fs" title="WriteFile" node="fs.writeFile(path, data)" import-path="github.com/sahilkhaire/gox/fs" />
## Overview

WriteFile writes data to path (fs.promises.writeFile).

**Node.js equivalent:** `fs.writeFile(path, data)`

## Signature

<div class="signature-block">

```go
func WriteFile(ctx context.Context, path string, data []byte, perm os.FileMode) error
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
await fs.promises.writeFile('out.txt', data);
```

```go [Standard Go]
err := os.WriteFile("out.txt", data, 0644)
```

```go [gox]
import "github.com/sahilkhaire/gox/fs"

err := fs.WriteFile(ctx, "out.txt", data)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/fs/copy">Copy</a><a class="related-chip" href="/packages/fs/exists">Exists</a><a class="related-chip" href="/packages/fs/mkdir">Mkdir</a>
</div>

← [Back to fs package overview](/packages/fs/)
