---
title: "Collection"
package: "mongo"
import: "github.com/sahilkhaire/gox/mongo"
gox-doc-version: "11"
---

<SymbolHeader pkg="mongo" title="Collection" node="mongoose" import-path="github.com/sahilkhaire/gox/mongo" />
## Overview

Collection wraps mongo.Collection with helper methods.

Part of the **`mongo`** package — Node.js analog: *mongoose*.

`Collection` is a type exported by gox. Methods on this type are documented separately.

## Signature

<div class="signature-block">

```go
type Collection struct {
	// contains filtered or unexported fields
}
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical mongoose pattern in Node.js
```

```go [Standard Go]
client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
```

```go [gox]
import "github.com/sahilkhaire/gox/mongo"

_ = mongo.Collection
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/mongo"

_ = mongo.Collection
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/mongo/eq">Eq</a><a class="related-chip" href="/packages/mongo/gt">Gt</a><a class="related-chip" href="/packages/mongo/in">In</a>
</div>

← [Back to mongo package overview](/packages/mongo/)
