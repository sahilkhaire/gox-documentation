---
title: "Client.Connect"
package: "mongo"
import: "github.com/sahilkhaire/gox/mongo"
node: "mongoose.connect(uri)"
gox-doc-version: "11"
---

<SymbolHeader pkg="mongo" title="Client.Connect" node="mongoose.connect(uri)" import-path="github.com/sahilkhaire/gox/mongo" />
## Overview

Connect dials uri.

If you are coming from Node.js, the closest pattern is **`mongoose.connect(uri)`**.

Method on **`Client`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func Connect(ctx context.Context, uri string) (*Client, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
await mongoose.connect(uri);
```

```go [Standard Go]
client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
```

```go [gox]
import "github.com/sahilkhaire/gox/mongo"

client, err := mongo.Connect(ctx, uri)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/mongo"

client, err := mongo.Connect(ctx, uri)
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to mongo package overview](/packages/mongo/)
