---
title: "ZipExtract"
package: "archive"
import: "github.com/sahilkhaire/gox/archive"
gox-doc-version: "10"
---

<SymbolHeader pkg="archive" title="ZipExtract" node="archiver" import-path="github.com/sahilkhaire/gox/archive" />
## Overview

ZipExtract writes zip data into destDir.

## Signature

<div class="signature-block">

```go
func ZipExtract(data []byte, destDir string) error
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
// archive/zip or archive/tar from stdlib
```

```go [gox]
import "github.com/sahilkhaire/gox/archive"

// archive
_ = archive.ZipExtract(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/archive/tar-create">TarCreate</a><a class="related-chip" href="/packages/archive/tar-extract">TarExtract</a><a class="related-chip" href="/packages/archive/zip-create">ZipCreate</a>
</div>

← [Back to archive package overview](/packages/archive/)
