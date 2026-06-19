---
title: "ZipCreateEntries"
package: "archive"
import: "github.com/sahilkhaire/gox/archive"
gox-doc-version: "11"
---

<SymbolHeader pkg="archive" title="ZipCreateEntries" node="archiver" import-path="github.com/sahilkhaire/gox/archive" />
## Overview

ZipCreateEntries builds a zip archive from file entries.

Part of the **`archive`** package — Node.js analog: *archiver*.

## Signature

<div class="signature-block">

```go
func ZipCreateEntries(entries []FileEntry) ([]byte, error)
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

// archive
_ = archive.ZipCreateEntries(/* args */)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/archive"

// archive
_ = archive.ZipCreateEntries(/* args */)
```

## Tips

Import `github.com/sahilkhaire/gox/archive` and call `ZipCreateEntries` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/archive/tar-create">TarCreate</a><a class="related-chip" href="/packages/archive/tar-extract">TarExtract</a><a class="related-chip" href="/packages/archive/zip-create">ZipCreate</a>
</div>

← [Back to archive package overview](/packages/archive/)
