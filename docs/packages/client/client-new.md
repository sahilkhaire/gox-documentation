---
title: "Client.New"
package: "client"
import: "github.com/sahilkhaire/gox/client"
node: "axios.create()"
gox-doc-version: "10"
---

<SymbolHeader pkg="client" title="Client.New" node="axios.create()" import-path="github.com/sahilkhaire/gox/client" />
## Overview

New creates a Client with optional timeout.

**Node.js equivalent:** `axios.create()`

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

← [Back to client package overview](/packages/client/)
