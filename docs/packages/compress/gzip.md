---
title: "Gzip"
package: "compress"
import: "github.com/sahilkhaire/gox/compress"
node: "gzipSync"
gox-doc-version: "11"
---

<SymbolHeader pkg="compress" title="Gzip" node="gzipSync" import-path="github.com/sahilkhaire/gox/compress" />
## Overview

Gzip compresses data with gzip.

If you are coming from Node.js, the closest pattern is **`gzipSync`**.

## Signature

<div class="signature-block">

```go
func Gzip(data []byte) ([]byte, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
zlib.gzipSync(data);
```

```go [Standard Go]
// compress/gzip NewWriter / NewReader
```

```go [gox]
import "github.com/sahilkhaire/gox/compress"

compressed, err := compress.Gzip(data)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/compress"

compressed, err := compress.Gzip(data)
```

## Tips

Import `github.com/sahilkhaire/gox/compress` and call `Gzip` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/compress/gunzip">Gunzip</a><a class="related-chip" href="/packages/compress/gzip-reader">GzipReader</a><a class="related-chip" href="/packages/compress/gzip-writer">GzipWriter</a>
</div>

← [Back to compress package overview](/packages/compress/)
