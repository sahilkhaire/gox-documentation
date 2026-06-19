---
title: "Client"
package: "mongo"
import: "github.com/sahilkhaire/gox/mongo"
gox-doc-version: "14"
---

<SymbolHeader pkg="mongo" title="Client" node="mongoose" import-path="github.com/sahilkhaire/gox/mongo" />
## Overview

Client wraps mongo.Client.

Part of the **`mongo`** package — Node.js analog: *mongoose*.

`Client` is a type exported by gox. Methods on this type are documented separately.

## Signature

<div class="signature-block">

```go
type Client struct {
	MC *mongo.Client
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

client, err := mongo.Connect(ctx, "mongodb://localhost:27017")
if err != nil {
    return err
}
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/mongo"

client, err := mongo.Connect(ctx, "mongodb://localhost:27017")
if err != nil {
    return err
}
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
