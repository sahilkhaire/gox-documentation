---
title: "TarExtract"
package: "archive"
import: "github.com/sahilkhaire/gox/archive"
gox-doc-version: "14"
---

<SymbolHeader pkg="archive" title="TarExtract" node="archiver" import-path="github.com/sahilkhaire/gox/archive" />
## Overview

TarExtract extracts a gzip-compressed tar archive into destDir.

Part of the **`archive`** package — Node.js analog: *archiver*.

## Signature

<div class="signature-block">

```go
func TarExtract(data []byte, destDir string) error
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

err := archive.TarExtract(ctx, r, "./out")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/archive"

err := archive.TarExtract(ctx, r, "./out")
```

## Tips

Import `github.com/sahilkhaire/gox/archive` and call `TarExtract` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/archive/tar-create">TarCreate</a><a class="related-chip" href="/packages/archive/zip-create">ZipCreate</a><a class="related-chip" href="/packages/archive/zip-create-entries">ZipCreateEntries</a>
</div>

← [Back to archive package overview](/packages/archive/)
