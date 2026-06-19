---
title: "Collection"
package: "mongo"
import: "github.com/sahilkhaire/gox/mongo"
gox-doc-version: "10"
---

<SymbolHeader pkg="mongo" title="Collection" node="mongoose" import-path="github.com/sahilkhaire/gox/mongo" />
## Overview

Collection wraps mongo.Collection with helper methods.

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
// See package overview
```

```go [Standard Go]
client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
```

```go [gox]
import "github.com/sahilkhaire/gox/mongo"

_ = mongo.Collection
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/mongo/eq">Eq</a><a class="related-chip" href="/packages/mongo/gt">Gt</a><a class="related-chip" href="/packages/mongo/in">In</a>
</div>

← [Back to mongo package overview](/packages/mongo/)
