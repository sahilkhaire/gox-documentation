---
title: "ZipExtract"
package: "archive"
import: "github.com/sahilkhaire/gox/archive"
gox-doc-version: "14"
---

<SymbolHeader pkg="archive" title="ZipExtract" node="archiver" import-path="github.com/sahilkhaire/gox/archive" />
## Overview

ZipExtract writes zip data into destDir.

Part of the **`archive`** package — Node.js analog: *archiver*.

## Signature

<div class="signature-block">

```go
func ZipExtract(data []byte, destDir string) error
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical archiver pattern in Node.js
```

```go [Standard Go]
// archive/zip or archive/tar from stdlib
```

```go [gox]
import "github.com/sahilkhaire/gox/archive"

err := archive.ZipExtract(ctx, "archive.zip", "./out")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/archive"

err := archive.ZipExtract(ctx, "archive.zip", "./out")
```

## Tips

Import `github.com/sahilkhaire/gox/archive` and call `ZipExtract` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/archive/tar-create">TarCreate</a><a class="related-chip" href="/packages/archive/tar-extract">TarExtract</a><a class="related-chip" href="/packages/archive/zip-create">ZipCreate</a>
</div>

← [Back to archive package overview](/packages/archive/)
