---
title: "Response.Fetch"
package: "client"
import: "github.com/sahilkhaire/gox/client"
node: "fetch(url)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: fetch(url)</span><span class="api-badge import">github.com/sahilkhaire/gox/client</span></div>
# Response.Fetch

## Overview

Fetch is a fetch()-style helper using the package default client.

**Node.js equivalent:** `fetch(url)`

## Signature

```go
func Fetch(ctx context.Context, url string, opts *RequestOpts) (*Response, error)
```

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

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

← [Back to client package overview](/packages/client/)
