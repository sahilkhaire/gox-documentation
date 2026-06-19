---
title: "SetDefaultClient"
package: "client"
import: "github.com/sahilkhaire/gox/client"
gox-doc-version: "10"
---

<SymbolHeader pkg="client" title="SetDefaultClient" node="axios, fetch" import-path="github.com/sahilkhaire/gox/client" />
## Overview

SetDefaultClient replaces the client used by Fetch.

## Signature

<div class="signature-block">

```go
func SetDefaultClient(c *Client)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
resp, err := http.NewRequestWithContext(ctx, method, url, body)
client.Do(req)
```

```go [gox]
import "github.com/sahilkhaire/gox/client"

// client
_ = client.SetDefaultClient(/* args */)
```

:::

← [Back to client package overview](/packages/client/)
