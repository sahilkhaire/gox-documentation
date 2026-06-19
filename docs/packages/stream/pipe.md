---
title: "Pipe"
package: "stream"
import: "github.com/sahilkhaire/gox/stream"
node: "src.pipe(dst)"
gox-doc-version: "11"
---

<SymbolHeader pkg="stream" title="Pipe" node="src.pipe(dst)" import-path="github.com/sahilkhaire/gox/stream" />
## Overview

Pipe copies from src to dst until EOF or error.
Node: src.pipe(dst)

If you are coming from Node.js, the closest pattern is **`src.pipe(dst)`**.

## Signature

<div class="signature-block">

```go
func Pipe(src io.Reader, dst io.Writer) (int64, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
src.pipe(dst)
```

```go [Standard Go]
io.Copy(dst, src)
// or io.ReadAll(r)
```

```go [gox]
import "github.com/sahilkhaire/gox/stream"

stream.Pipe(src, dst)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/stream"

stream.Pipe(src, dst)
```

## Tips

Import `github.com/sahilkhaire/gox/stream` and call `Pipe` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/stream/read-all">ReadAll</a><a class="related-chip" href="/packages/stream/tee-reader">TeeReader</a><a class="related-chip" href="/packages/stream/transform">Transform</a>
</div>

← [Back to stream package overview](/packages/stream/)
