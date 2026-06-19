---
title: "mongo Cookbook"
package: "mongo"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: mongoose</span><span class="api-badge import">github.com/sahilkhaire/gox/mongo</span></div>
# mongo Cookbook

mongoose-style collections and filters.

```go
import "github.com/sahilkhaire/gox/mongo"

client, err := mongo.Connect(ctx, uri)
coll := client.DB("app").Collection("users")

var doc bson.M
err = coll.FindOne(ctx, mongo.Eq("email", email)).Decode(&doc)
```

← [Back to mongo overview](/packages/mongo/)
