---
title: "GzipWriter"
package: "compress"
import: "github.com/sahilkhaire/gox/compress"
gox-doc-version: "14"
---

<SymbolHeader pkg="compress" title="GzipWriter" node="zlib" import-path="github.com/sahilkhaire/gox/compress" />
## Overview

GzipWriter wraps dst with a gzip writer; caller must Close the writer.

Part of the **`compress`** package — Node.js analog: *zlib*.

## Signature

<div class="signature-block">

```go
func GzipWriter(dst io.Writer) *gzip.Writer
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical zlib pattern in Node.js
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches.
```

```go [gox]
import "github.com/sahilkhaire/gox/compress"

compress.GzipWriter(w)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/compress"

compress.GzipWriter(w)
```

## Tips

Import `github.com/sahilkhaire/gox/compress` and call `GzipWriter` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/compress/gunzip">Gunzip</a><a class="related-chip" href="/packages/compress/gzip">Gzip</a><a class="related-chip" href="/packages/compress/gzip-reader">GzipReader</a>
</div>

← [Back to compress package overview](/packages/compress/)
