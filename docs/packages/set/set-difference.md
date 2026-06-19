---
title: "Set.Difference"
package: "set"
import: "github.com/sahilkhaire/gox/set"
gox-doc-version: "11"
---

<SymbolHeader pkg="set" title="Set.Difference" node="Set" import-path="github.com/sahilkhaire/gox/set" />
## Overview

Difference returns elements in a but not in b.

Part of the **`set`** package — Node.js analog: *Set*.

Method on **`Set`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func Difference[T comparable](a, b Set[T]) Set[T]
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical Set pattern in Node.js
```

```go [Standard Go]
// Iterate maps/sets manually
```

```go [gox]
import "github.com/sahilkhaire/gox/set"

var v Set
v.Difference(/* args */)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/set"

var v Set
v.Difference(/* args */)
```

## Tips

Obtain a `Set` value first (see constructors on the package overview), then call `Difference`.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/set/set-intersection">Set.Intersection</a><a class="related-chip" href="/packages/set/set-new">Set.New</a><a class="related-chip" href="/packages/set/set-union">Set.Union</a>
</div>

← [Back to set package overview](/packages/set/)
