---
title: "ZipCreate"
package: "archive"
import: "github.com/sahilkhaire/gox/archive"
gox-doc-version: "10"
---

<SymbolHeader pkg="archive" title="ZipCreate" node="archiver" import-path="github.com/sahilkhaire/gox/archive" />
## Overview

ZipCreate builds a zip archive from a map of path -&gt; contents.

## Signature

<div class="signature-block">

```go
func ZipCreate(files map[string][]byte) ([]byte, error)
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
_ = archive.ZipCreate(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/archive/tar-create">TarCreate</a><a class="related-chip" href="/packages/archive/tar-extract">TarExtract</a><a class="related-chip" href="/packages/archive/zip-create-entries">ZipCreateEntries</a>
</div>

← [Back to archive package overview](/packages/archive/)
