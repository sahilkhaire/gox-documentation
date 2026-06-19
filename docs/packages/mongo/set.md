---
title: "Set"
package: "mongo"
import: "github.com/sahilkhaire/gox/mongo"
gox-doc-version: "11"
---

<SymbolHeader pkg="mongo" title="Set" node="mongoose" import-path="github.com/sahilkhaire/gox/mongo" />
## Overview

Set returns bson.M for $set updates.

Part of the **`mongo`** package — Node.js analog: *mongoose*.

## Signature

<div class="signature-block">

```go
func Set(fields bson.M) bson.M
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
_ = mongo.Set(/* args */)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/mongo"

// mongo
_ = mongo.Set(/* args */)
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
