---
title: "MustParseUUID"
package: "id"
import: "github.com/sahilkhaire/gox/id"
gox-doc-version: "14"
---

<SymbolHeader pkg="id" title="MustParseUUID" node="uuid, nanoid" import-path="github.com/sahilkhaire/gox/id" />
## Overview

MustParseUUID parses s or panics.

Part of the **`id`** package — Node.js analog: *uuid, nanoid*.

## Signature

<div class="signature-block">

```go
func MustParseUUID(s string) uuid.UUID
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical uuid, nanoid pattern in Node.js
```

```go [Standard Go]
id := uuid.New()
u, err := uuid.Parse(s)
```

```go [gox]
import "github.com/sahilkhaire/gox/id"

id := id.MustParseUUID("550e8400-e29b-41d4-a716-446655440000")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/id"

id := id.MustParseUUID("550e8400-e29b-41d4-a716-446655440000")
```

## Tips

Import `github.com/sahilkhaire/gox/id` and call `MustParseUUID` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
id := uuid.New()
u, err := uuid.Parse(s)
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/id/new-nano-id">NewNanoID</a><a class="related-chip" href="/packages/id/new-uuid">NewUUID</a><a class="related-chip" href="/packages/id/parse-uuid">ParseUUID</a>
</div>

← [Back to id package overview](/packages/id/)
