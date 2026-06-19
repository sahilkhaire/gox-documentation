---
title: "NewNanoID"
package: "id"
import: "github.com/sahilkhaire/gox/id"
node: "nanoid(size)"
gox-doc-version: "11"
---

<SymbolHeader pkg="id" title="NewNanoID" node="nanoid(size)" import-path="github.com/sahilkhaire/gox/id" />
## Overview

NewNanoID returns a URL-safe ID using the default nanoid alphabet.

If you are coming from Node.js, the closest pattern is **`nanoid(size)`**.

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

## Example

```go
import "github.com/sahilkhaire/gox/id"

id := id.NewNanoID(21)
```

## Tips

Import `github.com/sahilkhaire/gox/id` and call `NewNanoID` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/id/must-parse-uuid">MustParseUUID</a><a class="related-chip" href="/packages/id/new-uuid">NewUUID</a><a class="related-chip" href="/packages/id/parse-uuid">ParseUUID</a>
</div>

← [Back to id package overview](/packages/id/)
