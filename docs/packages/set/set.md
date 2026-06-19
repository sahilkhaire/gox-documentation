---
title: "Set"
package: "set"
import: "github.com/sahilkhaire/gox/set"
gox-doc-version: "11"
---

<SymbolHeader pkg="set" title="Set" node="Set" import-path="github.com/sahilkhaire/gox/set" />
## Overview

Set is a set of comparable values.

Part of the **`set`** package — Node.js analog: *Set*.

`Set` is a type exported by gox. Methods on this type are documented separately.

## Signature

<div class="signature-block">

```go
type Set[T comparable] map[T]struct{}
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical Set pattern in Node.js
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

## Example

```go
import "github.com/sahilkhaire/gox/set"

_ = set.Set
```

## Tips

Browse methods on this type in the sidebar for handler-style APIs and options structs.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to set package overview](/packages/set/)
