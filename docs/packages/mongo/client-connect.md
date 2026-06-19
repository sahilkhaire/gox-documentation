---
title: "Client.Connect"
package: "mongo"
import: "github.com/sahilkhaire/gox/mongo"
node: "mongoose.connect(uri)"
gox-doc-version: "10"
---

<SymbolHeader pkg="mongo" title="Client.Connect" node="mongoose.connect(uri)" import-path="github.com/sahilkhaire/gox/mongo" />
## Overview

Connect dials uri.

**Node.js equivalent:** `mongoose.connect(uri)`

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

← [Back to mongo package overview](/packages/mongo/)
