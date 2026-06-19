---
title: "MemoryStore.NewMemoryStore"
package: "http"
import: "github.com/sahilkhaire/gox/http"
gox-doc-version: "10"
---

<SymbolHeader pkg="http" title="MemoryStore.NewMemoryStore" node="express, cors, helmet, morgan" import-path="github.com/sahilkhaire/gox/http" />
## Overview

NewMemoryStore creates a memory session store.

## Signature

<div class="signature-block">

```go
func NewMemoryStore() *MemoryStore
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
func handler(w http.ResponseWriter, r *http.Request) {
    // chi or net/http
}
```

```go [gox]
import "github.com/sahilkhaire/gox/http"

var v MemoryStore
v.NewMemoryStore(/* args */)
```

:::

← [Back to http package overview](/packages/http/)
