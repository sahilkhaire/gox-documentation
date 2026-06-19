---
title: "MemoryStore"
package: "http"
import: "github.com/sahilkhaire/gox/http"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: express, cors, helmet, morgan</span><span class="api-badge import">github.com/sahilkhaire/gox/http</span></div>
# MemoryStore

## Overview

MemoryStore is an in-memory session store (express-session memory).

## Signature

```go
type MemoryStore struct {
	// contains filtered or unexported fields
}
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
func handler(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(data)
}
```

```go [gox]
import "github.com/sahilkhaire/gox/http"

_ = http.MemoryStore
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [SaveUploadedFile](/packages/http/save-uploaded-file)

← [Back to http package overview](/packages/http/)
