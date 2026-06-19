---
title: "Buffer.FromString"
package: "buffer"
import: "github.com/sahilkhaire/gox/buffer"
node: "Buffer.from(str)"
gox-doc-version: "10"
---

<SymbolHeader pkg="buffer" title="Buffer.FromString" node="Buffer.from(str)" import-path="github.com/sahilkhaire/gox/buffer" />
## Overview

FromString creates a buffer from a string.

**Node.js equivalent:** `Buffer.from(str)`

## Signature

<div class="signature-block">

```go
func FromString(s string) Buffer
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
Buffer.from(str)
```

```go [Standard Go]
b := append([]byte(nil), parts...)
```

```go [gox]
import "github.com/sahilkhaire/gox/buffer"

buffer.FromString(str)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/buffer/buffer-concat">Buffer.Concat</a><a class="related-chip" href="/packages/buffer/buffer-from">Buffer.From</a>
</div>

← [Back to buffer package overview](/packages/buffer/)
