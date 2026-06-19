---
title: "Eq"
package: "mongo"
import: "github.com/sahilkhaire/gox/mongo"
gox-doc-version: "11"
---

<SymbolHeader pkg="mongo" title="Eq" node="mongoose" import-path="github.com/sahilkhaire/gox/mongo" />
## Overview

Eq matches equality on field.

Part of the **`mongo`** package — Node.js analog: *mongoose*.

## Signature

<div class="signature-block">

```go
func Eq(field string, value any) bson.M
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

// mongo
_ = mongo.Eq(/* args */)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/mongo"

// mongo
_ = mongo.Eq(/* args */)
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/mongo/gt">Gt</a><a class="related-chip" href="/packages/mongo/in">In</a><a class="related-chip" href="/packages/mongo/set">Set</a>
</div>

← [Back to mongo package overview](/packages/mongo/)
