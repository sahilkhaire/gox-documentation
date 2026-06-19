---
title: "NewUUID"
package: "id"
import: "github.com/sahilkhaire/gox/id"
node: "uuid.v4()"
gox-doc-version: "14"
---

<SymbolHeader pkg="id" title="NewUUID" node="uuid.v4()" import-path="github.com/sahilkhaire/gox/id" />
## Overview

NewUUID returns a random UUID string (uuid.v4).

If you are coming from Node.js, the closest pattern is **`uuid.v4()`**.

## Signature

<div class="signature-block">

```go
func NewUUID() string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
uuid.v4();
```

```go [Standard Go]
id := uuid.New()
u, err := uuid.Parse(s)
```

```go [gox]
import "github.com/sahilkhaire/gox/id"

id := id.NewUUID()
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/id"

id := id.NewUUID()
```

## Tips

Import `github.com/sahilkhaire/gox/id` and call `NewUUID` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
id := uuid.New()
u, err := uuid.Parse(s)
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/id/must-parse-uuid">MustParseUUID</a><a class="related-chip" href="/packages/id/new-nano-id">NewNanoID</a><a class="related-chip" href="/packages/id/parse-uuid">ParseUUID</a>
</div>

← [Back to id package overview](/packages/id/)
