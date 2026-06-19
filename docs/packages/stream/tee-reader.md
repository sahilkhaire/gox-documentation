---
title: "TeeReader"
package: "stream"
import: "github.com/sahilkhaire/gox/stream"
gox-doc-version: "10"
---

<SymbolHeader pkg="stream" title="TeeReader" node="Node stream" import-path="github.com/sahilkhaire/gox/stream" />
## Overview

TeeReader returns a Reader that writes to w what it reads from r.

## Signature

<div class="signature-block">

```go
func TeeReader(r io.Reader, w io.Writer) io.Reader
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
_ = stream.TeeReader(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/stream/pipe">Pipe</a><a class="related-chip" href="/packages/stream/read-all">ReadAll</a><a class="related-chip" href="/packages/stream/transform">Transform</a>
</div>

← [Back to stream package overview](/packages/stream/)
