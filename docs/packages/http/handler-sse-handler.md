---
title: "Handler.SSEHandler"
package: "http"
import: "github.com/sahilkhaire/gox/http"
gox-doc-version: "14"
---

<SymbolHeader pkg="http" title="Handler.SSEHandler" node="express, cors, helmet, morgan" import-path="github.com/sahilkhaire/gox/http" />
## Overview

SSEHandler runs fn with an EventStream (Express-style handler).

Part of the **`http`** package — Node.js analog: *express, cors, helmet, morgan*.

Method on **`Handler`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func SSEHandler(fn func(*EventStream) error) Handler
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

:::

## Example

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

Stack `Logger`, `Recover`, and `Security` middleware the way you would morgan + helmet in Express.

## Standard library alternative

Use the standard library directly:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // chi or net/http
}
```

← [Back to http package overview](/packages/http/)
