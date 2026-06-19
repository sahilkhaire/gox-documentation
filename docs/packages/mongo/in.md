---
title: "In"
package: "mongo"
import: "github.com/sahilkhaire/gox/mongo"
gox-doc-version: "14"
---

<SymbolHeader pkg="mongo" title="In" node="mongoose" import-path="github.com/sahilkhaire/gox/mongo" />
## Overview

In matches field in values.

Part of the **`mongo`** package — Node.js analog: *mongoose*.

## Signature

<div class="signature-block">

```go
func In(field string, values ...any) bson.M
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

filter := mongo.In("role", "admin", "editor")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/mongo"

filter := mongo.In("role", "admin", "editor")
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
<a class="related-chip" href="/packages/mongo/eq">Eq</a><a class="related-chip" href="/packages/mongo/gt">Gt</a><a class="related-chip" href="/packages/mongo/set">Set</a>
</div>

← [Back to mongo package overview](/packages/mongo/)
