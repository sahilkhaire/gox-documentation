---
title: "TarExtract"
package: "archive"
import: "github.com/sahilkhaire/gox/archive"
gox-doc-version: "10"
---

<SymbolHeader pkg="archive" title="TarExtract" node="archiver" import-path="github.com/sahilkhaire/gox/archive" />
## Overview

TarExtract extracts a gzip-compressed tar archive into destDir.

## Signature

<div class="signature-block">

```go
func TarExtract(data []byte, destDir string) error
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
_ = archive.TarExtract(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/archive/tar-create">TarCreate</a><a class="related-chip" href="/packages/archive/zip-create">ZipCreate</a><a class="related-chip" href="/packages/archive/zip-create-entries">ZipCreateEntries</a>
</div>

← [Back to archive package overview](/packages/archive/)
