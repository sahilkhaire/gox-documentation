---
title: "NewNanoID"
package: "id"
import: "github.com/sahilkhaire/gox/id"
node: "nanoid(size)"
gox-doc-version: "10"
---

<SymbolHeader pkg="id" title="NewNanoID" node="nanoid(size)" import-path="github.com/sahilkhaire/gox/id" />
## Overview

NewNanoID returns a URL-safe ID using the default nanoid alphabet.

**Node.js equivalent:** `nanoid(size)`

## Signature

<div class="signature-block">

```go
func NewNanoID(size int) (string, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
nanoid();
```

```go [Standard Go]
// generate random URL-safe string
```

```go [gox]
import "github.com/sahilkhaire/gox/id"

id := id.NewNanoID(21)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/id/must-parse-uuid">MustParseUUID</a><a class="related-chip" href="/packages/id/new-uuid">NewUUID</a><a class="related-chip" href="/packages/id/parse-uuid">ParseUUID</a>
</div>

← [Back to id package overview](/packages/id/)
