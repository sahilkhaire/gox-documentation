---
title: "Transform"
package: "stream"
import: "github.com/sahilkhaire/gox/stream"
gox-doc-version: "10"
---

<SymbolHeader pkg="stream" title="Transform" node="Node stream" import-path="github.com/sahilkhaire/gox/stream" />
## Overview

Transform returns a Reader that applies fn to each chunk from reader.

## Signature

<div class="signature-block">

```go
func Transform(reader io.Reader, fn func([]byte) ([]byte, error)) io.Reader
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
_ = stream.Transform(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/stream/pipe">Pipe</a><a class="related-chip" href="/packages/stream/read-all">ReadAll</a><a class="related-chip" href="/packages/stream/tee-reader">TeeReader</a>
</div>

← [Back to stream package overview](/packages/stream/)
