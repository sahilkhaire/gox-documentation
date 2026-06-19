---
title: "NewUUID"
package: "id"
import: "github.com/sahilkhaire/gox/id"
node: "uuid.v4()"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: uuid.v4()</span><span class="api-badge import">github.com/sahilkhaire/gox/id</span></div>
# NewUUID

## Overview

NewUUID returns a random UUID string (uuid.v4).

**Node.js equivalent:** `uuid.v4()`

## Signature

```go
func NewUUID() string
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
uuid.v4();
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/id"

id := id.NewUUID()
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [MustParseUUID](/packages/id/must-parse-uuid)
- [NewNanoID](/packages/id/new-nano-id)
- [ParseUUID](/packages/id/parse-uuid)

← [Back to id package overview](/packages/id/)
