---
title: "Gunzip"
package: "compress"
import: "github.com/sahilkhaire/gox/compress"
node: "gunzipSync"
gox-doc-version: "10"
---

<SymbolHeader pkg="compress" title="Gunzip" node="gunzipSync" import-path="github.com/sahilkhaire/gox/compress" />
## Overview

Gunzip decompresses gzip data.

**Node.js equivalent:** `gunzipSync`

## Signature

<div class="signature-block">

```go
func Gunzip(data []byte) ([]byte, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
gunzipSync
```

```go [Standard Go]
// compress/gzip NewWriter / NewReader
```

```go [gox]
import "github.com/sahilkhaire/gox/compress"

compress.Gunzip(data)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/compress/gzip">Gzip</a><a class="related-chip" href="/packages/compress/gzip-reader">GzipReader</a><a class="related-chip" href="/packages/compress/gzip-writer">GzipWriter</a>
</div>

← [Back to compress package overview](/packages/compress/)
