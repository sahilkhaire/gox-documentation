---
title: "MustParseUUID"
package: "id"
import: "github.com/sahilkhaire/gox/id"
gox-doc-version: "11"
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

// id
_ = id.MustParseUUID(/* args */)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/id"

// id
_ = id.MustParseUUID(/* args */)
```

## Tips

Import `github.com/sahilkhaire/gox/id` and call `MustParseUUID` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/id/new-nano-id">NewNanoID</a><a class="related-chip" href="/packages/id/new-uuid">NewUUID</a><a class="related-chip" href="/packages/id/parse-uuid">ParseUUID</a>
</div>

← [Back to id package overview](/packages/id/)
