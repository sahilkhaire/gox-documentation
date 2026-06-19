---
title: "SetDefaultClient"
package: "client"
import: "github.com/sahilkhaire/gox/client"
gox-doc-version: "14"
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

client.SetDefaultClient(c)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/client"

client.SetDefaultClient(c)
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

Use the standard library directly:

```go
resp, err := http.NewRequestWithContext(ctx, method, url, body)
client.Do(req)
```

← [Back to client package overview](/packages/client/)
