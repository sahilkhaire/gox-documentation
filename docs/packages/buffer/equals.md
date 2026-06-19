---
title: "Equals"
package: "buffer"
import: "github.com/sahilkhaire/gox/buffer"
node: "buf.equals(other)"
gox-doc-version: "11"
---

<SymbolHeader pkg="buffer" title="Equals" node="buf.equals(other)" import-path="github.com/sahilkhaire/gox/buffer" />
## Overview

Equals reports whether a and b are equal (timing-safe enough for content equality).

If you are coming from Node.js, the closest pattern is **`buf.equals(other)`**.

## Signature

<div class="signature-block">

```go
func Equals(a, b Buffer) bool
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
buf.equals(other)
```

```go [Standard Go]
ok := bytes.Equal(a, b)
```

```go [gox]
import "github.com/sahilkhaire/gox/buffer"

buffer.Equals(a, b)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/buffer"

buffer.Equals(a, b)
```

## Tips

Import `github.com/sahilkhaire/gox/buffer` and call `Equals` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/buffer/compare">Compare</a>
</div>

← [Back to buffer package overview](/packages/buffer/)
