---
title: "ParseUUID"
package: "id"
import: "github.com/sahilkhaire/gox/id"
node: "uuid.parse(s)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: uuid.parse(s)</span><span class="api-badge import">github.com/sahilkhaire/gox/id</span></div>
# ParseUUID

## Overview

Maps the Node.js pattern `uuid.parse(s)` to gox `id.ParseUUID(s)`.

**Node.js equivalent:** `uuid.parse(s)`

## Signature

```go
func ParseUUID(s string) (uuid.UUID, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
uuid.parse(s)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/id"

id.ParseUUID(s)
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
- [NewUUID](/packages/id/new-uuid)

← [Back to id package overview](/packages/id/)
