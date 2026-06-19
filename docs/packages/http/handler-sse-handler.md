---
title: "Handler.SSEHandler"
package: "http"
import: "github.com/sahilkhaire/gox/http"
gox-doc-version: "10"
---

<SymbolHeader pkg="http" title="Handler.SSEHandler" node="express, cors, helmet, morgan" import-path="github.com/sahilkhaire/gox/http" />
## Overview

SSEHandler runs fn with an EventStream (Express-style handler).

## Signature

<div class="signature-block">

```go
func SSEHandler(fn func(*EventStream) error) Handler
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

← [Back to http package overview](/packages/http/)
