---
title: "SetDefaultClient"
package: "client"
import: "github.com/sahilkhaire/gox/client"
gox-doc-version: "11"
---

<SymbolHeader pkg="client" title="SetDefaultClient" node="axios, fetch" import-path="github.com/sahilkhaire/gox/client" />
## Overview

SetDefaultClient replaces the client used by Fetch.

Part of the **`client`** package — Node.js analog: *axios, fetch*.

## Signature

<div class="signature-block">

```go
func SetDefaultClient(c *Client)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical axios, fetch pattern in Node.js
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

## Example

```go
import "github.com/sahilkhaire/gox/client"

// client
_ = client.SetDefaultClient(/* args */)
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to client package overview](/packages/client/)
