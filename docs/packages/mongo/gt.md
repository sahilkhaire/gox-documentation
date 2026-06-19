---
title: "Gt"
package: "mongo"
import: "github.com/sahilkhaire/gox/mongo"
gox-doc-version: "10"
---

<SymbolHeader pkg="mongo" title="Gt" node="mongoose" import-path="github.com/sahilkhaire/gox/mongo" />
## Overview

Gt matches field &gt; value.

## Signature

<div class="signature-block">

```go
func Gt(field string, value any) bson.M
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
_ = mongo.Gt(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/mongo/eq">Eq</a><a class="related-chip" href="/packages/mongo/in">In</a><a class="related-chip" href="/packages/mongo/set">Set</a>
</div>

← [Back to mongo package overview](/packages/mongo/)
