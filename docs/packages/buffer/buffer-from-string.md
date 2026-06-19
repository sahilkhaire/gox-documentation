---
title: "Buffer.FromString"
package: "buffer"
import: "github.com/sahilkhaire/gox/buffer"
node: "Buffer.from(str)"
gox-doc-version: "11"
---

<SymbolHeader pkg="buffer" title="Buffer.FromString" node="Buffer.from(str)" import-path="github.com/sahilkhaire/gox/buffer" />
## Overview

FromString creates a buffer from a string.

If you are coming from Node.js, the closest pattern is **`Buffer.from(str)`**.

Method on **`Buffer`** — call it on a value of that type after constructing or receiving one from a constructor.

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

## Example

```go
import "github.com/sahilkhaire/gox/buffer"

buffer.FromString(str)
```

## Tips

Obtain a `Buffer` value first (see constructors on the package overview), then call `FromString`.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/buffer/buffer-concat">Buffer.Concat</a><a class="related-chip" href="/packages/buffer/buffer-from">Buffer.From</a>
</div>

← [Back to buffer package overview](/packages/buffer/)
