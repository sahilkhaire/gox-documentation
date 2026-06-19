---
title: "NewNanoID"
package: "id"
import: "github.com/sahilkhaire/gox/id"
node: "nanoid(size)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: nanoid(size)</span><span class="api-badge import">github.com/sahilkhaire/gox/id</span></div>
# NewNanoID

## Overview

NewNanoID returns a URL-safe ID using the default nanoid alphabet.

**Node.js equivalent:** `nanoid(size)`

## Signature

```go
func NewNanoID(size int) (string, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
nanoid();
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/id"

id := id.NewNanoID(21)
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
- [NewUUID](/packages/id/new-uuid)
- [ParseUUID](/packages/id/parse-uuid)

← [Back to id package overview](/packages/id/)
