---
title: "Set"
package: "set"
import: "github.com/sahilkhaire/gox/set"
gox-doc-version: "10"
---

<SymbolHeader pkg="set" title="Set" node="Set" import-path="github.com/sahilkhaire/gox/set" />
## Overview

Set is a set of comparable values.

## Signature

<div class="signature-block">

```go
type Set[T comparable] map[T]struct{}
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
import "github.com/sahilkhaire/gox/set"

_ = set.Set
```

:::

← [Back to set package overview](/packages/set/)
