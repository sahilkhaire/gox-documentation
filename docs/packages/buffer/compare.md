---
title: "Compare"
package: "buffer"
import: "github.com/sahilkhaire/gox/buffer"
node: "buf.compare(other)"
gox-doc-version: "10"
---

<SymbolHeader pkg="buffer" title="Compare" node="buf.compare(other)" import-path="github.com/sahilkhaire/gox/buffer" />
## Overview

Compare compares two buffers lexicographically (Buffer.compare).

**Node.js equivalent:** `buf.compare(other)`

## Signature

<div class="signature-block">

```go
func Compare(a, b Buffer) int
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
buf.compare(other)
```

```go [Standard Go]
ok := bytes.Equal(a, b)
```

```go [gox]
import "github.com/sahilkhaire/gox/buffer"

buffer.Compare(a, b)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/buffer/equals">Equals</a>
</div>

← [Back to buffer package overview](/packages/buffer/)
