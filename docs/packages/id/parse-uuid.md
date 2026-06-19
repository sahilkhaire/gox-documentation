---
title: "ParseUUID"
package: "id"
import: "github.com/sahilkhaire/gox/id"
node: "uuid.parse(s)"
gox-doc-version: "10"
---

<SymbolHeader pkg="id" title="ParseUUID" node="uuid.parse(s)" import-path="github.com/sahilkhaire/gox/id" />
## Overview

ParseUUID parses a UUID string.

**Node.js equivalent:** `uuid.parse(s)`

## Signature

<div class="signature-block">

```go
func ParseUUID(s string) (uuid.UUID, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
uuid.parse(s)
```

```go [Standard Go]
id := uuid.New()
u, err := uuid.Parse(s)
```

```go [gox]
import "github.com/sahilkhaire/gox/id"

id.ParseUUID(s)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/id/must-parse-uuid">MustParseUUID</a><a class="related-chip" href="/packages/id/new-nano-id">NewNanoID</a><a class="related-chip" href="/packages/id/new-uuid">NewUUID</a>
</div>

← [Back to id package overview](/packages/id/)
