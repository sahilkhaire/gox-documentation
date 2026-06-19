---
title: "FileEntry"
package: "archive"
import: "github.com/sahilkhaire/gox/archive"
gox-doc-version: "11"
---

<SymbolHeader pkg="archive" title="FileEntry" node="archiver" import-path="github.com/sahilkhaire/gox/archive" />
## Overview

FileEntry is a named file for archive creation.

Part of the **`archive`** package — Node.js analog: *archiver*.

`FileEntry` is a type exported by gox. Methods on this type are documented separately.

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
// Typical archiver pattern in Node.js
```

```go [Standard Go]
// archive/zip or archive/tar from stdlib
```

```go [gox]
import "github.com/sahilkhaire/gox/archive"

_ = archive.FileEntry
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/archive"

_ = archive.FileEntry
```

## Tips

Browse methods on this type in the sidebar for handler-style APIs and options structs.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/archive/tar-create">TarCreate</a><a class="related-chip" href="/packages/archive/tar-extract">TarExtract</a><a class="related-chip" href="/packages/archive/zip-create">ZipCreate</a>
</div>

← [Back to archive package overview](/packages/archive/)
