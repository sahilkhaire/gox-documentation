---
title: "Buffer.Concat"
package: "buffer"
import: "github.com/sahilkhaire/gox/buffer"
node: "Buffer.concat(list)"
gox-doc-version: "10"
---

<SymbolHeader pkg="buffer" title="Buffer.Concat" node="Buffer.concat(list)" import-path="github.com/sahilkhaire/gox/buffer" />
## Overview

Concat joins buffers (Buffer.concat).

**Node.js equivalent:** `Buffer.concat(list)`

## Signature

<div class="signature-block">

```go
func Concat(parts ...Buffer) Buffer
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
Buffer.concat([buf1, buf2]);
```

```go [Standard Go]
b := append([]byte(nil), parts...)
```

```go [gox]
import "github.com/sahilkhaire/gox/buffer"

combined := buffer.Concat(buf1, buf2)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/buffer/buffer-from">Buffer.From</a><a class="related-chip" href="/packages/buffer/buffer-from-string">Buffer.FromString</a>
</div>

← [Back to buffer package overview](/packages/buffer/)
