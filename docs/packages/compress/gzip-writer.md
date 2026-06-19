---
title: "GzipWriter"
package: "compress"
import: "github.com/sahilkhaire/gox/compress"
gox-doc-version: "10"
---

<SymbolHeader pkg="compress" title="GzipWriter" node="zlib" import-path="github.com/sahilkhaire/gox/compress" />
## Overview

GzipWriter wraps dst with a gzip writer; caller must Close the writer.

## Signature

<div class="signature-block">

```go
func GzipWriter(dst io.Writer) *gzip.Writer
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
_ = compress.GzipWriter(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/compress/gunzip">Gunzip</a><a class="related-chip" href="/packages/compress/gzip">Gzip</a><a class="related-chip" href="/packages/compress/gzip-reader">GzipReader</a>
</div>

← [Back to compress package overview](/packages/compress/)
