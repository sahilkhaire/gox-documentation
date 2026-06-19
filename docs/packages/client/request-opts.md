---
title: "RequestOpts"
package: "client"
import: "github.com/sahilkhaire/gox/client"
gox-doc-version: "10"
---

<SymbolHeader pkg="client" title="RequestOpts" node="axios, fetch" import-path="github.com/sahilkhaire/gox/client" />
## Overview

RequestOpts configures a single request.

## Signature

<div class="signature-block">

```go
type RequestOpts struct {
	Method  string
	URL     string
	Body    io.Reader
	Headers http.Header
	Query   url.Values
	JSON    any
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

_ = client.RequestOpts
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/client/set-default-client">SetDefaultClient</a>
</div>

← [Back to client package overview](/packages/client/)
