---
title: "ReadAll"
package: "stream"
import: "github.com/sahilkhaire/gox/stream"
gox-doc-version: "10"
---

<SymbolHeader pkg="stream" title="ReadAll" node="Node stream" import-path="github.com/sahilkhaire/gox/stream" />
## Overview

ReadAll reads r until EOF or ctx is cancelled.

## Signature

<div class="signature-block">

```go
func ReadAll(ctx context.Context, r io.Reader) ([]byte, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
io.Copy(dst, src)
// or io.ReadAll(r)
```

```go [gox]
import "github.com/sahilkhaire/gox/stream"

// stream
_ = stream.ReadAll(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/stream/pipe">Pipe</a><a class="related-chip" href="/packages/stream/tee-reader">TeeReader</a><a class="related-chip" href="/packages/stream/transform">Transform</a>
</div>

← [Back to stream package overview](/packages/stream/)
