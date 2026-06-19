---
title: "Set.Difference"
package: "set"
import: "github.com/sahilkhaire/gox/set"
gox-doc-version: "10"
---

<SymbolHeader pkg="set" title="Set.Difference" node="Set" import-path="github.com/sahilkhaire/gox/set" />
## Overview

Difference returns elements in a but not in b.

## Signature

<div class="signature-block">

```go
func Difference[T comparable](a, b Set[T]) Set[T]
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
v.Difference(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/set/set-intersection">Set.Intersection</a><a class="related-chip" href="/packages/set/set-new">Set.New</a><a class="related-chip" href="/packages/set/set-union">Set.Union</a>
</div>

← [Back to set package overview](/packages/set/)
