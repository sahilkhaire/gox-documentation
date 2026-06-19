---
title: "Buffer.Concat"
package: "buffer"
import: "github.com/sahilkhaire/gox/buffer"
node: "Buffer.concat(list)"
gox-doc-version: "11"
---

<SymbolHeader pkg="buffer" title="Buffer.Concat" node="Buffer.concat(list)" import-path="github.com/sahilkhaire/gox/buffer" />
## Overview

Concat joins buffers (Buffer.concat).

If you are coming from Node.js, the closest pattern is **`Buffer.concat(list)`**.

Method on **`Buffer`** — call it on a value of that type after constructing or receiving one from a constructor.

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

## Example

```go
import "github.com/sahilkhaire/gox/buffer"

combined := buffer.Concat(buf1, buf2)
```

## Tips

Obtain a `Buffer` value first (see constructors on the package overview), then call `Concat`.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/buffer/buffer-from">Buffer.From</a><a class="related-chip" href="/packages/buffer/buffer-from-string">Buffer.FromString</a>
</div>

← [Back to buffer package overview](/packages/buffer/)
