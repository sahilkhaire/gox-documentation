---
title: "MustParseUUID"
package: "id"
import: "github.com/sahilkhaire/gox/id"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: uuid, nanoid</span><span class="api-badge import">github.com/sahilkhaire/gox/id</span></div>
# MustParseUUID

## Overview

MustParseUUID parses s or panics.

## Signature

```go
func MustParseUUID(s string) uuid.UUID
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/id"

// id
_ = id.MustParseUUID(/* args */)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [NewNanoID](/packages/id/new-nano-id)
- [NewUUID](/packages/id/new-uuid)
- [ParseUUID](/packages/id/parse-uuid)

← [Back to id package overview](/packages/id/)
