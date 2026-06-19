---
title: "Handler.SSEHandler"
package: "http"
import: "github.com/sahilkhaire/gox/http"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: express, cors, helmet, morgan</span><span class="api-badge import">github.com/sahilkhaire/gox/http</span></div>
# Handler.SSEHandler

## Overview

SSEHandler runs fn with an EventStream (Express-style handler).

## Signature

```go
func SSEHandler(fn func(*EventStream) error) Handler
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

var v Handler
v.SSEHandler(/* args */)
```

:::

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

```go
import "github.com/sahilkhaire/gox/http"

app := New()
app.Get("/events", SSEHandler(func(es *EventStream) error {
	if err := es.Send("msg", "hello"); err != nil {
		return err
	}
	return es.SendData(`{"n":1}`)
}))
req := httptest.NewRequest(http.MethodGet, "/events", nil)
rec := httptest.NewRecorder()
app.ServeHTTP(rec, req)
sc := bufio.NewScanner(rec.Body)
```

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

← [Back to http package overview](/packages/http/)
