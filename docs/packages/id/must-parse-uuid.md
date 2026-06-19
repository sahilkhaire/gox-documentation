---
title: "MustParseUUID"
package: "id"
import: "github.com/sahilkhaire/gox/id"
gox-doc-version: "10"
---

<SymbolHeader pkg="id" title="MustParseUUID" node="uuid, nanoid" import-path="github.com/sahilkhaire/gox/id" />
## Overview

MustParseUUID parses s or panics.

## Signature

<div class="signature-block">

```go
func MustParseUUID(s string) uuid.UUID
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/id/new-nano-id">NewNanoID</a><a class="related-chip" href="/packages/id/new-uuid">NewUUID</a><a class="related-chip" href="/packages/id/parse-uuid">ParseUUID</a>
</div>

← [Back to id package overview](/packages/id/)
