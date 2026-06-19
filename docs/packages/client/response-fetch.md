---
title: "Response.Fetch"
package: "client"
import: "github.com/sahilkhaire/gox/client"
node: "fetch(url)"
gox-doc-version: "10"
---

<SymbolHeader pkg="client" title="Response.Fetch" node="fetch(url)" import-path="github.com/sahilkhaire/gox/client" />
## Overview

Fetch is a fetch()-style helper using the package default client.

**Node.js equivalent:** `fetch(url)`

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

res, err := client.Fetch(ctx, url)
```

:::

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

```go
import "github.com/sahilkhaire/gox/client"

srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("ok"))
}))
defer srv.Close()
resp, err := Fetch(context.Background(), srv.URL, nil)
b, err := resp.Bytes()
```

← [Back to client package overview](/packages/client/)
