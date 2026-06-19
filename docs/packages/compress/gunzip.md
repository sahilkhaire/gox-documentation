---
title: "Gunzip"
package: "compress"
import: "github.com/sahilkhaire/gox/compress"
node: "gunzipSync"
gox-doc-version: "14"
---

<SymbolHeader pkg="compress" title="Gunzip" node="gunzipSync" import-path="github.com/sahilkhaire/gox/compress" />
## Overview

Gunzip decompresses gzip data.

If you are coming from Node.js, the closest pattern is **`gunzipSync`**.

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

## Example

```go
import "github.com/sahilkhaire/gox/compress"

compress.Gunzip(data)
```

## Tips

Import `github.com/sahilkhaire/gox/compress` and call `Gunzip` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/compress/gzip">Gzip</a><a class="related-chip" href="/packages/compress/gzip-reader">GzipReader</a><a class="related-chip" href="/packages/compress/gzip-writer">GzipWriter</a>
</div>

← [Back to compress package overview](/packages/compress/)
