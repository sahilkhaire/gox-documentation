---
title: "ReadAll"
package: "stream"
import: "github.com/sahilkhaire/gox/stream"
gox-doc-version: "14"
---

<SymbolHeader pkg="stream" title="ReadAll" node="Node stream" import-path="github.com/sahilkhaire/gox/stream" />
## Overview

ReadAll reads r until EOF or ctx is cancelled.

Part of the **`stream`** package — Node.js analog: *Node stream*.

## Signature

<div class="signature-block">

```go
func ReadAll(ctx context.Context, r io.Reader) ([]byte, error)
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

data, err := stream.ReadAll(ctx, r)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/stream"

data, err := stream.ReadAll(ctx, r)
```

## Tips

Import `github.com/sahilkhaire/gox/stream` and call `ReadAll` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
io.Copy(dst, src)
// or io.ReadAll(r)
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/stream/pipe">Pipe</a><a class="related-chip" href="/packages/stream/tee-reader">TeeReader</a><a class="related-chip" href="/packages/stream/transform">Transform</a>
</div>

← [Back to stream package overview](/packages/stream/)
