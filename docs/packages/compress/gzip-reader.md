---
title: "GzipReader"
package: "compress"
import: "github.com/sahilkhaire/gox/compress"
gox-doc-version: "10"
---

<SymbolHeader pkg="compress" title="GzipReader" node="zlib" import-path="github.com/sahilkhaire/gox/compress" />
## Overview

GzipReader wraps a gzip reader over src.

## Signature

<div class="signature-block">

```go
func GzipReader(src io.Reader) (*gzip.Reader, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/compress/gunzip">Gunzip</a><a class="related-chip" href="/packages/compress/gzip">Gzip</a><a class="related-chip" href="/packages/compress/gzip-writer">GzipWriter</a>
</div>

← [Back to compress package overview](/packages/compress/)
