---
title: "MemoryStore.NewMemoryStore"
package: "http"
import: "github.com/sahilkhaire/gox/http"
gox-doc-version: "11"
---

<SymbolHeader pkg="http" title="MemoryStore.NewMemoryStore" node="express, cors, helmet, morgan" import-path="github.com/sahilkhaire/gox/http" />
## Overview

NewMemoryStore creates a memory session store.

Part of the **`http`** package — Node.js analog: *express, cors, helmet, morgan*.

Method on **`MemoryStore`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func NewMemoryStore() *MemoryStore
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical express, cors, helmet, morgan pattern in Node.js
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

## Example

```go
import "github.com/sahilkhaire/gox/http"

var v MemoryStore
v.NewMemoryStore(/* args */)
```

## Tips

Stack `Logger`, `Recover`, and `Security` middleware the way you would morgan + helmet in Express.

## Standard library alternative

Use `net/http` with handler functions `func(w http.ResponseWriter, r *http.Request)` or a router like chi/echo directly.

← [Back to http package overview](/packages/http/)
