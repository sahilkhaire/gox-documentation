---
title: "Collection"
package: "mongo"
import: "github.com/sahilkhaire/gox/mongo"
gox-doc-version: "14"
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

coll := db.Collection("users")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/mongo"

coll := db.Collection("users")
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

Use the standard library directly:

```go
client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/mongo/eq">Eq</a><a class="related-chip" href="/packages/mongo/gt">Gt</a><a class="related-chip" href="/packages/mongo/in">In</a>
</div>

← [Back to mongo package overview](/packages/mongo/)
