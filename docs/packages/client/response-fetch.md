---
title: "Response.Fetch"
package: "client"
import: "github.com/sahilkhaire/gox/client"
node: "fetch(url)"
gox-doc-version: "14"
---

<SymbolHeader pkg="client" title="Response.Fetch" node="fetch(url)" import-path="github.com/sahilkhaire/gox/client" />
## Overview

Fetch is a fetch()-style helper using the package default client.

If you are coming from Node.js, the closest pattern is **`fetch(url)`**.

Method on **`Response`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func Fetch(ctx context.Context, url string, opts *RequestOpts) (*Response, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const res = await fetch(url);
```

```go [Standard Go]
resp, err := http.Get(url)
```

```go [gox]
import "github.com/sahilkhaire/gox/client"

srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("ok"))
}))
defer srv.Close()
resp, err := Fetch(context.Background(), srv.URL, nil)
b, err := resp.Bytes()
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/client"

srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("ok"))
}))
defer srv.Close()
resp, err := Fetch(context.Background(), srv.URL, nil)
b, err := resp.Bytes()
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

Use the standard library directly:

```go
resp, err := http.Get(url)
```

← [Back to client package overview](/packages/client/)
