---
title: "Client"
package: "client"
import: "github.com/sahilkhaire/gox/client"
gox-doc-version: "10"
---

<SymbolHeader pkg="client" title="Client" node="axios, fetch" import-path="github.com/sahilkhaire/gox/client" />
## Overview

Client is a reusable HTTP client with defaults (axios instance).

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
// See package overview
```

```go [Standard Go]
resp, err := http.NewRequestWithContext(ctx, method, url, body)
client.Do(req)
```

```go [gox]
import "github.com/sahilkhaire/gox/client"

_ = client.Client
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/client/set-default-client">SetDefaultClient</a>
</div>

← [Back to client package overview](/packages/client/)
