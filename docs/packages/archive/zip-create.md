---
title: "ZipCreate"
package: "archive"
import: "github.com/sahilkhaire/gox/archive"
gox-doc-version: "14"
---

<SymbolHeader pkg="archive" title="ZipCreate" node="archiver" import-path="github.com/sahilkhaire/gox/archive" />
## Overview

ZipCreate builds a zip archive from a map of path -&gt; contents.

Part of the **`archive`** package — Node.js analog: *archiver*.

## Signature

<div class="signature-block">

```go
func ZipCreate(files map[string][]byte) ([]byte, error)
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

err := archive.ZipCreate(ctx, "out.zip", entries)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/archive"

err := archive.ZipCreate(ctx, "out.zip", entries)
```

## Tips

Import `github.com/sahilkhaire/gox/archive` and call `ZipCreate` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/archive/tar-create">TarCreate</a><a class="related-chip" href="/packages/archive/tar-extract">TarExtract</a><a class="related-chip" href="/packages/archive/zip-create-entries">ZipCreateEntries</a>
</div>

← [Back to archive package overview](/packages/archive/)
