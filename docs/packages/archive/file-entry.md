---
title: "FileEntry"
package: "archive"
import: "github.com/sahilkhaire/gox/archive"
gox-doc-version: "10"
---

<SymbolHeader pkg="archive" title="FileEntry" node="archiver" import-path="github.com/sahilkhaire/gox/archive" />
## Overview

FileEntry is a named file for archive creation.

## Signature

<div class="signature-block">

```go
type FileEntry struct {
	Name string
	Data []byte
}
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

_ = archive.FileEntry
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/archive/tar-create">TarCreate</a><a class="related-chip" href="/packages/archive/tar-extract">TarExtract</a><a class="related-chip" href="/packages/archive/zip-create">ZipCreate</a>
</div>

← [Back to archive package overview](/packages/archive/)
