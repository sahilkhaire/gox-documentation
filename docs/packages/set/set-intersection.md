---
title: "Set.Intersection"
package: "set"
import: "github.com/sahilkhaire/gox/set"
gox-doc-version: "10"
---

<SymbolHeader pkg="set" title="Set.Intersection" node="Set" import-path="github.com/sahilkhaire/gox/set" />
## Overview

Intersection returns elements in both a and b.

## Signature

<div class="signature-block">

```go
func Intersection[T comparable](a, b Set[T]) Set[T]
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
// Iterate maps/sets manually
```

```go [gox]
import "github.com/sahilkhaire/gox/set"

var v Set
v.Intersection(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/set/set-difference">Set.Difference</a><a class="related-chip" href="/packages/set/set-new">Set.New</a><a class="related-chip" href="/packages/set/set-union">Set.Union</a>
</div>

← [Back to set package overview](/packages/set/)
