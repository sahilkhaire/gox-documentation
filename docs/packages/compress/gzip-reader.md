---
title: "GzipReader"
package: "compress"
import: "github.com/sahilkhaire/gox/compress"
gox-doc-version: "11"
---

<SymbolHeader pkg="compress" title="GzipReader" node="zlib" import-path="github.com/sahilkhaire/gox/compress" />
## Overview

GzipReader wraps a gzip reader over src.

Part of the **`compress`** package — Node.js analog: *zlib*.

## Signature

<div class="signature-block">

```go
func GzipReader(src io.Reader) (*gzip.Reader, error)
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

// compress
_ = compress.GzipReader(/* args */)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/compress"

// compress
_ = compress.GzipReader(/* args */)
```

## Tips

Import `github.com/sahilkhaire/gox/compress` and call `GzipReader` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/compress/gunzip">Gunzip</a><a class="related-chip" href="/packages/compress/gzip">Gzip</a><a class="related-chip" href="/packages/compress/gzip-writer">GzipWriter</a>
</div>

← [Back to compress package overview](/packages/compress/)
