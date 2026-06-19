---
title: "TeeReader"
package: "stream"
import: "github.com/sahilkhaire/gox/stream"
gox-doc-version: "14"
---

<SymbolHeader pkg="stream" title="TeeReader" node="Node stream" import-path="github.com/sahilkhaire/gox/stream" />
## Overview

TeeReader returns a Reader that writes to w what it reads from r.
Node: passThrough with tap

Part of the **`stream`** package — Node.js analog: *Node stream*.

## Signature

<div class="signature-block">

```go
func TeeReader(r io.Reader, w io.Writer) io.Reader
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical Node stream pattern in Node.js
```

```go [Standard Go]
io.Copy(dst, src)
// or io.ReadAll(r)
```

```go [gox]
import "github.com/sahilkhaire/gox/stream"

reader := stream.TeeReader(r, auditWriter)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/stream"

reader := stream.TeeReader(r, auditWriter)
```

## Tips

Import `github.com/sahilkhaire/gox/stream` and call `TeeReader` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
io.Copy(dst, src)
// or io.ReadAll(r)
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/stream/pipe">Pipe</a><a class="related-chip" href="/packages/stream/read-all">ReadAll</a><a class="related-chip" href="/packages/stream/transform">Transform</a>
</div>

← [Back to stream package overview](/packages/stream/)
