---
title: "Buffer"
package: "buffer"
import: "github.com/sahilkhaire/gox/buffer"
gox-doc-version: "10"
---

<SymbolHeader pkg="buffer" title="Buffer" node="Buffer" import-path="github.com/sahilkhaire/gox/buffer" />
## Overview

Buffer is a byte slice with helper methods.

## Signature

<div class="signature-block">

```go
type Buffer []byte
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches.
```

```go [gox]
import "github.com/sahilkhaire/gox/buffer"

_ = buffer.Buffer
```

:::

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

```go
import "github.com/sahilkhaire/gox/buffer"

a := FromString("ab")
b := From([]byte("ab"))
c := Concat(FromString("a"), FromString("b"))
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/buffer/compare">Compare</a><a class="related-chip" href="/packages/buffer/equals">Equals</a>
</div>

← [Back to buffer package overview](/packages/buffer/)
