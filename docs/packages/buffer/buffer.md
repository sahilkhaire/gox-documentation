---
title: "Buffer"
package: "buffer"
import: "github.com/sahilkhaire/gox/buffer"
gox-doc-version: "14"
---

<SymbolHeader pkg="buffer" title="Buffer" node="Buffer" import-path="github.com/sahilkhaire/gox/buffer" />
## Overview

Buffer is a byte slice with helper methods.

Part of the **`buffer`** package — Node.js analog: *Buffer*.

`Buffer` is a type exported by gox. Methods on this type are documented separately.

## Signature

<div class="signature-block">

```go
type Buffer []byte
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical Buffer pattern in Node.js
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches.
```

```go [gox]
import "github.com/sahilkhaire/gox/buffer"

a := FromString("ab")
b := From([]byte("ab"))
c := Concat(FromString("a"), FromString("b"))
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/buffer"

a := FromString("ab")
b := From([]byte("ab"))
c := Concat(FromString("a"), FromString("b"))
```

## Tips

Browse methods on this type in the sidebar for handler-style APIs and options structs.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/buffer/compare">Compare</a><a class="related-chip" href="/packages/buffer/equals">Equals</a>
</div>

← [Back to buffer package overview](/packages/buffer/)
