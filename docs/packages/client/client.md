---
title: "Client"
package: "client"
import: "github.com/sahilkhaire/gox/client"
gox-doc-version: "14"
---

<SymbolHeader pkg="client" title="Client" node="axios, fetch" import-path="github.com/sahilkhaire/gox/client" />
## Overview

Client is a reusable HTTP client with defaults (axios instance).

Part of the **`client`** package — Node.js analog: *axios, fetch*.

`Client` is a type exported by gox. Methods on this type are documented separately.

## Signature

<div class="signature-block">

```go
type Client struct {
	HTTP           *http.Client
	BaseURL        string
	DefaultHeaders http.Header
	RetryOn5xx     bool
}
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

c := client.New()
resp, err := c.Get(ctx, "https://api.example.com/users", nil)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/client"

c := client.New()
resp, err := c.Get(ctx, "https://api.example.com/users", nil)
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

Use the standard library directly:

```go
resp, err := http.NewRequestWithContext(ctx, method, url, body)
client.Do(req)
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/client/set-default-client">SetDefaultClient</a>
</div>

← [Back to client package overview](/packages/client/)
