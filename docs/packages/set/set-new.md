---
title: "Set.New"
package: "set"
import: "github.com/sahilkhaire/gox/set"
node: "new Set(arr)"
gox-doc-version: "14"
---

<SymbolHeader pkg="set" title="Set.New" node="new Set(arr)" import-path="github.com/sahilkhaire/gox/set" />
## Overview

New creates a set from items.

If you are coming from Node.js, the closest pattern is **`new Set(arr)`**.

Method on **`Set`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func New[T comparable](items ...T) Set[T]
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const s = new Set([1, 2, 3]);
```

```go [Standard Go]
s := make(map[T]struct{})
for _, v := range items { s[v] = struct{}{} }
```

```go [gox]
import "github.com/sahilkhaire/gox/set"

s := set.New(1, 2, 3)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/set"

s := set.New(1, 2, 3)
```

## Tips

Obtain a `Set` value first (see constructors on the package overview), then call `New`.

## Standard library alternative

Use the standard library directly:

```go
s := make(map[T]struct{})
for _, v := range items { s[v] = struct{}{} }
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/set/set-difference">Set.Difference</a><a class="related-chip" href="/packages/set/set-intersection">Set.Intersection</a><a class="related-chip" href="/packages/set/set-union">Set.Union</a>
</div>

← [Back to set package overview](/packages/set/)
