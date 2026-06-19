---
title: "Client.New"
package: "client"
import: "github.com/sahilkhaire/gox/client"
node: "axios.create()"
gox-doc-version: "11"
---

<SymbolHeader pkg="client" title="Client.New" node="axios.create()" import-path="github.com/sahilkhaire/gox/client" />
## Overview

New creates a Client with optional timeout.

If you are coming from Node.js, the closest pattern is **`axios.create()`**.

Method on **`Client`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func New() *Client
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
axios.create()
```

```go [Standard Go]
resp, err := http.NewRequestWithContext(ctx, method, url, body)
client.Do(req)
```

```go [gox]
import "github.com/sahilkhaire/gox/client"

client.New()
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/client"

client.New()
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to client package overview](/packages/client/)
