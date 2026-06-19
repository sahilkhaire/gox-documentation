---
title: "Buffer.From"
package: "buffer"
import: "github.com/sahilkhaire/gox/buffer"
node: "Buffer.from(data)"
gox-doc-version: "11"
---

<SymbolHeader pkg="buffer" title="Buffer.From" node="Buffer.from(data)" import-path="github.com/sahilkhaire/gox/buffer" />
## Overview

From allocates a copy of data (Buffer.from).

If you are coming from Node.js, the closest pattern is **`Buffer.from(data)`**.

Method on **`Buffer`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func From(data []byte) Buffer
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
Buffer.from([1, 2, 3]);
```

```go [Standard Go]
b := append([]byte(nil), parts...)
```

```go [gox]
import "github.com/sahilkhaire/gox/buffer"

b := buffer.From([]byte{1, 2, 3})
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/buffer"

b := buffer.From([]byte{1, 2, 3})
```

## Tips

Obtain a `Buffer` value first (see constructors on the package overview), then call `From`.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/buffer/buffer-concat">Buffer.Concat</a><a class="related-chip" href="/packages/buffer/buffer-from-string">Buffer.FromString</a>
</div>

← [Back to buffer package overview](/packages/buffer/)
