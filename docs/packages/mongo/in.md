---
title: "In"
package: "mongo"
import: "github.com/sahilkhaire/gox/mongo"
gox-doc-version: "10"
---

<SymbolHeader pkg="mongo" title="In" node="mongoose" import-path="github.com/sahilkhaire/gox/mongo" />
## Overview

In matches field in values.

## Signature

<div class="signature-block">

```go
func In(field string, values ...any) bson.M
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
```

```go [gox]
import "github.com/sahilkhaire/gox/mongo"

// mongo
_ = mongo.In(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/mongo/eq">Eq</a><a class="related-chip" href="/packages/mongo/gt">Gt</a><a class="related-chip" href="/packages/mongo/set">Set</a>
</div>

← [Back to mongo package overview](/packages/mongo/)
